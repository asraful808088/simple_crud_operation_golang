package deleteOne

import (
	"errors"
	"fmt"
	getOne "server/pack/db/getone"
	"server/pack/db/setup"
)
func DeleteOne(userEmail string) (error){
	db:=setup.Init()
	defer db.Close()
	result,_ := getOne.GetOne(userEmail)
	if result.Id == 0{
		return errors.New("user not found");
	}
	queryString:= fmt.Sprintf("DELETE FROM GOLANG WHERE EMAIL='%v'",userEmail)
	db.Query(queryString)
	return nil
}