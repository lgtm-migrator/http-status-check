// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package healthcheck

import "testing"

func TestJoinUrl(t *testing.T) {
	tests := []struct {
		ipPort 	  string
		httpPath  string
		want      string
		isError   bool
	}{
		{
			ipPort: "http://10.0.0.1:8765",
			httpPath: "/",
			want: "http://10.0.0.1:8765/",
			isError: false,
		},
		{
			ipPort: "http://10.0.0.1",
			httpPath: "//index",
			want: "http://10.0.0.1/index",
			isError: false,
		},
		{
			ipPort: "my.domain",
			httpPath: "/hello",
			want: "my.domain/hello",
			isError: false,
		},
	}

	for _, test := range tests {
		got, err := JoinURL(test.ipPort, test.httpPath)
		if (err != nil) != test.isError {
			t.Fatal("JoinUrl error. error expected")
		}
		if (test.want != got) {
			t.Fatalf("JoinUrl error. Got (%s) want (%s)", got, test.want)
		}
	}
}
