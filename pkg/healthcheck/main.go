// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package healthcheck

import (
	"context"
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/sighupio/service-endpoints-check/pkg/client"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getService(kc *client.KubernetesClient,
	svcName string,
	namespace string) (*corev1.Service, error) {
	service, err := kc.Client.CoreV1().Services(namespace).Get(context.TODO(),
		svcName, metav1.GetOptions{})

	return service, err

}

func getEndpoints(kc *client.KubernetesClient, service *corev1.Service,
	namespace string) (*corev1.Endpoints, error) {

	// Retrive all the endpoints corresponding to the service
	// Name of the endpoint will always match that of the svc
	endpoint, err := kc.Client.CoreV1().Endpoints(namespace).Get(
		context.TODO(), service.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return endpoint, err
}

func epAddress(endpoint *corev1.Endpoints) ([]string, []int32) {
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

func JoinURL(base string, paths ...string) string {
	p := path.Join(paths...)
	return fmt.Sprintf("%s/%s", strings.TrimRight(base, "/"), strings.TrimLeft(p, "/"))
}

func makehttpCall(url string) (*http.Response, error) {
	/* #nosec G107: Url generation*/
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func CallServiceHTTPEndpoint(client *client.KubernetesClient,
	serviceName string, namespace string, httpPath string) (map[string]int, error) {
	service, err := getService(client, serviceName, namespace)
	if err != nil {
		return nil, err
	}

	endpoints, err := getEndpoints(client, service, namespace)
	if err != nil {
		return nil, err
	}
	statusCodes := make(map[string]int)
	addrs, ports := epAddress(endpoints)
	if len(addrs) == 0 || len(ports) == 0 {
		return nil, fmt.Errorf("No endpoint addresses were found service  "+
			"%v (namespace %v)", serviceName, namespace)

	}
	for _, addr := range addrs {
		for _, port := range ports {
			url := JoinURL(fmt.Sprintf("http://%v:%v", addr, port), httpPath)
			resp, err := makehttpCall(url)
			if err != nil {
				return nil, err
			}
			statusCodes[url] = resp.StatusCode
		}
	}
	return statusCodes, nil
}
