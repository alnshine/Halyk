package main

import (
	"fmt"
	"hal10/flatjson"
)

type OutStruct struct {
	Name                                string
	Age                                 int64
	EducationDegree                     string
	educationAverageGrade               float32 // Приватное поле должно остаться пустым
	EducationUniversity                 string
	EducationFacultyName                string
	EducationFacultyDepartment          string
	EducationFacultyAdviserFirstName    string
	EducationFacultyAdviserLastName     string
	EducationFacultyAdviserDegree       string
	EducationFacultyAdviserArticleCount int32
}

func main() {
	data := `
	{
		"Name": "Alibek",
		"Age": 21,
		"Education": {
			"Degree": "bachelor",
			"AverageGrade": 4.4,
			"University": "ENU",
			"Faculty": {
				"Name": "Mechmath",
				"Department": "Mathematical and computer modeling",
				"Adviser": {
					"FirstName": "Ivanov",
					"LastName": "Ivan",
					"Degree": "PhD",
					"ArticleCount": 30
				}
			}
		}
	}
	`

	out := &OutStruct{}
	err := flatjson.Unmarshal([]byte(data), out)
	if err != nil {
		fmt.Printf("Unmarshal error: %v\n", err)
		return
	}
	fmt.Println(out)
}

// type OutStruct struct {
// 	FirstField  string
// 	SecondField bool
// 	NestedName  string
// }

// func main() {
// 	data := `{
// 		"FirstField": "hi world",
// 		"SecondField": true,
// 		"Nested": {
// 			"Name": "123"
// 		}
// 	}`
// 	out := &OutStruct{}
// 	err := flatjson.Unmarshal([]byte(data), out)
// 	if err != nil {
// 		fmt.Printf("Unmarshal error: %v\n", err)
// 		return
// 	}
// 	fmt.Println(out)
// }
