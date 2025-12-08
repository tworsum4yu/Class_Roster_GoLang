package controller

import (
	"database/sql"
	"fmt"
	"myapp/dao"
	"myapp/dto"
	"myapp/service"
	"myapp/ui"

	_ "github.com/go-sql-driver/mysql"
)

func Run(db *sql.DB) {

	for {

		menuOption := ui.GetMenuOption()

		switch menuOption {
		case 1:
			listRecords(db)
		case 2:
			getIndividualRecord(db)
		case 3:
			createRecords(db)
		case 4:
			updateRecords(db)
		case 5:
			removeRecords(db)
		case 6:
			return
		default:
			ui.PrintInvalidInput()
		}

	}

}

func listRecords(db *sql.DB) {

	option := ui.GetListOption()

	switch option {
	case 1:
		students, err := dao.GetStudentList(db)

		if err != nil {
			fmt.Println(err)
			return
		}

		for _, student := range students {
			ui.PrintStudent(student)
		}

	case 2:
		teachers, err := dao.GetTeacherList(db)

		if err != nil {
			fmt.Println(err)
			return
		}

		for _, teacher := range teachers {
			ui.PrintTeacher(teacher)
		}
	case 3:
		courses, err := dao.GetCourseList(db)

		if err != nil {
			fmt.Println(err)
			return
		}

		for _, course := range courses {
			ui.PrintCourse(course)
		}
	default:
		ui.PrintInvalidInput()
	}

	// fmt.Println("List")
}

func getIndividualRecord(db *sql.DB) {
	option := ui.GetListOption()

	switch option {
	case 1:

		id := ui.GetID("Enter the id to find: ")

		student, err := dao.GetStudent(db, id)

		if err != nil {
			fmt.Println(err)
			return
		}

		ui.PrintStudent(*student)

	case 2:

		id := ui.GetID("Enter the id to find: ")

		teacher, err := dao.GetTeacher(db, id)

		if err != nil {
			fmt.Println(err)
			return
		}

		ui.PrintTeacher(*teacher)

	case 3:

		id := ui.GetID("Enter the id to find: ")

		course, err := dao.GetCourse(db, id)

		if err != nil {
			fmt.Println(err)
			return
		}

		ui.PrintCourse(*course)

	default:
		ui.PrintInvalidInput()
	}

}

func createRecords(db *sql.DB) {
	option := ui.GetListOption()

	switch option {
	case 1:

		student := &dto.Student{}

		student.FirstName, student.LastName = ui.GetStudentInfo()

		if err := dao.CreateStudentRecord(db, student); err != nil {
			fmt.Println(err)
			return
		}

	case 2:
		teacher := &dto.Teacher{}

		teacher.FirstName, teacher.LastName, teacher.Title, teacher.Office, teacher.Department = ui.GetTeacherInfo()

		if err := dao.CreateTeacherRecord(db, teacher); err != nil {
			fmt.Println(err)
			return
		}

	case 3:
		course := &dto.Course{}

		course.Name = ui.GetCourseInfo()

		if err := dao.CreateCourseRecord(db, course); err != nil {
			fmt.Println(err)
			return
		}

	default:
		ui.PrintInvalidInput()
	}
}

func updateRecords(db *sql.DB) {
	option := ui.GetListOption()

	switch option {
	case 1:

		id := ui.GetID("Enter the id to update: ")

		newStudent := &dto.Student{}
		newStudent.ID = id

		currentStudent, err := dao.GetStudent(db, id)

		if err != nil {
			fmt.Println(err)
			return
		}

		newStudent.FirstName, newStudent.LastName = ui.GetStudentInfo()

		if err := dao.UpdateStudentRecord(db, service.UpdateStudentRecordFix(currentStudent, newStudent)); err != nil {
			fmt.Println(err)
			return
		}

	case 2:

		id := ui.GetID("Enter the id to update: ")

		newTeacher := &dto.Teacher{}
		newTeacher.ID = id

		currentTeacher, err := dao.GetTeacher(db, id)

		if err != nil {
			fmt.Println(err)
			return
		}

		newTeacher.FirstName, newTeacher.LastName, newTeacher.Title, newTeacher.Office, newTeacher.Department = ui.GetTeacherInfo()

		if err := dao.UpdateTeacherRecord(db, service.UpdateTeacherRecordFix(currentTeacher, newTeacher)); err != nil {
			fmt.Println(err)
			return
		}

	case 3:

		id := ui.GetID("Enter the id to update: ")

		newCourse := &dto.Course{}
		newCourse.ID = id

		currentCourse, err := dao.GetCourse(db, id)

		if err != nil {
			fmt.Println(err)
			return
		}

		newCourse.Name = ui.GetCourseInfo()

		if err := dao.UpdateCourseRecord(db, service.UpdatCourseRecordFix(currentCourse, newCourse)); err != nil {
			fmt.Println(err)
			return
		}
	default:
		ui.PrintInvalidInput()
	}
}

func removeRecords(db *sql.DB) {
	option := ui.GetListOption()

	switch option {
	case 1:

		id := ui.GetID("Enter the id to remove: ")

		if err := dao.DeleteStudentRecord(db, id); err != nil {
			fmt.Println(err)
			return
		}

	case 2:

		id := ui.GetID("Enter the id to remove: ")

		if err := dao.DeleteTeacherRecord(db, id); err != nil {
			fmt.Println(err)
			return
		}

	case 3:

		id := ui.GetID("Enter the id to remove: ")

		if err := dao.DeleteCourseRecord(db, id); err != nil {
			fmt.Println(err)
			return
		}

	default:
		ui.PrintInvalidInput()
	}
}
