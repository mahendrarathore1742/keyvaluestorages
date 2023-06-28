package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverAddress string
var cfgFile string
var key string

var rootCMD = &cobra.Command{
	Use:   "Key",
	Short: "CLI client for a server-side key value store with subscription feature",
	Long: `This is a CLI client to the Key value storage server.
	You can perform the following operations,
	GET all the key and value pairs:      KeyvalueCli --server-address <address> get --all
	GET a given key-value pair:   KeyvalueCli --server-address <address> get --key <key>
	SET a given key-value pair:   KeyvalueCli --server-address <address> set --key <key> --value <value>
	Watch the server for changes: KeyvalueCli --server-address <address> watch
	`,
}

func Execute() {
	cobra.CheckErr(rootCMD.Execute())

}

func init() {
	cobra.OnInitialize(initConfig)
	rootCMD.PersistentFlags().StringVar(&serverAddress, "server-address", ":3000", "The address of the key values datastore")

}

func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)

	} else {
		// Find the home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config fine in the  home directory with name ".cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cli")

		viper.AutomaticEnv() // read in environment variables that match

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	}

}
