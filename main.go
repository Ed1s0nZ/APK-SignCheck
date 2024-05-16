package main

import (
	"bytes"
	"html/template"
	"io"
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
	http.ListenAndServe(":8080", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// 解析上传的文件
	r.ParseMultipartForm(500 << 20) // 限制上传大小为500MB
	file, _, err := r.FormFile("apkfile")
	if err != nil {
		http.Error(w, "Error Retrieving the File", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// 创建临时文件存放APK
	tempFile, err := os.CreateTemp("", "upload-*.apk")
	if err != nil {
		http.Error(w, "Error Creating Temp File", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	// 读取上传的文件内容到临时文件
	// fileBytes, err := io.ReadAll(file)
	// if err != nil {
	// 	http.Error(w, "Error Reading File Content", http.StatusInternalServerError)
	// 	return
	// }
	// tempFile.Write(fileBytes)

	// 使用 io.Copy() 代替读后写的操作
	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "Error Writing File Content", http.StatusInternalServerError)
		return
	}

	// 调用keytool命令检查签名
	cmd := exec.Command("keytool", "-printcert", "-jarfile", tempFile.Name())
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {

		http.Error(w, "Error Running Keytool Command", http.StatusInternalServerError)
		return
	}

	// 显示结果
	data := PageData{
		Signature: out.String(),
	}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, data)
}
