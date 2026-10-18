package blockfile

import (
	"errors"
	"os"
	"path/filepath"

	"ipfs-server/cid/validator"
)

func Put(base, cid string, data []byte) error {
	path, err := pathForCID(base, cid)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(base, 0755); err != nil {
		return err
	}

	file, err := os.CreateTemp(base, ".block-*")
	if err != nil {
		return err
	}
	tempPath := file.Name()
	defer os.Remove(tempPath)

	if err := file.Chmod(0644); err != nil {
		file.Close()
		return err
	}
	if _, err := file.Write(data); err != nil {
		file.Close()
		return err
	}
	if err := file.Sync(); err != nil {
		file.Close()
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	return os.Rename(tempPath, path)
}

func Get(base, cid string) ([]byte, error) {
	path, err := pathForCID(base, cid)
	if err != nil {
		return nil, err
	}
	return os.ReadFile(path)
}

func Delete(base, cid string) error {
	path, err := pathForCID(base, cid)
	if err != nil {
		return err
	}
	return os.Remove(path)
}

func Exists(base, cid string) bool {
	path, err := pathForCID(base, cid)
	if err != nil {
		return false
	}
	_, err = os.Stat(path)
	return err == nil
}

func pathForCID(base, cid string) (string, error) {
	if !validator.IsValid(cid) {
		return "", errors.New("invalid CID")
	}
	return filepath.Join(base, cid), nil
}
