package core

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// Upload provides a method to upload an OVA file to the AM repository
func Upload(templates string, path string) (r *Reply, err error) {
	fd, err := os.Open(path)
	if err != nil {
		return
	}
	defer fd.Close()

	buffer := new(bytes.Buffer)
	writer := multipart.NewWriter(buffer)
	form, err := writer.CreateFormFile("diskFile", filepath.Base(path))
	if err != nil {
		return
	}

	io.Copy(form, fd)
	writer.Close()
	req, err := http.NewRequest("POST", templates, buffer)
	if err != nil {
		return
	}

	const maxMemory = 256 * 1024 * 1024
	req.ParseMultipartForm(maxMemory)
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return
	}

	r, err = newReply(res, nil)
	return
}
