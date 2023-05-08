package database

import (
	"database/sql"
	"fmt"
	"go/models"

	_ "github.com/lib/pq"
)

func CreateDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://rxqqfizt:We3yatlHlIbLHP-6StNyZ4m1w3fqv2te@horton.db.elephantsql.com/rxqqfizt")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS pfc (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL,
			height INTEGER NOT NULL,
			weight INTEGER NOT NULL,
			calories INTEGER NOT NULL,
			protein INTEGER NOT NULL,
			fats INTEGER NOT NULL,
			carbohydrates INTEGER NOT NULL,
			sex TEXT NOT NULL,
			age INTEGER NOT NULL,
			waitingForWeight BOOL NOT NULL,
			waitingForHeight BOOL NOT NULL,
			waitingForAge BOOL NOT NULL,
			waitingForSex BOOL NOT NULL,
			waitingForActivity BOOL NOT NULL,
			ACTIVITY DOUBLE PRECISION NOT NULL
		);
	`

	if _, err := db.Exec(createTableQuery); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func InsertPFC(db *sql.DB, pfc models.PFC) error {
	db, err := sql.Open("postgres", "postgres://rxqqfizt:We3yatlHlIbLHP-6StNyZ4m1w3fqv2te@horton.db.elephantsql.com/rxqqfizt")
	if err != nil {
		panic(err)
	}
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM pfc WHERE user_id = $1", pfc.User_id).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("record with user_id %v already exists", pfc.User_id)
	}

	insertQuery := `
	INSERT INTO pfc (USER_ID, HEIGHT, WEIGHT, CALORIES, PROTEIN, FATS, CARBOHYDRATES, SEX, AGE, WAITINGFORAGE, WAITINGFORHEIGHT, WAITINGFORWEIGHT, WAITINGFORSEX, WAITINGFORACTIVITY, ACTIVITY)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
	`

	_, err = db.Exec(insertQuery, pfc.User_id, pfc.Weight, pfc.Height, pfc.Calories, pfc.Protein, pfc.Fats, pfc.Carbohydrates, pfc.Sex, pfc.Age, pfc.WaitingForAge, pfc.WaitingForHeight, pfc.WaitingForWeight, pfc.WaitingForSex, pfc.WaitingForActivity, pfc.Activity)
	if err != nil {
		return err
	}

	return nil
}

func UpdatePFCCell(db *sql.DB, userId int, cellName string, cellValue interface{}) error {
	_, err := CheckIfPFCExists(db, userId)
	if err != nil {
		panic(err)
	}

	db, err = sql.Open("postgres", "postgres://rxqqfizt:We3yatlHlIbLHP-6StNyZ4m1w3fqv2te@horton.db.elephantsql.com/rxqqfizt")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := fmt.Sprintf("UPDATE pfc SET %s = $1 WHERE user_id = $2", cellName)
	_, err = db.Exec(query, cellValue, userId)
	return err
}

func GetPFCCellValue(db *sql.DB, userId int, cellName string) (interface{}, error) {
	_, err := CheckIfPFCExists(db, userId)
	if err != nil {
		panic(err)
	}
	db, err = sql.Open("postgres", "postgres://rxqqfizt:We3yatlHlIbLHP-6StNyZ4m1w3fqv2te@horton.db.elephantsql.com/rxqqfizt")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var cellValue interface{}
	query := fmt.Sprintf("SELECT %s FROM pfc WHERE user_id = $1", cellName)
	err = db.QueryRow(query, userId).Scan(&cellValue)
	if err != nil {
		return nil, err
	}
	return cellValue, nil
}

func Calculate(db *sql.DB, userId int) (int, int, int, string) {
	_, err := CheckIfPFCExists(db, userId)
	if err != nil {
		panic(err)
	}
	db, err = sql.Open("postgres", "postgres://rxqqfizt:We3yatlHlIbLHP-6StNyZ4m1w3fqv2te@horton.db.elephantsql.com/rxqqfizt")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	r, err := GetPFCCellValue(db, userId, "age")
	if err != nil {
		panic(err)
	}
	age := r.(int)
	r, err = GetPFCCellValue(db, userId, "weight")
	if err != nil {
		panic(err)
	}
	weight := r.(int)
	r, err = GetPFCCellValue(db, userId, "height")
	if err != nil {
		panic(err)
	}
	height := r.(int)
	r, err = GetPFCCellValue(db, userId, "sex")
	if err != nil {
		panic(err)
	}
	sex := r.(string)
	return age, weight, height, sex
}

func CheckIfPFCExists(db *sql.DB, userId int) (bool, error) {
	db, err := sql.Open("postgres", "postgres://rxqqfizt:We3yatlHlIbLHP-6StNyZ4m1w3fqv2te@horton.db.elephantsql.com/rxqqfizt")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var exists bool
	query := "SELECT exists(SELECT 1 FROM pfc WHERE user_id=$1)"
	err = db.QueryRow(query, userId).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func GetPFCByUserID(db *sql.DB, userID int) (models.PFC, error) {
	db, err := sql.Open("postgres", "postgres://rxqqfizt:We3yatlHlIbLHP-6StNyZ4m1w3fqv2te@horton.db.elephantsql.com/rxqqfizt")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := `
		SELECT user_id, calories, protein, fats, carbohydrates, sex, age, waitingForAge,
		waitingForHeight, waitingForWeight, activity, weight, height, waitingForSex, waitingForActivity
		FROM pfc
		WHERE user_id = $1
		LIMIT 1;
	`
	row := db.QueryRow(query, userID)
	pfc := models.PFC{}
	err = row.Scan(
		&pfc.User_id, &pfc.Calories, &pfc.Protein, &pfc.Fats, &pfc.Carbohydrates,
		&pfc.Sex, &pfc.Age, &pfc.WaitingForAge, &pfc.WaitingForHeight, &pfc.WaitingForWeight,
		&pfc.Activity, &pfc.Weight, &pfc.Height, &pfc.WaitingForSex, &pfc.WaitingForActivity,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.PFC{}, nil
		}
		return models.PFC{}, err
	}
	return pfc, nil
}

func UpdatePFC(db *sql.DB, pfc models.PFC) error {
	db, err := sql.Open("postgres", "postgres://rxqqfizt:We3yatlHlIbLHP-6StNyZ4m1w3fqv2te@horton.db.elephantsql.com/rxqqfizt")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	updateQuery := `
        UPDATE pfc
        SET calories = $1, protein = $2, fats = $3, carbohydrates = $4,
            sex = $5, age = $6, waitingForAge = $7, waitingForHeight = $8,
            waitingForWeight = $9, activity = $10, weight = $11, height = $12,
			waitingForSex = $13, waitingForActivity = $14
        WHERE user_id = $15;
    `
	_, err = db.Exec(updateQuery, pfc.Calories, pfc.Protein, pfc.Fats, pfc.Carbohydrates,
		pfc.Sex, pfc.Age, pfc.WaitingForAge, pfc.WaitingForHeight, pfc.WaitingForWeight,
		pfc.Activity, pfc.Weight, pfc.Height, pfc.WaitingForSex, pfc.WaitingForActivity, pfc.User_id)
	if err != nil {
		return err
	}
	return nil
}
