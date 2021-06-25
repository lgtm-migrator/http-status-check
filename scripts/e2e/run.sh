#!/bin/bash -e
# Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.


CLUSTER_ID=e2e-$(LC_ALL=C tr -dc 'a-z0-9' </dev/urandom | head -c 8 ; echo)
KUBECONFIG=$(pwd)/kubeconfig

function cleanup {
  echo "  Destroying the cluster ${CLUSTER_ID}"
  kind delete cluster --name "${CLUSTER_ID}"
}

echo "  Creating the cluster ${CLUSTER_ID}"
kind create cluster --name "${CLUSTER_ID}" --kubeconfig "${KUBECONFIG}"
trap cleanup EXIT
echo "  Loading the container image ${1} in the cluster ${CLUSTER_ID}"
kind load docker-image --name "${CLUSTER_ID}" "${1}"
echo "  Waiting the cluster ${CLUSTER_ID} to become ready"
kubectl --kubeconfig "${KUBECONFIG}" wait --timeout=180s --for=condition=ready pod --all -n kube-system
echo "  Executing e2e tests"
KUBECONFIG=$(pwd)/kubeconfig CONTAINER_IMAGE=${1} bats -t ./scripts/e2e/tests.sh
