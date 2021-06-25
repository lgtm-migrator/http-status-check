#!/bin/bash -e
# Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

GITHUB_OWNER="sighupio"
GITHUB_PROJECT=${1}
GITHUB_TOKEN=${2}
DRONE_TOKEN=${3}
REGISTRY=${4}
REGISTRY_USER=${5}
REGISTRY_PASSWORD=${6}

echo "Initializing repository for ${GITHUB_PROJECT}"
echo " Configuring drone project @ ci.sighup.io"
echo "  Waiting 30s to discover the newest repositories"
curl --fail -s -X POST -H "Authorization: Bearer ${DRONE_TOKEN}" "https://ci.sighup.io/api/user/repos?async=true"
sleep 30
echo "  Enabling project"
curl --fail -s -X POST -H "Authorization: Bearer ${DRONE_TOKEN}" "https://ci.sighup.io/api/repos/${GITHUB_OWNER}/${GITHUB_PROJECT}"
echo "  Marking project as trusted"
curl --fail -s -X PATCH -H "Content-Type: application/json" -H "Authorization: Bearer ${DRONE_TOKEN}" --data '{"trusted":true}' "https://ci.sighup.io/api/repos/${GITHUB_OWNER}/${GITHUB_PROJECT}"
echo "  Creating required secrets"
curl --fail -s -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${DRONE_TOKEN}" --data '{"name":"GITHUB_TOKEN","data":"'"${GITHUB_TOKEN}"'", "pull_request": false}' "https://ci.sighup.io/api/repos/${GITHUB_OWNER}/${GITHUB_PROJECT}/secrets"
curl --fail -s -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${DRONE_TOKEN}" --data '{"name":"REGISTRY","data":"'"${REGISTRY}"'", "pull_request": false}' "https://ci.sighup.io/api/repos/${GITHUB_OWNER}/${GITHUB_PROJECT}/secrets"
curl --fail -s -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${DRONE_TOKEN}" --data '{"name":"REGISTRY_USER","data":"'"${REGISTRY_USER}"'", "pull_request": false}' "https://ci.sighup.io/api/repos/${GITHUB_OWNER}/${GITHUB_PROJECT}/secrets"
curl --fail -s -X POST -H "Content-Type: application/json" -H "Authorization: Bearer ${DRONE_TOKEN}" --data '{"name":"REGISTRY_PASSWORD","data":"'"${REGISTRY_PASSWORD}"'", "pull_request": false}' "https://ci.sighup.io/api/repos/${GITHUB_OWNER}/${GITHUB_PROJECT}/secrets"
echo "  Removing this script"
rm -rf "${0}"

cat << EOF
Successfully initialized. Feel free to:

$ git add .
$ git commit --amend -m "Initial commit"
$ git push -f

In order to have a clean repo starting point
EOF
