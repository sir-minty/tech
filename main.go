package main

//
import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sir-minty/tech/views"
)

func main() {
	ssl := flag.Bool("ssl", false, "flag enable ssl")
	port := flag.String("port", os.Getenv("PORT"), "port to run server on")
	host := flag.String("host", os.Getenv("HOST"), "host to run server on")
	keyPem := flag.String("key", "key.pem", "location of your key.pem file")
	certPem := flag.String("cert", "cert.pem", "location of your cert.pem file")

	dbname := flag.String("db-name", "foobar", "db name")
	dbport := flag.String("db-port", "3306", "db port")
	dbhost := flag.String("db-host", "172.17.0.2", "db host url")
	dbusername := flag.String("db-username", "root", "db username")
	dbpassword := flag.String("db-password", "root", "db password")

	flag.Parse()

	dbURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		*dbusername,
		*dbpassword,
		*dbhost,
		*dbport,
		*dbname,
	)
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	c := views.NewContext(db)

	r := mux.NewRouter()
	r.HandleFunc("/login", c.LoginHandler)

	router := handlers.LoggingHandler(os.Stdout, r)
	log.Printf("Serving on port %s:%s with ssl %t", *host, *port, *ssl)

	if *ssl {
		http.ListenAndServeTLS(fmt.Sprintf("%s:%s", *host, *port), *certPem, *keyPem, router)
	} else {
		http.ListenAndServe(fmt.Sprintf("%s:%s", *host, *port), router)
	}
}
