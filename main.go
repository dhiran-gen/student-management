package main

import (
	"database/sql"
	"log"
	"net/http"

	httpstudent "github.com/student-management/handler/student"
	servicestudent "github.com/student-management/service/student"
	storestudent "github.com/student-management/store/student"

	"github.com/gorilla/mux"
	"github.com/student-management/driver"
	"github.com/student-management/middileware"
)

func main() {
	// define the mysql configuration
	sqlConf := driver.MySQLConfig{
		Driver:   "mysql",
		Host:     "localhost",
		User:     "tuya",
		Password: "1234",
		Port:     "3306",
		Db:       "students",
	}

	// handle opening sql connection
	db, err := driver.ConnectToMySQL(sqlConf)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("error closing connection to sql %v", err)
		}
	}(db)
	if err != nil {
		log.Printf("error connecting to sql server %v", err)
	}

	// define each layer handlers
	s := storestudent.New(db)
	sv := servicestudent.New(s)
	usrHandler := httpstudent.New(sv)

	// define mux and routes with their handlers
	r := mux.NewRouter()
	r.Handle("/students", middleware.Authentication(http.HandlerFunc(usrHandler.GetAllstudentHandler))).Methods(http.MethodGet)
	r.Handle("/students/{id}", middleware.Authentication(http.HandlerFunc(usrHandler.GetstudentByIdHandler))).Methods(http.MethodGet)
	r.Handle("/students", middleware.Authentication(http.HandlerFunc(usrHandler.CreatestudentHandler))).Methods(http.MethodPost)
	r.Handle("/students/{id}", middleware.Authentication(http.HandlerFunc(usrHandler.UpdatestudentHandler))).Methods(http.MethodPut)
	r.Handle("/students/{id}", middleware.Authentication(http.HandlerFunc(usrHandler.DeletestudentHandler))).Methods(http.MethodDelete)

	// Run the server
	log.Printf("Listening on port 8000...")
	err = http.ListenAndServe(":8000", r)
	if err != nil {
		log.Printf("error creating server: %v", err)
	}
}
