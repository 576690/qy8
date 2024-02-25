package router

import (
	"net/http"
	"path/filepath"

	"github.com/576690/qy8/backend/api"
	"github.com/576690/qy8/backend/views"
)

func Router() {
	//页面映射
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/blog 主页/index.html", views.HTML.Index)
	currPath, _ := filepath.Abs("../") //获取上一级路径

	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/blog 用户登陆系统/login.html", views.HTML.Login)
	http.HandleFunc("/blog 留言功能/留言.html", views.HTML.Comment)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(currPath+"/前端所有文件/blog 主页/css/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir(currPath+"/前端所有文件/blog 主页/js/"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(currPath+"/前端所有文件/blog 主页/images/"))))
	http.Handle("/resourses/", http.StripPrefix("/resourses/", http.FileServer(http.Dir(currPath+"/前端所有文件/blog 主页/resourses/")))) // 静态文件服务
}
