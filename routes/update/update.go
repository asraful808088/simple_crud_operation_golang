package update

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	userupdate "server/pack/db/update"
	userinfo "server/pack/model/userInfoModel"
)

func GetHtmlRes(res http.ResponseWriter,req *http.Request){
	data:=userinfo.UserInfo{}

	if req.Method =="POST"{
	email:=req.FormValue("email")
	name:=req.FormValue("name")
	oldemail:=req.FormValue("oldemail")
	admin:=req.FormValue("admin")
	staff:=req.FormValue("staff")
	data.Email=email
	data.Name=name
	if staff =="staff"{
		data.Staff=true
	}else{
		data.Staff=false
	}
	if admin =="admin"{
		data.Admin=true
	}else{
		data.Admin=false
	}
	userupdate.UpdateUser(oldemail,&data)
	}
	fileTemplates:=path.Join("templates","update.html")
	temp,err:=template.ParseFiles(fileTemplates)
	if err != nil{
		fmt.Println(err)
	}
	temp.Execute(res,data)
}