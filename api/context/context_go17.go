// +build go1.7

package context

import (
	"context"
	"net/http"
)

// Context carries a deadline, a cancelation signal, and other values across
// API boundaries.
type Context interface {
	context.Context
}

func doRequest(ctx Context, req *http.Request) (*http.Response, error) {
	return doRequestWithClient(ctx, http.DefaultClient, req)
}

func doRequestWithClient(
	ctx Context,
	client *http.Client,
	req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	return client.Do(req)
}
