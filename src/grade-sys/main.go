package main

import (
	"grade-sys/models"
	"grade-sys/pkg/logging"
	"grade-sys/utils"
	"sync"
)

var wg sync.WaitGroup

func main(){
	// var wg sync.WaitGroup

	stus := models.GetAllStudents(2018,0)
	logging.Info(len(stus))

	wg.Add(len(stus))
	
	for _,stu := range stus{
		go HandleStu(stu)
	}

	wg.Wait()

	logging.Info("finish")
}

func HandleStu(stu models.GradeNotify){
	defer wg.Done()
	
	xn := stu.Xn
	xq := stu.Xq + 1
	stuid := stu.StuID
	hdjwpass := stu.HdjwPassword
	ptpass  := stu.PtPassword
	force := 1

	var resp models.Resp
	resp,err := utils.GetGrade(xn ,xq ,stuid ,hdjwpass ,ptpass ,force )

	if err != nil{
		logging.Info(stu.StuID,err)
		return
	}

	rowCount,err := resp.Check()
	if err != nil {
		logging.Info(stu.StuID,err)
		return
	}

	if stu.CourseCount < rowCount {
		logging.Info(stu.StuID," has new grade")
	}else{
		logging.Info(stu.StuID," don't haa new grade")
	}
}