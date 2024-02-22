package main

import (
	"account"
	"fmt"      //formato de entrada y salida
	"log"      //datos en consola
	"net/http" //mostrar web
)

func main() {
	http.HandleFunc("/account/", account.Host)
	http.HandleFunc("/account/insert", account.Insert)
	http.HandleFunc("/account/add", account.Add)
	http.HandleFunc("/account/delete", account.Remove)
	http.HandleFunc("/account/update", account.Update)
	http.HandleFunc("/account/mod", account.Mod)
	log.Println("Servidor corriendo en localhost:8080")
	fmt.Println("---")
	http.ListenAndServe(":8080", nil)
}
