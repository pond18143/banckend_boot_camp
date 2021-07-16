package loger

import (
	"time"
)

const (
	InfoLevel = "info"
	DebugLevel = "debug"
	ErrorLevel = "error"
)

type logModel struct {
	SeriesId	string		`json:"series_id"`
	CompanyId	int64		`json:"company_id"`
	Endpoint	string		`json:"endpoint"`
	LogLevel	string		`json:"log_level"`
	Message		string		`json:"message"`
	DateTime	time.Time	`json:"date_time"`
}


type MessageResponse struct {
	Status             int    `json:"status"`
	MessageCode        string `json:"message_code"`
	MessageDescription string `json:"message_description"`
}
