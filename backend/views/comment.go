package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func (*HTMLApi) Comment(w http.ResponseWriter, r *http.Request) {
	index := IndexData{"Comment"}
	t := template.New("comment.html")
	currPath, _ := filepath.Abs("../") //获取上一级路径
	t.ParseFiles(currPath + "/前端所有文件/blog 留言功能/留言.html")
	t.Execute(w, index)
}
