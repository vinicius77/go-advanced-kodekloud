package main

import "fmt"

// Declare the Expense struct here
// your code goes here

type Expense struct {
	Name  string
	Value float64
	Date  string
}

func (expense Expense) getName() string {
	return expense.Name
}

// Implement the Total method to calculate the total amount spent
// your code goes here
func Total(expenses []Expense) float64 {
	var res float64

	for _, expense := range expenses {
		res += expense.Value
	}

	return res

}

// Implement the getName method on the Expense struct here
// your code goes here

func main() {
	expenses := []Expense{
		{"Grocery", 50.0, "2022-01-01"},
		{"Gas", 30.0, "2022-01-02"},
		{"Restaurant", 40.0, "2022-01-03"},
	}

	fmt.Println(Total(expenses))
	fmt.Println(expenses[0].getName())
}
