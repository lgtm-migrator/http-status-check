// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sighupio/http-status-check/internal/healthcheck"
	"github.com/sighupio/service-endpoints-check/pkg/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var cfgFile string

var namespace string
var httpEp string
var serviceName string

const envPrefix = "HSC"

var rootCmd = &cobra.Command{
	Use:   "http-status-check",
	Short: "Health check to monitor the http endpoints of a service",
	Run: func(cmd *cobra.Command, args []string) {
		clientSet := initClient()
		err := healthcheck.ValidateHTTPEndpoint(clientSet, serviceName, namespace, httpEp)
		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		log.Infof("HTTP path %v of Service %v in namespace %v responded with 200",
			httpEp, serviceName, namespace)
		os.Exit(0)
	},
}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}

func initClient() *client.KubernetesClient {
	config, err := client.Config("")
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	return &client.KubernetesClient{Client: config}
}

func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Environment variables can't have dashes in them, so bind them to their equivalent
		// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			err := v.BindEnv(f.Name, fmt.Sprintf("%s_%s", envPrefix, envVarSuffix))
			if err != nil {
				log.Fatal(err)
				os.Exit(-1)

			}

		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			err := cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
			if err != nil {
				log.Fatal(err)
				os.Exit(-1)

			}
		}
	})
}

func init() {
	v := initConfig()

	rootCmd.Flags().StringVarP(&serviceName, "service", "s", "", "Name of the service to monitor (required)")
	rootCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "Namespace of the service to monitor")
	rootCmd.Flags().StringVarP(&httpEp, "http-path", "p", "/", "HTTP Path to monitor")
	bindFlags(rootCmd, v)
	err := rootCmd.MarkFlagRequired("service")
	if err != nil {
		log.WithError(err).Fatal("error in the cli. Exiting")
		os.Exit(1)
	}
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file "+
		"(default is $HOME/.http-status-check.yaml)")
}

func initConfig() *viper.Viper {
	v := viper.New()
	if cfgFile != "" {
		// Use config file from the flag.
		v.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name (without extension).
		v.AddConfigPath(home)
		v.SetConfigType("yaml")
		v.SetConfigName(".http-status-check")
	}

	// Attempt to read the config file, gracefully ignoring errors
	// caused by a config file not being found. Return an error
	// if we cannot parse the config file.
	if err := v.ReadInConfig(); err != nil {
		// It's okay if there isn't a config file
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Info(err)
		}
	}

	v.SetEnvPrefix(envPrefix)

	// Bind to environment variables
	// Works great for simple config names, but needs help for names
	// like --favorite-color which we fix in the bindFlags function
	v.AutomaticEnv()
	return v
}
