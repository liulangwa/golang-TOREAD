package main

import (
	"errors"
	"fmt"
)

//自定义err
type RunError struct {
	status string
	Err    error
}

func (runErr RunError) Error() string {
	return runErr.status + runErr.Err.Error()
}

func throw() error {
	return errors.New("抛出异常\n")
	// return fmt.Errorf("%s", "errormessage")
}

func throwInfo() error {

	return RunError{status: "启动中", Err: errors.New("遇到严重错误")}
}

func main() {

	err := throw()
	if err != nil {
		fmt.Printf(err.Error())
	}

	err = throwInfo()
	if err != nil {
		fmt.Printf(err.Error())
	}

}
