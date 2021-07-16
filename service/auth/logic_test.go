package auth

import (
	"fmt"
	"testing"
)

const (
	password        = "P@ssw0rd"
	loginUuId		= "e05f8145-5f6a-4d2c-9440-a52fc98864ca"
)

func Test_resetPassword(t *testing.T) {
	fmt.Println("*****************Test_resetPassword**************")
	result :=  encryptPassword(compilePassword(password,loginUuId))
	outputHash := "0232296c53ef9d84b048a4a0ba0b07d551fd75cc2709ac1649f2247fd13ccb52"
	if result != outputHash {
		fmt.Println("Test_resetPassword test Fail")
		t.FailNow()
	}
}

//func Test_compilePassword (t *testing.T) {
//	fmt.Println("*****************Test_rePassword**************")
//	result :=  compilePassword(password,loginUuId)
//	outputCompilePassword := "P@ssw0rdd2c-9"
//	if result != outputCompilePassword {
//		fmt.Println("Test_rePassword test Fail")
//		t.FailNow()
//	}
//}


//func Test_hashPassword (t *testing.T) {
//	fmt.Println("*****************Test_hashPassword**************")
//	result :=  encryptPassword(passwordCompile)
//	output := "0232296c53ef9d84b048a4a0ba0b07d551fd75cc2709ac1649f2247fd13ccb52"
//	if result != output{
//		fmt.Println("Test_hashPassword test Fail")
//		t.FailNow()
//	}
//}
