package dpfm_api_input_reader

type EC_MC struct {
}

type SDC struct {
	ConnectionKey       string             `json:"connection_key"`
	Result              bool               `json:"result"`
	RedisKey            string             `json:"redis_key"`
	Filepath            string             `json:"filepath"`
	APIStatusCode       int                `json:"api_status_code"`
	RuntimeSessionID    string             `json:"runtime_session_id"`
	BusinessPartnerID   *int               `json:"business_partner"`
	ServiceLabel        string             `json:"service_label"`
	PaymentRequisition  PaymentRequisition `json:"PaymentRequisition"`
	APISchema           string             `json:"api_schema"`
	Accepter            []string           `json:"accepter"`
	OrderID             *int               `json:"order_id"`
	Deleted             bool               `json:"deleted"`
	SQLUpdateResult     *bool              `json:"sql_update_result"`
	SQLUpdateError      string             `json:"sql_update_error"`
	SubfuncResult       *bool              `json:"subfunc_result"`
	SubfuncError        string             `json:"subfunc_error"`
	ExconfResult        *bool              `json:"exconf_result"`
	ExconfError         string             `json:"exconf_error"`
	APIProcessingResult *bool              `json:"api_processing_result"`
	APIProcessingError  string             `json:"api_processing_error"`
}

type PaymentRequisition struct {
	Payer                          *int      `json:"Payer"`
	PayerPaymentDate               *string   `json:"PayerPaymentDate"`
	PayerPaymentRequisitionID      *int      `json:"PayerPaymentRequisitionID"`
	PaymentReqnStatus              string    `json:"PaymentReqnStatus"`
	AcceptanceNoByFinInst          string    `json:"AcceptanceNoByFinInst"`
	PaytReqnAmtInTransCrcy         float32   `json:"PaytReqnAmtInTransCrcy"`
	Currency                       string    `json:"Currency"`
	PaymentMethod                  string    `json:"PaymentMethod"`
	PayerPostingDate               *string   `json:"PayerPostingDate"`
	PayerHouseBank                 string    `json:"PayerHouseBank"`
	PayerHouseBankAccount          string    `json:"PayerHouseBankAccount"`
	PayerFinInstCountry            string    `json:"PayerFinInstCountry"`
	PayerFinInstCode               string    `json:"PayerFinInstCode"`
	PayerFinInstBranchCode         string    `json:"PayerFinInstBranchCode"`
	PayerFinInstFullCode           string    `json:"PayerFinInstFullCode"`
	PayerFinInstSWIFTCode          string    `json:"PayerFinInstSWIFTCode"`
	PayerInternalFinInstCustomerID *int      `json:"PayerInternalFinInstCustomerID"`
	PayerInternalFinInstAccountID  *int      `json:"PayerInternalFinInstAccountID"`
	PayerFinInstControlKey         string    `json:"PayerFinInstControlKey"`
	PayerFinInstAccount            string    `json:"PayerFinInstAccount"`
	PayerFinInstAccountName        string    `json:"PayerFinInstAccountName"`
	PayerFinInstName               string    `json:"PayerFinInstName"`
	PayerFinInstBranchName         string    `json:"PayerFinInstBranchName"`
	PaymentApplicantCode           string    `json:"PaymentApplicantCode"`
	CreationDateTime               *string   `json:"CreationDateTime"`
	ChangedOnDateTime              *string   `json:"ChangedOnDateTime"`
	HeaderPDF                      HeaderPDF `json:"HeaderPDF"`
	Item                           Item      `json:"Item"`
}

type HeaderPDF struct {
	BusinessPartner      *int    `json:"BusinessPartner"`
	PaymentDate          *string `json:"PaymentDate"`
	PaymentRequisitionID *int    `json:"PaymentRequisitionID"`
	DocType              string  `json:"DocType"`
	DocVersionID         *int    `json:"DocVersionID"`
	DocID                string  `json:"DocID"`
	FileName             string  `json:"FileName"`
}

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
