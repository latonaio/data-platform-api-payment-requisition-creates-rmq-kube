package dpfm_api_output_formatter

import (
	"data-platform-api-payment-requisition-creates-rmq-kube/sub_func_complementer"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) *HeaderCreates {
	data := subfuncSDC.Message.Header

	headerCreate := &HeaderCreates{
		PayerPaymentRequisitionID:       data.PayerPaymentRequisitionID,
		Payer:                           data.Payer,
		PayerPaymentDate:                data.PayerPaymentDate,
		PaymentReqnStatus:               data.PaymentReqnStatus,
		AcceptanceNoByFinInst:           data.AcceptanceNoByFinInst,
		PaytReqnAmtInTransCrcy:          data.PaytReqnAmtInTransCrcy,
		Currency:                        data.Currency,
		PaymentMethod:                   data.PaymentMethod,
		PayerPostingDate:                data.PayerPostingDate,
		PayerHouseBank:                  data.PayerHouseBank,
		PayerHouseBankAccount:           data.PayerHouseBankAccount,
		PayerFinInstCountry:             data.PayerFinInstCountry,
		PayerFinInstCode:                data.PayerFinInstCode,
		PayerFinInstBranchCode:          data.PayerFinInstBranchCode,
		PayerFinInstFullCode:            data.PayerFinInstFullCode,
		PayerFinInstSWIFTCode:           data.PayerFinInstSWIFTCode,
		PayerInternalFinInstCustomerID:  data.PayerInternalFinInstCustomerID,
		PayerInternalFinInstAccountID:   data.PayerInternalFinInstAccountID,
		PayerFinInstControlKey:          data.PayerFinInstControlKey,
		PayerFinInstAccount:             data.PayerFinInstAccount,
		PayerFinInstAccountName:         data.PayerFinInstAccountName,
		PayerFinInstName:                data.PayerFinInstName,
		PayerFinInstBranchName:          data.PayerFinInstBranchName,
		PayerFinInstCustomerIDByFinInst: data.PayerFinInstCustomerIDByFinInst,
		PaymentRequisitionIsCanceled:    data.PaymentRequisitionIsCanceled,
		CreationDateTime:                data.CreationDateTime,
		ChangedOnDateTime:               data.ChangedOnDateTime,
	}

	return headerCreate
}

// func ConvertToHeaderUpdates(headerUpdates dpfm_api_input_reader.HeaderUpdates) *HeaderUpdates {
// 	data := headerUpdates

// 	return &HeaderUpdates{}
// }

func ConvertToHeaderDoc(subfuncSDC *sub_func_complementer.SDC) *[]HeaderDoc {
	var headerdoc []HeaderDoc

	for _, data := range *subfuncSDC.Message.HeaderDoc {
		headerdoc = append(headerdoc, HeaderDoc{
			BusinessPartner:          data.BusinessPartner,
			PaymentDate:              data.PaymentDate,
			DocType:                  data.DocType,
			DocVersionID:             data.DocVersionID,
			DocID:                    data.DocID,
			FileExtension:            data.FileExtension,
			FileName:                 data.FileName,
			FilePath:                 data.FilePath,
			DocIssuerBusinessPartner: data.DocIssuerBusinessPartner,
		})
	}

	return &headerdoc
}

func ConvertToItem(subfuncSDC *sub_func_complementer.SDC) *[]Item {
	var item []Item

	for _, data := range *subfuncSDC.Message.Item {
		item = append(item, Item{
			PayerPaymentRequisitionItem:    data.PayerPaymentRequisitionItem,
			Payer:                          data.Payer,
			PayerPaymentDate:               data.PayerPaymentDate,
			Payee:                          data.Payee,
			BillFromParty:                  data.BillFromParty,
			BillToParty:                    data.BillToParty,
			Buyer:                          data.Buyer,
			Seller:                         data.Seller,
			InvoiceDocument:                data.InvoiceDocument,
			PayeeFinInstCountry:            data.PayeeFinInstCountry,
			PayeeFinInstCode:               data.PayeeFinInstCode,
			PayeeFinInstBranchCode:         data.PayeeFinInstBranchCode,
			PayeeFinInstFullCode:           data.PayeeFinInstFullCode,
			PayeeFinInstSWIFTCode:          data.PayeeFinInstSWIFTCode,
			PaytReqnItemAmtInTransCrcy:     data.PaytReqnItemAmtInTransCrcy,
			PayeeInternalFinInstCustomerID: data.PayeeInternalFinInstCustomerID,
			PayeeInternalFinInstAccountID:  data.PayeeInternalFinInstAccountID,
			PayeeFinInstControlKey:         data.PayeeFinInstControlKey,
			PayeeFinInstAccount:            data.PayeeFinInstAccount,
			PayeeFinInstAccountName:        data.PayeeFinInstAccountName,
			PayeeFinInstName:               data.PayeeFinInstName,
			PayeeFinInstBranchName:         data.PayeeFinInstBranchName,
			PayerAccountingDocument:        data.PayerAccountingDocument,
			PayerAccountingDocumentItem:    data.PayerAccountingDocumentItem,
			CreationDateTime:               data.CreationDateTime,
			ChangedOnDateTime:              data.ChangedOnDateTime,
		})
	}

	return &item
}
