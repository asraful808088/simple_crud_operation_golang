package userupdate

import (
	"errors"
	"fmt"
	"server/pack/db/getone"
	"server/pack/db/setup"
	userinfo "server/pack/model/userInfoModel"
)
func UpdateUser(userEmail string, changeInfo *userinfo.UserInfo)(error){
	fmt.Println(changeInfo)
	db:=setup.Init()
	defer db.Close()
	result,_ := getOne.GetOne(userEmail)
	if result.Id == 0{
		return errors.New("user not found");
	}
	fmt.Println(result)
	if changeInfo.Name != ""&&result.Name != changeInfo.Name{
		queryString:= fmt.Sprintf("UPDATE GOLANG SET NAME='%v' WHERE EMAIL='%v'",changeInfo.Name,userEmail)
		db.Query(queryString)
	}
	if result.Staff != changeInfo.Staff{
		queryString:= fmt.Sprintf("UPDATE GOLANG SET Staff='%v' WHERE EMAIL='%v'",changeInfo.Staff,userEmail)
		db.Query(queryString)
	}
	if result.Admin != changeInfo.Admin{
		queryString:= fmt.Sprintf("UPDATE GOLANG SET Admin='%v' WHERE EMAIL='%v'",changeInfo.Admin,userEmail)
		db.Query(queryString)
	}
	if changeInfo.Email != ""&& result.Email != changeInfo.Email{
		result,_ := getOne.GetOne(changeInfo.Email)
		if result.Id == 0{
		queryString:= fmt.Sprintf("UPDATE GOLANG SET EMAIL='%v' WHERE EMAIL='%v'",changeInfo.Email,userEmail)
		db.Query(queryString)
		}else{
			return errors.New("email already used");
		}
	}
return nil
}