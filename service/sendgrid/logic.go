package sendgrid

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"time"

	"github.com/spf13/viper"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"go-api-game-boot-camp/service/loger"
)

var log *loger.Loger

func InitSendgrid(configPath string) {
	log = loger.NewLogController()

	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Errorf("Fatal Error Config File: ", err)
		panic("Fatal Error Config File")
	}
}

func SendMail() (string, error) {
	log.Info("Send Mail : SendGrid")
	// read message_queue
	mails, err := readMessageQueue()
	if err != nil {
		return "", err
	}

	if len(mails) != 0 {
		for _, m := range mails {
			from := mail.NewEmail(m.SenderName, m.SenderMail)
			subject := m.Subject
			to := mail.NewEmail(m.ReceiverName, m.ReceiverMail)
			plainTextContent := m.Message
			htmlContent := m.Message

			message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
			client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
			response, err := client.Send(message)
			if err != nil {
				log.Errorf("Error while sending mail: ", err)
				return "", err
			}

			log.Infof("StatusCode : %d", response.StatusCode)
			// log.Debug(response.Body)
			// log.Debug(response.Headers)
			if response.StatusCode != 202 {
				return "", errors.New("Problems occured while sending mail")
			}
			log.Debug("Send mail successfully")
		}
		return "Sent mail successfully", nil
	}

	return "", nil
}

func ParseTemplateDocument(action string, senderUID string, senderComp int64, docNum string, docType int64, sellerId int64, buyerId int64) (string, error) {
	log.Info("Parse HTML Template Document : SendGrid")

	// get mail detail (model for parse template)
	detail, err := getMailDetail(docNum, docType, sellerId, buyerId)
	if err != nil {
		return "", err
	}

	// format CreateDate
	createDateTime, _ := time.Parse(time.RFC3339, detail.CreateDate)
	detail.CreateDate = createDateTime.Format("02 January 2006 15:04")

	detail.CancelDate = time.Now().Format("02 January 2006 15:04")

	compInfo, err := getCompanyType(senderComp)
	if err != nil {
		log.Error("Error while getting company type")
		return "", err
	}

	// assign filePath, mailSubject
	var filePath string
	var mailSubject string
	if action == "submit" {
		mailSubject = fmt.Sprintf("Submitted Document %s", docNum)
		if compInfo.CompanyType == "Buyer" { // buyer
			filePath = viper.GetString("templatePath.submitBuyer")
		} else if compInfo.CompanyType == "Seller" { // seller
			filePath = viper.GetString("templatePath.submitSeller")
		}
	} else if action == "cancel" {
		mailSubject = fmt.Sprintf("Cancelled Document %s", docNum)
		if compInfo.CompanyType == "Buyer" { // buyer
			filePath = viper.GetString("templatePath.cancelBuyer")
		} else if compInfo.CompanyType == "Seller" { // seller
			filePath = viper.GetString("templatePath.cancelSeller")
		}
	}
	log.Debugf("filePath: ", filePath)
	log.Debugf("Subject: ", mailSubject)

	// parse HTML template
	t, err := template.ParseFiles(filePath)
	if err != nil {
		log.Error("Error while parsing template")
		return "", err
	}

	// store parsed html template in buf
	buf := new(bytes.Buffer)
	err = t.Execute(buf, detail)
	if err != nil {
		log.Error("Error while storing parsed template")
		return "", err
	}

	// get mail header (sender, receiver)
	mail, err := getMailHeader(senderUID, docNum, docType, sellerId, buyerId)

	// assign Subject, Message, IsSent
	mail.Subject = mailSubject
	mail.Message = buf.String()
	mail.IsSent = false

	// change sender's mail receiver's mail
	mail.SenderMail = "siriya013@gmail.com"
	mail.ReceiverMail = "pingsiriya@gmail.com"

	err = storeMessageQueue(mail)
	if err != nil {
		return "", err
	}

	return "Parse HTML template successfully", nil
}

func ParseTemplateWelcome(username string, password string, name string, email string) (string, error) {
	log.Info("Parse HTML Template Welcome : SendGrid")

	user := userInfo{
		Username: username,
		Password: password,
	}

	filePath := viper.GetString("templatePath.welcomeUser")

	// parse HTML template
	t, err := template.ParseFiles(filePath)
	if err != nil {
		log.Error("Error while parsing template")
		return "", err
	}

	// store parsed html template in buf
	buf := new(bytes.Buffer)
	err = t.Execute(buf, user)
	if err != nil {
		log.Error("Error while storing parsed template")
		return "", err
	}

	fmt.Println("ReceiverMail:", email)

	mail := inputMail{
		SenderName:   "Malar admin",
		SenderMail:   "siriya013@gmail.com",
		ReceiverName: name,
		ReceiverMail: "pingsiriya@gmail.com",
		Subject:      "Welcome to Malar",
		Message:      buf.String(),
		IsSent:       false,
	}

	err = storeMessageQueue(mail)
	if err != nil {
		return "", err
	}

	return "Parse HTML template successfully", nil
}
