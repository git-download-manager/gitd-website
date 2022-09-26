package cmd

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/segmentio/ksuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:     "clear",
	Short:   "clear expire uid folders",
	Long:    `delete expires folders inside git download manager temp folder`,
	Example: "./dmcli clear --config .clear.prod.yaml",
	RunE: func(cmd *cobra.Command, args []string) error {

		var counter int

		tempDir := viper.GetString("temp-dir")
		d, err := os.Open(tempDir)
		if err != nil {
			fmt.Printf("error: temp dir: %v\n", err)
			os.Exit(1)
		}
		defer d.Close()

		files, err := d.ReadDir(-1)
		if err != nil {
			fmt.Printf("error: read temp dir: %v\n", err)
			os.Exit(1)
		}

		lifetime := viper.GetInt64("temp-dir-lifetime")
		for _, f := range files {
			if !f.IsDir() {
				continue
			}

			if validUID(f.Name(), lifetime) {
				continue
			}

			//fmt.Println(f.Name())
			err := os.RemoveAll(path.Join(tempDir, f.Name()))
			if err != nil {
				return err
			}

			counter++
		}

		if counter > 0 {
			fmt.Printf("Deleted: %d / Total: %d\n", counter, len(files))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}

func validUID(uid string, lifetime int64) bool {
	if uid == "" {
		return false
	}

	ksuid, err := ksuid.Parse(uid)
	if err != nil {
		return false
	}

	now := time.Now().UnixMilli()
	past := ksuid.Time().UnixMilli()
	ttl := past + lifetime + lifetime // double lifetime

	//fmt.Println(now, past, ttl, g.Lifetime, now < ttl)

	return now < ttl
}
