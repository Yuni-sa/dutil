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

// rminsCmd represents the rmins command
var rminsCmd = &cobra.Command{
	Use:   "rmins",
	Short: "Remove insecure registry",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			fmt.Println("Not enough arguments. Specify a registry to Remove.")
		case 1:
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
		default:
			fmt.Println("Too many arguments. Remove one hostname at a time")
		}
	},
}

func init() {
	rootCmd.AddCommand(rminsCmd)
	rminsCmd.PersistentFlags().StringP("daemon-file", "f", "/etc/docker/daemon.json", "A custom daemon.json file path other than /etc/docker/daemon.json")
	rminsCmd.PersistentFlags().Uint16P("port", "p", 0, "A custom port to access the registry")
}
