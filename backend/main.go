package main

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

type IndexData struct {
	Title string
	Desc  string
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	currPath, _ := filepath.Abs("../") //获取上一级路径
	t.ParseFiles(currPath + "/前端所有文件/blog 主页/index.html")
	//log.Println(currPath + "/前端所有文件/blog 主页/index.html")
	t.Execute(w, index)
} //index主页

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	//页面映射
	http.HandleFunc("/", index)

	error := server.ListenAndServe()
	if error != nil {
		log.Println(error)
	}
	// 启动服务器
}
