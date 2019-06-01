package models

import (
	"grade-sys/pkg/logging"
	"database/sql"
)

type Student struct {
	StuID      string
	PtPassword string
	HdjwPassword string
}

type GradeNotify struct {
	Student
	Xn    	   int
	Xq   	   int
	CourseCount   int
	Email      string
	IsFinish   int
	Phone      sql.NullString
	Type       sql.NullInt64
}


func GetAllStuId(xn,xq int) []GradeNotify {
	sql := "SELECT stuID,xn,xq,courseCount,email,isFinish,phone,type FROM grade_notify_tmp where xn=? AND xq=? AND isFinish=0 limit 8"
	rows, err := Db.Query(sql,xn,xq)

	checkErr(err)

	stus := []GradeNotify{}

	stu  := GradeNotify{}


	for rows.Next() {
		err = rows.Scan(&stu.StuID, &stu.Xn,&stu.Xq,&stu.CourseCount,&stu.Email,&stu.IsFinish,&stu.Phone,&stu.Type)
		checkErr(err)
		logging.Info(stu)
		stus = append(stus,stu)
	}

	return stus
}

func GetAllStudents(xn,xq int) []GradeNotify {
	sql := "SELECT s.stuID,s.stuPASS,s.hdjwPass,g.xn,g.xq,g.courseCount,g.email,g.isFinish,g.phone,g.type FROM mini_bind s JOIN grade_notify g ON s.stuID=g.stuID AND s.mode=1 AND g.xn=? AND g.xq=? AND isFinish=0 LIMIT 800"
	rows, err := Db.Query(sql,xn,xq)

	checkErr(err)

	stus := []GradeNotify{}

	stu  := GradeNotify{}


	for rows.Next() {
		err = rows.Scan(&stu.StuID,&stu.PtPassword,&stu.HdjwPassword, &stu.Xn,&stu.Xq,&stu.CourseCount,&stu.Email,&stu.IsFinish,&stu.Phone,&stu.Type)
		checkErr(err)
		// logging.Info(stu)
		stus = append(stus,stu)
	}

	return stus
}

