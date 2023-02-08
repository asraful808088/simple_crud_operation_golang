package insert

import (
	"fmt"
	"server/pack/db/setup"
	userinfo "server/pack/model/userInfoModel"
)
func Insert(user *userinfo.UserInfo)  {
	queryString:=fmt.Sprintf(`INSERT INTO golang("email","name","staff","admin")values('%v','%v','%t','%t')`,user.Email,user.Name,user.Staff,user.Admin)
	db:=setup.Init()
   	defer db.Close()
	_,err:=db.Query(queryString)
	if err!=nil{
		fmt.Println(err)
		return;
	}
}