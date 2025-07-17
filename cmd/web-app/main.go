package main

import (
	"database/sql"
	"fmt"
	"github.com/matizaj/url-shortner-app/internal/controllers"
	"github.com/matizaj/url-shortner-app/internal/db"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	slite, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	if err := db.CreateTable(slite); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	defer slite.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			controllers.ShowIndex(w, r)
		} else {
			controllers.Proxy(slite)(w, r)
		}
	})
	http.HandleFunc("/shorten", controllers.Shorten(slite))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
