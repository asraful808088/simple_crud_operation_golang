package deleteOne

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"server/pack/db/deleteOne"
)
func DeleteOne(res http.ResponseWriter,req *http.Request){
	deleteOne.DeleteOne(req.FormValue("email"))
	fileTemplates:=path.Join("templates","delete.html")
	temp,err:=template.ParseFiles(fileTemplates)
	if err != nil{
		fmt.Println(err)
	}
	temp.Execute(res,nil)
}