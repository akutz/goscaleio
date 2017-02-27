// +build go1.7

package api

import (
	"net/http"

	"github.com/codedellemc/goscaleio/api/context"
)

func doRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	return doRequestWithClient(ctx, http.DefaultClient, req)
}

func doRequestWithClient(
	ctx context.Context,
	client *http.Client,
	req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	return client.Do(req)
}
