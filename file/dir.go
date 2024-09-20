package file

import (
	"errors"
	"os"
)

func MkdirWithModePerm(dirPath string, mode os.FileMode) (err error) {
	if len(dirPath) == 0 {
		return errors.New("the dir path is empty")
	}
	if _, err = os.Stat(dirPath); err != nil {
		if !os.IsExist(err) {
			err = os.MkdirAll(dirPath, mode)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Mkdir create a specific dir
func Mkdir(dirPath string) error {
	return MkdirWithModePerm(dirPath, 0o755)
}

func PathCreate(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

// IsDirExist 判断目录是否存在
func IsDirExist(addr string) bool {
	s, err := os.Stat(addr)
	if err != nil {
		return os.IsExist(err)
	}
	return s.IsDir()
}
