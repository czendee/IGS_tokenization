package main

import (
//	"fmt"
	//"log"
    utilito "banwire/services/file_tokenizer/util"


	modelito "banwire/services/file_tokenizer/model"
	miu "banwire/services/file_tokenizer/util"
	 _ "github.com/lib/pq"   //use go get github.com/lib/pq
	 
    "strings"
//    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "net/url"
)


///////////////////////////////////////business logic functions

//////////////////////////////////web


   func logicProcesspaymentWeb(requestData modelito.RequestPayment, errorGeneral string) (modelito.ExitoData,string) {
   	
   	     var resultadoPayment modelito.ExitoData  //response.go

	    valoresbanwire := url.Values{
		"method": {"payment"},
//		"user": {"pruebasbw"},  //this value was mentioned by Charly  dec 05,2018
		"user": {Config_WS_crbanwire_pass},  //this value needs to be configurable to Move to production. 22 Jan 2019        
		"reference": {requestData.Paymentreference}, 
		"token": {requestData.Token},         
		"amount": {strings.Replace(requestData.Amount, ",", "", -1)  },          
        "cvv": {requestData.Cvv},  
	}
utilito.LevelLog(Config_env_log, "3", "web api el amount 6:"+strings.Replace(requestData.Amount, ",", "", -1))    
utilito.LevelLog(Config_env_log, "3", "web api el cvv"+requestData.Cvv)
//	    response,err := http.PostForm("https://cr.banwire.com/?action=card",
	    response,err := http.PostForm(Config_WS_crbanwire_url+"/?action=card",

	valoresbanwire)
	
	
	    if err != nil {
	        utilito.LevelLog(Config_env_log, "3", "The HTTP request failed with error "+ err.Error() +"\n")
	    } else {
	        data, _ := ioutil.ReadAll(response.Body)
	        utilito.LevelLog(Config_env_log, "3", string(data))
	        if strings.Contains(string(data), "error") {
	         	errorGeneral ="Error returned: "+string(data)
	        }else{
		        	var str WebResponsePayment
					_ = json.Unmarshal(data, &str)
					
					if str.Description == "Payment on Demand" {
					    // Do Stuff
					    resultadoPayment.Token  =requestData.Token
					    resultadoPayment.PaymentReference = str.Reference
					    resultadoPayment.Authcode =str.Authcode
					    resultadoPayment.Idtransaction =str.Idtransaction

					}else{
	         			errorGeneral ="Error in JSON returned/internal webservice"

					}
	        	
	        }
	    }
	    utilito.LevelLog(Config_env_log, "3", "web api Terminating the application...")
	    
   	  return  resultadoPayment, errorGeneral
   }
   
   
   func logicGeneratetokenizedWeb(requestData modelito.RequestTokenized, errorGeneral string) (modelito.ExitoDataTokenized,string) {
		////////////////////////////////////////////////process db steps
	   //START    

      var resultadoService modelito.ExitoDataTokenized

     month,year, errorGeneral := miu.ConvertMMYYintoMonthYear(requestData.Exp)   //utils.go

      
	   if errorGeneral==""{
		    valoresbanwire := url.Values{
			"method": {"add"},
//		"user": {"pruebasbw"},  //this value was mentioned by Charly dec 05,2018
		"user": {Config_WS_crbanwire_pass},  //this value needs to be configurable to Move to production. 22 Jan 2019        
			"email": {"generalseguros@genearlseguros.com"},
			"number": {requestData.Card},  //
			"exp_month": {month},              //requestData.Exp  solo mes
			"exp_year": {year},             //requestData.Exp  solo año u de 4 digitos
//05Dec2018	"cvv": {requestData.Cvv},                   //iissue necesita el cvv
			"cvv": {"000"},                   //Charly V indicated to send 000
			"name": {"generalseguros"}, 
			"address": {"generalseguros"},
			"postal_code": {"06000"},
			}	

//		    response,err := http.PostForm("https://cr.banwire.com/?action=card&exists=1",
		    response,err := http.PostForm(Config_WS_crbanwire_url+"/?action=card&exists=1",

			valoresbanwire)
			
		    if err != nil {
		        utilito.LevelLog(Config_env_log, "3", "The HTTP request failed with error "+ err.Error() +"\n")
		         errorGeneral ="The HTTP request failed with error "+err.Error()
		    } else {
		        data, _ := ioutil.ReadAll(response.Body)
		        utilito.LevelLog(Config_env_log, "3", "webservice response:" +string(data))
		        if strings.Contains(string(data), "error") {
		         	errorGeneral ="Error returned: "+string(data)
		        }else{
		        	utilito.LevelLog(Config_env_log, "3", "logicGeneratetokenizedWeb 02")

		        	var str WebResponseAdd
                      
					_ = json.Unmarshal(data, &str)
		        	utilito.LevelLog(Config_env_log, "3", "logicGeneratetokenizedWeb 03")
   					utilito.LevelLog(Config_env_log, "3", "logicGeneratetokenizedWeb token:"+str.Token)
   					utilito.LevelLog(Config_env_log, "3", "logicGeneratetokenizedWeb task:"+str.Task)
   					utilito.LevelLog(Config_env_log, "3", "logicGeneratetokenizedWeb type:"+str.Card.Type+"::")
   					utilito.LevelLog(Config_env_log, "3", "logicGeneratetokenizedWeb cat:"+str.Card.Category+"::")
				    resultadoService.Token  =str.Token
				    resultadoService.Type  =str.Card.Type
				    resultadoService.Category  =str.Card.Category
///revisar con charly y toño estas reglas
/// cuando si insertar y cuando no
				    
					if str.Task == "add" {
		        		utilito.LevelLog(Config_env_log, "3", "logicGeneratetokenizedWeb 04")
						if str.Result {
		        			utilito.LevelLog(Config_env_log, "3", "logicGeneratetokenizedWeb 05")
							if str.Exists  {
								//error
								 // the card already exists
		         				//aun asi, debe de registrarse en DB,
		         				//errorGeneral ="Error : this card already exists and was tokenized previously"
							    resultadoService.Token  =str.Token
							    resultadoService.Type  =str.Card.Type
							    resultadoService.Category  =str.Card.Category
   					        	utilito.LevelLog(Config_env_log, "3", "logicGeneratetokenizedWeb 07:"+resultadoService.Token)
   					        	utilito.LevelLog(Config_env_log, "3", "logicGeneratetokenizedWeb 07:"+resultadoService.Type)

							}else{ // this is the flow we care
   					        	utilito.LevelLog(Config_env_log, "3", "logicGeneratetokenizedWeb 06")
					    // Do Stuff
							    resultadoService.Token  =str.Token
							    resultadoService.Type  =str.Card.Type
							    resultadoService.Category  =str.Card.Category
   					        	utilito.LevelLog(Config_env_log, "3", "logicGeneratetokenizedWeb 07:"+resultadoService.Token)
   					        	utilito.LevelLog(Config_env_log, "3", "logicGeneratetokenizedWeb 07:"+resultadoService.Type)
							}
							
						}else{
							// the card tokenization fail
		         			errorGeneral ="Error : the result of the tokenization is FALSE"
						}


					}else{
						// the task for the card needs to be add
		         		errorGeneral ="Error : intenal configuration task needs to be ADD"
					}
					
		        }
		    }
	   	
	   }
     
	    utilito.LevelLog(Config_env_log, "3", "web api Terminating the application...")
   	  return  resultadoService, errorGeneral
   }
   
   
 type WebResponseAdd struct {
//    ID string `json:"id"`
    Token string `json:"token"`
    Task string `json:"task"`
    Result bool `json:"result"`
    Exists bool `json:"exists"`
    Card card `json:"card"`
   
}

 type card struct {
    Type string `json:"type"`
    Category string `json:"category"`    
}

 type WebResponsePayment struct {
    ID string `json:"id"`
    Idtransaction string `json:"id_transaction"`
    Authcode string `json:"auth_code"`
    Reference string `json:"reference"`
    Description string `json:"description"`
    Amount string `json:"amount"`
}