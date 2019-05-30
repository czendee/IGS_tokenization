package main

import (
	"net/http"
	//"log"
    utilito "banwire/services/file_tokenizer/util"
	"banwire/services/file_tokenizer/db"
	modelito "banwire/services/file_tokenizer/model"
//	"time"
//	"encoding/json"
	 _ "github.com/lib/pq"   //use go get github.com/lib/pq

)


    
func ProcessGettokenizedcards(w http.ResponseWriter,  requestData modelito.RequestTokenizedCards) (string,string) {

    var errorGeneral string
    var errorGeneralNbr string
    	var result string
   var valoresParaResponder  []modelito.Card

    errorGeneral=""

    if errorGeneral=="" {


		utilito.LevelLog(Config_env_log, "3", "CZ   STEP Validate Parms")

		/// START
	        result ="OK realizarpago"+requestData.Cardreference
		    utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening fetch:"+result)
		    
		     utilito.LevelLog(Config_env_log, "3", "CZ   STEP Validate paramters request")
		    errorGeneral= validaReqFetchCards(requestData)
		
		
		/// END

    }				    
		        
    if errorGeneral!="" && errorGeneralNbr=="" {
    	//prepare response with error 300
    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 300. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:300 -Missing parameter"	+errorGeneral
    	errorGeneralNbr="300"
    }

	////////////////////////////////////////////////DB	
	//	    resultado,errfetchDB:= fetchFromDB ()
	if errorGeneral==""{//continue next step

       	    utilito.LevelLog(Config_env_log, "3", "CZ   STEP Consume DB")
         valoresParaResponder,errorGeneral =logicDBGettokenizedcardsV2(requestData, errorGeneral) 


    }				    
    if errorGeneral!="" && errorGeneralNbr==""{
    	//prepare response with error 310
    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 310. Error obtaining cards:"+errorGeneral)
    	errorGeneral="ERROR:310 -  Error obtaining cards -"	+errorGeneral
	    errorGeneralNbr="310"
    }

		        
	//response
    utilito.LevelLog(Config_env_log, "3", "CZ    handler DB Listening test gettokenizedcards  2")					

	//////////    format the response
    if errorGeneral==""{//continue next step
		utilito.LevelLog(Config_env_log, "3", "CZ   STEP Validate Parms")
			/// START
		fieldDataBytesJson,err := getJsonResponseV2(valoresParaResponder)
		
		utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test gettokenizedcards  3")	
		
		result ="OK gettokenizedcards"+requestData.Cardreference+"resultado"
		//////////    write the response
		w.Header().Set("Content-Type", "application/json")
		 w.Write(fieldDataBytesJson)
		 
		 utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test gettokenizedcards  4"+"<html><body>"+ result+"</body></html>")
			         
        if err!=nil{
        	utilito.LevelLog(Config_env_log, "3", "Eror en generando response")
            errorGeneral= err.Error()
        }		
		
		/// END

    }				    
		 
    if errorGeneral!="" && errorGeneralNbr==""{//continue next step
    	utilito.LevelLog(Config_env_log, "3", "CZ   prepare the JSON response for ERROR")

	    //  START 
	    errorGeneral="ERROR:330 -Error preparing the response"	+errorGeneral
	    errorGeneralNbr="330"
	    //  END
    }

     return errorGeneral, errorGeneralNbr
}



