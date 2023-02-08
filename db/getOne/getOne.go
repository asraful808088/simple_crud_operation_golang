package getOne

import (
	"fmt"
	"server/pack/db/setup"
	userinfo "server/pack/model/userInfoModel"
)
func GetOne(email string) (userinfo.UserInfo,error){
	var user userinfo.UserInfo; 
	db:=setup.Init()
	defer db.Close()
	queryString:=fmt.Sprintf(`SELECT * FROM GOLANG WHERE EMAIL='%v'`,email)
	result,err:=  db.Query(queryString)
	if err!=nil{
		fmt.Println(err)
		return user,err
	}
	for result.Next(){
		var id int64
		var email string
		var name string
		var staff bool
		var admin bool
		result.Scan(&id,&email,&name,&staff,&admin)
		user=userinfo.UserInfo{Id: id,Email: email,Name: name,Staff: staff,Admin: admin}
		break
	}
	return user,nil
}