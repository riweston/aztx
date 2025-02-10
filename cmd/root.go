/*
Copyright © 2024 Richard Weston

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/riweston/aztx/pkg/profile"
	"github.com/riweston/aztx/pkg/state"
	"github.com/riweston/aztx/pkg/storage"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aztx",
	Short: "Azure Tenant Context Switcher",
	Long: `aztx is a command line tool that helps you switch between Azure tenants and subscriptions.
It provides a fuzzy finder interface to select subscriptions and remembers your last context.`,
	Args: cobra.MaximumNArgs(1),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cfg := state.ViperAdapter{Viper: viper.GetViper()}
		lc := state.NewStateReaderWriter(&cfg)

		reader := storage.FileAdapter{}
		if err := reader.FetchDefaultPath("/.azure/azureProfile.json"); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		writer := storage.FileAdapter{}
		if err := writer.FetchDefaultPath("/.azure/azureProfile.json"); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		c := profile.NewConfigurationAdapter(reader, writer)

		if len(args) > 0 && args[0] == "-" {
			if err := c.SetPreviousContext(lc); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			os.Exit(0)
		}

		sub, err := c.SelectWithFinder()
		if errors.Is(err, fuzzyfinder.ErrAbort) {
			os.Exit(0)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if err := c.SetContext(sub.ID); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
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
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		os.Exit(1)
	}

	// Search config in home directory with name ".aztx" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigType("yml")
	viper.SetConfigName(".aztx")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		os.Exit(1)
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; creating default
			if err := viper.SafeWriteConfigAs(home + "/.aztx.yml"); err != nil {
				fmt.Println("Can't write config:", err)
				os.Exit(1)
			}
		} else {
			// Config file was found but another error was produced
			fmt.Println("Error reading config:", err)
			os.Exit(1)
		}
	}
}
