// +build !go1.7

package context

import (
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
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
	return ctxhttp.Do(ctx, client, req)
}
