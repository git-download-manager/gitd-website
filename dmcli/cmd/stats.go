package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:     "stats",
	Short:   "generate clone repositories top list",
	Long:    `clone repositories analysis everyday and collects stats datas.`,
	Example: "./dmcli stats --config .stats.prod.yaml",
	RunE: func(cmd *cobra.Command, args []string) error {

		statsLogFile := viper.GetString("stats-log-file")

		currentPath, err := os.Getwd()
		if err != nil {
			fmt.Printf("error: pwd command: %v\n", err)
			panic(err)
		}

		// calculate stats from log files
		gcmd := exec.Command("/bin/bash", "./sh/clone_stats.sh", statsLogFile)
		gcmd.Dir = currentPath

		fmt.Printf("Run Command: %s\n", gcmd.String())

		out, err := gcmd.CombinedOutput()
		if err != nil {
			fmt.Printf("stats out: %v\n%s\n", out, out)
			fmt.Printf("stats error: %v\n", err)
			panic(err)
		}

		fmt.Printf("clone stats command output: %s\n", out)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
