package main

import (
	//"log"
    "strings"
    utilito "banwire/services/file_tokenizer/util"
	modelito "banwire/services/file_tokenizer/model"
	 _ "github.com/lib/pq"   //use go get github.com/lib/pq
	 "time"
	 "encoding/json"
)

///////////////////////////////////////////////////////////////7/////version 2

///fetchCards
func getJsonResponseV2( cardsFound []modelito.Card)([]byte, error) {
	
	if(cardsFound !=nil){
		
		mainStruct := modelito.AutoGenerated{StatusMessage: "Success" ,Status:"1"}
     	for _, d := range cardsFound {
     		utilito.LevelLog(Config_env_log, "3", "el registor trae:"+d.Token+" "+d.Bin)
			w := modelito.CardData{time.Now(), d.Token,d.Last,d.Brand,d.Valid,d.Bin,d.Score,d.Type}   //request.go
			mainStruct.Cards = append(mainStruct.Cards, w)
 		}
 		return json.MarshalIndent(mainStruct, "", "  ")

/* 		
		for i := 0; i < 5; i++ {
			w := modelito.CardData{time.Now(), "G#$$%ytoywteouwytr","1234","VISA","2501","781234","1"}
			mainStruct.Cards = append(mainStruct.Cards, w)
		}
*/
	}


	return nil,nil
}

//processpayment
func getJsonResponsePaymentV2(datoPayment modelito.ExitoData)([]byte, error) {
	
									utilito.LevelLog(Config_env_log, "3", " getJsonResponsePaymentV2 token:!\n"+datoPayment.Token)
									utilito.LevelLog(Config_env_log, "3", " getJsonResponsePaymentV2 bin:!\n"+datoPayment.Bin)
									utilito.LevelLog(Config_env_log, "3", " getJsonResponsePaymentV2 last:!\n"+datoPayment.LastDigits)
									
	mainStruct := modelito.ResponsePayment{StatusMessage: "Success",Status:"1"}        //response.go
    w := modelito.ExitoData{ datoPayment.Token,datoPayment.PaymentReference,datoPayment.Authcode,datoPayment.Idtransaction,datoPayment.Marca,datoPayment.Bin,datoPayment.LastDigits,datoPayment.Type}   //response.go
    mainStruct.SucessData=w
	return json.MarshalIndent(mainStruct, "", "  ")
}



/////////response for tokenize


func getJsonResponseTokenizeV2(cardTokenized modelito.Card)([]byte, error) {
	
	mainStruct := modelito.ResponseTokenize{StatusMessage: "Success",Status:"1"}
    w := modelito.CardData{time.Now(), cardTokenized.Token,cardTokenized.Last, cardTokenized.Brand,cardTokenized.Valid,cardTokenized.Bin,cardTokenized.Score,cardTokenized.Type}  
    mainStruct.Card = w


	return json.MarshalIndent(mainStruct, "", "  ")
}

func getJsonResponseValidateFileV2(cardTokenized modelito.Card)([]byte, error) {
	
	mainStruct := modelito.ResponseTokenize{StatusMessage: "Success",Status:"1"}
//    w := modelito.CardData{time.Now(), cardTokenized.Token,cardTokenized.Last, cardTokenized.Brand,cardTokenized.Valid,cardTokenized.Bin,cardTokenized.Score,cardTokenized.Type}  
//    mainStruct.Card = w


	return json.MarshalIndent(mainStruct, "", "  ")
}

////////////////////////ERROR response


func getJsonResponseError(errorMsg, errorNumber string )([]byte, error) {
	
	mainStruct := modelito.ResponseError{StatusMessage: errorMsg,Status:errorNumber}

	return json.MarshalIndent(mainStruct, "", "  ")
}

func getJsonResponseErrorValidateFile(fileStatusMsg, fileStatusNumber string, linesStatus []modelito.ExitoDataValidaLine  )([]byte, error) {
	
    mainStruct :=modelito.ResponseTokenFile{StatusMessage: fileStatusMsg ,Status:fileStatusNumber}




     	for _, d := range linesStatus {
     		utilito.LevelLog(Config_env_log, "3"," getting json ready - line:"+d.StatusMessage )
			w := modelito.ExitoDataValidaLine{d.Line, d.StatusMessage,d.Status}   //request.go
			mainStruct.SucessValidacion = append(mainStruct.SucessValidacion, w)
 		}


	return json.MarshalIndent(mainStruct, "", "  ")
}

func getJsonResponseTokenFile(fileStatusMsg, fileStatusNumber string, validaLinesStatus []modelito.ExitoDataValidaLine, tokenLinesStatus []modelito.ExitoDataTokenLine  )([]byte, error) {
	
    mainStruct :=modelito.ResponseTokenFile{StatusMessage: fileStatusMsg ,Status:fileStatusNumber, CardsToken:"4"}


     	for _, d := range validaLinesStatus {
     		utilito.LevelLog(Config_env_log, "3"," getting json ready - line:"+d.StatusMessage )
			w := modelito.ExitoDataValidaLine{d.Line, d.StatusMessage,d.Status}   //request.go
			mainStruct.SucessValidacion = append(mainStruct.SucessValidacion, w)
 		}

        for _, d := range tokenLinesStatus {
     		utilito.LevelLog(Config_env_log, "3"," getting json ready - line:"+d.StatusMessage )
             d.StatusMessage =  strings.Replace(d.StatusMessage, ":", "", -1)
			w := modelito.ExitoDataTokenLine{d.Line, d.StatusMessage,d.Status, d.Date, d.Token, d.LastDigits, d.Marca, d.Vigencia, d.Bin, d.Score, d.Type}   //request.go
			mainStruct.SucessDataEachRowToken = append(mainStruct.SucessDataEachRowToken, w)
 		}


	return json.MarshalIndent(mainStruct, "", "  ")
}

func getJsonResponsePaymentFile(fileStatusMsg, fileStatusNumber string, validaLinesStatus []modelito.ExitoDataValidaLine, paymentLinesStatus []modelito.ExitoDataPayLine  )([]byte, error) {
	
    mainStruct :=modelito.ResponsePayFile{StatusMessage: fileStatusMsg ,Status:fileStatusNumber, Payments:"4"}


     	for _, d := range validaLinesStatus {
     		utilito.LevelLog(Config_env_log, "3"," getting json ready - line:"+d.StatusMessage )
			w := modelito.ExitoDataValidaLine{d.Line, d.StatusMessage,d.Status}   //request.go
			mainStruct.SucessValidacion = append(mainStruct.SucessValidacion, w)
 		}

        for _, d := range paymentLinesStatus {
     		utilito.LevelLog(Config_env_log, "3"," getting json ready - line:"+d.StatusMessage )
             d.StatusMessage =  strings.Replace(d.StatusMessage, ":", "", -1)
			w := modelito.ExitoDataPayLine{d.Line, d.StatusMessage,d.Status,d.Token,d.PaymentReference,d.Authcode,d.Idtransaction,d.Marca,d.Bin,d.LastDigits,d.Type}   //request.go
			mainStruct.SucessDataEachRowPay = append(mainStruct.SucessDataEachRowPay, w)
 		}
         

	return json.MarshalIndent(mainStruct, "", "  ")
}