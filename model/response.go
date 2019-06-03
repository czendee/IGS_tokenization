package model



///////////////////response for validate token
type ResponseTokenFile struct {    
	StatusMessage string       `json:"status_message_file"`
	Status       string       `json:"status_file"`
	SucessDataEachRowToken   []ExitoDataTokenLine     `json:"sucess_data_each_row"`
    SucessDataEachRowProcess   []ExitoDataTokenLine     `json:"sucess_data_each_row_process"`
}

type ExitoDataTokenLine struct {    
    Line          string       `json:"line"`
	StatusMessage string       `json:"status_message"`
	Status       string       `json:"status"`
	
}


///////////////////response for payment
type ResponsePayment struct {    
	StatusMessage string       `json:"status_message"`
	Status       string       `json:"status"`
	SucessData   ExitoData     `json:"sucess_data"`
}

type ExitoData struct {
	Token  string   `json:"card_token"`
    PaymentReference  string   `json:"paymentreference"`
	Authcode  string   `json:"authcode"`
	Idtransaction  string   `json:"idtransaction"`
	Marca  string   `json:"card_brand"`
	Bin  string   `json:"card_bin"`
    LastDigits  string   `json:"card_last"`
	Type  string   `json:"type_card"`
}

type ExitoDataTokenized struct {
	Token  string   `json:"token"`
    Type  string   `json:"type"`
	Category  string   `json:"category"`
}
