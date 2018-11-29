package main

import (
	"html/template"
	"net/http"
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"fmt"
)

type customer struct {
	ID int
	First string
}

var tpl *template.Template
var db *sql.DB

func init() {
	var err error

	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	connStr := "dbname=thanksgiving user=turkey password=gravy host=localhost port=5432 sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("**** NO OPEN ****", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("**** NO PING WORKING ****", err)
	} else {
		fmt.Println("Ping successful")
	}
}

func main() {
	defer db.Close()

	http.HandleFunc("/", read)
	http.HandleFunc("/process", create)
	http.HandleFunc("/processtoo", ptoo)
	http.HandleFunc("/update", up)
	http.HandleFunc("/delete", del)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./resources"))))
	http.ListenAndServe(":8080", nil)
}

//READ
func read(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM customers;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	xc := []customer{}

	for rows.Next() {
		c := customer{}
		rows.Scan(&c.ID, &c.First)
		xc = append(xc, c)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", xc)
}

// CREATE
func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fn := r.FormValue("fffnnnaaammmeee")
	fmt.Println(fn)

	result, err := db.Exec("INSERT INTO customers (cfirst) VALUES ($1);", fn)
	if err != nil {
		log.Fatal(err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	fmt.Println("rows affected", n)

	http.Redirect(w, r, "/", http.StatusFound)
}


// UPDATE
func up(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	customerid := r.FormValue("recordid")

	row, err := db.Query("SELECT * FROM customers WHERE cid = $1", customerid)
	if err != nil {
		http.Error(w, "coudn't retrieve record in up", http.StatusInternalServerError)
		return
	}

	c := customer{}
	for row.Next() {
		row.Scan(&c.ID, &c.First)
	}

	tpl.ExecuteTemplate(w, "update.gohtml", c)
}


func ptoo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	id := r.FormValue("customerid")
	fn := r.FormValue("fffnnnaaammmeee")
	fmt.Println(id)
	fmt.Println(fn)

	result, err := db.Exec("UPDATE customers SET cfirst = $2 WHERE cid = $1;", id, fn)
	if err != nil {
		log.Fatal(err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	fmt.Println("rows affected", n)

	http.Redirect(w, r, "/", http.StatusFound)
}


// DELETE
func del(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	customerid := r.FormValue("recordid")


	result, err := db.Exec("DELETE FROM customers WHERE cid = $1;", customerid)
	if err != nil {
		http.Error(w, "didn't delete", http.StatusInternalServerError)
		return
	}

	n, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "didn't show rows affected", http.StatusInternalServerError)
		return
	}
	fmt.Println("rows affected", n)

	http.Redirect(w, r, "/", http.StatusFound)
}
