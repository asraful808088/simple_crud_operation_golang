package main

import (
	"fmt"
	"log"
	"net/http"
	"server/pack/routes/routes"
)

func main(){
	PORT:=":8000"
	fmt.Printf("server start on %v",PORT)
	if err:=http.ListenAndServe(PORT,routes.GetRoutes());err!=nil{
		log.Fatal(err)
	}
}