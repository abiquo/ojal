package core

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const maxMemory = 256 * 1024 * 1024

// Upload provides a method to upload an OVA file to the AM repository
func Upload(templates string, path string) (r *Reply, err error) {
	var f *os.File
	var req *http.Request
	var res *http.Response
	f, err = os.Open(path)
	if err == nil {
		defer f.Close()
		var form io.Writer
		buffer := new(bytes.Buffer)
		writer := multipart.NewWriter(buffer)
		if form, err = writer.CreateFormFile("diskFile", filepath.Base(path)); err == nil {
			io.Copy(form, f)
			writer.Close()
			if req, err = http.NewRequest("POST", templates, buffer); err == nil {
				req.ParseMultipartForm(maxMemory)
				req.Header.Set("Accept", "application/json, text/plain, */*")
				req.Header.Set("Content-Type", writer.FormDataContentType())
				if res, err = client.Do(req); err == nil {
					r, err = newReply(res, nil)
				}
			}
		}
	}
	return
}
