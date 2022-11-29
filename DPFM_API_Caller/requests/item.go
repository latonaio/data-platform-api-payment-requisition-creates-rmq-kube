package requests

type Item struct {
	Payer                          *int     `json:"Payer"`
	PayerPaymentDate               *string  `json:"PayerPaymentDate"`
	PayerPaymentRequisitionID      *int     `json:"PayerPaymentRequisitionID"`
	PayerPaymentRequisitionItem    *int     `json:"PayerPaymentRequisitionItem"`
	Payee                          *int     `json:"Payee"`
	PayeeFinInstCountry            string   `json:"PayeeFinInstCountry"`
	PayeeFinInstCode               string   `json:"PayeeFinInstCode"`
	PayeeFinInstBranchCode         string   `json:"PayeeFinInstBranchCode"`
	PayeeFinInstFullCode           string   `json:"PayeeFinInstFullCode"`
	PayeeFinInstSWIFTCode          string   `json:"PayeeFinInstSWIFTCode"`
	PaytReqnItemAmtInTransCrcy     *float32 `json:"PaytReqnItemAmtInTransCrcy"`
	PayeeInternalFinInstCustomerID *int     `json:"PayeeInternalFinInstCustomerID"`
	PayeeInternalFinInstAccountID  *int     `json:"PayeeInternalFinInstAccountID"`
	PayeeFinInstControlKey         string   `json:"PayeeFinInstControlKey"`
	PayeeFinInstAccount            string   `json:"PayeeFinInstAccount"`
	PayeeFinInstAccountName        string   `json:"PayeeFinInstAccountName"`
	PayeeFinInstName               string   `json:"PayeeFinInstName"`
	PayeeFinInstBranchName         string   `json:"PayeeFinInstBranchName"`
	PayerAccountingDocument        *int     `json:"PayerAccountingDocument"`
	PayerAccountingDocumentItem    *int     `json:"PayerAccountingDocumentItem"`
	CreationDateTime               *string  `json:"CreationDateTime"`
	ChangedOnDateTime              *string  `json:"ChangedOnDateTime"`
}
