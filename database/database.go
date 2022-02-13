package database

import (
	"fmt"
	"log"
	"strconv"
)

func PsqlConnection(host string, port string, user string, password string, db string) string {
	p, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal("Error occurred when converting port into integer")
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, p, user, password, db)
}
