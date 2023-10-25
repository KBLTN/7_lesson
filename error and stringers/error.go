package main

import "fmt"

type appError struct {
	Err    error
	Custom string
	Field  int
}

//type sppError interface {
//	Error() string
//	Unwrap() error
//}

func (e *appError) Error() string {
	return fmt.Sprintf("%s error", e.Custom)
}

func (e *appError) Unwrap() error {
	return e.Err

}

func main() {
	err := method1()
	if err != nil {
		fmt.Println(err.Unwrap())
		return
	}
	fmt.Println("success")
}

func method1() *appError {
	return method2()

	//return &AppError{
	//	Err:    fmt.Errorf("my error"),
	//	Custom: "value here",
	//	Field:  89,
}

func method2() *appError {
	return method3()
}

func method3() *appError {
	return &appError{
		Err:    fmt.Errorf("internal error"),
		Custom: "something goes wrong.",
	}
}
