package main

import (
	"github.com/e61983/buyla-buy-la/buyla"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	data := Buyla.NewMetaData()

	bot, err := Buyla.NewBot(
		os.Getenv("ChannelSecret"),
		os.Getenv("ChannelAccessToken"),
		os.Getenv("TEST_URL"),
		data,
	)

	if err != nil {
		log.Fatal(err)
	}

	api := Buyla.NewApi(data)

	//staticFileServer := http.FileServer(http.Dir("static"))
	//http.HandleFunc("/static/", http.StripPrefix("/static/", staticFileServer).ServeHTTP)

	//downloadFileServer := http.FileServer(http.Dir(bot.DownloadDir))
	//http.HandleFunc("/downloaded/", http.StripPrefix("/downloaded/", downloadFileServer).ServeHTTP)

	r := mux.NewRouter()
	v1 := r.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/{gid}/orders", api.HandleGetOrders).Methods(http.MethodGet)
	v1.HandleFunc("/{gid}/order/{uid}", api.HandleGetOrder).Methods(http.MethodGet)
	v1.HandleFunc("/{gid}/order/{uid}", api.HandlePutOrder).Methods(http.MethodPut)
	v1.HandleFunc("/{gid}/order/{uid}", api.HandlePostOrder).Methods(http.MethodPost)
	v1.HandleFunc("/{gid}/order/{uid}", api.HandleDeleteOrder).Methods(http.MethodDelete)
	v1.HandleFunc("/{gid}/order/{uid}", api.HandlePatchOrder).Methods(http.MethodPatch)

	r.HandleFunc("/callback", bot.Callback)
	log.Println("Listen", os.Getenv("TEST_URL"), os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
