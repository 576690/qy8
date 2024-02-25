package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type IndexData struct {
	Title string
}

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := IndexData{"index"}
	t := template.New("index.html")
	currPath, _ := filepath.Abs("../") //获取上一级路径
	t.ParseFiles(currPath + "/前端所有文件/blog 主页/index.html")
	//log.Println(currPath + "/前端所有文件/blog 主页/index.html")
	t.Execute(w, index)
}
