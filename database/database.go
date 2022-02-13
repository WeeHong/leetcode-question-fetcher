package database

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func PsqlConnection(host string, port string, user string, password string, db string, ssl *bool) string {
	p, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal("Error occurred when converting port into integer")
	}
	var connectionString strings.Builder

	fmt.Fprintf(&connectionString, "host=%s port=%d user=%s password=%s dbname=%s", host, p, user, password, db)

	if *ssl {
		fmt.Fprintf(&connectionString, " sslmode=verify-full sslrootcert=%s", "./cert/ca-certificate.crt")
	} else {
		fmt.Fprintf(&connectionString, " sslmode=disable")
	}

	return connectionString.String()
}
