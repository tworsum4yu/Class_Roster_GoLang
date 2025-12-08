package dao

import (
	"database/sql"

	"myapp/dto"
)

func GetTeacherList(db *sql.DB) ([]dto.Teacher, error) {

	var listRecords []dto.Teacher

	rows, err := db.Query(`SELECT * FROM teacher`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var teacher dto.Teacher
		if err := rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Title, &teacher.Office, &teacher.Department); err != nil {
			return nil, err
		}
		listRecords = append(listRecords, teacher)
	}

	return listRecords, nil
}

func CreateTeacherRecord(db *sql.DB, teacher *dto.Teacher) error {

	_, err := db.Query(
		`INSERT INTO student (firstName, lastName, title, office, department) VALUES (?, ?, ?, ?, ?)`,
		teacher.FirstName,
		teacher.LastName,
		teacher.Title,
		teacher.Office,
		teacher.Department,
	)

	if err != nil {
		return err
	}

	return nil
}
