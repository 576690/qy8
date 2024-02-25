package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	index := IndexData{"login"}
	t := template.New("login.html")
	currPath, _ := filepath.Abs("../") //获取上一级路径
	t.ParseFiles(currPath + "/前端所有文件/blog 用户登陆系统/login.html")
	t.Execute(w, index)
}
