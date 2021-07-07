module github.com/sighupio/http-status-check

go 1.16

require (
	github.com/evanphx/json-patch v4.9.0+incompatible // indirect
	github.com/google/addlicense v0.0.0-20210428195630-6d92264d7170 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/sighupio/poc-service-endpoints-check v0.0.0-20210625121503-9255bf1f0c0b // indirect
	github.com/sighupio/service-endpoints-check v0.0.0-20210625150411-0fbbcb8982ef
	github.com/sirupsen/logrus v1.2.0
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/tools v0.1.3 // indirect
	k8s.io/api v0.18.14
	k8s.io/apimachinery v0.18.14
	k8s.io/client-go v0.18.4
)

replace k8s.io/client-go => k8s.io/client-go v0.18.14
