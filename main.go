package main

import (
	"crypto/sha256"
	"encoding/hex"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type PageData struct {
	Title          string
	Name           string
	SuccessMessage string
	Sha256         string
	BcryptPass     string
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title: "Password Generator",
		Name:  "Henry",
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		plainPass := r.FormValue("plain_pass")
		plainPassHashed := sha256.Sum256([]byte(plainPass))

		bcryptPass, err := encryptPassword(hex.EncodeToString(plainPassHashed[:]))

		if err != nil {
			http.Error(w, "Could not encrypt password", http.StatusInternalServerError)
		}

		data.BcryptPass = bcryptPass
		data.Sha256 = hex.EncodeToString(plainPassHashed[:])
	}

	tmpl.Execute(w, data)
}

func encryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))),
	)
	http.HandleFunc("/", templateHandler)
	http.ListenAndServe(":3001", nil)

}
