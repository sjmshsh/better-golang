package main

import "os"

type File struct {
	path    string
	content string
}

func writeManyFiles(files []File) error {
	for _, file := range files {
		f, err := os.Open(file.path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.WriteString(file.content)
		if err != nil {
			return err
		}

		err = f.Sync()
		if err != nil {
			return err
		}
	}
	return nil
}

func shouldWriteManyFiles(files []File) error {
	for _, file := range files {
		if err := func() error {
			f, err := os.Open(file.path)
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = f.WriteString(file.content)
			if err != nil {
				return err
			}

			return f.Sync()
		}(); err != nil {
			return err
		}
	}
	return nil
}
