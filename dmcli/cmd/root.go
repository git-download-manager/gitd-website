package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dmcli",
	Short: "gitd manager cli tool",
	Long:  `gitd manager command line tools`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	rootCmd.MarkPersistentFlagRequired("config")

	//rootCmd.PersistentFlags().StringP("webaddr", "", "localhost:81", "url for website")
	//rootCmd.PersistentFlags().StringP("filepath", "", "./", "file path to data write")

	//viper.BindPFlag("webAddr", rootCmd.PersistentFlags().Lookup("webaddr"))
	//viper.BindPFlag("filePath", rootCmd.PersistentFlags().Lookup("filepath"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		fmt.Println("Config file not found!")
		os.Exit(1)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("config file not found")
		} else {
			// Config file was found but another error was produced
			fmt.Println("config file found but something wrong")
		}

		fmt.Println("config file reader error:", err)
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
