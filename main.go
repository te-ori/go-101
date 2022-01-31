package main

import (
	"fmt"
	"scholl/graduate"
	"scholl/lesson"
	"scholl/student"
)

func main() {
	calculate()
}

func calculate() {
	var result []*student.Student

	result = studentList(result)

	graduateResult1 := graduate.GetGraduateList("Math I", result, graduate.GraduateCalculationByStandartAverage(50))

	printGraduateResult(graduateResult1)

	graduateResult2 := graduate.GetGraduateList("Math I", result, graduate.GraduateCalculationByWeightedAverage("Math I", result, 50))

	printGraduateResult(graduateResult2)

	customCalculator := func(scores []lesson.Score) (float32, bool) {
		return float32(scores[2]), scores[2] > 60
	}
	graduateResult3 := graduate.GetGraduateList("Math I", result, customCalculator)

	printGraduateResult(graduateResult3)

	avesomeCalc := func(scores []lesson.Score) (float32, bool) {
		return 100, true
	}
	graduateResult4 := graduate.GetGraduateList("Math I", result, avesomeCalc)

	printGraduateResult(graduateResult4)
}

func studentList(result []*student.Student) []*student.Student {
	result = append(result, generateStudent("Ayşe", generateMath(100, 100, 100)))
	result = append(result, generateStudent("Veli", generateMath(60, 80, 80)))
	result = append(result, generateStudent("Ali", generateMath(10, 45, 60)))
	result = append(result, generateStudent("Fatma", generateMath(100, 50, 50)))
	result = append(result, generateStudent("Ahmet", generateMath(86, 86, 86)))
	result = append(result, generateStudent("Nuri", generateMath(20, 10, 5)))
	result = append(result, generateStudent("Merve", generateMath(70, 75, 60)))
	result = append(result, generateStudent("Saniye", generateMath(100, 100, 100)))
	result = append(result, generateStudent("Murat", generateMath(100, 100, 100)))
	result = append(result, generateStudent("Berke", generateMath(100, 100, 100)))
	return result
}

func printGraduateResult(graduates []*graduate.GraduateResult) {
	for _, g := range graduates {
		fmt.Printf("%s %v %s %v\n", g.StudentName, g.Final, g.Final.Grade(), g.IsGraduate)
	}
	fmt.Println()
}

func generateStudent(name string, lesson *lesson.Lesson) *student.Student {
	student := student.NewStudent(name)
	student.InsertLesson(lesson)

	return &student
}

func generateMath(firstExam, seconfExam, thirdExam lesson.Score) *lesson.Lesson {
	math1 := lesson.NewLesson("Math I")
	math1.SetScoreOf(0, firstExam)
	math1.SetScoreOf(1, seconfExam)
	math1.SetScoreOf(2, thirdExam)

	return &math1
}

func misc() {
	math1 := lesson.NewLesson("Math I")
	fmt.Println("", math1.Notes())

	var e error
	_, e = math1.SetScoreOf(0, 120)
	fmt.Println("err: ", e)
	fmt.Println(math1.Notes())

	_, e = math1.SetScoreOf(0, 90)
	fmt.Println("err: ", e)
	fmt.Println(math1.Notes())

	s, e := math1.GetScoreOf(12)
	fmt.Println("err: ", e)
	fmt.Println(s)

	s, e = math1.GetScoreOf(0)
	fmt.Println("err: ", e)
	fmt.Println(s)
	fmt.Println(s.Grade())

	math2 := lesson.NewLesson("Math II")

	student1 := student.NewStudent("Zeki")

	_, e = student1.InsertLesson(&math1)
	fmt.Println("err: ", e)

	_, e = student1.InsertLesson(&math2)
	fmt.Println("err: ", e)

	_, e = student1.InsertLesson(&math2)
	fmt.Println("err: ", e)

	biology := lesson.NewLesson("Biology")
	student1.InsertLesson(&biology)

	physic := lesson.NewLesson("Physic")
	student1.InsertLesson(&physic)

	chemistry := lesson.NewLesson("Chemistery")
	student1.InsertLesson(&chemistry)

	lessonsOfStudents := student1.ListOfLessons()
	fmt.Println("Lessons: ", lessonsOfStudents)

	scoresOfEconomy := student1.ScoresOfLesson("Economy")
	fmt.Println("Scores of Economy: ", scoresOfEconomy)

	scoresOfMath1 := student1.ScoresOfLesson("Math I")
	fmt.Println("Scores of Math I: ", scoresOfMath1)

}
