package utils
import "regexp"

const (
	TestFormat = `[\w*]`

)

func checkEnglishNumber(text string) bool {
	words,_ := regexp.Compile(`^[a-zA-Z0-9]*$`)
	return words.MatchString(text)
}

func checkEnglish(text string) bool {
	words,_ := regexp.Compile(`^[a-zA-Z]*$`)
	return words.MatchString(text)
}

func checkNumber(text string) bool {
	words,_ := regexp.Compile(`^[0-9]*$`)
	return words.MatchString(text)
}

func checkEnglishNumberSpecialCharacters (text string) bool {
	words,_ := regexp.Compile(`^[a-zA-Z0-9!@#$%^&*()_]*$`)
	return words.MatchString(text)
}

// 1lower&upper , 1num  , 1symbol
func checkPassword (text string) bool {

	wordsNum,_ := regexp.Compile(`[0-9]{1}`)
	if wordsNum.MatchString(text) == false {
		return false
	}

	wordsLower,_ := regexp.Compile(`[a-z]{1}`)
	if wordsLower.MatchString(text) == false {
		return false
	}

	wordsUpper,_ := regexp.Compile(`[A-Z]{1}`)
	if wordsUpper.MatchString(text) == false {
		return false
	}

	wordsSymbol,_ := regexp.Compile(`[!@#$%^&*()_]{1}`)
	if wordsSymbol.MatchString(text) == false {
		return false
	}
	return true
}