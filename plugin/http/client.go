package pphttp

import (
	"bytes"
	"context"
	"net/http"

	"github.com/pinpoint-apm/pinpoint-go-agent"
)

// NewHttpClientTracer is deprecated. Use WrapClient or DoClient.
func NewHttpClientTracer(tracer pinpoint.Tracer, operationName string, req *http.Request) pinpoint.Tracer {
	return before(tracer, operationName, req)
}

func before(tracer pinpoint.Tracer, operationName string, req *http.Request) pinpoint.Tracer {
	if tracer == nil {
		return tracer
	}

	tracer.NewSpanEvent(operationName)
	tracer.SpanEvent().SetEndPoint(req.Host)
	tracer.SpanEvent().SetDestination(req.Host)
	tracer.SpanEvent().SetServiceType(pinpoint.ServiceTypeGoHttpClient)

	if tracer.IsSampled() {
		var b bytes.Buffer
		b.WriteString(req.Method)
		b.WriteString(" ")
		b.WriteString(req.URL.String())
		tracer.SpanEvent().Annotations().AppendString(pinpoint.AnnotationHttpUrl, b.String())

		a := tracer.SpanEvent().Annotations()
		recordClientHttpRequestHeader(a, req.Header)
		recordClientHttpCookie(a, req.Cookies())
	}

	tracer.Inject(req.Header)
	return tracer
}

// EndHttpClientTracer is deprecated.
func EndHttpClientTracer(tracer pinpoint.Tracer, resp *http.Response, err error) {
	after(tracer, resp, err)
}

func after(tracer pinpoint.Tracer, resp *http.Response, err error) {
	if tracer == nil {
		return
	}

	tracer.SpanEvent().SetError(err)
	if resp != nil && tracer.IsSampled() {
		a := tracer.SpanEvent().Annotations()
		a.AppendInt(pinpoint.AnnotationHttpStatusCode, int32(resp.StatusCode))
		recordClientHttpResponseHeader(a, resp.Header)
	}
	tracer.EndSpanEvent()
}

// DoClient instruments and executes a given doFunc.
// It is necessary to pass the context containing the pinpoint.Tracer to the http.Request.
//
//	req, _ := http.NewRequestWithContext(pinpoint.NewContext(context.Background(), tracer), "GET", url, nil)
//	pphttp.DoClient(http.DefaultClient.Do, req)
func DoClient(doFunc func(req *http.Request) (*http.Response, error), req *http.Request) (*http.Response, error) {
	tracer := before(pinpoint.TracerFromRequestContext(req), "http/Client.Do()", req)
	resp, err := doFunc(req)
	after(tracer, resp, err)

	return resp, err
}

type roundTripper struct {
	original http.RoundTripper
	ctx      context.Context
}

// WrapClient returns a new *http.Client ready to instrument.
// It is necessary to pass the context containing the pinpoint.Tracer to the http.Request.
//
//	req, _ := http.NewRequestWithContext(pinpoint.NewContext(context.Background(), tracer), "GET", url, nil)
//	client := pphttp.WrapClient(&http.Client{})
//	client.Do(req)
func WrapClient(client *http.Client) *http.Client {
	if client == nil {
		client = http.DefaultClient
	}

	c := *client
	c.Transport = wrapRoundTripper(nil, c.Transport)
	return &c
}

// WrapClientWithContext returns a new *http.Client ready to instrument.
// It is possible to trace only when the given context contains a pinpoint.Tracer.
//
//	client := pphttp.WrapClientWithContext(pinpoint.NewContext(context.Background(), tracer), &http.Client{})
//	client.Get(external_url)
func WrapClientWithContext(ctx context.Context, client *http.Client) *http.Client {
	if client == nil {
		client = http.DefaultClient
	}

	c := *client
	c.Transport = wrapRoundTripper(ctx, c.Transport)
	return &c
}

func wrapRoundTripper(ctx context.Context, original http.RoundTripper) http.RoundTripper {
	if original == nil {
		original = http.DefaultTransport
	}

	return &roundTripper{
		original: original,
		ctx:      ctx,
	}
}

func (r *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	var tracer pinpoint.Tracer

	if r.ctx != nil {
		tracer = pinpoint.FromContext(r.ctx)
	} else {
		tracer = pinpoint.FromContext(req.Context())
	}

	// By the specification of http.RoundTripper, it requires that the given Request is not changed.
	// We make a copy of the Request because pinpoint headers need to be added.
	clone := *req
	clone.Header = make(http.Header, len(req.Header))
	for k, v := range req.Header {
		clone.Header[k] = v
	}
	req = &clone

	tracer = before(tracer, "http/Client.Do()", req)
	resp, err := r.original.RoundTrip(req)
	after(tracer, resp, err)

	return resp, err
}
