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
package dutil

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/buger/jsonparser"
)

func AddInsecure(daemonfile, hostname string) error {
	// Declare registry list
	var registries []string

	// Check if the file exists
	if err := checkFile(daemonfile); err != nil {
		return err
	}

	// Read the file
	file, err := os.ReadFile(daemonfile)
	if err != nil {
		return err
	}

	// Check if the "insecure-registries" key exists and adds it if not

	if err = checkRegistryKey(&file); err != nil {
		return err
	}

	// Go over each value in the array

	if _, err = jsonparser.ArrayEach(file, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if err != nil {
			fmt.Println(err)
			return
		}
		if string(value) == hostname {
			err = fmt.Errorf("hostname already exists")
			fmt.Println(err)
			return
		}

		// Add it to registry list
		registries = append(registries, string(value))
	}, "insecure-registries"); err != nil {
		return err
	}

	// Add the new registry
	registries = append(registries, hostname)

	// Marshal the new registry list
	registryBytes, _ := json.Marshal(registries)

	// Set the new registry list
	file, err = jsonparser.Set(file, registryBytes, "insecure-registries")
	if err != nil {
		return err
	}

	// Finally make the changes to the file
	if err = os.WriteFile(daemonfile, file, 0644); err != nil {
		return err
	}

	return nil
}
