package main

import (
	"fmt"
)

type Account struct {
	name    string
	balance float64
}

type Depositor interface {
	Deposit(amount float64)
	String() string
}

type Withdrawer interface {
	Withdraw(amount float64)
}

func (a Account) String() string {
	return fmt.Sprintf("Account %s Balance is %f", a.name, a.balance)
}

//Pointer Receivers..
func (a *Account) Deposit(amount float64) {
	a.balance += amount
}

func (a *Account) Withdraw(amount float64) {
	a.balance -= amount
}

func main() {
	a := &Account{"aaaaaaaa", 1000} ////Pointer Receivers, therefore &Acc...
	test(a)
}

func test(with_draw Withdrawer) {

	// Type assertion pull the methods from the concrete type...
	// respond_to && send
	if wd, ok := with_draw.(Withdrawer); ok {
		wd.Withdraw(20000)
	}

	if depo, ok := with_draw.(Depositor); ok {
		depo.Deposit(111111111111)
		fmt.Println(depo.String())
	}

}

//func test(de Depositor) {
//	de.Deposit(10000)

//	type stringer interface {
//		String() string
//	}

//	if with_draw, ok := de.(Withdrawer); ok {
//		fmt.Println("Withdawer............")
//		if v, ok := de.(stringer); ok {
//			fmt.Println(v.String())
//		}
//		with_draw.Withdraw(200)
//		if v, ok := de.(stringer); ok {
//			fmt.Println(v.String())
//		}

//		fmt.Println("Withdawer............")
//	}

//	if v, ok := de.(stringer); ok {
//		fmt.Println(v.String())
//	}
//}
