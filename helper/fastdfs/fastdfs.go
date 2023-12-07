package fastdfs

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
)

type FileInfo struct {
	Domain  string `json:"domain"`
	Md5     string `json:"md5"`
	Path    string `json:"path"`
	Retcode int    `json:"retcode"`
	Retsg   string `json:"retmsg"`
	Scene   string `json:"scene"`
	Scenes  string `json:"scenes"`
	Src     string `json:"src"`
	URL     string `json:"url"`
}

func Upload(file *multipart.FileHeader) (err error, fileInfo FileInfo) {
	// 打开文件
	src, err := file.Open()
	if err != nil {
		return err, fileInfo
	}
	defer src.Close()

	// 创建一个buffer来存储文件内容
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	_ = writer.WriteField("output", "json")
	part, err := writer.CreateFormFile("file", file.Filename)
	if err != nil {
		return err, fileInfo
	}
	_, err = io.Copy(part, src)
	if err != nil {
		return err, fileInfo
	}
	writer.Close()

	// 发送POST请求到go-fastdfs服务器
	resp, err := http.Post("http://192.168.100.142:8080/upload", writer.FormDataContentType(), &buf)
	if err != nil {
		return err, fileInfo
	}

	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, fileInfo
	}

	return json.Unmarshal(body, &fileInfo), fileInfo
}
