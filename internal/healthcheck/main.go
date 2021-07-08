// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package healthcheck

import (
	"context"
	"fmt"

	config "github.com/sighupio/http-status-check/internal/config"
	pkg "github.com/sighupio/http-status-check/pkg/healthcheck"
	log "github.com/sirupsen/logrus"
)

func ValidateHTTPEndpoint(ctx *context.Context, cfg *config.HscConfig) error {
	const successStatusCode = 200

	statusCodes, err := callServiceHTTPEndpoint(ctx, cfg)
	if err != nil {
		return err
	}

	for url, code := range statusCodes {
		if code != successStatusCode {
			return fmt.Errorf("Endpoint %v of service %v "+
				"(namespace %v) responded with %v (expected 200)", url,
				cfg.ServiceName, cfg.Namespace, code)
		}
	}

	return nil
}

func callServiceHTTPEndpoint(ctx *context.Context, cfg *config.HscConfig) (map[string]int, error) {
	service, err := cfg.KubeClient.GetService(ctx, cfg.ServiceName, cfg.Namespace)
	if err != nil {
		return nil, err
	}

	endpoints, err := cfg.KubeClient.GetEndpoints(ctx, service, cfg.Namespace)
	if err != nil {
		return nil, err
	}

	statusCodes := make(map[string]int)

	addrs, ports := pkg.EpAddress(endpoints)
	if len(addrs) == 0 || len(ports) == 0 {
		return nil, fmt.Errorf("No endpoint addresses were found service  "+
			"%v (namespace %v)", cfg.ServiceName, cfg.Namespace)
	}

	for _, addr := range addrs {
		for _, port := range ports {
			url, err := pkg.JoinURL(fmt.Sprintf("http://%v:%v", addr, port), cfg.HTTPPath)
			if err != nil {
				log.Fatalf("IP parsing error service %v address: IP %v port %d",
					cfg.ServiceName, addr, port)
			}

			resp, err := pkg.MakehttpCall(url)
			if err != nil {
				return nil, err
			}

			statusCodes[url] = resp.StatusCode
		}
	}

	return statusCodes, nil
}