// Generatetokenized for receive and handle the request from client
func ProcessGeneratetokenized(w http.ResponseWriter, requestData modelito.RequestTokenized) (string,string) {
	defer func() {
		db.Connection.Close(nil)
	}()
	  var result string

     var errorGeneral string
     var errorGeneralNbr string
     
     var resultCardTokenized modelito.Card
     
     var obtainedDataWebservice modelito.ExitoDataTokenized
     
    errorGeneral=""


	////////////////////////////////////////////////validate parms
	/// START
    if errorGeneral==""{//continue next step
		    result ="OK realizarpago"+requestData.Clientreference+"    :    " +requestData.Paymentreference+"    :    " +requestData.Card+"    :    " +requestData.Exp+"    :    " +requestData.Cvv
    		utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test handleGeneratetokenized:"+result)
		    utilito.LevelLog(Config_env_log, "3", "CZ   STEP Validate paramters request")
		    errorGeneral= validaReqGenerateTokenized(requestData)	
		/// END

	}	
		              
    if errorGeneral!="" && errorGeneralNbr=="" {
    	//prepare response with error 800
    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 200. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR :200 -Missing parameter "	+errorGeneral
		errorGeneralNbr="200"
    }




	////////////////////////////////////////////////consume internal websrvice banwire
	//////////////////            tokenization 

    if errorGeneral==""{//continue next step
				/// START
				obtainedDataWebservice, errorGeneral =logicGeneratetokenizedWeb(requestData, errorGeneral)
				
				/// END
	}	

    if errorGeneral!="" && errorGeneralNbr==""{
    	//prepare response with error 210
    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 210. Error processing payment:"+errorGeneral)
    	errorGeneral="ERROR:210 -Error processing payment:"	+errorGeneral
		errorGeneralNbr="210"
    }

				
				
	////////////////////////////////////////////////DB	
	//	    insert new record in Card , if customer doesn't exist, insert a new one?
	//  Update if exist, if not insert in Customer

    if errorGeneral==""{//continue next stepjhlkjg 
        	utilito.LevelLog(Config_env_log, "3", "CZ   el  token:"+obtainedDataWebservice.Token)
    			resultCardTokenized, errorGeneral =logicGeneratetokenizedDBV2(requestData,obtainedDataWebservice , errorGeneral)
    						
	}					

    if errorGeneral!="" && errorGeneralNbr==""{
    	//prepare response with error 220
    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 220. Error generating token:"+errorGeneral)
    	errorGeneral="ERROR:220 -Error generating token:"	+errorGeneral
		errorGeneralNbr="220"
    }

	//response
    if errorGeneral==""{//continue next step
		utilito.LevelLog(Config_env_log, "3", "CZ   STEP Post the response JSON ready")
		
			/// START
		fieldDataBytesJsonTokenize,err := getJsonResponseTokenizeV2(resultCardTokenized)
			
		utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test realizarpago  3")	
	    
	    w.Header().Set("Content-Type", "application/json")
	    w.Write(fieldDataBytesJsonTokenize)
		utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test handleGeneratetokenized  4"+"<html><body>"+ result+"</body></html>")
        if err!=nil{
        	utilito.LevelLog(Config_env_log, "3", "Eror en generando response")
	        errorGeneral= err.Error()
        }
				
		/// END
	}	

    if errorGeneral!="" && errorGeneralNbr==""{
    	//prepare response with error 230
    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 230. Error generating Response Tokenized:"+errorGeneral)
    	errorGeneral="ERROR:230 -Error generating Response Tokenized:"	+errorGeneral
		errorGeneralNbr="230"
    }
    
	 utilito.LevelLog(Config_env_log, "3", "CZ  ends func tokenized")
	 
	return errorGeneral, errorGeneralNbr
}





func GetCardType(number string) string {
	return "VISA"
}


/////////////////////////v4
/////////////////////////v4

