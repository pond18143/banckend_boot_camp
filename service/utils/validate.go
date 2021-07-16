package utils

import (
	"net/http"
	"regexp"
)

func RegxTest(input string) (reg bool, msgReg messageResponse, err error) {
	cop, err := regexp.Compile(TestFormat)
	if err != nil {
		return
	}

	reg = cop.MatchString(input)
	//log.Infof("%s",reg)
	if reg != true {
		msgReg = messageResponse{
			Status:             http.StatusBadRequest,
			MessageCode:        "0000",
			MessageDescription: "missFormat"}
		return
	}

	return
}

func Validate(input string, miniMum int, maxiMum int, format string) (reg bool, msgReg messageResponse) {
	if input == "" {
		msgReg = messageResponse{
			Status:             http.StatusBadRequest,
			MessageCode:        "0000",
			MessageDescription: "not input"}
		return
	}

	if len(input) < miniMum || len(input) > maxiMum {
		msgReg = messageResponse{
			Status:             http.StatusBadRequest,
			MessageCode:        "0000",
			MessageDescription: "length miss"}
		return
	}

	switch format {
	case EnglishNumber:
		check := checkEnglishNumber(input)
		if check != true {
			msgReg = messageResponse{
				Status:             http.StatusBadRequest,
				MessageCode:        "0000",
				MessageDescription: "input must be string & number"}
			return
		}

	case English:
		check := checkEnglish(input)
		if check != true {
			msgReg = messageResponse{
				Status:             http.StatusBadRequest,
				MessageCode:        "0000",
				MessageDescription: "input must be string"}
			return
		}

	case Number:
		check := checkNumber(input)
		if check != true {
			msgReg = messageResponse{
				Status:             http.StatusBadRequest,
				MessageCode:        "0000",
				MessageDescription: "input must be number"}
			return
		}

	case EnglishNumberSpecialCharacters:
		check := checkEnglishNumberSpecialCharacters(input)
		if check != true {
			msgReg = messageResponse{
				Status:             http.StatusBadRequest,
				MessageCode:        "0000",
				MessageDescription: "input must be string & number & specialCharacter"}
			return
		}

	case Password:
		check := checkPassword(input)
		if check != true {
			msgReg = messageResponse{
				Status:             http.StatusBadRequest,
				MessageCode:        "0000",
				MessageDescription: "input must require stringUpperCase , stringLowerCase , number , specialCharacter At least 1"}
			return
		}
	}

	msgReg = messageResponse{
		MessageDescription: "format correct"}

	return true ,msgReg
}
