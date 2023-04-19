package dutil

import (
	"os"

	"github.com/buger/jsonparser"
)

func checkFile(file string) error {
	_, err := os.Stat(file)
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
