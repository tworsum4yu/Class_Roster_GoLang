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

func GetTeacher(db *sql.DB, teacherID int) (*dto.Teacher, error) {
	teacher := &dto.Teacher{}

	err := db.QueryRow(
		`SELECT * FROM teacher WHERE teacherID = ?`,
		teacherID,
	).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Title, &teacher.Office, &teacher.Department)

	if err != nil {
		return nil, err
	}

	return teacher, nil

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

func UpdateTeacherRecord(db *sql.DB, teacher *dto.Teacher) error {

	_, err := db.Query(
		`UPDATE teacher
		SET firstName = ?, lastName = ?, title = ?, office = ?, department = ?
		WHERE teacherID = ?`,
		teacher.FirstName,
		teacher.LastName,
		teacher.Title,
		teacher.Office,
		teacher.Department,
		teacher.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteTeacherRecord(db *sql.DB, teacherID int) error {

	_, err := db.Query(
		`DELETE FROM teacher WHERE teacherID = ?`,
		teacherID,
	)

	if err != nil {
		return err
	}

	return nil
}
