package models

type Student struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type Test struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Question struct {
	Id       string `json:"id"`
	Answer   string `json:"answer"`
	Question string `json:"question"`
	TestId   string `json:"test_id"`
}

type Enrollment struct {
	TestId    string `json:"test_id"`
	StudentId string `json:"student_id"`
}
