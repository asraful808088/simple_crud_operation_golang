package get_all_query

import (
	"fmt"
	"server/pack/db/setup"
	userinfo "server/pack/model/userInfoModel"
)
func GetAllQuery(userdata *[]userinfo.UserInfo) {
	db:=setup.Init()
	defer db.Close()
	result,err:=db.Query(`SELECT * FROM "golang"`)
	if err!=nil{
		fmt.Println(err)
	}
	for result.Next(){
		var  id    int64
		var  name  string
		var  email string
		var  staff bool
		var  admin bool
		result.Scan(&id,&email,&name,&staff,&admin)
		*userdata = append(*userdata, userinfo.UserInfo{Id: id,Name: name,Email: email,Staff: staff,Admin: admin})
	}
	
}