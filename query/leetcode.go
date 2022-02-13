package query

import (
	"database/sql"
	"fmt"
	"io"

	"github.com/weehong/leetcode-tracker/model"
)

func FetchLatestRecord(db *sql.DB, w io.Writer) int {
	var total int
	q := `SELECT total FROM records ORDER BY created_at ASC LIMIT 1`
	err := db.QueryRow(q).Scan(&total)
	if err != nil {
		fmt.Fprintf(w, "Error occurred on Record during fetching the latest record: %s\n", err.Error())
	}
	return total
}

func UpdateOrCreateRecord(db *sql.DB, total int, w io.Writer) {
	q := `INSERT INTO records(total) VALUES ($1)`
	_, err := db.Exec(q, total)
	if err != nil {
		fmt.Fprintf(w, "Error occurred on Records table during insert new total record: %s\n", err.Error())
	}
	fmt.Printf("Latest record has been updated\n")
}

func InsertTag(db *sql.DB, tag model.TopicTag, w io.Writer) {
	q := `INSERT INTO tags(id, name) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := db.Exec(q, tag.ID, tag.Name)
	if err != nil {
		fmt.Fprintf(w, "Error occurred on Tag table during insertion at %s, %s: %s\n", tag.ID, tag.Name, err.Error())
	}
	fmt.Printf("Insert Tag record: %s\n", tag.Name)
}

func InsertQuestion(db *sql.DB, question model.Question, w io.Writer) {
	q := `INSERT INTO questions(id, title, slug, difficulty) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING`
	_, err := db.Exec(q, question.FrontendQuestionID, question.Title, question.TitleSlug, question.Difficulty)
	if err != nil {
		fmt.Fprintf(w, "Error occurred on Question table during insertion at No. - %s, with title - %s: %s\n", question.FrontendQuestionID, question.Title, err.Error())
	}
	fmt.Printf("Insert Question record: %s - %s\n", question.FrontendQuestionID, question.Title)
}

func InsertQuestionTag(db *sql.DB, questionId int, tag string, w io.Writer) {
	q := `INSERT INTO questions_tags(question_id, tag_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := db.Exec(q, questionId, tag)
	if err != nil {
		fmt.Fprintf(w, "Error occurred on QuestionTag table during insertion at Question No. - %d, and Tag No. - %s: %s\n", questionId, tag, err.Error())
	}
	fmt.Printf("Insert QuestionTag record: %d - %s\n", questionId, tag)
}
