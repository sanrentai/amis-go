package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/sanrentai/amis-go"
	"github.com/sanrentai/amis-go/form"
)

func handle(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./amis.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	data := amis.Page(
		amis.Page_title("表单页面"),
		amis.Page_body(
			amis.Form(
				amis.Form_mode("horizontal"),
				amis.Form_body(
					form.InputText(
						form.FormItem_label("名称"),
						form.FormItem_name("name"),
					),
					form.InputPassword(
						form.FormItem_label("密码"),
						form.FormItem_name("password"),
					),
				),
			),
		),
	)

	tmpl.Execute(w, amis.JSON(data, true))
}

func main() {
	http.HandleFunc("/", handle)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP SERVER failed,err:", err)
		return
	}
}
