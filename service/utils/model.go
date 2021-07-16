package utils

const (
	EnglishNumber = `EnglishNumber`

	English = `English`

	Number = `Number`

	EnglishNumberSpecialCharacters = `EnglishNumberSpecialCharacters`

	Password = `checkPassword`
)


type messageResponse struct {
	Status             int    `json:"status"`
	MessageCode        string `json:"message_code"`
	MessageDescription string `json:"message_description"`
}