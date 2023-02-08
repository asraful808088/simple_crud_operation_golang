package insertHtml

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	insert "server/pack/db/insertUser"
	userinfo "server/pack/model/userInfoModel"
)

func GetHtmlRes(res http.ResponseWriter,req *http.Request){
	data:=userinfo.UserInfo{}
	if req.Method == "POST"{
		email:=req.FormValue("email")
		name:=req.FormValue("name")
		admin:=req.FormValue("admin")
		staff:=req.FormValue("staff")
		data.Email=email
		data.Name=name
		if staff =="staff"{
			data.Staff=true
		}
		if admin =="admin"{
			data.Admin=true
		}
		insert.Insert(&data)
			fmt.Println(staff)
	}
	
	
	fileTemplates:=path.Join("templates","insert.html")
	temp,err:=template.ParseFiles(fileTemplates)
	if err != nil{
		fmt.Println(err)
	}
	temp.Execute(res,data)
}