package deleteAll

import (
	"server/pack/db/setup"
)
func DeleteAll(){
	db:=setup.Init()
	defer db.Close()
	db.Query("DELETE FROM GOLANG ")
}