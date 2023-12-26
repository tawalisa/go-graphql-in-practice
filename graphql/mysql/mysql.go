package mysql

import (
	"database/sql"
	"go-grapohql-in-practice/graphql/mySchema"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//var db *sql.DB

func initDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/mockserver")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetScoreByID(id int) (mySchema.Score, error) {
	var db = initDB()

	defer db.Close()
	var score mySchema.Score

	query := `SELECT * FROM Scores WHERE id = ?`

	err := db.QueryRow(query, id).Scan(&score.ID, &score.CompanyID, &score.Score, &score.CalculateDate, &score.ScoreGrade)
	if err != nil {
		return score, err
	}

	return score, nil
}
