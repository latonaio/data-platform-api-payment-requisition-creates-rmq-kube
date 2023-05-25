package requests

type Header struct {
	Payer                           *int      `json:"Payer"`
	PayerPaymentDate                *string   `json:"PayerPaymentDate"`
	PayerPaymentRequisitionID       int       `json:"PayerPaymentRequisitionID"`
	PaymentReqnStatus               *string   `json:"PaymentReqnStatus"`
	AcceptanceNoByFinInst           *string   `json:"AcceptanceNoByFinInst"`
	PaytReqnAmtInTransCrcy          *float32  `json:"PaytReqnAmtInTransCrcy"`
	Currency                        *string   `json:"Currency"`
	PaymentMethod                   *string   `json:"PaymentMethod"`
	PayerPostingDate                *string   `json:"PayerPostingDate"`
	PayerHouseBank                  *string   `json:"PayerHouseBank"`
	PayerHouseBankAccount           *string   `json:"PayerHouseBankAccount"`
	PayerFinInstCountry             *string   `json:"PayerFinInstCountry"`
	PayerFinInstCode                *string   `json:"PayerFinInstCode"`
	PayerFinInstBranchCode          *string   `json:"PayerFinInstBranchCode"`
	PayerFinInstFullCode            *string   `json:"PayerFinInstFullCode"`
	PayerFinInstSWIFTCode           *string   `json:"PayerFinInstSWIFTCode"`
	PayerInternalFinInstCustomerID  *int      `json:"PayerInternalFinInstCustomerID"`
	PayerInternalFinInstAccountID   *int      `json:"PayerInternalFinInstAccountID"`
	PayerFinInstControlKey          *string   `json:"PayerFinInstControlKey"`
	PayerFinInstAccount             *string   `json:"PayerFinInstAccount"`
	PayerFinInstAccountName         *string   `json:"PayerFinInstAccountName"`
	PayerFinInstName                *string   `json:"PayerFinInstName"`
	PayerFinInstBranchName          *string   `json:"PayerFinInstBranchName"`
	PayerFinInstCustomerIDByFinInst *string   `json:"PayerFinInstCustomerIDByFinInst"`
	PaymentRequisitionIsCanceled    *bool     `json:"PaymentRequisitionIsCanceled"`
	CreationDateTime                *string   `json:"CreationDateTime"`
	ChangedOnDateTime               *string   `json:"ChangedOnDateTime"`
	HeaderPDF                       HeaderPDF `json:"HeaderPDF"`
	Item                            Item      `json:"Item"`
}
