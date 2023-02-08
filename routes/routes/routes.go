package routes

import (
	deletehtml "server/pack/routes/delete"
	"server/pack/routes/deleteAll"
	"server/pack/routes/deleteOne"
	_ "server/pack/routes/deleteAll"
	goHtml "server/pack/routes/getHtml"
	insertHtml "server/pack/routes/insert"
	update "server/pack/routes/update"

	"github.com/gorilla/mux"
)
func GetRoutes() *mux.Router {

	router:=mux.NewRouter()
	router.HandleFunc("/",goHtml.GetHtmlRes)
	router.HandleFunc("/post",insertHtml.GetHtmlRes)
	router.HandleFunc("/update",update.GetHtmlRes)
	router.HandleFunc("/delete",deletehtml.DeleteHtmlRes)
	router.HandleFunc("/deleteAll",deleteAll.DeleteAll)
	router.HandleFunc("/deleteOne",deleteOne.DeleteOne)
	return router
}