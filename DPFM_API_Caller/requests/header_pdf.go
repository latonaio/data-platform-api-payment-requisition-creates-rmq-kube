package requests

type HeaderPDF struct {
	DocType      string  `json:"DocType"`
	DocVersionID int     `json:"DocVersionID"`
	DocID        string  `json:"DocID"`
	FileName     *string `json:"FileName"`
}
