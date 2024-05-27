package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
)

type PageData struct {
	Signature string
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/upload", UploadFile)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/favicon.ico")
	})
	http.ListenAndServe(":808", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// 解析上传的文件
	file, tempFile, err := parseAndSaveFile(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	defer tempFile.Close()

	// 检查签名
	signature, err := checkSignature(tempFile)
	if err != nil {
		http.Error(w, "Error Running Keytool Command", http.StatusInternalServerError)
		return
	}

	// 显示结果
	displayResult(w, signature)
}

func parseAndSaveFile(r *http.Request) (multipart.File, *os.File, error) {
	r.ParseMultipartForm(500 << 20) // 限制上传大小为500MB
	file, _, err := r.FormFile("apkfile")
	if err != nil {
		return nil, nil, fmt.Errorf("error retrieving the file")
	}

	tempFile, err := os.CreateTemp("", "upload-*.apk")
	if err != nil {
		return nil, nil, fmt.Errorf("error creating temp file")
	}

	_, err = io.Copy(tempFile, file)
	if err != nil {
		return nil, nil, fmt.Errorf("error writing file content")
	}

	return file, tempFile, nil
}

func checkSignature(tempFile *os.File) (string, error) {
	cmd := exec.Command("keytool", "-printcert", "-jarfile", tempFile.Name())
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error running keytool command")
	}
	return out.String(), nil
}

func displayResult(w http.ResponseWriter, signature string) {
	data := PageData{
		Signature: signature,
	}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, data)
}
