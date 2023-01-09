package main

import (
	"database/sql"
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

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	//password := GetMD5Hash(p)

	db, err := openDB("web:ukraine@/somedb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt := "SELECT * FROM users WHERE username = ? AND password = ?"

	row := db.QueryRow(stmt, username, password)

	var id int
	tmpl, err := template.ParseGlob("./ui/static/html/*")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		tmpl.ExecuteTemplate(w, "Base", nil)
	default:
		tmpl.ExecuteTemplate(w, "Success", nil)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

//func GetMD5Hash(pass string) string {
//	hasher := md5.New()
//	hasher.Write([]byte(pass))
//	return hex.EncodeToString(hasher.Sum(nil))
//}