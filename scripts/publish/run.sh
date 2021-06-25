#!/bin/bash -e
# Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

IMAGE=${1}
REGISTRY=${2}
REGISTRY_USER=${3}
REGISTRY_PASSWORD=${4}
IMAGE_NAME=${5}
IMAGE_TAG=${6}

NEW_IMAGE=${REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG}

echo "Starting release process"
echo "  Login into the registry"
docker login "${REGISTRY}" -u "${REGISTRY_USER}" -p "${REGISTRY_PASSWORD}"
echo "  Tag the built image"
docker tag "${IMAGE}" "${NEW_IMAGE}"
echo "  Pushing the image ${NEW_IMAGE}"
docker push "${NEW_IMAGE}"
echo "${NEW_IMAGE} Released"
