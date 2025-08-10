package main

import "fmt"

// declaring a typr struct
type employee struct {
	name     string
	deadline map[string]float64
	salary   int
}

func makeEmployee(name string) employee {

	b := employee{
		name:     name,
		deadline: map[string]float64{"play": 3},
		salary:   45,
	}

	return b
}

// This is a receiver function, this is like a method in classes
func (t employee) task() string {
	task := "Employee Tasks: \n"
	//%-25v allocates the whole space to the variable hence a table can be drawn
	// negative value allocated space on right and positive allocates on the left
	task += fmt.Sprintf("%25v %25v \n", "TASK", "HOURS ALLOCATED")
	var total float64 = 0

	for k, v := range t.deadline {
		task += fmt.Sprintf("%25v %25v \n", k, v)
		total += v
	}

	task += fmt.Sprintf("%25v %25v", "Total hours alocated", total)

	return task
}

func (e *employee) incSalary(s int) {
	e.salary = s
}

func (e employee) addTask(n string, h float64) {
	e.deadline[n] = h
}
