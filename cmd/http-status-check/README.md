# http-status-check CLI usage

The tool basically expects 3 input information to function:

* Name of the service to monitor

* Namespace of the service

* HTTP Path to monitor

These values can be provided in three different ways:

## As flags to the command

For eg, the following command can be used to check if the endpoint `"/"`
of the service `nginx` in default namespace responded with `200 OK`.

``` sh
$ http-status-check --service nginx --http-path "/"
# output

INFO[0000] HTTP path "/" of Service nginx in namespace default responded with 200

```

`http-path`(http path) and `namespace` flags are optional and will get
default values `/` and `default` respectively.

## As environment variables

If no flags are provided, the tool looks for environment variables for getting
the information. To avoid possible confusion, the environment variables has the
prefix `HSC`. The usage can be as follows:

``` sh
$ export SEC_SERVICE=nginx # ENV var for service
$ export SEC_HTTP_PATH="/hello" # ENV var for http-path,
                                 # notice that `-` becomes `_`

$ http-status-check --KUBECONFIG ~/.kubeconfig

#output
FATA[0000] HTTP Endpoint "/hello" of service nginx
(namespace: default) did not respond
exit status 1

```

The default values for namespace and http-path remains the same.

## As configuration file

If the tool couldn't find the values in the above two cases, it looks for a
configuration file with the data. By default, it looks for the file
`$HOME/.http-status-check.yaml` if a flag `--config` with the right file
is not provided.

An example usage is:

``` sh
$ cat /home/username/.config
service: nginx
namespace: dev
http-path: "/app/"

$ http-status-check --config /home/username/.config
FATA[0000] services "nginx" not found
exit status 1

```

## Monitoring a remote cluster

By default, the tool expects the Kuberentes cluster is present in the local and
hence look for a kubeconfig `~/.kube/config`. To override this and connect to a
remote cluster, one could use the flag `--KUBECONFIG` (or env var
`HSC_KUBECONFIG` or provide in the config file being used under a key
`KUBECONFIG`).

An example usage is:

``` sh
$ export KUBECONFIG=/home/username/kubebin/kubeconfig
$ http-status-check --service nginx
INFO[0000] HTTP path "/" of Service nginx in namespace default responded with 200

```
