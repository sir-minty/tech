package main

//
import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

	dbname := flag.String("db-name", os.Getenv("DB_NAME"), "db name")
	dbport := flag.String("db-port", os.Getenv("DB_PORT"), "db port")
	dbhost := flag.String("db-host", os.Getenv("DB_HOST"), "db host")
	dbusername := flag.String("db-username", os.Getenv("DB_USERNAME"), "db username")
	dbpassword := flag.String("db-password", os.Getenv("DB_PASSWORD"), "db password")

	flag.Parse()

	dbURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		*dbusername, *dbpassword,
		*dbhost, *dbport,
		*dbname,
	)
	log.Printf("Connecting to DB at %s", dbURL)
	var db *sql.DB
	var err error
	backoff := 200
	for attempt := 0; attempt < 5; attempt++ {
		db, err = sql.Open("mysql", dbURL)
		if err != nil {
			log.Fatal(err)
			continue
		}
		// Check that we can actually ping the DB
		if err := db.Ping(); err != nil {
			log.Printf("Ping attempt %d, sleeping for %d: %s", attempt+1, backoff, err.Error())
			time.Sleep(time.Millisecond * time.Duration(backoff))
			backoff = backoff * 2
			continue
		}
		break
	}

	log.Println("Ping success")
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
