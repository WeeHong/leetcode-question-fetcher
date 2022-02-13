package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/weehong/leetcode-tracker/database"
	"github.com/weehong/leetcode-tracker/graphql"
	"github.com/weehong/leetcode-tracker/query"
)

var host, port, user, password, dbname string
var ssl *bool

func init() {
	ssl = flag.Bool("ssl", false, "Require SSL to connect database")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load env file")
	}

	host = os.Getenv("POSTGRES_HOST")
	port = os.Getenv("POSTGRES_PORT")
	user = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname = os.Getenv("POSTGRES_DB")
}

func main() {
	p := "/temp/leetcode-fetcher"
	if _, err := os.Stat(p); os.IsNotExist(err) {
		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	f, err := os.OpenFile("/temp/leetcode-fetcher/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	dt := time.Now()
	fmt.Fprintln(w, "Script Triggered: ", dt.String())

	resp, err := graphql.Query()
	if err != nil {
		fmt.Fprintf(w, "GraphQL Error: %s", err.Error())
	}

	db, err := sql.Open("postgres", database.PsqlConnection(host, port, user, password, dbname, ssl))
	if err != nil {
		fmt.Fprintf(w, "Failed to open connection to database: %s", err.Error())
		os.Exit(-1)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Fprintf(w, "Failed to ping database: %s", err.Error())
		os.Exit(-1)
	}

	currentRecord := query.FetchLatestRecord(db, w)

	// Only update the database when there's new record
	if currentRecord != resp.ProblemsetQuestionList.Total {

		for i, question := range resp.ProblemsetQuestionList.Questions {
			query.InsertQuestion(db, question, w)
			if questionId, err := strconv.Atoi(question.FrontendQuestionID); err == nil {
				for _, tag := range question.TopicTags {
					query.InsertTag(db, tag, w)
					query.InsertQuestionTag(db, questionId, tag.ID, w)
				}
			}
			fmt.Fprintf(w, "%d record(s) inserted.\n", i+1)
		}

		query.UpdateOrCreateRecord(db, resp.ProblemsetQuestionList.Total, w)
		fmt.Fprintf(w, "Completed.\n")
	}
	w.Flush()
}
