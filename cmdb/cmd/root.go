package cmd

import (
	"fmt"
	"os"
	"errors"
	"github.com/spf13/cobra"
)

var (
	// pusher service config option
	confType string
	confFile string
	confETCD string
)

var vers bool

var RootCmd = &cobra.Command{
	Use:   "demo-api",
	Short: "",
	Long:  "",
	RunE: func(*cobra.Command, []string) error {
		return errors.New("no flags find")
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
func init() {
	RootCmd.PersistentFlags().StringVarP(&confType, "config-type", "t", "file", "the service config type [file/env/etcd]")
	RootCmd.PersistentFlags().StringVarP(&confFile, "config-file", "f", "etc/demo.toml", "the service config from file")
	RootCmd.PersistentFlags().StringVarP(&confETCD, "config-etcd", "e", "127.0.0.1:2379", "the service config from etcd")
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "the demo version")
}
