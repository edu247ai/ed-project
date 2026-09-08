package files

import (
	"io"
	"os"
)

func SaveLocal(path string, file io.Reader) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	return err
}
