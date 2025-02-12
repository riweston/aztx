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
// Package cmd provides the command-line interface for the aztx application.
// It implements the core functionality for switching between Azure tenants and subscriptions
// using a fuzzy finder interface.
package cmd

import (
	"errors"
	"os"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
	pkgerrors "github.com/riweston/aztx/pkg/errors"
	"github.com/riweston/aztx/pkg/profile"
	"github.com/riweston/aztx/pkg/state"
	"github.com/riweston/aztx/pkg/storage"
	"github.com/riweston/aztx/pkg/subscription"
	"github.com/riweston/aztx/pkg/tenant"
	"github.com/riweston/aztx/pkg/types"

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
	RunE: func(cmd *cobra.Command, args []string) error {
		stateManager := state.NewViperStateManager(viper.GetViper())
		storage := storage.FileAdapter{}
		if err := storage.FetchDefaultPath("/.azure/azureProfile.json"); err != nil {
			return pkgerrors.ErrFileOperation("fetching default profile path", err)
		}

		logger := profile.NewLogger(viper.GetString("log-level"))
		cfg, err := storage.ReadConfig()
		if err != nil {
			return pkgerrors.ErrReadingConfiguration(err)
		}

		if len(args) > 0 && args[0] == "-" {
			adapter := profile.NewConfigurationAdapter(&storage, logger)
			if err := adapter.SetPreviousContext(stateManager); err != nil {
				return pkgerrors.ErrSettingPreviousContext(err)
			}
			return nil
		}

		// Check if tenant selection is requested
		if viper.GetBool("by-tenant") {
			tenantManager := tenant.Manager{BaseManager: types.BaseManager{Configuration: cfg}}
			selectedTenant, err := tenantManager.FindTenantIndex()
			if err != nil {
				if errors.Is(err, fuzzyfinder.ErrAbort) {
					return nil
				}
				return pkgerrors.ErrTenantOperation("selecting tenant", err)
			}

			subManager := subscription.Manager{BaseManager: types.BaseManager{Configuration: cfg}}
			sub, err := subManager.FindSubscriptionIndexByTenant(selectedTenant.ID)
			if err != nil {
				if errors.Is(err, fuzzyfinder.ErrAbort) {
					return nil
				}
				return pkgerrors.ErrSelectingSubscription(err)
			}

			adapter := profile.NewConfigurationAdapter(&storage, logger)
			if err := adapter.SetContext(sub.ID); err != nil {
				return pkgerrors.ErrOperation("setting context", err)
			}
			return nil
		}

		// Default subscription selection
		adapter := profile.NewConfigurationAdapter(&storage, logger)
		sub, err := adapter.SelectWithFinder()
		if err != nil {
			if errors.Is(err, fuzzyfinder.ErrAbort) {
				return nil
			}
			return pkgerrors.ErrSelectingSubscription(err)
		}

		if err := adapter.SetContext(sub.ID); err != nil {
			return pkgerrors.ErrOperation("setting context", err)
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// It is called by main.main() and only needs to happen once to the rootCmd.
// Returns an error if the command execution fails.
func Execute() error {
	return rootCmd.Execute()
}

// init initializes the command configuration by setting up flags and binding them to viper.
// It is automatically called by cobra during command initialization.
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().String("log-level", "info", "Set log level (debug, info, warn, error)")
	rootCmd.Flags().Bool("by-tenant", false, "Select tenant before choosing subscription")

	// Bind flags to viper and check for errors
	if err := viper.BindPFlag("log-level", rootCmd.PersistentFlags().Lookup("log-level")); err != nil {
		logger := profile.NewLogger("error")
		logger.Error("Failed to bind log-level flag: %v", err)
		os.Exit(1)
	}
	if err := viper.BindPFlag("by-tenant", rootCmd.Flags().Lookup("by-tenant")); err != nil {
		logger := profile.NewLogger("error")
		logger.Error("Failed to bind by-tenant flag: %v", err)
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
// It looks for a .aztx.yml file in the user's home directory and creates one if it doesn't exist.
// The function will exit with status code 1 if there are any errors accessing the home directory
// or handling the configuration file.
func initConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		logger := profile.NewLogger("error")
		logger.Error("Failed to get home directory: %v", err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.SetConfigType("yml")
	viper.SetConfigName(".aztx")
	viper.SetEnvPrefix("AZTX")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	// Create config if it doesn't exist
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := viper.SafeWriteConfigAs(home + "/.aztx.yml"); err != nil {
				logger := profile.NewLogger("error")
				logger.Error("Failed to write config: %v", err)
				os.Exit(1)
			}
		} else {
			logger := profile.NewLogger("error")
			logger.Error("Failed to read config: %v", err)
			os.Exit(1)
		}
	}
}
