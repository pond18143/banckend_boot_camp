package sendgrid

import "time"

type outputModel struct {
	Message  string    `json:"message"`
	DateTime time.Time `json:"date_time"`
}

type inputMail struct {
	SenderName   string `json:"sender_name"`
	SenderMail   string `json:"sender_mail"`
	ReceiverName string `json:"receiver_name"`
	ReceiverMail string `json:"receiver_mail"`
	Subject      string `json:"subject"`
	Message      string `json:"message"`
	IsSent       bool   `json:"is_sent"`
}

type mailDetail struct {
	ReceiverCompany string  `json:"receiver_company"`
	SenderCompany   string  `json:"sender_company"`
	DocumentNumber  string  `json:"document_number"`
	CreateDate      string  `json:"create_date"`
	CancelDate      string  `json:"cancel_date"`
	GrandTotal      float64 `json:"grand_total"`
	Currency        string  `json:"currency"`
}

type docType struct {
	DocumentType int64
}

type compType struct {
	CompanyName string `json:"company_name"`
	CompanyType string `json:"company_type"`
}

type userInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
