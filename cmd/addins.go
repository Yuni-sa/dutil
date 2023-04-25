/*
Copyright © 2023 Yonatan Sasson <yonatanxd72@gmail.com>

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
	"fmt"

	dutil "dutil/pkg"

	"github.com/spf13/cobra"
)

// addinsCmd represents the addins command
var addinsCmd = &cobra.Command{
	Use:   "addins",
	Short: "Add an insecure registry",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			fmt.Println("Not enough arguments. Specify a registry to add.")
		case 1:
			daemonfile, _ := cmd.Flags().GetString("daemon-file")
			hostname := args[0]
			port, _ := cmd.Flags().GetUint16("port")
			if port != 0 {
				hostname = fmt.Sprintf("%v:%v", hostname, port)
			}
			fmt.Printf("Adding registry %v to %v... \n", hostname, daemonfile)
			err := dutil.AddInsecure(daemonfile, hostname)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Insecure registry successfully added, Please restart the docker service.")
			}
		default:
			fmt.Println("Too many arguments. Add one hostname at a time")
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
