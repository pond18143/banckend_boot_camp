package sendgrid

import (
	"go-api-game-boot-camp/app"
)

func getDocType(docNum string) (dt docType, err error) {
	if err = app.GameBootCamp.DB.
		Table("malar.dbo.document_header").
		Select("document_type").
		Where("document_number = ?", docNum).
		Find(&dt).Error; err != nil {
		return
	}

	return
}

func getCompanyType(senderComp int64) (ct compType, err error) {
	if err = app.GameBootCamp.DB.
		Table("malar.dbo.company").
		Select("company.company_name, config.description AS company_type").
		Joins("INNER JOIN malar.dbo.config ON company.id = ? AND company.company_type = config.status_code", senderComp).
		Find(&ct).Error; err != nil {
		return
	}

	return
}

func getMailHeader(senderUID string, docNum string, docType int64, sellerId int64, buyerId int64) (header inputMail, err error) {
	if docType == 110 { // PO
		if err = app.GameBootCamp.DB.
			Table("malar.dbo.login user1, malar.dbo.company comp1, malar.dbo.login user2, malar.dbo.company comp2, malar.dbo.document_header").
			Select("comp1.company_name AS sender_name, user1.email AS sender_mail, comp2.company_name AS receiver_name, user2.email AS receiver_mail").
			Where("user1.login_uuid = ? AND comp1.id = user1.company_id AND (document_header.document_number = ? AND document_header.document_type = ? AND document_header.seller_company_id = ? AND document_header.buyer_company_id = ?) AND comp2.id = document_header.seller_company_id AND comp2.id = user2.company_id", senderUID, docNum, docType, sellerId, buyerId).
			Order("sender_name").Limit(1).Find(&header).Error; err != nil {
			return
		}
	} else if docType == 210 { // INV
		if err = app.GameBootCamp.DB.
			Table("malar.dbo.login user1, malar.dbo.company comp1, malar.dbo.login user2, malar.dbo.company comp2, malar.dbo.document_header").
			Select("comp1.company_name AS sender_name, user1.email AS sender_mail, comp2.company_name AS receiver_name, user2.email AS receiver_mail").
			Where("user1.login_uuid = ? AND comp1.id = user1.company_id AND (document_header.document_number = ? AND document_header.document_type = ? AND document_header.seller_company_id = ? AND document_header.buyer_company_id = ?) AND comp2.id = document_header.buyer_company_id AND comp2.id = user2.company_id", senderUID, docNum, docType, sellerId, buyerId).
			Order("sender_name").Limit(1).Find(&header).Error; err != nil {
			return
		}
	}

	return
}

func getMailDetail(docNum string, docType int64, sellerId int64, buyerId int64) (detail mailDetail, err error) {
	if docType == 110 { // PO
		if err = app.GameBootCamp.DB.
			Table("malar.dbo.document_header, malar.dbo.company comp1, malar.dbo.company comp2").
			Select("comp1.company_name AS receiver_company, comp2.company_name AS sender_company, document_header.document_number, document_header.create_date, document_header.grand_total, document_header.currency").
			Where("(document_header.document_number = ? AND document_header.document_type = ? AND document_header.seller_company_id = ? AND document_header.buyer_company_id = ?) AND comp1.id = ? AND comp2.id = ?", docNum, docType, sellerId, buyerId, sellerId, buyerId).
			Find(&detail).Error; err != nil {
			return
		}
	} else if docType == 210 { // INV
		if err = app.GameBootCamp.DB.
			Table("malar.dbo.document_header, malar.dbo.company comp1, malar.dbo.company comp2").
			Select("comp1.company_name AS receiver_company, comp2.company_name AS sender_company, document_header.document_number, document_header.create_date, document_header.grand_total, document_header.currency").
			Where("(document_header.document_number = ? AND document_header.document_type = ? AND document_header.seller_company_id = ? AND document_header.buyer_company_id = ?) AND comp1.id = ? AND comp2.id = ?", docNum, docType, sellerId, buyerId, buyerId, sellerId).
			Find(&detail).Error; err != nil {
			return
		}
	}

	return
}

func storeMessageQueue(mail inputMail) (err error) {
	// insert mail
	if err = app.GameBootCamp.DB.
		Table("malar.dbo.message_queue").
		Create(&mail).Error; err != nil {
		return
	}

	return
}

func readMessageQueue() (mails []inputMail, err error) {
	// true = 1, false = 0
	// select mail
	if err = app.GameBootCamp.DB.
		Table("malar.dbo.message_queue").
		Where("is_sent = ?", 0).
		Find(&mails).Error; err != nil {
		return
	}

	// update mail sent status
	if err = app.GameBootCamp.DB.
		Table("malar.dbo.message_queue").
		Where("is_sent = ?", 0).
		Update("is_sent", 1).Error; err != nil {
		return
	}

	return mails, nil
}
