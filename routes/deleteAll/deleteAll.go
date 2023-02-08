package deleteAll

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"server/pack/db/deleteAll"
)
func DeleteAll(res http.ResponseWriter,req *http.Request){
	deleteAll.DeleteAll()
	fileTemplates:=path.Join("templates","delete.html")
	temp,err:=template.ParseFiles(fileTemplates)
	if err != nil{
		fmt.Println(err)
	}
	temp.Execute(res,nil)
}