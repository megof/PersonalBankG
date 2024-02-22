package account

import (
	"bd"
	"net/http"
	"text/template"
)

type Account struct {
	Id      int
	Name    string
	Amount  int
	Date    string
	Comment string
	User    int
}

var templates = template.Must(template.ParseGlob("templates/*"))

func Redir(temp *template.Template) {

}

func Host(w http.ResponseWriter, r *http.Request) {
	Connectionok := bd.ConnectionDB()
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
		Connectionok := bd.ConnectionDB()
		ins, err := Connectionok.Prepare("INSERT INTO acount (name, amount, date, comment, user) VALUES (?, ?, '2024-02-01', ?, 1)")
		if err != nil {
			panic(err.Error())
		}
		ins.Exec(name, amount, comment)
		http.Redirect(w, r, "/account", 301)
	}
}

func Remove(w http.ResponseWriter, r *http.Request) {
	idaccount := r.URL.Query().Get("id")
	Connectionok := bd.ConnectionDB()
	del, err := Connectionok.Prepare("DELETE FROM acount WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	del.Exec(idaccount)
	http.Redirect(w, r, "/account", 301)

}

func Update(w http.ResponseWriter, r *http.Request) {
	idaccount := r.URL.Query().Get("id")
	Connectionok := bd.ConnectionDB()
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
		Connectionok := bd.ConnectionDB()
		upt, err := Connectionok.Prepare("UPDATE acount SET name = ?, amount = ?, comment = ? WHERE Id = ?")
		if err != nil {
			panic(err.Error())
		}

		upt.Exec(name, amount, comment, idaccount)
		http.Redirect(w, r, "/account", 301)
	}

}
