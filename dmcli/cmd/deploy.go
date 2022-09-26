package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:     "deploy",
	Short:   "deploy app",
	Long:    `deploy app on the server side`,
	Example: "./dmcli deploy --config .deploy.prod.yaml",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("deploy called")

		// Deploy Server Config
		ipAddr := viper.GetString("ipaddr")
		domain := viper.GetString("domain")

		currentPath, err := os.Getwd()
		if err != nil {
			fmt.Printf("error: go build: %v\n", err)
			panic(err)
		}

		// Services
		services := viper.GetStringSlice("services")

		if len(services) > 0 {
			for _, service := range services {

				// Service Parameters
				serviceParams := viper.GetStringMapString(service)

				// Deploy:
				gcmd := exec.Command("/bin/bash", "./sh/deploy.sh", ipAddr, domain, serviceParams["httpaddr"], serviceParams["foldername"], serviceParams["filename"], serviceParams["subdomain"])
				gcmd.Dir = currentPath
				fmt.Printf("Run Command: %s\n", gcmd.String())

				out, err := gcmd.CombinedOutput()
				if err != nil {
					fmt.Printf("deploy out: %v\n", out)
					fmt.Printf("deploy error: %v\n", err)
					panic(err)
				}

				fmt.Printf("%s service init command output: %s\n", service, out)

				time.Sleep(1 * time.Second)
			}
		}

		// Deploy After:
		time.Sleep(1 * time.Second)

		gcmdAfter := exec.Command("/bin/bash", "./sh/deploy_after.sh", ipAddr, domain)
		gcmdAfter.Dir = currentPath
		fmt.Printf("Run Command: %s\n", gcmdAfter.String())

		out, err := gcmdAfter.CombinedOutput()
		if err != nil {
			fmt.Printf("deploy out: %v\n%s\n", out, out)
			fmt.Printf("deploy error: %v\n", err)
			panic(err)
		}

		time.Sleep(1 * time.Second)
		fmt.Printf("Deploy after command output: %s\n", out)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
