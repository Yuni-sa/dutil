/*
Copyright © 2023 Yonatan Sasson <yonatanxd72@gmail.com>
*/
package cmd

import (
	"fmt"

	addins "dutil/pkg"

	"github.com/spf13/cobra"
)

// addinsCmd represents the addins command
var addinsCmd = &cobra.Command{
	Use:   "addins",
	Short: "Add an insecure registry",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("Add an insecure registry in /etc/docker/daemon.json")
		//fmt.Println("Or where the users configured")
		if len(args) > 1 {
			fmt.Println("Too many arguments. Add one hostname at a time")
		} else {
			daemonfile, _ := cmd.Flags().GetString("daemon-file")
			hostname := args[0]
			port, _ := cmd.Flags().GetUint16("port")
			if port != 0 {
				hostname = fmt.Sprintf("%v:%v", hostname, port)
			}
			err := addins.AddInsecure(daemonfile, hostname)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Insecure registry successfully added")
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(addinsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	addinsCmd.PersistentFlags().StringP("daemon-file", "f", "/etc/docker/daemon.json", "A custom daemon.json file path other than /etc/docker/daemon.json")
	addinsCmd.PersistentFlags().Uint16P("port", "p", 0, "A custom port to access the registry")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addinsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
