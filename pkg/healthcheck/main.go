// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package healthcheck

import (
	"net/http"
	"net/url"
	"path"

	corev1 "k8s.io/api/core/v1"
)

func EpAddress(endpoint *corev1.Endpoints) ([]string, []int32) {
	var epAddrs []string

	var epPorts []int32

	for _, subset := range endpoint.Subsets {
		for _, address := range subset.Addresses {
			epAddrs = append(epAddrs, address.IP)
		}

		for _, port := range subset.Ports {
			if port.Protocol == "TCP" {
				epPorts = append(epPorts, port.Port)
			}
		}
	}

	return epAddrs, epPorts
}

func JoinURL(base string, httpPath string) (string, error) {
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, httpPath)

	return u.String(), nil
}

func MakehttpCall(url string) (*http.Response, error) {
	resp, err := http.Get(url) // nolint:gosec // G107: Url generation
	if err != nil {
		return nil, err
	}

	return resp, nil
}
