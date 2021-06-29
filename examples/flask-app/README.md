# HTTP status check on service endpoints

The flask deployment has the following characteristics:

```text
Name of the deployment: simple-flask-app
Namespace: default
Name of the service: simple-flask-app
```

## Requirement

The flask app serves the names of certain individuals at the endpoint
"/names". Our system has to monitor it at all times so that we know the
service is serving requests at all times.  We get the kustomization,
role-binding, cronjob and the env template files from `deployments`
directory. To start with let us define the env vars in the env file. We name
our env file `flask.env` for cleanliness reasons and add the necessary data
as follows:

``` sh
$ cat flask.env
HSC_SERVICE=simple-flask-app
HSC_NAMESPACE=default
HSC_HTTP_PATH=/names
```

Now we edit the kustomization file to create the deployment of the flask
app, the role
binding and service account, the cronjob to monitor and the config map that
injects flask.env to the jobs pod. All these resources are to be created under
`default` namespace which can be specified in the kustomization file.
The file goes as follows:

The serviceAccount is created along with the role and role binding in the file
`role-binding.yaml`. This serviceAccount is used in the CronJob resource file.
The deployment configuration for the aforementioned deployment can be found in
the file `deployment.yaml` and the service configuration in
`service.yaml`.

The file goes as follows:

``` yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
namespace: default

resources:
- deploymant.yaml
- service.yaml
- role-binding.yaml
- cronjob.yaml
configMapGenerator:
- name: hsc-envs
  env: flask.env

```

To apply this configuration to the cluster run the command:

``` sh
$ kustomize build . | kubectl apply -f -
# Output will show the resources being created
```