// v4Processpayment  receive and handle the request from client, access DB
func v4ProcessProcessPayment(w http.ResponseWriter, requestData modelito.RequestPayment) (string,string){
	defer func() {
		db.Connection.Close(nil)
	}()
    var result string
    var errorGeneral string
    var	errorGeneralNbr string

    var resultadoPayment modelito.ExitoData
    errorGeneral=""

	////////////////////////////////////////////////validate parms
	/// START
    
    if errorGeneral==""{//continue next step
    	utilito.LevelLog(Config_env_log, "3", "CZ   STEP Validate Parms")

		/// START
	        result ="OK realizarpago"+requestData.Clientreference+"    :    " +requestData.Paymentreference+"    :    " +requestData.Token+"    :    " +requestData.Cvv+"    :    " +requestData.Amount+"    :    "
		    utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test realizarpago:"+result)
		    
		    utilito.LevelLog(Config_env_log, "3", "CZ   STEP Validate paramters request")
		    errorGeneral= validaReqProcessPayment(requestData)
		
		
		/// END

    }				    
		        
    if errorGeneral!="" {
    	//prepare response with error 100
    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 100. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:100 - Missing parameter"	+errorGeneral
		errorGeneralNbr="100"
    }
//////////////////////////////////////////DB verify if less payments for the same card
//////////////////////////////////////////in the same day

	////////////////////////////////////////////////DB	
	//	    resultado,errfetchDB:= fetchFromDB ()
 var valoresParaResponder  []modelito.Payment
	if errorGeneral==""{//continue next step

     	 utilito.LevelLog(Config_env_log, "3", "CZ   STEP Consume DB to check if more payments cvan be done today for this card")
         valoresParaResponder,errorGeneral =logicDBCheckNumberOfPaymentsToday(requestData, errorGeneral) 

    }				    
    if errorGeneral!="" && errorGeneralNbr==""{
    	//prepare response with error 105
    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 105. Error Max payments today for this card exceeded:"+errorGeneral)
    	errorGeneral="ERROR 105 -  Error Max payments today for this card exceeded -"	+errorGeneral
	    errorGeneralNbr="105"
    }
     if valoresParaResponder == nil{
         
     }
		        
	//response
    utilito.LevelLog(Config_env_log, "3", "CZ    handler DB Listening test gettokenizedcards  2")

	////////////////////////////////////////////////consume internal websrvice banwire
	//////////////////            process payment
	    if errorGeneral==""{//continue next step
	    	utilito.LevelLog(Config_env_log, "3", "CZ   STEP Consume internal websrvice banwire")

			/// START
			
			resultadoPayment, errorGeneral= logicProcesspaymentWeb(requestData , errorGeneral )  
			/// END

	    }				    
	    if errorGeneral!="" && errorGeneralNbr==""{
	    	//prepare response with error 110
	    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 110. Error processing payment:"+errorGeneral)
	    	errorGeneral="ERROR:110 - Error processing payment"	+errorGeneral
			errorGeneralNbr="110"

	    }

            ///START: New rule: 25/01/2019 - Remove Card if failed first payment 
            //    rule 30/01/2019  , or parameters for first payment are missing or wrong
            //sino hay pago registrado antes, entonces DELETE la card

	    if errorGeneral!="" && errorGeneralNbr!="105"{//except max 3 payments today 

             var valoresParaResponder  string
             var errorGeneralRemoveCard string

            utilito.LevelLog(Config_env_log, "3", "CZ   STEP Consume DB to Remove Card if failed first payment")
            valoresParaResponder,errorGeneralRemoveCard =logicDBRemoveCardIfNotPreviousPayment(requestData, errorGeneral) 

            if errorGeneral!="" && errorGeneralRemoveCard==""{
                //alredy  response, conitnue  with error 110,100
                utilito.LevelLog(Config_env_log, "3", "CZ   NO Prepare Response. Continue with 100,110. Error removing card:"+errorGeneralRemoveCard)
            }
            if valoresParaResponder == ""{
                
            }else{
                utilito.LevelLog(Config_env_log, "3", "CZ   Continue with error. Remove card for failed 1st payment:"+valoresParaResponder)
            }

	    }//end if error first payment failed, or parameters for first payment are missing or wrong
            ///END: New rule: 25/01/2019 - Remove Card if failed first payment
            //    rule 30/01/2019  , or parameters for first payment are missing or wrong
			

	////////////////////////////////////////////////DB	
	//      update the score field: increase by 1
	//      for this card
	//	    
	var  dataObtainedCard  modelito.Card
	    if errorGeneral==""{//continue next step
	    	utilito.LevelLog(Config_env_log, "3", "CZ   STEP  update the score field: increase by 1")
			requestData, dataObtainedCard, errorGeneral= logicProcesspaymentDBV4(requestData , errorGeneral )  

									utilito.LevelLog(Config_env_log, "3", " medio token:!\n"+dataObtainedCard.Token)
									utilito.LevelLog(Config_env_log, "3", " medio bin:!\n"+dataObtainedCard.Bin)
									utilito.LevelLog(Config_env_log, "3", " medio last:!\n"+dataObtainedCard.Last)
		    resultadoPayment.Marca = dataObtainedCard.Brand
		    resultadoPayment.Bin = dataObtainedCard.Bin
		    resultadoPayment.LastDigits= dataObtainedCard.Last
		    resultadoPayment.Type = dataObtainedCard.Type
		    
	    }				    

	    if errorGeneral!="" && errorGeneralNbr==""{
	    	//prepare response with error 120
	    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 120. Error recording results in DB:"+errorGeneral)
	    	errorGeneral="ERROR: 120 - Error recording results in DB"	+errorGeneral
			errorGeneralNbr="120"
	    }

    		    
	//response
	////////////////////////////////////////////////http response	
	//      prepare the JSON response
	//	    
	    if errorGeneral==""{//continue next step
	    	utilito.LevelLog(Config_env_log, "3", "CZ   STEP  prepare the JSON response for SUCCESS")

		    //  START 

		    fieldDataBytesJsonPayment,err := getJsonResponsePaymentV2(resultadoPayment)					
		        w.Header().Set("Content-Type", "application/json")
		        w.Write(fieldDataBytesJsonPayment)
				utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test handleProcesspayment  4"+"<html><body>"+ result+"</body></html>")
                if err!=nil{
                	utilito.LevelLog(Config_env_log, "3", "Eror en generando response")
                    errorGeneral= err.Error()
                }
		    //  END
        }

	    if errorGeneral!="" && errorGeneralNbr=="" {//continue next step
	    	utilito.LevelLog(Config_env_log, "3", "CZ   prepare the JSON response for ERROR")

		    //  START 
		    errorGeneral="ERROR:130 -Error preparing the response"	+errorGeneral
			errorGeneralNbr="130"
		    //  END
        }
 utilito.LevelLog(Config_env_log, "3", "CZ  END   handler Listening DB  realizarpago  2")
     return errorGeneral, errorGeneralNbr
}
