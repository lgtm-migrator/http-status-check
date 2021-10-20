# Deployment with Kustomization

In this file, there is the `kustomization` file which can be deployed using the
following command:

```sh
$ cp env_template .env
# Add the corresponding values in the .env file (explained below)
$ kustomize build . | kubectl apply -f -
# Output will show the resources being created
```

The contents of the `kustomization` file are as follows:

```yaml
---
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

commonLabels:
  app.kubernetes.io/created-by: healthcheck-controller
  kfip.sighup.io/group: SampleGroup
  kfip.sighup.io/target: SampleTarget
  kfip.sighup.io/check: http-status-check

images:
  - name: registry.sighup.io/poc/http-status-check
    newName: registry.sighup.io/poc/http-status-check
    newTag: unstable

resources:
  - cronjob.yaml

configMapGenerator:
  - name: hsc-envs
    envs:
      - .env
```

Like one can see, the two important resources created are
the cronjob doing the health check *(through [`cronjob.yaml`](cronjob.yaml) file)*
and the configmap created using the `configMapGenerator`.
Let us look into each to understand the configuration.

We can specify a namespace in which
all the resources are created by kustomize in the `Kustomization` file. If the
namespace does not exist we will have to add a `yaml` to create the namespace in
the resources sections. If namespace is not explicitly defined, it will be `default`.

## CronJob

In the file [cronjob](./http-status-check-cronjob)
we create a CronJob with the image of our tool. If you look under the containers
section you will notice that the environmental variables are injected from a
configMap created from the later section by kustomize. This environmental
variables are very important since they are used by our tool to decide which
service is to be monitored in which namespace and the number of endpoints
expected.

The entrypoint of the image used is /http-status-check which is
the binary build from the cmd/. This binary, as explained in the CLI usage guide,
expects flags, env vars, or configuration files. We use environment variables
while creating a job since it is the cleanest way *(open to a heated discussion)*
to inject data into a pod.

The environment variables necessary for the pod to execute are:

```yaml
HSC_HTTP_URL
HSC_LOG_LEVEL # this is optional
```

Refer the [CLI usage guide for detailed review of
this](../../cmd/http-status-check/README.md).

This environment variable data is expected inside the configMap. This configMap
can be injected like this in the job file:

``` yaml
            env:
              - name: HSC_HTTP_URL
                valueFrom:
                  configMapKeyRef:
                    name: hsc-envs
                    key: HSC_HTTP_URL
              - name: HSC_LOG_LEVEL
                valueFrom:
                  configMapKeyRef:
                    name: hsc-envs
                    key: HSC_LOG_LEVEL
                    optional: true
```

## ConfigMap

The configMap is created using `configMapGenerator` of kustomize. In the CronJob
file, we can see that a configMap of name hsc-envs is expected to hold the
aforementioned environmental variables. To create the configMap the following
kustomize section is used:

``` yaml
configMapGenerator:
- name: hsc-envs
  env: .env
```

Here a hsc-envs configMap Kubernetes resource is created from an environment
file called `.env`. A template to this file is provided as a file
`env_template`. The file has two keys defined with values left empty. So the
first step would be to rename it as `.env` since that is what `kustomization`
expects.

``` yaml
$ cp env_template .env
$ cat .env
HSC_HTTP_URL=https://sighup.io
HSC_LOG_LEVEL=info
```

Add the values for the above 3 environment variables. An example could be:

```yaml
$ cat .env
HSC_HTTP_URL=https://sighup.io
HSC_LOG_LEVEL=info
```

