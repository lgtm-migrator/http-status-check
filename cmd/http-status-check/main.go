// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	config "github.com/sighupio/http-status-check/internal/config"
	internal "github.com/sighupio/http-status-check/internal/healthcheck"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	// nolint:typecheck
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var cfg = config.NewHscConfig() // nolint:gochecknoglobals
const envPrefix = "HSC"         // nolint:gochecknoglobals

var rootCmd = &cobra.Command{ // nolint:gochecknoglobals
	PersistentPreRunE: cmdConfig,
	Use:               "http-status-check",
	Short:             "Health check to monitor the http endpoints of a service",
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.ValidateHTTPEndpoint(cfg)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		log.Infof("HTTP path %v of Service %v in namespace %v responded with 200",
			cfg.HttpPath, cfg.ServiceName, cfg.Namespace)
		os.Exit(0)
	},
}

func cmdConfig(cmd *cobra.Command, args []string) error {
	lvl, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.WithField("log-level", cfg.LogLevel).Fatal("incorrect log level")

		return fmt.Errorf("incorrect log level")
	}

	log.SetLevel(lvl)
	log.WithField("log-level", cfg.LogLevel).Debug("log level configured")

	ctx := context.Background()
	err = cfg.KubeClient.Init(&ctx)

	if err != nil {
		log.WithField("kubeconfig", cfg.KubeClient.KubeConfig).Fatal("incorrect kubeconfig configuration")

		return fmt.Errorf("incorrect kubeconfig configuration")
	}

	return nil
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

	rootCmd.Flags().StringVarP(&cfg.ServiceName, "service", "s", "",
		"Name of the service to monitor (required)")
	rootCmd.Flags().StringVarP(&cfg.Namespace, "namespace", "n",
		"default", "Namespace of the service to monitor")
	rootCmd.Flags().StringVarP(&cfg.HttpPath, "http-path", "p", "/",
		"HTTP Path to monitor")
	rootCmd.Flags().StringVar(&cfg.KubeClient.KubeConfig, "kubeconfig", "",
		"kubeconfig file. default: in-cluster configuration, "+
			"Fallback $HOME/.kube/config")
	bindFlags(rootCmd, v)

	err := rootCmd.MarkFlagRequired("service")
	if err != nil {
		log.WithError(err).Fatal("error in the cli. Exiting")
		os.Exit(1)
	}

	rootCmd.PersistentFlags().StringVar(&cfg.CfgFile, "config", "", "config file "+
		"(default is $HOME/.http-status-check.yaml)")
}

func initConfig() *viper.Viper {
	v := viper.New()
	if cfg.CfgFile != "" {
		// Use config file from the flag.
		v.SetConfigFile(cfg.CfgFile)
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

func main() {
	cobra.CheckErr(rootCmd.Execute())
}
