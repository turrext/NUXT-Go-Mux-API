package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Response struct {
	Persons []Person1 `json:"persons"`
}
type Person1 struct {
	Id         int       `json:"uid"`
	FirstName  string    `json:"username"`
	Departname string    `json:"departname"`
	Created    time.Time `json:"created"`
}

func main() {
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/persons", Persons).Methods("GET")
	router.HandleFunc("/add", addPersons).Methods("GET")
	router.HandleFunc("/register/{username}/{departname}/{password}", RegisterPerson).Methods("POST")
	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)
}

func RegisterPerson(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	fmt.Println("Register Person Used")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	username := params["username"]
	departname := params["departname"]
	password := params["password"]

	persons := prepareResponse()
	var exists bool
	exists = false
	for i := 0; i < len(persons); i++ {
		if persons[i].FirstName == username {
			exists = true
		}
	}
	if exists == false {
		fmt.Println(username + " Depart: " + departname)

		// insert
		stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
		checkErr(err)
		dt := time.Now()
		res, err := stmt.Exec(username, departname, dt)
		checkErr(err)

		id, err := res.LastInsertId()
		checkErr(err)

		fmt.Println(id)

		stmt, err = db.Prepare("INSERT INTO users(username, password) values(?,?)")
		checkErr(err)
		res, err = stmt.Exec(username, password)
		checkErr(err)

		id, err = res.LastInsertId()
		checkErr(err)

		fmt.Println(id)
		fmt.Fprintf(w, "Success. User Created!")
	} else {
		fmt.Fprintf(w, "This username already exists! Please use a different username!")
	}

}

func addPersons(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
	RunSqlite(w, r)
}

func Persons(w http.ResponseWriter, r *http.Request) {
	log.Println("entering persons end point")
	var response Response
	persons := prepareResponse()

	response.Persons = persons

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func prepareResponse() []Person1 {
	var persons []Person1
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	// query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	var person Person1
	for rows.Next() {
		err = rows.Scan(&person.Id, &person.FirstName, &person.Departname, &person.Created)
		checkErr(err)
		fmt.Println(person.Id)
		fmt.Println(person.FirstName)
		fmt.Println(person.Departname)
		fmt.Println(person.Created)
		persons = append(persons, person)
	}

	rows.Close() //good habit to close
	fmt.Println(persons)
	fmt.Println("=====-----=====")
	return persons
}

func RunSqlite(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	rows.Close() //good habit to close

	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/*
   CREATE TABLE `userinfo` (
       `uid` INTEGER PRIMARY KEY AUTOINCREMENT,
       `username` VARCHAR(64) NULL,
       `departname` VARCHAR(64) NULL,
       `created` DATE NULL
   );

   CREATE TABLE `users` (
       `uid` INTEGER PRIMARY KEY AUTOINCREMENT,
       `username` VARCHAR(64) NULL,
       `password` VARCHAR(64) NULL
   );
*/
