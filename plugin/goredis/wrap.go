// Package ppgoredis instruments the go-redis/redis package (https://github.com/go-redis/redis).
//
// This package instruments the go-redis calls.
// Use the NewClient as the redis.NewClient.
//
//	rc = ppgoredis.NewClient(redisOpts)
//
// It is necessary to pass the context containing the pinpoint.Tracer to Client using Client.WithContext.
//
//	rc = rc.WithContext(pinpoint.NewContext(context.Background(), tracer))
//	rc.Pipeline()
package ppgoredis

import (
	"bytes"
	"context"
	"strings"

	"github.com/go-redis/redis"
	"github.com/pinpoint-apm/pinpoint-go-agent"
)

// Client wraps redis.Client.
type Client struct {
	*redis.Client
	endpoint string
}

// NewClient returns a new Client ready to instrument.
func NewClient(opt *redis.Options) *Client {
	return &Client{Client: redis.NewClient(opt), endpoint: opt.Addr}
}

// WithContext sets the context.
// It is possible to trace only when the given context contains a pinpoint.Tracer.
func (c *Client) WithContext(ctx context.Context) *Client {
	c.Client = c.Client.WithContext(ctx)
	c.WrapProcess(process(ctx, c.endpoint))
	c.WrapProcessPipeline(processPipeline(ctx, c.endpoint))
	return c
}

// ClusterClient wraps redis.ClusterClient.
type ClusterClient struct {
	*redis.ClusterClient
	endpoint string
}

// NewClusterClient returns a new ClusterClient ready to instrument.
func NewClusterClient(opt *redis.ClusterOptions) *ClusterClient {
	endpoint := strings.Join(opt.Addrs, ",")
	return &ClusterClient{ClusterClient: redis.NewClusterClient(opt), endpoint: endpoint}
}

// WithContext sets the context.
// It is possible to trace only when the given context contains a pinpoint.Tracer.
func (c *ClusterClient) WithContext(ctx context.Context) *ClusterClient {
	c.ClusterClient = c.ClusterClient.WithContext(ctx)
	c.WrapProcess(process(ctx, c.endpoint))
	c.WrapProcessPipeline(processPipeline(ctx, c.endpoint))
	return c
}

func process(ctx context.Context, endpoint string) func(oldProcess func(cmd redis.Cmder) error) func(cmd redis.Cmder) error {
	return func(oldProcess func(cmd redis.Cmder) error) func(cmd redis.Cmder) error {
		return func(cmd redis.Cmder) error {
			tracer := pinpoint.FromContext(ctx)
			if tracer == nil {
				return oldProcess(cmd)
			}

			tracer.NewSpanEvent(makeMethodName("Cmd", []redis.Cmder{cmd}))
			defer tracer.EndSpanEvent()

			span := tracer.SpanEvent()
			span.SetServiceType(pinpoint.ServiceTypeRedis)
			span.SetDestination("REDIS")
			span.SetEndPoint(endpoint)

			err := oldProcess(cmd)
			if err != nil {
				span.SetError(err)
			}

			return err
		}
	}
}

func processPipeline(ctx context.Context, endpoint string) func(oldProcess func(cmds []redis.Cmder) error) func(cmds []redis.Cmder) error {
	return func(oldProcess func(cmds []redis.Cmder) error) func(cmds []redis.Cmder) error {
		return func(cmds []redis.Cmder) error {
			tracer := pinpoint.FromContext(ctx)
			if tracer == nil {
				return oldProcess(cmds)
			}

			tracer.NewSpanEvent(makeMethodName("Pipeline", cmds))
			defer tracer.EndSpanEvent()

			span := tracer.SpanEvent()
			span.SetServiceType(pinpoint.ServiceTypeRedis)
			span.SetDestination("REDIS")
			span.SetEndPoint(endpoint)

			err := oldProcess(cmds)
			if err != nil {
				span.SetError(err)
			}

			return err
		}
	}
}

func makeMethodName(operation string, cmds []redis.Cmder) string {
	var buf bytes.Buffer

	buf.WriteString("goredis.")
	buf.WriteString(operation)
	buf.WriteString("(")
	for i, cmd := range cmds {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString("'")
		buf.WriteString(strings.ToLower(cmd.Name()))
		buf.WriteString("'")
	}
	buf.WriteString(")")

	return buf.String()
}
