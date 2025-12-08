package dao

import (
	"database/sql"

	"myapp/dto"
)

func GetStudentList(db *sql.DB) ([]dto.Student, error) {

	var listRecords []dto.Student

	rows, err := db.Query(`SELECT * FROM student`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var student dto.Student
		if err := rows.Scan(&student.ID, &student.FirstName, &student.LastName); err != nil {
			return nil, err
		}
		listRecords = append(listRecords, student)
	}

	return listRecords, nil
}

func GetStudent(db *sql.DB, studentID int) (*dto.Student, error) {
	student := &dto.Student{}

	err := db.QueryRow(
		`SELECT * FROM student WHERE studentID = ?`,
		studentID,
	).Scan(&student.ID, &student.FirstName, &student.LastName)

	if err != nil {
		return nil, err
	}

	return student, nil

}

func CreateStudentRecord(db *sql.DB, student *dto.Student) error {

	_, err := db.Query(
		`INSERT INTO student (firstName, lastName) VALUES (?, ?)`,
		student.FirstName,
		student.LastName,
	)

	if err != nil {
		return err
	}

	return nil
}

func UpdateStudentRecord(db *sql.DB, student *dto.Student) error {

	_, err := db.Query(
		`UPDATE student
		SET firstName = ?, lastName = ?
		WHERE studentID = ?`,
		student.FirstName,
		student.LastName,
		student.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteStudentRecord(db *sql.DB, studentID int) error {

	_, err := db.Query(
		`DELETE FROM student WHERE studentID = ?`,
		studentID,
	)

	if err != nil {
		return err
	}

	return nil
}
