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

func GetScoreByCompanyID(companyID int) ([]mySchema.Score, error) {
	var db = initDB()

	defer db.Close()
	var scores []mySchema.Score

	// 定义查询语句
	query := `SELECT * FROM Scores WHERE company_id = ?`

	// 执行查询
	rows, err := db.Query(query, companyID)
	if err != nil {
		return scores, err
	}
	defer rows.Close()

	// 遍历查询结果并将数据添加到scores切片中
	for rows.Next() {
		var score mySchema.Score
		err := rows.Scan(&score.ID, &score.CompanyID, &score.Score, &score.CalculateDate, &score.ScoreGrade)
		if err != nil {
			return scores, err
		}
		scores = append(scores, score)
	}

	return scores, nil
}

func GetCompanyByID(id int) (mySchema.Company, error) {
	var db = initDB()

	defer db.Close()
	var company mySchema.Company

	query := `SELECT * FROM company WHERE id = ?`

	err := db.QueryRow(query, id).Scan(&company.ID, &company.Address, &company.Name)

	if err != nil {
		return company, err
	}
	return company, nil
}
