package utils
import (
	"testing"
	"fmt"
)

func TestGoodCaseValidateEnglish (t *testing.T){
	fmt.Println("TestGoodCaseValidateEnglish")
	boolean,result := Validate("hgjkAS",1,8,"English")
	expected := "ok"
	if boolean != true && result.MessageDescription != expected{
		fmt.Println("TestGoodCaseValidateEnglish test fail")
		fmt.Println("MessageDescription : ",result.MessageDescription)
		t.FailNow()
	}
}
func TestGoodCaseValidateEnglishHasNoText (t *testing.T){
	fmt.Println("TestGoodCaseValidateEnglishHasNoText")
	boolean,result := Validate("",1,8,"English")
	expected := "not input"
	if boolean != false && result.MessageDescription != expected{
		fmt.Println("TestGoodCaseValidateEnglishHasNoText test fail")
		fmt.Println("MessageDescription : ",result.MessageDescription)
		t.FailNow()
	}
}
func TestGoodCaseValidateEnglishMoreLen (t *testing.T){
	fmt.Println("TestGoodCaseValidateEnglishMoreLen")
	boolean,result := Validate("asdsadasAAA",1,8,"English")
	expected := "length miss"
	if boolean != false && result.MessageDescription != expected{
		fmt.Println("TestGoodCaseValidateEnglishMoreLen test fail")
		fmt.Println("MessageDescription : ",result.MessageDescription)
		t.FailNow()
	}
}
func TestGoodCaseValidateEnglishLessLen (t *testing.T){
	fmt.Println("TestGoodCaseValidateEnglishLessLen")
	boolean,result := Validate("asdsadasAAA",1,8,"English")
	expected := "length miss"
	if boolean != false && result.MessageDescription != expected{
		fmt.Println("TestGoodCaseValidateEnglishLessLen test fail")
		fmt.Println("MessageDescription : ",result.MessageDescription)
		t.FailNow()
	}
}
func TestGoodCaseValidateEnglishHasNumber (t *testing.T){
	fmt.Println("TestGoodCaseValidateEnglishLessLen")
	boolean,result := Validate("asds1AA",1,8,"English")
	expected := "input must be string"
	if boolean != false && result.MessageDescription != expected{
		fmt.Println("TestGoodCaseValidateEnglishLessLen test fail")
		fmt.Println("MessageDescription : ",result.MessageDescription)
		t.FailNow()
	}
}
func TestGoodCaseValidateEnglishHasSC (t *testing.T){
	fmt.Println("TestGoodCaseValidateEnglishLessLen")
	boolean,result := Validate("asds@AA",1,8,"English")
	expected := "input must be string"
	if boolean != false && result.MessageDescription != expected{
		fmt.Println("TestGoodCaseValidateEnglishLessLen test fail")
		fmt.Println("MessageDescription : ",result.MessageDescription)
		t.FailNow()
	}
}