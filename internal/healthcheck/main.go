// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package healthcheck

import (
	"context"
	"fmt"
	"net/http"
)

const successStatusCode = 200

func ValidateHTTPEndpoint(ctx *context.Context, url string) error {
	statusCode, err := curlHTTPEndpoint(ctx, url)
	if err != nil {
		return err
	}

	if statusCode != successStatusCode {
		return fmt.Errorf("HTTP status check on %v "+
			"responded with status code %v (expected %d)", url,
			statusCode, successStatusCode)
	}

	return nil
}

func curlHTTPEndpoint(ctx *context.Context, url string) (int, error) {
	req, err := http.NewRequest("GET", url, nil)
	// req = req.WithContext(ctx)

	if err != nil {
		return 0, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()

	return res.StatusCode, err
}
