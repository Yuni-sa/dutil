/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	dutil "dutil/pkg"

	"github.com/spf13/cobra"
)

// rminsCmd represents the rmins command
var rminsCmd = &cobra.Command{
	Use:   "rmins",
	Short: "Remove insecure registry",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("Too many arguments. Remove one hostname at a time")
		} else {
			daemonfile, _ := cmd.Flags().GetString("daemon-file")
			hostname := args[0]
			port, _ := cmd.Flags().GetUint16("port")
			if port != 0 {
				hostname = fmt.Sprintf("%v:%v", hostname, port)
			}
			fmt.Printf("Removing registry %v from %v... \n", hostname, daemonfile)
			err := dutil.RemoveInsecure(daemonfile, hostname)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Insecure registry successfully removed, Please restart the docker service.")
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(rminsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	rminsCmd.PersistentFlags().StringP("daemon-file", "f", "/etc/docker/daemon.json", "A custom daemon.json file path other than /etc/docker/daemon.json")
	rminsCmd.PersistentFlags().Uint16P("port", "p", 0, "A custom port to access the registry")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rminsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
