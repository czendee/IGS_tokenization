package model

import (
	"time"

	)


///////////////////response for validate token
type ResponseTokenFile struct {    
	StatusMessage string       `json:"status_message_file"`
	Status       string       `json:"status_file"`
    CardsToken string       `json:"cards_token"`
	SucessDataEachRowToken   []ExitoDataTokenLine     `json:"data_row"`
    SucessValidacion   []ExitoDataValidaLine     `json:"data_val_row"`
}

type ResponsePayFile struct {    
	StatusMessage string       `json:"status_message_file"`
	Status       string       `json:"status_file"`
    Payments string       `json:"payments"`
    SucessDataEachRowPay   []ExitoDataPayLine     `json:"data_row"`
    SucessValidacion   []ExitoDataValidaLine     `json:"data_val_row"`
}


type ResponsePaymentsToken struct {    
	StatusMessage string       `json:"status_message"`
	Status       string       `json:"status"`
    SucessDataEachRowPay   []ExitoDataPaymentsTokenLine     `json:"data_row"`

}

type ResponseTokensPerCustRef struct {    
	StatusMessage string       `json:"status_message"`
	Status       string       `json:"status"`
    SucessDataEachRowPay   []ExitoDataTokensPerCustLine     `json:"data_row"`

}


type ExitoDataTokenLine struct {    
    Line          string       `json:"line"`
	StatusMessage string       `json:"status_message"`
	Status       string       `json:"status"`
	
    Date time.Time `json:"card_date"`
	Token  string   `json:"card_token"`
    LastDigits  string   `json:"card_last"`
	Marca  string   `json:"card_brand"`
	Vigencia  string   `json:"card_exp"`
	Bin  string   `json:"card_bin"`
	Score  string   `json:"card_score"`
	Type  string   `json:"type_card"`
    
}

type ExitoDataPayLine struct {    
    Line          string       `json:"line"`
	StatusMessage string       `json:"status_message"`
	Status       string       `json:"status"`
	
    Token  string   `json:"card_token"`
    PaymentReference  string   `json:"paymentreference"`
	Authcode  string   `json:"authcode"`
	Idtransaction  string   `json:"idtransaction"`
	Marca  string   `json:"card_brand"`
	Bin  string   `json:"card_bin"`
    LastDigits  string   `json:"card_last"`
	Type  string   `json:"type_card"`
}

type ExitoDataValidaLine struct {    
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



type ExitoDataPaymentsTokenLine struct {    
    Line          string       `json:"line"`
	StatusMessage string       `json:"status_message"`
	Status       string       `json:"status"`
	
    Token  string   `json:"payment_token"`
    Date  string `json:"payment_date"`
    Amount  string   `json:"payment_amount"`
}

type ExitoDataTokensPerCustLine struct {    
    Line          string       `json:"line"`
	StatusMessage string       `json:"status_message"`
	Status       string       `json:"status"`
	
    Date string `json:"card_date"`
	Token  string   `json:"card_token"`
    LastDigits  string   `json:"card_last"`
	Marca  string   `json:"card_brand"`
	Vigencia  string   `json:"card_exp"`
	Bin  string   `json:"card_bin"`
	Score  string   `json:"card_score"`
	Type  string   `json:"type_card"`
}
