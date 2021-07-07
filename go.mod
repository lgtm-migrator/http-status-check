module github.com/sighupio/http-status-check

go 1.16

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/sighupio/fip-commons v0.1.4
	github.com/sirupsen/logrus v1.2.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.0
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	k8s.io/api v0.18.14
	k8s.io/apimachinery v0.18.14
)

replace k8s.io/client-go => k8s.io/client-go v0.18.14
