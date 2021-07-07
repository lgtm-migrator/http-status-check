// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import "github.com/sighupio/fip-commons/pkg/kube"

// HscConfig hosts the configuration of the application.
type HscConfig struct {
	CfgFile     string
	Namespace   string
	HTTPPath    string
	ServiceName string
	LogLevel    string
	KubeClient  kube.KubernetesClient
}

// NewHscConfig creates an empty configuration struct.
func NewHscConfig() *HscConfig {
	return &HscConfig{}
}
