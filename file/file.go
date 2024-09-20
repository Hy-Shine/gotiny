package file

import (
	"bytes"
	"errors"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func FileCreate(content bytes.Buffer, name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content.String())
	return err
}

// FileDelete removes the named file.
func FileDelete(filePath string) error {
	return os.Remove(filePath)
}

func FileIsExists(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}

// GetExt returns the extension of a given file name.
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

func FileName(filePath string) (name string, extension string) {
	ext := GetExt(filePath)
	return strings.TrimSuffix(filepath.Base(filePath), ext), ext
}

// 获取文件大小
func GetFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}

// FileSize 获取文件大小
func FileSize(filePath string, fmt byte, bit int) (float64, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}

	size := float64(info.Size()) // bytes
	var f float64
	switch fmt {
	case 'b', 'B':
		f = size
	case 'k', 'K':
		f = size / 1024
	case 'm', 'M':
		f = size / (1024 * 1024)
	case 'g', 'G':
		f = size / (1024 * 1024 * 1024)
	case 't', 'T':
		f = size / (1024 * 1024 * 1024 * 1024)
	default:
		return 0, errors.New("unknown unit")
	}

	if bit < 0 {
		bit = 0
	}
	fs := strconv.FormatFloat(f, 'f', bit, 64)
	f, _ = strconv.ParseFloat(fs, 64)
	return f, nil
}
