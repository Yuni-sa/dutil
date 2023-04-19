/*
Copyright © 2023 Yonatan Sasson <yonatanxd72@gmail.com>
*/
package addins

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/buger/jsonparser"
)

// Registry list
var registries []string

func AddInsecure(daemonfile, hostname string) error {

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
			os.Exit(0)
		}

		// Add it to registry list
		registries = append(registries, string(value))
	}, "insecure-registries"); err != nil {
		return err
	}
	// Get the registry list
	//value, _, _, err := jsonparser.Get(file, "insecure-registries")

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

func checkFile(daemonfile string) error {
	_, err := os.Stat(daemonfile)
	if os.IsNotExist(err) {
		return err
	}
	return nil
}

func checkRegistryKey(file *[]byte) error {
	_, _, _, err := jsonparser.Get(*file, "insecure-registries")
	if err == jsonparser.KeyPathNotFoundError {
		*file, err = jsonparser.Set(*file, []byte("[]"), "insecure-registries")
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}
