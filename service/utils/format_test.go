package utils
import (
	"testing"
	"fmt"
)

func TestGoodCaseNumber (t *testing.T){
	fmt.Println("TestGoodCaseNumber")
	result := checkNumber("12356")
	if result != true{
		fmt.Println("TestGoodCaseNumber test fail")
		t.FailNow()
	}
}

func TestBadCaseNumberHasEng (t *testing.T){
	fmt.Println("TestBadCaseNumberHasEng")
	result := checkNumber("12a356")
	if result != false{
		fmt.Println("TestBadCaseNumberHasEng test fail")
		t.FailNow()
	}
}

func TestBadCaseNumberHasSC (t *testing.T){
	fmt.Println("TestBadCaseNumberHasSC")
	result := checkNumber("123@56")
	if result != false{
		fmt.Println("TestBadCaseNumberHasSC test fail")
		t.FailNow()
	}
}

func TestBadCaseNumberHasSpaceBar (t *testing.T){
	fmt.Println("TestBadCaseNumberHasSpaceBar")
	result := checkNumber("123 56")
	if result != false{
		fmt.Println("TestBadCaseNumberHasSpaceBar test fail")
		t.FailNow()
	}
}

func TestGoodCaseCheckEnglish (t *testing.T){
	fmt.Println("TestGoodCasecheckEnglish")
	result := checkEnglish("abcABC")
	if result != true{
		fmt.Println("TestGoodCasecheckEnglish test fail")
		t.FailNow()
	}
}

func TestBadCaseCheckEnglishHasNumber (t *testing.T){
	fmt.Println("TestBadCasecheckEnglishHasNumber")
	result := checkEnglish("abcABC1")
	if result != false{
		fmt.Println("TestBadCasecheckEnglishHasNumber test fail")
		t.FailNow()
	}
}

func TestBadCaseCheckEnglishHasSC (t *testing.T){
	fmt.Println("TestBadCasecheckEnglishHasSC")
	result := checkEnglish("@abcABC")
	if result != false{
		fmt.Println("TestBadCasecheckEnglishHasSC test fail")
		t.FailNow()
	}
}

func TestBadCaseCheckEnglishHasSpaceBar (t *testing.T){
	fmt.Println("TestBadCasecheckEnglishHasSpaceBar")
	result := checkEnglish("a bcABC")
	if result != false{
		fmt.Println("TestBadCasecheckEnglishHasSpaceBar test fail")
		t.FailNow()
	}
}

func TestGoodCaseCheckEnglishNumber (t *testing.T){
	fmt.Println("TestGoodCaseCheckEnglishNumber")
	result := checkEnglishNumber("aaa123")
	if result != true{
		fmt.Println("TestGoodCaseCheckEnglishNumber test fail")
		t.FailNow()
	}
}

func TestGoodCaseCheckEnglishNumberHasEngOnly (t *testing.T){
	fmt.Println("TestGoodCaseCheckEnglishNumberHasEngOnly")
	result := checkEnglishNumber("aaAa")
	if result != true{
		fmt.Println("TestGoodCaseCheckEnglishNumberHasEngOnly test fail")
		t.FailNow()
	}
}

func TestGoodCaseCheckEnglishNumberHasnumberOnly (t *testing.T){
	fmt.Println("TestGoodCaseCheckEnglishNumberHasnumberOnly")
	result := checkEnglishNumber("1236")
	if result != true{
		fmt.Println("TestGoodCaseCheckEnglishNumberHasnumberOnly test fail")
		t.FailNow()
	}
}

func TestBadCaseCheckEnglishNumberHasSC (t *testing.T){
	fmt.Println("TestBadCaseCheckEnglishNumberHasSC")
	result := checkEnglishNumber("1a23b6%")
	if result != false{
		fmt.Println("TestBadCaseCheckEnglishNumberHasSC test fail")
		t.FailNow()
	}
}

func TestBadCaseCheckEnglishNumberHasSpaceBar (t *testing.T){
	fmt.Println("TestBadCaseCheckEnglishNumberHasSpaceBar")
	result := checkEnglishNumber("1a23b6%")
	if result != false{
		fmt.Println("TestBadCaseCheckEnglishNumberHasSpaceBar test fail")
		t.FailNow()
	}
}

func TestGoodCaseCheckEnglishNumberSpecialCharacters (t *testing.T){
	fmt.Println("TestGoodCaseCheckEnglishNumberSpecialCharacters")
	result := checkEnglishNumberSpecialCharacters("1a23B6%")
	if result != true{
		fmt.Println("TestGoodCaseCheckEnglishNumberSpecialCharacters test fail")
		t.FailNow()
	}
}

func TestGoodCaseCheckEnglishNumberSpecialCharactersHasEngOnly (t *testing.T){
	fmt.Println("TestGoodCaseCheckEnglishNumberSpecialCharactersHasEngOnly")
	result := checkEnglishNumberSpecialCharacters("abDF")
	if result != true{
		fmt.Println("TestGoodCaseCheckEnglishNumberSpecialCharactersHasEngOnly test fail")
		t.FailNow()
	}
}

func TestGoodCaseCheckEnglishNumberSpecialCharactersHasNumberOnly (t *testing.T){
	fmt.Println("TestGoodCaseCheckEnglishNumberSpecialCharactersHasNumberOnly")
	result := checkEnglishNumberSpecialCharacters("123")
	if result != true{
		fmt.Println("TestGoodCaseCheckEnglishNumberSpecialCharactersHasNumberOnly test fail")
		t.FailNow()
	}
}

func TestGoodCaseCheckEnglishNumberSpecialCharactersHasSCOnly (t *testing.T){
	fmt.Println("TestGoodCaseCheckEnglishNumberSpecialCharactersHasSCOnly")
	result := checkEnglishNumberSpecialCharacters("!@#$%^&*()_")
	if result != true{
		fmt.Println("TestGoodCaseCheckEnglishNumberSpecialCharactersHasSCOnly test fail")
		t.FailNow()
	}
}

func TestGoodCaseCheckEnglishNumberSpecialCharactersHasSpaceBar (t *testing.T){
	fmt.Println("TestGoodCaseCheckEnglishNumberSpecialCharactersHasSCOnly")
	result := checkEnglishNumberSpecialCharacters("asd123A!@#$% ^&*()_")
	if result != false{
		fmt.Println("TestGoodCaseCheckEnglishNumberSpecialCharactersHasSCOnly test fail")
		t.FailNow()
	}
}