// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/m3ninx/index/segment (interfaces: Segment,MutableSegment)

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

// Package segment is a generated GoMock package.
package segment

import (
	"reflect"

	"github.com/m3db/m3/src/m3ninx/doc"
	"github.com/m3db/m3/src/m3ninx/index"

	"github.com/golang/mock/gomock"
)

// MockSegment is a mock of Segment interface
type MockSegment struct {
	ctrl     *gomock.Controller
	recorder *MockSegmentMockRecorder
}

// MockSegmentMockRecorder is the mock recorder for MockSegment
type MockSegmentMockRecorder struct {
	mock *MockSegment
}

// NewMockSegment creates a new mock instance
func NewMockSegment(ctrl *gomock.Controller) *MockSegment {
	mock := &MockSegment{ctrl: ctrl}
	mock.recorder = &MockSegmentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSegment) EXPECT() *MockSegmentMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockSegment) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockSegmentMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSegment)(nil).Close))
}

// ContainsID mocks base method
func (m *MockSegment) ContainsID(arg0 []byte) (bool, error) {
	ret := m.ctrl.Call(m, "ContainsID", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainsID indicates an expected call of ContainsID
func (mr *MockSegmentMockRecorder) ContainsID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainsID", reflect.TypeOf((*MockSegment)(nil).ContainsID), arg0)
}

// Fields mocks base method
func (m *MockSegment) Fields() ([][]byte, error) {
	ret := m.ctrl.Call(m, "Fields")
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fields indicates an expected call of Fields
func (mr *MockSegmentMockRecorder) Fields() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fields", reflect.TypeOf((*MockSegment)(nil).Fields))
}

// Reader mocks base method
func (m *MockSegment) Reader() (index.Reader, error) {
	ret := m.ctrl.Call(m, "Reader")
	ret0, _ := ret[0].(index.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reader indicates an expected call of Reader
func (mr *MockSegmentMockRecorder) Reader() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reader", reflect.TypeOf((*MockSegment)(nil).Reader))
}

// Size mocks base method
func (m *MockSegment) Size() int64 {
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockSegmentMockRecorder) Size() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockSegment)(nil).Size))
}

// Terms mocks base method
func (m *MockSegment) Terms(arg0 []byte) ([][]byte, error) {
	ret := m.ctrl.Call(m, "Terms", arg0)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Terms indicates an expected call of Terms
func (mr *MockSegmentMockRecorder) Terms(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Terms", reflect.TypeOf((*MockSegment)(nil).Terms), arg0)
}

// MockMutableSegment is a mock of MutableSegment interface
type MockMutableSegment struct {
	ctrl     *gomock.Controller
	recorder *MockMutableSegmentMockRecorder
}

// MockMutableSegmentMockRecorder is the mock recorder for MockMutableSegment
type MockMutableSegmentMockRecorder struct {
	mock *MockMutableSegment
}

// NewMockMutableSegment creates a new mock instance
func NewMockMutableSegment(ctrl *gomock.Controller) *MockMutableSegment {
	mock := &MockMutableSegment{ctrl: ctrl}
	mock.recorder = &MockMutableSegmentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMutableSegment) EXPECT() *MockMutableSegmentMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockMutableSegment) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockMutableSegmentMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMutableSegment)(nil).Close))
}

// ContainsID mocks base method
func (m *MockMutableSegment) ContainsID(arg0 []byte) (bool, error) {
	ret := m.ctrl.Call(m, "ContainsID", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainsID indicates an expected call of ContainsID
func (mr *MockMutableSegmentMockRecorder) ContainsID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainsID", reflect.TypeOf((*MockMutableSegment)(nil).ContainsID), arg0)
}

// Fields mocks base method
func (m *MockMutableSegment) Fields() ([][]byte, error) {
	ret := m.ctrl.Call(m, "Fields")
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fields indicates an expected call of Fields
func (mr *MockMutableSegmentMockRecorder) Fields() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fields", reflect.TypeOf((*MockMutableSegment)(nil).Fields))
}

// Insert mocks base method
func (m *MockMutableSegment) Insert(arg0 doc.Document) ([]byte, error) {
	ret := m.ctrl.Call(m, "Insert", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockMutableSegmentMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockMutableSegment)(nil).Insert), arg0)
}

// InsertBatch mocks base method
func (m *MockMutableSegment) InsertBatch(arg0 index.Batch) error {
	ret := m.ctrl.Call(m, "InsertBatch", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertBatch indicates an expected call of InsertBatch
func (mr *MockMutableSegmentMockRecorder) InsertBatch(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBatch", reflect.TypeOf((*MockMutableSegment)(nil).InsertBatch), arg0)
}

// IsSealed mocks base method
func (m *MockMutableSegment) IsSealed() bool {
	ret := m.ctrl.Call(m, "IsSealed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSealed indicates an expected call of IsSealed
func (mr *MockMutableSegmentMockRecorder) IsSealed() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSealed", reflect.TypeOf((*MockMutableSegment)(nil).IsSealed))
}

// Reader mocks base method
func (m *MockMutableSegment) Reader() (index.Reader, error) {
	ret := m.ctrl.Call(m, "Reader")
	ret0, _ := ret[0].(index.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reader indicates an expected call of Reader
func (mr *MockMutableSegmentMockRecorder) Reader() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reader", reflect.TypeOf((*MockMutableSegment)(nil).Reader))
}

// Seal mocks base method
func (m *MockMutableSegment) Seal() (Segment, error) {
	ret := m.ctrl.Call(m, "Seal")
	ret0, _ := ret[0].(Segment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Seal indicates an expected call of Seal
func (mr *MockMutableSegmentMockRecorder) Seal() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seal", reflect.TypeOf((*MockMutableSegment)(nil).Seal))
}

// Size mocks base method
func (m *MockMutableSegment) Size() int64 {
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockMutableSegmentMockRecorder) Size() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockMutableSegment)(nil).Size))
}

// Terms mocks base method
func (m *MockMutableSegment) Terms(arg0 []byte) ([][]byte, error) {
	ret := m.ctrl.Call(m, "Terms", arg0)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Terms indicates an expected call of Terms
func (mr *MockMutableSegmentMockRecorder) Terms(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Terms", reflect.TypeOf((*MockMutableSegment)(nil).Terms), arg0)
}
