package mySchema

type Score struct {
	ID            int
	CompanyID     int
	Score         float64
	CalculateDate string
	ScoreGrade    string
}

type Company struct {
	ID      int
	Name    string
	Address string
	Scores  []Score
}
