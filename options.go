package consumer

import (
	"log/slog"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/prometheus/client_golang/prometheus"
)

// Option is used to override defaults when creating a new Consumer
type Option func(*Consumer)

// WithGroup overrides the default storage
func WithGroup(group Group) Option {
	return func(c *Consumer) {
		c.group = group
	}
}

// WithStore overrides the default storage
func WithStore(store Store) Option {
	return func(c *Consumer) {
		c.store = store
	}
}

// WithLogger overrides the default logger
func WithLogger(logger *slog.Logger) Option {
	return func(c *Consumer) {
		c.logger = logger
	}
}

// WithCounter overrides the default counter.
// Deprecated. Will be removed in favor of WithMetricRegistry in a future release.
func WithCounter(counter Counter) Option {
	return func(c *Consumer) {
		c.counter = counter
	}
}

// WithMetricRegistry specifies a registry to add the prometheus metrics to. Defaults to nil.
func WithMetricRegistry(registry prometheus.Registerer) Option {
	return func(c *Consumer) {
		c.metricRegistry = registry
	}
}

// WithClient overrides the default client
func WithClient(client kinesisClient) Option {
	return func(c *Consumer) {
		c.client = client
	}
}

// WithShardIteratorType overrides the starting point for the consumer
func WithShardIteratorType(t string) Option {
	return func(c *Consumer) {
		c.initialShardIteratorType = types.ShardIteratorType(t)
	}
}

// WithTimestamp overrides the starting point for the consumer
func WithTimestamp(t time.Time) Option {
	return func(c *Consumer) {
		c.initialTimestamp = &t
	}
}

// WithScanInterval overrides the scan interval for the consumer
func WithScanInterval(d time.Duration) Option {
	return func(c *Consumer) {
		c.scanInterval = d
	}
}

// WithMaxRecords overrides the maximum number of records to be
// returned in a single GetRecords call for the consumer (specify a
// value of up to 10,000)
func WithMaxRecords(n int64) Option {
	return func(c *Consumer) {
		c.maxRecords = n
	}
}

// WithGetRecordsOptions passes the given option functions to the
// kinesis client's GetRecords call
func WithGetRecordsOptions(opts ...func(*kinesis.Options)) Option {
	return func(c *Consumer) {
		c.getRecordsOpts = opts
	}
}

// WithAggregation overrides the default option for aggregating records
func WithAggregation(a bool) Option {
	return func(c *Consumer) {
		c.isAggregated = a
	}
}

// WithParallelProcessing sets the size of the Worker Pool that processes incoming requests. Defaults to 1
func WithParallelProcessing(numWorkers int) Option {
	return func(c *Consumer) {
		c.numWorkers = numWorkers
	}
}

// WithShardClosedHandler defines a custom handler for closed shards.
func WithShardClosedHandler(h ShardClosedHandler) Option {
	return func(c *Consumer) {
		c.shardClosedHandler = h
	}
}

// ShardClosedHandler is a handler that will be called when the consumer has reached the end of a closed shard.
// No more records for that shard will be provided by the consumer.
// An error can be returned to stop the consumer.
type ShardClosedHandler = func(streamName, shardID string) error
