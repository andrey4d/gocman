package helpers

import "os"

func MakeDirAll(name string) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		if err := os.MkdirAll(name, 0755); err != nil {
			return err
		}
	}
	return nil
}

func MakeDir(name string) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		if err := os.Mkdir(name, 0755); err != nil {
			return err
		}
	}
	return nil
}
