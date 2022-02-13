package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/weehong/leetcode-tracker/database"
	"github.com/weehong/leetcode-tracker/graphql"
	"github.com/weehong/leetcode-tracker/query"
)

func main() {
	resp, err := graphql.Query()
	if err != nil {
		log.Fatalf("GraphQL Error: %s", err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	db, err := sql.Open("postgres", database.PsqlConnection(host, port, user, password, dbname))
	if err != nil {
		log.Fatalf("Failed to open connection to database: %s", err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %s", err.Error())
	}

	currentRecord := query.FetchLatestRecord(db)

	// Only update the database when there's new record
	if currentRecord != resp.ProblemsetQuestionList.Total {

		for i, question := range resp.ProblemsetQuestionList.Questions {
			query.InsertQuestion(db, question)
			if questionId, err := strconv.Atoi(question.FrontendQuestionID); err == nil {
				for _, tag := range question.TopicTags {
					query.InsertTag(db, tag)
					query.InsertQuestionTag(db, questionId, tag.ID)
				}
			}
			fmt.Printf("--------- %d record(s) inserted ---------\n\n", i+1)
		}

		query.UpdateOrCreateRecord(db, resp.ProblemsetQuestionList.Total)
		fmt.Printf("--------- Completed ---------\n\n")
	}
}
