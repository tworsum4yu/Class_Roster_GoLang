package dao

import (
	"database/sql"

	"myapp/dto"
)

func GetCourseList(db *sql.DB) ([]dto.Course, error) {

	var listRecords []dto.Course

	rows, err := db.Query(`SELECT * FROM course`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var course dto.Course
		if err := rows.Scan(&course.ID, &course.Name); err != nil {
			return nil, err
		}
		listRecords = append(listRecords, course)
	}

	return listRecords, nil
}

func GetCourse(db *sql.DB, courseID int) (*dto.Course, error) {
	course := &dto.Course{}

	err := db.QueryRow(
		`SELECT * FROM course WHERE courseID = ?`,
		courseID,
	).Scan(&course.ID, &course.Name)

	if err != nil {
		return nil, err
	}

	return course, nil

}

func CreateCourseRecord(db *sql.DB, course *dto.Course) error {

	_, err := db.Query(
		`INSERT INTO course (courseName) VALUES (?)`,
		course.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

func UpdateCourseRecord(db *sql.DB, course *dto.Course) error {

	_, err := db.Query(
		`UPDATE course
		SET courseName = ?
		WHERE courseID = ?`,
		course.Name,
		course.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteCourseRecord(db *sql.DB, courseID int) error {

	_, err := db.Query(
		`DELETE FROM course WHERE courseID = ?`,
		courseID,
	)

	if err != nil {
		return err
	}

	return nil
}
