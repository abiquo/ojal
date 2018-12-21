package core

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Upload provides a method to upload templates to the AM repository
func Upload(templates, path, info string) (r *Reply, err error) {
	fd, err := os.Open(path)
	if err != nil {
		return
	}
	defer fd.Close()

	var form io.Writer
	buffer := new(bytes.Buffer)
	writer := multipart.NewWriter(buffer)

	if info != "" {
		form, err = writer.CreateFormField("diskInfo")
		if err != nil {
			return
		}
		io.Copy(form, strings.NewReader(info))
	}

	form, err = writer.CreateFormFile("diskFile", filepath.Base(path))
	if err != nil {
		return
	}
	io.Copy(form, fd)

	writer.Close()
	req, err := http.NewRequest("POST", templates, buffer)
	if err != nil {
		return
	}

	req.ParseMultipartForm(256 * 1024 * 1024)
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return
	}

	r, err = newReply(res, nil)
	return
}
