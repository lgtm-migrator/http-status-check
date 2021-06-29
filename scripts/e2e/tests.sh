#!/usr/bin/env bats
# Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.


info(){
  echo -e "${BATS_TEST_NUMBER}: ${BATS_TEST_DESCRIPTION}" >&3
}

test_envs_01(){
    export NAMESPACE=default
    export SVCNAME=nginx
    export SERVICE_ACCOUNT=sa-endpoint-check
    export CLUSTER_ROLE=service-reader
    export CLUSTER_ROLE_BINDING=service-reader-rb
    export JOB_NAME=http-status-check
}

test_envs_02(){
    export NAMESPACE=default
    export SVCNAME=nginx-deploy
    export HTTP_PATH="/index.html"
    export SERVICE_ACCOUNT=sa-endpoint-check
    export JOB_NAME=http-status-check-02
}

test_envs_03(){
    export NAMESPACE=default
    export SVCNAME=no-deploy
    export MIN_EP="/app/live"
    export SERVICE_ACCOUNT=sa-endpoint-check
    export JOB_NAME=http-status-check-03
}

@test "Requirements" {
    info
    test_envs_01
    deploy(){
        kubectl run $SVCNAME --image nginx:latest --port 80 --expose
    }
    run deploy
    [ "$status" -eq 0 ]
}

@test "Requirements Check" {
    info
    test_envs_01
    deploy(){
        kubectl wait --timeout=180s --for=condition=ready pod -l run=$SVCNAME
    }
    run deploy
    [ "$status" -eq 0 ]
}

@test "Create RBAC for the job " {
    info
    test_envs_01
    create_svc_account() {
        kubectl create sa "$SERVICE_ACCOUNT"
    }
    run create_svc_account
    [ "$status" -eq 0 ]

    create_cluster_role() {
         kubectl create clusterrole "$CLUSTER_ROLE"  --verb=get,list,watch --resource=services,pods,endpoints
     }
    run create_cluster_role
    [ "$status" -eq 0 ]

    create_cluster_role_binding() {
        kubectl create clusterrolebinding "$CLUSTER_ROLE_BINDING" --clusterrole "$CLUSTER_ROLE" --serviceaccount="$NAMESPACE:$SERVICE_ACCOUNT"
    }
    run create_cluster_role_binding
    [ "$status" -eq 0 ]

}

@test "Deploy  job" {
    info
    test_envs_01
    deploy(){
        envsubst < "scripts/e2e/test_job.template" | kubectl apply -f -
    }

    run deploy
    [ "$status" -eq 0 ]
}


@test "Check the job completed successfully" {
    info
    test_envs_01

    check(){
        kubectl wait --for=condition=complete --timeout=180s "job/$JOB_NAME" -n "$NAMESPACE"
    }
    run check
    [ "$status" -eq 0 ]
}

@test "Create deployment with 3 replicas" {
    info
    test_envs_02
    deploy(){
         kubectl create deploy $SVCNAME --image nginx:latest --replicas 3 --port 80
    }
    run deploy
    [ "$status" -eq 0 ]
}

@test "Wait for replicas to come up" {
    info
    test_envs_02
    deploy(){
        kubectl wait --timeout=180s --for=condition=ready pod -l app=$SVCNAME
    }
    run deploy
    [ "$status" -eq 0 ]
}

@test "Create Service for the deployment" {
    info
    test_envs_02
    deploy(){
        kubectl create service clusterip $SVCNAME --tcp=8080:80
    }
    run deploy
    [ "$status" -eq 0 ]
}

@test "Deploy monitoring job 02" {
    info
    test_envs_02
    deploy(){
        envsubst < "scripts/e2e/test_job.template" | kubectl apply -f -
    }

    run deploy
    [ "$status" -eq 0 ]
}


@test "Check the job 02 completed successfully" {
    info
    test_envs_02

    check() {
    kubectl wait --for=condition=complete --timeout=180s "job/$JOB_NAME" -n "$NAMESPACE"
    }
    run check
    [ "$status" -eq 0 ]
}

@test "Create Service for non existing pods" {
    info
    test_envs_03
    deploy(){
        kubectl create service clusterip $SVCNAME --tcp=8080:80
    }
    run deploy
    [ "$status" -eq 0 ]
}

@test "Deploy  job 03" {
    info
    test_envs_03
    deploy(){
        envsubst < "scripts/e2e/test_job.template" | kubectl apply -f -
    }

    run deploy
    [ "$status" -eq 0 ]
}


@test "Check the job 03 did not complete successfully" {
    info
    test_envs_03
    deploy() {
        kubectl wait --for=condition=failed --timeout=90s "job/$JOB_NAME" -n "$NAMESPACE"
    }
    run deploy

    [ "$status" -eq 0 ]
}
