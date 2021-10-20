// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package healthcheck

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCurlHTTPEndpoint(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	}))

	tests := []struct {
		args    context.Context
		want    int
		wantErr bool
		url     string
	}{
		{
			want:    200,
			url:     ts.URL,
			wantErr: false,
		},
		{
			want:    0,
			wantErr: true,
			url:     "https://idonotexist/bleh",
		},
		{
			want:    200,
			wantErr: false,
			url:     "https://google.com",
		},
		{
			want:    404,
			wantErr: false,
			url:     "https://www.google.com/error",
		},
	}

	for _, test := range tests {
		got, err := curlHTTPEndpoint(&test.args, test.url)
		if (err != nil) != test.wantErr {
			t.Errorf("curlHTTPEndpoint() error = %v, wantErr %v", err, test.wantErr)
			return
		}

		if err == nil {
			if got != test.want {
				t.Errorf("curlHTTPEndpoint() = %v, want %v", got, test.want)
			}
		}
	}
}
