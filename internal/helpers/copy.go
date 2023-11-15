/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package helpers

import "os"

func Copy(source, dest string) error {
	in, err := os.ReadFile(source)
	if err != nil {
		return err
	}

	if err := os.WriteFile(dest, in, 0644); err != nil {
		return err
	}
	return nil
}
