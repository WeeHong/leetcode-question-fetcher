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
	"github.com/weehong/leetcode-tracker/model"
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
	checkDirectory()

	f := openFile()
	defer f.Close()

	w := bufio.NewWriter(f)

	dt := time.Now()
	fmt.Fprintln(w, "Script Triggered: ", dt.String())

	fmt.Println("Fetching Leetcode question(s) ...")
	resp := fetchLeetCodeQuestions(w)
	fmt.Println("LeetCode question(s) has been fetched successfully.")

	db := databaseConnection(w)
	defer db.Close()

	currentRecord := query.FetchLatestRecord(db, w)
	fmt.Println("Database connection has been fetched successfully.")

	if currentRecord != resp.ProblemsetQuestionList.Total {
		createRecord(currentRecord, resp, db, w)
	}

	w.Flush()
}

func checkDirectory() {
	p := "/temp/leetcode-fetcher"
	if _, err := os.Stat(p); os.IsNotExist(err) {
		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func openFile() *os.File {
	f, err := os.OpenFile("/temp/leetcode-fetcher/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func fetchLeetCodeQuestions(w *bufio.Writer) *model.LeetCode {
	resp, err := graphql.Query()
	if err != nil {
		fmt.Fprintf(w, "GraphQL Error: %s", err.Error())
	}
	return resp
}

func databaseConnection(w *bufio.Writer) *sql.DB {
	db, err := sql.Open("postgres", database.PsqlConnection(host, port, user, password, dbname, ssl))
	if err != nil {
		fmt.Fprintf(w, "Failed to open connection to database: %s", err.Error())
		os.Exit(-1)
	}

	err = db.Ping()
	if err != nil {
		fmt.Fprintf(w, "Failed to ping database: %s", err.Error())
		os.Exit(-1)
	}
	return db
}

func createRecord(currentRecord int, resp *model.LeetCode, db *sql.DB, w *bufio.Writer) {
	fmt.Fprintf(w, "Inserting data ...\n")
	for _, question := range resp.ProblemsetQuestionList.Questions {
		query.InsertQuestion(db, question, w)
		if questionId, err := strconv.Atoi(question.FrontendQuestionID); err == nil {
			for _, tag := range question.TopicTags {
				query.InsertTag(db, tag, w)
				query.InsertQuestionTag(db, questionId, tag.ID, w)
			}
		}
	}

	query.UpdateOrCreateRecord(db, resp.ProblemsetQuestionList.Total, w)
	fmt.Fprintf(w, "Data insertion completed.\n")
}
