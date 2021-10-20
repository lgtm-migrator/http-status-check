// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

// HscConfig hosts the configuration of the application.
type HscConfig struct {
	CfgFile  string
	HTTPPath string
	LogLevel string
}

// NewHscConfig creates an empty configuration struct.
func NewHscConfig() *HscConfig {
	return &HscConfig{}
}
