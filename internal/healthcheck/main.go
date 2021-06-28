// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package healthcheck

import (
	"fmt"
	"errors"
	"github.com/sighupio/service-endpoints-check/pkg/client"

	"github.com/sighupio/http-status-check/pkg/healthcheck"
)

func ValidateHTTPEndpoint(client *client.KubernetesClient,
	service string,
	namespace string, httpPath string) error {
	statusCodes, err := healthcheck.CallServiceHTTPEndpoint(client, service, namespace, httpPath)
	if err != nil {
		return err
	}

	for url, code := range statusCodes {
		if code != 200 {
			return errors.New(fmt.Sprintf("Endpoint %v of service %v " +
				"(namespace %v) responded with %v (expected 200)", url,
				service, namespace, code))
		}
	}
	return nil
}
