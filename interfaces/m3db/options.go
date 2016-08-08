	"os"
	"github.com/m3db/m3x/log"
	"github.com/m3db/m3x/metrics"

	tchannel "github.com/uber/tchannel-go"
	// EncodingTsz sets tsz encoding and returns a new DatabaseOptions
	EncodingTsz() DatabaseOptions

	Logger(value xlog.Logger) DatabaseOptions
	GetLogger() xlog.Logger
	MetricsScope(value xmetrics.Scope) DatabaseOptions
	GetMetricsScope() xmetrics.Scope
	// CommitLogFlushSize sets the commit log flush size and returns a new DatabaseOptions
	CommitLogFlushSize(value int) DatabaseOptions

	// GetCommitLogFlushSize returns the commit log flush size
	GetCommitLogFlushSize() int

	// CommitLogFlushInterval sets the commit log flush interval and returns a new DatabaseOptions
	CommitLogFlushInterval(value time.Duration) DatabaseOptions

	// GetCommitLogFlushInterval returns the commit log flush interval
	GetCommitLogFlushInterval() time.Duration

	// CommitLogBacklogQueueSize sets the commit log backlog queue size and returns a new DatabaseOptions
	CommitLogBacklogQueueSize(value int) DatabaseOptions

	// GetCommitLogBacklogQueueSize returns the commit log backlog queue size
	GetCommitLogBacklogQueueSize() int

	// CommitLogStrategy sets the commit log strategy and returns a new DatabaseOptions
	CommitLogStrategy(value CommitLogStrategy) DatabaseOptions

	// GetCommitLogStrategy returns the commit log strategy
	GetCommitLogStrategy() CommitLogStrategy

	// SegmentReaderPool sets the segment reader pool.
	SegmentReaderPool(value SegmentReaderPool) DatabaseOptions

	// GetSegmentReaderPool returns the segment reader pool.
	GetSegmentReaderPool() SegmentReaderPool
	// ReaderIteratorPool sets the readerIteratorPool and returns a new DatabaseOptions
	ReaderIteratorPool(value ReaderIteratorPool) DatabaseOptions
	// GetReaderIteratorPool returns the readerIteratorPool
	GetReaderIteratorPool() ReaderIteratorPool

	// MultiReaderIteratorPool sets the multiReaderIteratorPool and returns a new DatabaseOptions
	// GetMultiReaderIteratorPool returns the multiReaderIteratorPool
	// FileWriterOptions sets the file writer options.
	FileWriterOptions(value FileWriterOptions) DatabaseOptions

	// GetFileWriterOptions returns the file writer options.
	GetFileWriterOptions() FileWriterOptions


	// NewPersistenceManagerFn sets the function for creating a new persistence manager.
	NewPersistenceManagerFn(value NewPersistenceManagerFn) DatabaseOptions

	// GetNewPersistenceManagerFn returns the function for creating a new persistence manager.
	GetNewPersistenceManagerFn() NewPersistenceManagerFn

	// WriterBufferSize sets the buffer size for writing TSDB files.
	WriterBufferSize(value int) DatabaseOptions

	// GetWriterBufferSize returns the buffer size for writing TSDB files.
	GetWriterBufferSize() int

	// ReaderBufferSize sets the buffer size for reading TSDB files.
	ReaderBufferSize(value int) DatabaseOptions

	// GetReaderBufferSize returns the buffer size for reading TSDB files.
	GetReaderBufferSize() int
	// Validate validates the options
	Validate() error

	// EncodingTsz sets tsz encoding and returns a new ClientOptions
	EncodingTsz() ClientOptions

	Logger(value xlog.Logger) ClientOptions
	GetLogger() xlog.Logger
	MetricsScope(value xmetrics.Scope) ClientOptions
	GetMetricsScope() xmetrics.Scope
	// ClusterConnectConsistencyLevel sets the clusterConnectConsistencyLevel and returns a new ClientOptions
	ClusterConnectConsistencyLevel(value ConsistencyLevel) ClientOptions

	// GetClusterConnectConsistencyLevel returns the clusterConnectConsistencyLevel
	GetClusterConnectConsistencyLevel() ConsistencyLevel

	// FetchRequestTimeout sets the fetchRequestTimeout and returns a new ClientOptions
	FetchRequestTimeout(value time.Duration) ClientOptions

	// GetFetchRequestTimeout returns the fetchRequestTimeout
	GetFetchRequestTimeout() time.Duration

	// FetchBatchOpPoolSize sets the fetchBatchOpPoolSize and returns a new ClientOptions
	FetchBatchOpPoolSize(value int) ClientOptions

	// GetFetchBatchOpPoolSize returns the fetchBatchOpPoolSize
	GetFetchBatchOpPoolSize() int

	// FetchBatchSize sets the fetchBatchSize and returns a new ClientOptions
	// NB(r): for a fetch only application load this should match the host
	// queue ops flush size so that each time a host queue is flushed it can
	// fit the entire flushed fetch ops into a single batch.
	FetchBatchSize(value int) ClientOptions

	// GetFetchBatchSize returns the fetchBatchSize
	GetFetchBatchSize() int


	// SeriesIteratorPoolSize sets the seriesIteratorPoolSize and returns a new ClientOptions
	SeriesIteratorPoolSize(value int) ClientOptions

	// GetSeriesIteratorPoolSize returns the seriesIteratorPoolSize
	GetSeriesIteratorPoolSize() int

	// SeriesIteratorArrayPoolBuckets sets the seriesIteratorArrayPoolBuckets and returns a new ClientOptions
	SeriesIteratorArrayPoolBuckets(value []PoolBucket) ClientOptions

	// GetSeriesIteratorArrayPoolBuckets returns the seriesIteratorArrayPoolBuckets
	GetSeriesIteratorArrayPoolBuckets() []PoolBucket

	// ReaderIteratorAllocate sets the readerIteratorAllocate and returns a new ClientOptions
	ReaderIteratorAllocate(value ReaderIteratorAllocate) ClientOptions

	// GetReaderIteratorAllocate returns the readerIteratorAllocate
	GetReaderIteratorAllocate() ReaderIteratorAllocate
// TopologyTypeOptions is a set of static topology type options
type TopologyTypeOptions interface {
	// ShardScheme sets the shardScheme and returns a new TopologyTypeOptions
	ShardScheme(value ShardScheme) TopologyTypeOptions
	// Replicas sets the replicas and returns a new TopologyTypeOptions
	Replicas(value int) TopologyTypeOptions
	// HostShardSets sets the hostShardSets and returns a new TopologyTypeOptions
	HostShardSets(value []HostShardSet) TopologyTypeOptions

// FileWriterOptions is a set of file writing options for a file writer
type FileWriterOptions interface {
	// NewFileMode sets the new file mode.
	NewFileMode(value os.FileMode) FileWriterOptions

	// GetNewFileMode returns the new file mode.
	GetNewFileMode() os.FileMode

	// NewDirectoryMode sets the new directory mode.
	NewDirectoryMode(value os.FileMode) FileWriterOptions

	// GetNewDirectoryMode returns the new directory mode.
	GetNewDirectoryMode() os.FileMode
}