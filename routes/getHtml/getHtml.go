package goHtml

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	get_all_query "server/pack/db/getAllQuery"
	getOne "server/pack/db/getone"
	htmlResponseType "server/pack/model/htmlResponse"
	userinfo "server/pack/model/userInfoModel"
)
func Singledata(req *http.Request)userinfo.UserInfo{
	email:=req.FormValue("email")
	userdata,_:= getOne.GetOne(email)
	return userdata
}
func GetHtmlRes(res http.ResponseWriter,req *http.Request){
	var userdata  []userinfo.UserInfo;
	get_all_query.GetAllQuery(&userdata)
	data:=htmlResponseType.GetAllUser{Data:userdata,SingleData:[]userinfo.UserInfo{Singledata(req)}}
	fmt.Println(data)
	fileTemplates:=path.Join("templates","get.html")
	temp,err:=template.ParseFiles(fileTemplates)
	if err != nil{
		fmt.Println(err)
	}
	temp.Execute(res,data)
}