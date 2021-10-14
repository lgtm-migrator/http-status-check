# http-status-check CLI usage

> This CLI can only be used as a Kubernetes resource since it uses
> the internal Cluster IP to ping the endpoints

The tool basically expects 1 input information to function:

* HTTP URL of the endpoint to monitor

This value can be provided in three different ways:

## As flags to the command

For eg, the following command can be used to check if the service `nginx-svc` in
the same cluster as the `http-status-check` tool under `default` namespaces is is reachable
and responds with `200 OK`.

``` sh
$ http-status-check --http-url http://nginx-svc.default.svc.cluster.local --log-level info
# output

INFO[0000] HTTP URL http://nginx-svc.default.svc.cluster.local responded with 200

```

`--log-level` flag is optional and will get the default value `info` if not provided.

## As environment variables

If no flags are provided, the tool looks for environment variables for getting
the information. To avoid possible confusion, the environment variables has the
prefix `HSC`. The usage can be as follows:

``` sh
$ export HSC_HTTP_URL=http://fip-service.ingress.local # ENV var for http-url
$ export HSC_LOG_LEVEL=error # ENV var for log-level

$ http-status-check

#output
Error: HTTP status check on http://fip-service.ingress.local responded with status code 404 (expected 200)
FATA[0000] error in the cli. Exiting                     error="HTTP status check on http://fip-service.ingress.local responded with status code 404 (expected 200)"
exit status 1
```

## As configuration file

If the tool couldn't find the values in the above two cases, it looks for a
configuration file with the data. By default, it looks for the file
`$HOME/.http-status-check.yaml` if a flag `--config` with the right file
is not provided.

An example usage is:

``` sh
$ cat /home/username/.config
http-url: http://sighup.io/nginx
log-level: warn

$ http-status-check --config /home/username/.config
Error: HTTP status check on http://sighup.io/nginx responded with status code 404 (expected 200)
FATA[0000] error in the cli. Exiting                     error="HTTP status check on http://sighup.io/nginx responded with status code 404 (expected 200)"
exit status 1
```
