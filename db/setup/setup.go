package setup

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)
const  (
	host="localhost"
	port=5432
	user="postgres"
	password="1111"
	dbname="golang"
	sslmode="disable"
)
func Init() *sql.DB{
	connectionString:=fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s ",host,port,user,password,dbname,sslmode)
	db,err:=sql.Open("postgres",connectionString)
	if err!=nil{
		fmt.Println(err)
	}
	return db
}