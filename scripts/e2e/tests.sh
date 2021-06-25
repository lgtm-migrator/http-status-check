#!/usr/bin/env bats
# Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.


info(){
  echo -e "${BATS_TEST_NUMBER}: ${BATS_TEST_DESCRIPTION}" >&3
}

@test "Requirements" {
    info
    deploy(){
        kubectl run nginx --image nginx:latest --port 80 --expose
    }
    run deploy
    [ "$status" -eq 0 ]
}

@test "Requirements Check" {
    info
    deploy(){
        kubectl wait --timeout=180s --for=condition=ready pod -l run=nginx
    }
    run deploy
    [ "$status" -eq 0 ]
}

@test "Deploy" {
    info
    deploy(){
        kubectl create job service-endpoints-check --image="${CONTAINER_IMAGE}"
    }
    run deploy
    [ "$status" -eq 0 ]
}

@test "Check" {
    info
    check(){
        kubectl wait --for=condition=complete --timeout=180s job/service-endpoints-check
    }
    run check
    [ "$status" -eq 0 ]
}
