package main

import (
	"fmt"
	o "hm2/office"
	"io/ioutil"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("employeeActions")
	if err != nil {
		fmt.Errorf("can't read file")
	}
	res := strings.Split(string(file), "\r\n")
	for i := 0; i < len(res); i++ {
		fmt.Println(" the worker started his journey")
		err := CheckEl(res[i])
		if err != nil {
			fmt.Println("the worker seems to have made a mistake")
		} else {
			fmt.Println("the worker went well all the way")
		}
	}
}
func CheckEl(s string) error {
	res := strings.Split(s, ",")
	if len(res) < 2 {
		fmt.Errorf("not enough data to get started")
	}
	if res[1] != "office" {
		fmt.Errorf("How does an employee want to start his journey if he does not start from the office")
	}
	employee, err := o.NewEmployeeFactory(res[0])
	if err != nil {
		fmt.Errorf("something is wrong with the worker creation")
	}
	return rec(employee, res[1:])

}

func rec(employee o.Employee, array []string) error {
	var err error
	if len(array) == 0 {
		return err
	}
	if employee == nil {
		fmt.Errorf("no employee")
	}
	loc, err := o.NewLocationFactory(array[0])
	if err != nil {
		fmt.Errorf("rec:%w-%s", err, array[0])
	}
	err = employee.MoveToLocation(loc)
	if err != nil {
		err = fmt.Errorf("rec:%w-%s", err, array[0])
		fmt.Println(err)
		return fmt.Errorf("rec:%w-%s", err, array[0])

	}
	return rec(employee, array[1:])
}
