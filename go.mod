module github.com/sighupio/http-status-check

go 1.16

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/sirupsen/logrus v1.2.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.0
)

replace k8s.io/client-go => k8s.io/client-go v0.18.14
