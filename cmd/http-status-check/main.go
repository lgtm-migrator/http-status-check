// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	internal "github.com/sighupio/http-status-check/internal/http-status-check"
	pkg "github.com/sighupio/http-status-check/pkg/http-status-check"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "http-status-check",
	Short: "http-status-check TBD",
	Long:  "TBD",
	Run: func(cmd *cobra.Command, args []string) {
		// Do business logic
		internal.Hello()
		pkg.Hello()
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
