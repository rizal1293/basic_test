package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"github.com/jedib0t/go-pretty/table"
)

type students struct {
	StudentId int
	Name string
	finalScore float64
	Grade string

}

func main() {
	var count int
	
	var listStudent []students

	var studentID int
	var studentName string
	var midTermScore float64
	var semsterScore float64
	var attendance int

	passingStudents := 0
	failStudents := 0

	fmt.Print("Input the number of students : ")
	scCount := bufio.NewScanner(os.Stdin)
	if scCount.Scan() {
		res, err := strconv.Atoi(scCount.Text())
		if err != nil {
			println("input must be number")
			return
		} else {
			count = res
		}
	}

	print("\n")
	
	for i := 1; i <= count; i++ {

		fmt.Print("Student ID : ")
		scStudentID := bufio.NewScanner(os.Stdin)
		if scStudentID.Scan() {
			resCount, err := strconv.Atoi(scStudentID.Text())
			if err != nil {
				println("Student ID must be number")
				return
			} else {
				studentID = resCount
			}
		}

		fmt.Print("Name : ")
		scName := bufio.NewScanner(os.Stdin)
		if scName.Scan() {
			studentName = scName.Text()
		}

		fmt.Print("Mid Term Test Score : ")
		scMidTerm := bufio.NewScanner(os.Stdin)
		if scMidTerm.Scan() {
			resTermMid, err := strconv.ParseFloat(scMidTerm.Text(), 64)
			if err != nil {
					println("Mid Term Test Score be number or decimal")
					return
			} else {
				if resTermMid < 0 || resTermMid > 100 {
					println("Mid Term Test Score min 0 and max 100")
					return
				}
				midTermScore = resTermMid
			}
		}

		fmt.Print("Semester Test Score : ")
		scSemester := bufio.NewScanner(os.Stdin)
		if scSemester.Scan() {
			resSemester, err := strconv.ParseFloat(scSemester.Text(), 64)
			if err != nil {
				println("Semester Test Score be number or decimal")
				return
			} else {
				if resSemester < 0 || resSemester > 100 {
					println("Semester Test Score min 0 and max 100")
					return
				}
				semsterScore = resSemester
			}
		}

		fmt.Print("Attendance : ")
		scAttendance := bufio.NewScanner(os.Stdin)
		if scAttendance.Scan() {
			resAttedance, err := strconv.Atoi(scAttendance.Text())
			if err != nil {
				println("Attendance must be number")
				return
			} else {
				if resAttedance < 0 || resAttedance > 100 {
					println("Attendance min 0 and max 100")
					return
				}
				attendance = resAttedance
			}
		}
	
		print("\n")
		print("=============")
		print("\n")

		finalScore := (0.2 * float64(attendance)) + (0.4 * midTermScore) + (0.4 * semsterScore)
		grade := getGarde(finalScore)
	
		listStudent = append(listStudent, students{
			StudentId: studentID,
			Name: studentName,
			finalScore: finalScore,
			Grade: grade,
		})

		if grade == "E" || grade == "D" {
			failStudents++
		} else {
			passingStudents++
		}
	}

	showReport(listStudent)
	
	print("\n")
	
	fmt.Printf("Number of Students \t\t: %v \n", len(listStudent))
	fmt.Printf("Number of Passing Students \t: %v\n", passingStudents)
	fmt.Printf("Number of Failed Students \t: %v\n", failStudents)

	print("\n")
}

func showReport(students []students)  {
	t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"No", "Student ID", "Name", "Final Score", "Grade"})
	
	for i, row := range students {
		i++
		t.AppendRow([]interface{}{i, row.StudentId, row.Name, row.finalScore, row.Grade})
	}
	t.Render()
}

func getGarde(finalScore float64) (grade string) {
	if finalScore >= float64(85) && finalScore <= float64(100) {
		grade = "A"
	} else if finalScore >= float64(76) && finalScore <= float64(84) {
		grade = "B"
	} else if finalScore >= float64(61) && finalScore <= float64(75) {
		grade = "C"
	} else if finalScore >= float64(46) && finalScore <= float64(60) {
		grade = "D"
	} else if (finalScore >= float64(0) && finalScore <= float64(45)){
		grade = "E"
	} else {
		grade = "final score outside the specified range"
	}
	return
}