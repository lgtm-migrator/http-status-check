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

@test "Prepare deploy 01" {
    info
    prep(){
        cp deployments/kustomization/env_template deployments/kustomization/.env
    }
    run prep
    [ "$status" -eq 0 ]
}

@test "Deploy 01" {
    info
    deploy(){
        kustomize build deployments/kustomization | kubectl apply -f -
    }
    run deploy
    [ "$status" -eq 0 ]
}

@test "Use loaded image in the cronjob 01" {
    info
    mutate(){
        kubectl set image cronjob/http-status-check http-status-check="${CONTAINER_IMAGE}"
    }
    run mutate
    [ "$status" -eq 0 ]
}

@test "Create a job from the cronjob 01" {
    info
    mutate(){
        kubectl create job http-status-check-job-1 --from cronjob/http-status-check
    }
    run mutate
    [ "$status" -eq 0 ]
}

@test "Check if the job 01 completed" {
    info
    check(){
        kubectl wait --for=condition=complete --timeout=180s job/http-status-check-job-1
    }
    run check
    [ "$status" -eq 0 ]
}

@test "Clean up 01" {
    info
    cleanup(){
        kustomize build deployments/kustomization | kubectl delete -f -
    }
    run cleanup
    [ "$status" -eq 0 ]
}

@test "Create deployment with 3 replicas" {
    info
    deploy(){
         kubectl create deploy nginx-deploy --image nginx:latest --replicas 3 --port 80
    }
    run deploy
    [ "$status" -eq 0 ]
}

@test "Wait for replicas to come up" {
    info
    deploy(){
        kubectl wait --timeout=180s --for=condition=ready pod -l app=nginx-deploy
    }
    run deploy
    [ "$status" -eq 0 ]
}

@test "Create Service for the deployment" {
    info
    deploy(){
        kubectl create service clusterip nginx-deploy --tcp=8080:80
    }
    run deploy
    [ "$status" -eq 0 ]
}

@test "Prepare deploy 02" {
    info
    prep(){
        cp deployments/kustomization/env_template_2 deployments/kustomization/.env
    }
    run prep
    [ "$status" -eq 0 ]
}

@test "Deploy 02" {
    info
    deploy(){
        kustomize build deployments/kustomization | kubectl apply -f -
    }
    run deploy
    [ "$status" -eq 0 ]
}

@test "Use loaded image in the cronjob 02" {
    info
    mutate(){
        kubectl set image cronjob/http-status-check http-status-check="${CONTAINER_IMAGE}"
    }
    run mutate
    [ "$status" -eq 0 ]
}

@test "Create a job from the cronjob 02" {
    info
    mutate(){
        kubectl create job http-status-check-job-2 --from cronjob/http-status-check
    }
    run mutate
    [ "$status" -eq 0 ]
}

@test "Check if the job 02 completed" {
    info
    check(){
        kubectl wait --for=condition=complete --timeout=180s job/http-status-check-job-2
    }
    run check
    [ "$status" -eq 0 ]
}


@test "Create Service for non existing pods" {
    info
    deploy(){
        kubectl create service clusterip no-deploy --tcp=8080:80
    }
    run deploy
    [ "$status" -eq 0 ]
}

@test "Prepare deploy 03" {
    info
    prep(){
        cp deployments/kustomization/env_template_3 deployments/kustomization/.env
    }
    run prep
    [ "$status" -eq 0 ]
}

@test "Deploy 03" {
    info
    deploy(){
        kustomize build deployments/kustomization | kubectl apply -f -
    }
    run deploy
    [ "$status" -eq 0 ]
}

@test "Use loaded image in the cronjob 03" {
    info
    mutate(){
        kubectl set image cronjob/http-status-check http-status-check="${CONTAINER_IMAGE}"
    }
    run mutate
    [ "$status" -eq 0 ]
}

@test "Create a job from the cronjob 03" {
    info
    mutate(){
        kubectl create job http-status-check-job-3 --from cronjob/http-status-check
    }
    run mutate
    [ "$status" -eq 0 ]
}

@test "Check the job 03 did not complete successfully" {
    info
    deploy() {
        kubectl wait --for=condition=failed --timeout=90s job/http-status-check-job-3
    }
    run deploy

    [ "$status" -eq 0 ]
}
