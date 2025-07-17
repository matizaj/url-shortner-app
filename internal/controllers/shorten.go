package controllers

import (
	"database/sql"
	"github.com/matizaj/url-shortner-app/internal/db"
	"github.com/matizaj/url-shortner-app/internal/url"
	"net/http"
	"strings"
	"text/template"
)

func Shorten(sql *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		originalUrl := r.FormValue("url")
		if originalUrl == "" {
			http.Error(w, "Url can't be empty", http.StatusBadRequest)
			return
		}

		if !strings.HasPrefix(originalUrl, "http://") && !strings.HasPrefix(originalUrl, "https://") {
			originalUrl = "https://" + originalUrl
		}

		shortURL := url.ShortenUrl(originalUrl)
		if err := db.StoreUrl(sql, shortURL, originalUrl); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := map[string]string{"ShortURL": shortURL}
		t, err := template.ParseFiles("internal/views/shorten.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = t.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func Proxy(lite *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortUrl := r.URL.Path[1:]
		if shortUrl == "" {
			http.Error(w, "Url can't be empty", http.StatusBadRequest)
			return
		}

		originalUrl, err := db.GetOriginalUrl(lite, shortUrl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, originalUrl, http.StatusPermanentRedirect)
	}
}

func ShortenProxy(lite *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = db.GetAllUrls(lite)
		w.Write([]byte("ok"))
	}
}
