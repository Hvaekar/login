package main

import (
	"database/sql"
	"encoding/json"
	"github.com/Hvaekar/login/pkg/models"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Главная страница"))
}

func (app *app) login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	//password := GetMD5Hash(p)

	stmt := "SELECT * FROM users WHERE username = ? AND password = ?"

	row := app.DB.QueryRow(stmt, username, password)

	user := &models.User{}
	tmpl, err := template.ParseGlob("./ui/static/html/*")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	switch err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Name); err {
	case sql.ErrNoRows:
		tmpl.ExecuteTemplate(w, "Base", nil)
	default:
		acc := models.Account{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		}
		j, err := json.Marshal(acc)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		w.Write(j)
	}
}

//func GetMD5Hash(pass string) string {
//	hasher := md5.New()
//	hasher.Write([]byte(pass))
//	return hex.EncodeToString(hasher.Sum(nil))
//}
