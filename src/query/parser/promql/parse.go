// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package promql

import (
	"fmt"

	"github.com/m3db/m3/src/query/errors"
	"github.com/m3db/m3/src/query/parser"

	pql "github.com/prometheus/prometheus/promql"
)

type promParser struct {
	expr pql.Expr
}

// Parse takes a promQL string and converts parses it into a DAG
func Parse(q string) (parser.Parser, error) {
	expr, err := pql.ParseExpr(q)
	if err != nil {
		return nil, err
	}

	return &promParser{
		expr: expr,
	}, nil
}

func (p *promParser) DAG() (parser.Nodes, parser.Edges, error) {
	state := &parseState{}
	err := state.walk(p.expr)
	if err != nil {
		return nil, nil, err
	}

	return state.transforms, state.edges, nil
}

func (p *promParser) String() string {
	return p.expr.String()
}

type parseState struct {
	edges      parser.Edges
	transforms parser.Nodes
}

func (p *parseState) lastTransformID() parser.NodeID {
	if len(p.transforms) == 0 {
		return parser.NodeID(-1)
	}

	return p.transforms[len(p.transforms)-1].ID
}

func (p *parseState) transformLen() int {
	return len(p.transforms)
}

func (p *parseState) walk(node pql.Node) error {
	if node == nil {
		return nil
	}

	switch n := node.(type) {
	case *pql.AggregateExpr:
		err := p.walk(n.Expr)
		if err != nil {
			return err
		}

		op, err := NewAggregationOperator(n)
		if err != nil {
			return err
		}

		opTransform := parser.NewTransformFromOperation(op, p.transformLen())
		p.edges = append(p.edges, parser.Edge{
			ParentID: p.lastTransformID(),
			ChildID:  opTransform.ID,
		})
		p.transforms = append(p.transforms, opTransform)
		// TODO: handle labels, params
		return nil

	case *pql.MatrixSelector:
		operation, err := NewSelectorFromMatrix(n)
		if err != nil {
			return err
		}

		p.transforms = append(p.transforms, parser.NewTransformFromOperation(operation, p.transformLen()))
		return nil

	case *pql.VectorSelector:
		operation, err := NewSelectorFromVector(n)
		if err != nil {
			return err
		}

		p.transforms = append(p.transforms, parser.NewTransformFromOperation(operation, p.transformLen()))
		return nil

	case *pql.Call:
		expressions := n.Args
		argValues := make([]interface{}, 0, len(expressions))
		for _, expr := range expressions {
			switch e := expr.(type) {
			case *pql.NumberLiteral:
				argValues = append(argValues, e.Val)
				continue
			}

			err := p.walk(expr)
			if err != nil {
				return err
			}
		}

		op, err := NewFunctionExpr(n.Func.Name, argValues)
		if err != nil {
			return err
		}

		opTransform := parser.NewTransformFromOperation(op, p.transformLen())
		p.edges = append(p.edges, parser.Edge{
			ParentID: p.lastTransformID(),
			ChildID:  opTransform.ID,
		})
		p.transforms = append(p.transforms, opTransform)
		return nil

	case *pql.BinaryExpr:
		err := p.walk(n.LHS)
		if err != nil {
			return err
		}

		lhsID := p.lastTransformID()
		err = p.walk(n.RHS)
		if err != nil {
			return err
		}

		rhsID := p.lastTransformID()
		op, err := NewBinaryOperator(n, lhsID, rhsID)
		if err != nil {
			return err
		}

		opTransform := parser.NewTransformFromOperation(op, p.transformLen())
		p.edges = append(p.edges, parser.Edge{
			ParentID: lhsID,
			ChildID:  opTransform.ID,
		})
		p.edges = append(p.edges, parser.Edge{
			ParentID: rhsID,
			ChildID:  opTransform.ID,
		})
		p.transforms = append(p.transforms, opTransform)
		return nil

	case *pql.NumberLiteral:
		op := NewScalarOperator(n)
		opTransform := parser.NewTransformFromOperation(op, p.transformLen())
		p.transforms = append(p.transforms, opTransform)
		return nil

	case *pql.ParenExpr:
		// Evaluate inside of paren expressions
		return p.walk(n.Expr)

	default:
		return fmt.Errorf("promql.Walk: unhandled node type %T, %v", node, node)
	}

	// TODO: This should go away once all cases have been implemented
	return errors.ErrNotImplemented
}
