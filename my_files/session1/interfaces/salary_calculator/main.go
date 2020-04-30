package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicPay int
	pf       int
}

type Contract struct {
	empId    int
	basicPay int
}

func main() {
	pemp1 := Permanent{
		empId:    1,
		basicPay: 5000,
		pf:       20,
	}
	pemp2 := Permanent{
		empId:    2,
		basicPay: 6000,
		pf:       30,
	}
	pemp3 := Permanent{
		empId:    3,
		basicPay: 8000,
		pf:       70,
	}
	cemp1 := Contract{
		empId:    3,
		basicPay: 3000,
	}
	employees := []SalaryCalculator{pemp1, pemp2, pemp3, cemp1}
	fmt.Println(employees)
	fmt.Printf("type: %T ->> value: %v\n", employees, employees)
	fmt.Printf("type: %T ->> value: %v\n", employees[0], employees[0])
	totalExpense(employees)

}

func (p Permanent) CalculateSalary() int {
	return p.basicPay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicPay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total HR expense per month $%d\n", expense)
}
