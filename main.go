package main

import (
	"database/sql"
	"fmt"
	"log"
	// "time"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/pg", HelloServer)
	http.HandleFunc("/", ItWorks)
    http.ListenAndServe(":8080", nil)
}

func ItWorks(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "It works!, %s!", r.URL.Path[1:])
}


func HelloServer(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	
	db, err := sql.Open("postgres", "postgres://postgres:12345@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	var roleName string
	err = db.QueryRow("select role_name from information_schema.applicable_roles;").Scan(&roleName)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, roleName, r.URL.Path[1:])
}


func CreateMessage()  {
	// get message from web and insert it to database
}

func main2() {
	db, err := sql.Open("postgres", "postgres://postgres:12345@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	var roleName string
	err = db.QueryRow("select role_name from information_schema.applicable_roles;").Scan(&roleName)
	if err != nil {
		panic(err)
	}

	fmt.Print(roleName)

	// res, err := db.Exec(`DELETE FROM my_table WHERE expires_at = $1`, time.Now())
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Print(res)

	// var num int
	// start := time.Now()
	// for i := 0; i < 10; i++ {
	// 	err = db.QueryRow("select count(*) from pg_stat_activity where datname = $1", "postgres").Scan(&num)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// fmt.Print(time.Since(start))
}