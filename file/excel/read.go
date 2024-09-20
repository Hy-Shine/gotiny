package excel

import (
	"os"

	"github.com/hy-shine/gotiny/file"
	"github.com/xuri/excelize/v2"
)

// Read reads the contents of an Excel file and returns the data as a slice of slices of strings.
func Read(fileName string, sheetName ...string) ([][]string, error) {
	if !file.FileIsExists(fileName) {
		return nil, os.ErrNotExist
	}

	f, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var readSheet string
	if len(sheetName) > 0 {
		readSheet = sheetName[0]
	} else {
		readSheet = f.GetSheetName(0)
	}

	rows, err := f.GetRows(readSheet)
	return rows, err
}
