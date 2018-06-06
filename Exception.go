package main

import (
	"fmt"
)

// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// 实现 `error` 接口
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

func Devide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		divideError := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = divideError.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

func main() {
	if result, errorMsg := Devide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	if _, errorMsg := Devide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}
