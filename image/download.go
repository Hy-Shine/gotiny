package image

import (
	"net/http"
)

func Download(url, storePath string) error {
	rsp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	err = Write(rsp.Body, storePath)
	return err
}
