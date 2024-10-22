package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	scheduling "github.com/madflojo/tasks"
	"github.com/nurtai325/kaspi/mailing/internal/config"
	"github.com/nurtai325/kaspi/mailing/internal/db"
	"github.com/nurtai325/kaspi/mailing/internal/handlers"
	"github.com/nurtai325/kaspi/mailing/internal/tasks"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Ltime | log.Lshortfile)

	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	conf := config.New()

	closeDB, err := db.Connect(conf)
	defer closeDB()
	if err != nil {
		log.Fatal(err)
	}

	jobs := []*scheduling.Task{
		tasks.NewOrders(),
		tasks.CompletedOrders(),
	}
	stop, err := tasks.Start(conf, jobs)
	defer stop()
	if err != nil {
		log.Fatal(err)
	}

	err = handlers.ParseViews()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handlers.HandleClientsView)
	http.HandleFunc("/add/", handlers.HandleAddClientView)
	http.HandleFunc("/add/client", handlers.HandleAddClient)
	http.HandleFunc("/extend/client", handlers.HandleExtendClientDate)
	http.HandleFunc("/deactivate", handlers.HandleDeactivate)
	http.HandleFunc("/qrcode", handlers.HandleConnectQrcode)

	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	log.Println("starting web server")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
