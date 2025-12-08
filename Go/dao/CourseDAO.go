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
