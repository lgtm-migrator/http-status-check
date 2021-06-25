// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package check

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Dummy",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Hello()
		})
	}
}
