# http-status-check CLI usage

The tool expects TBD input information to function:

* first one
* second one
* next

These values can be provided in three different ways:

## As flags to the command

For eg, the following command can be used to check if TBD.

``` sh
$ http-status-check --flag-1 1 --flag-2 2
# output

INFO[0000] Info message with flag-1=1 and flag-2=2
```

`flag-3`(flag-3) is optional and will get default values `1` and `default` respectively.

## As environment variables

If no flags are provided, the tool looks for environment variables for getting
the information. To avoid possible confusion, the environment variables has the
prefix `TBD`. The usage can be as follows:

``` sh
$ export TBD_FLAG_1=1 # ENV var for flag-1, notice that `-` becomes `_`
$ export TBD_FLAG_2=2 # ENV var for flag-2, notice that `-` becomes `_`

$ http-status-check --KUBECONFIG ~/.kubeconfig

#output
INFO[0000] Info message with flag-1=1 and flag-2=2
```

`flag-3`(flag-3) is optional and will get default values `1` and `default` respectively.

## As configuration file

If the tool couldn't find the values in the above two cases, it looks for a
configuration file with the data. By default, it looks for the file
`$HOME/.http-status-check.yaml` if a flag `--config` with the right file
is not provided.

An example usage is:

``` sh
$ cat /home/username/.config
flag-1: 1
flag-2: 2
flag-3: 3

$ http-status-check --config /home/username/.config
INFO[0000] Info message with flag-1=1 and flag-2=2
```

## Monitoring a remote cluster

By default, the tool expects the Kuberentes cluster is present in the local and
hence look for a kubeconfig `~/.kube/config`. To override this and connect to a
remote cluster, one could use the flag `--KUBECONFIG` (or env var
`SEC_KUBECONFIG` or provide in the config file being used under a key `KUBECONFIG`).

An example usage is:

``` sh
$ export KUBECONFIG=/home/username/kubebin/kubeconfig
$ http-status-check --flag-1 1 --flag-2 2
INFO[0000] Info message with flag-1=1 and flag-2=2
```
