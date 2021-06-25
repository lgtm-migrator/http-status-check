# Deployments Guidelines

The suggested method of deployment in a Kubernetes cluster is as a
CronJob resource that keeps TBD.

To allow the job to access the deployment, we create a service account
bound to a role that can TBD.

To provide the TBD
for the functioning of the tool, we use the environment variables that can be
injected via configmaps.

## Helm

## Kustomize

We ship the kustomization file along with jobs and RBAC resources needed to
deploy the tool and configure it to monitor the service.

The specific information about the files involved and the configuration options
can found in the [`README`](kustomization/README.md) of the `kustomization` directory.
Head over to [that doc](./kustomization/) for detailed deployment info.

## Manifests
