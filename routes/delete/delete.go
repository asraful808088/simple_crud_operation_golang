package delete

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func DeleteHtmlRes(res http.ResponseWriter,req *http.Request){
	
	fileTemplates:=path.Join("templates","delete.html")
	temp,err:=template.ParseFiles(fileTemplates)
	if err != nil{
		fmt.Println(err)
	}
	temp.Execute(res,nil)
}