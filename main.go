package main

import (
	"account"
	"database/sql"
	"fmt"      //formato de entrada y salida
	"log"      //datos en consola
	"net/http" //mostrar web
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

// build ejeuta archivos
// run compila y ejecuta

var templates = template.Must(template.ParseGlob("templates/*"))

type Account struct {
	Id      int
	Name    string
	Amount  int
	Date    string
	Comment string
	User    int
}

func main() {
	fmt.Println(account.Hellow())
	http.HandleFunc("/", Host)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/add", Add)
	http.HandleFunc("/delete", Remove)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/mod", Mod)
	log.Println("Servidor corriendo en localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func ConnectionDB() (connection *sql.DB) {
	Driver := "mysql"
	User := "root"
	Password := ""
	Db := "personal"

	connection, err := sql.Open(Driver, User+":"+Password+"@tcp(127.0.0.1)/"+Db)

	if err != nil {
		panic(err.Error())
	}
	return connection
}

func Host(w http.ResponseWriter, r *http.Request) {
	Connectionok := ConnectionDB()
	cons, err := Connectionok.Query("SELECT * FROM acount")
	if err != nil {
		panic(err.Error())
	}
	account := Account{}
	accounts := []Account{}

	for cons.Next() {
		var id, amount, user int
		var name, comment, date string
		err = cons.Scan(&id, &name, &amount, &date, &comment, &user)
		if err != nil {
			panic(err.Error())
		}
		account.Id = id
		account.Amount = amount
		account.User = user
		account.Name = name
		account.Comment = comment
		account.Date = date
		accounts = append(accounts, account)
	}
	//fmt.Println(accounts)
	templates.ExecuteTemplate(w, "_index", accounts)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "_create", nil)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		amount := r.FormValue("amount")
		comment := r.FormValue("comment")
		Connectionok := ConnectionDB()
		ins, err := Connectionok.Prepare("INSERT INTO acount (name, amount, date, comment, user) VALUES (?, ?, '2024-02-01', ?, 1)")
		if err != nil {
			panic(err.Error())
		}
		ins.Exec(name, amount, comment)
		http.Redirect(w, r, "/", 301)
	}
}

func Remove(w http.ResponseWriter, r *http.Request) {
	idaccount := r.URL.Query().Get("id")
	Connectionok := ConnectionDB()
	del, err := Connectionok.Prepare("DELETE FROM acount WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	del.Exec(idaccount)
	http.Redirect(w, r, "/", 301)

}

func Update(w http.ResponseWriter, r *http.Request) {
	idaccount := r.URL.Query().Get("id")
	Connectionok := ConnectionDB()
	upt, err := Connectionok.Query("SELECT * FROM acount WHERE id = " + idaccount) //Id=?", idaccount
	if err != nil {
		panic(err.Error())
	}
	account := Account{}

	for upt.Next() {
		var id, amount, user int
		var name, comment, date string
		err = upt.Scan(&id, &name, &amount, &date, &comment, &user)
		if err != nil {
			panic(err.Error())
		}
		account.Id = id
		account.Amount = amount
		account.User = user
		account.Name = name
		account.Comment = comment
		account.Date = date
	}
	//fmt.Println(accounts)
	templates.ExecuteTemplate(w, "_update", account)
}

func Mod(w http.ResponseWriter, r *http.Request) {
	idaccount := r.URL.Query().Get("id")

	if r.Method == "POST" {
		name := r.FormValue("name")
		amount := r.FormValue("amount")
		comment := r.FormValue("comment")
		Connectionok := ConnectionDB()
		upt, err := Connectionok.Prepare("UPDATE acount SET name = ?, amount = ?, comment = ? WHERE Id = ?")
		if err != nil {
			panic(err.Error())
		}

		upt.Exec(name, amount, comment, idaccount)
		http.Redirect(w, r, "/", 301)
	}

}

func ejemplo1() {
	fmt.Println("asdasd")

	var name string           //definir tipo de variable value=""
	var last string = "Mega"  // value="Mega"
	age := 20                 //asignado e inferido de tipo
	var uno, dos, tres string //multidef
	var cua, cin, sei = "4", "5", 6
	name = "Alpha"
	fmt.Println(name, last, age, uno, dos, tres, cua, cin, sei)
}
