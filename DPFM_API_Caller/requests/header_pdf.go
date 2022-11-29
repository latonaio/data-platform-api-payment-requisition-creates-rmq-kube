package requests

type HeaderPDF struct {
	BusinessPartner      *int    `json:"BusinessPartner"`
	PaymentDate          *string `json:"PaymentDate"`
	PaymentRequisitionID *int    `json:"PaymentRequisitionID"`
	DocType              string  `json:"DocType"`
	DocVersionID         *int    `json:"DocVersionID"`
	DocID                string  `json:"DocID"`
	FileName             string  `json:"FileName"`
}

