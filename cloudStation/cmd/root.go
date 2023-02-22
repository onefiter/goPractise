package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version bool
)

var RootCmd = &cobra.Command{
	Use:     "cloud-station-cli",
	Long:    "cloud-station-cli 云中转站",
	Short:   "cloud-station-cli 云中转站",
	Example: "cloud-station-cli cmds",
	RunE: func(cmd *cobra.Command, args []string) error {
		if version {
			fmt.Println("cloud-station-cli v0.0.1")
		}
		return nil
	},
}

func init() {
	f := RootCmd.PersistentFlags()
	f.BoolVarP(&version, "verson", "v", false, "cloud station 版本信息")

}
