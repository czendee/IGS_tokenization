package main

import (
	"net/http"
    "strings"
    "strconv"
	//"log"
    utilito "banwire/services/file_tokenizer/util"
	"banwire/services/file_tokenizer/db"
	modelito "banwire/services/file_tokenizer/model"
//	"time"
//	"encoding/json"
	 _ "github.com/lib/pq"   //use go get github.com/lib/pq
    "io"
    "bytes"

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
    	errorGeneral="ERROR_300 -Missing parameter"	+errorGeneral
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
    	errorGeneral="ERROR_310 -  Error obtaining cards -"	+errorGeneral
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
	    errorGeneral="ERROR_330 -Error preparing the response"	+errorGeneral
	    errorGeneralNbr="330"
	    //  END
    }

     return errorGeneral, errorGeneralNbr
}



// Generatetokenized for receive and handle the request from client
func ProcessGeneratetokenized(w http.ResponseWriter, requestData modelito.RequestTokenized) (string,string,modelito.Card) {
	defer func() {
		db.Connection.Close(nil)
	}()
	  var result string

     var errorGeneral string
     var errorGeneralNbr string
     
     var resultCardTokenized modelito.Card
     
     var obtainedDataWebservice modelito.ExitoDataTokenized

     var resultadoTokenSingle modelito.Card
     //var resultadoToken modelito.ExitoData
     
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
    	errorGeneral="ERROR_200 -Missing parameter "	+errorGeneral
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
    	errorGeneral="ERROR_210 -Error processing payment:"	+errorGeneral
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
    	errorGeneral="ERROR_220 -Error generating token:"	+errorGeneral
		errorGeneralNbr="220"
    }

	//response
    if errorGeneral==""{//continue next step
		utilito.LevelLog(Config_env_log, "3", "CZ   STEP Post the response JSON ready")
		
			/// START
/*            		fieldDataBytesJsonTokenize,err := getJsonResponseTokenizeV2(resultCardTokenized) // logicresponse.go
			
		utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test realizarpago  3")	

		utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test handleGeneratetokenized  4"+"<html><body>"+ result+"</body></html>")
        if err!=nil{
        	utilito.LevelLog(Config_env_log, "3", "Eror en generando response")
	        errorGeneral= err.Error()
        }else{
            errorGeneral =string(fieldDataBytesJsonTokenize);
        }
*/				
		// END
        resultadoTokenSingle = resultCardTokenized
	}	

    if errorGeneral!="" && errorGeneralNbr==""{
    	//prepare response with error 230
    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 230. Error generating Response Tokenized:"+errorGeneral)
    	errorGeneral="ERROR_230 -Error generating Response Tokenized:"	+errorGeneral
		errorGeneralNbr="230"
    }
    
	 utilito.LevelLog(Config_env_log, "3", "CZ  ends func tokenized")
	 
	return errorGeneral, errorGeneralNbr, resultadoTokenSingle
}





func GetCardType(number string) string {
	return "VISA"
}


/////////////////////////v4
/////////////////////////v4

// v4Processpayment  receive and handle the request from client, access DB
func v4ProcessProcessPayment(w http.ResponseWriter, requestData modelito.RequestPayment) (string,string,modelito.ExitoData){
	defer func() {
		db.Connection.Close(nil)
	}()
    var result string
    var errorGeneral string
    var	errorGeneralNbr string
    var resultadoPaymentSingle modelito.ExitoData

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
    	errorGeneral="ERROR_100 - Missing parameter"	+errorGeneral
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
	    	errorGeneral="ERROR_110 - Error processing payment"	+errorGeneral
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
	    	errorGeneral="ERROR_ 120 - Error recording results in DB"	+errorGeneral
			errorGeneralNbr="120"
	    }

    		    
	//response
	////////////////////////////////////////////////http response	
	//      prepare the JSON response
	//	    
	    if errorGeneral==""{//continue next step
	    	utilito.LevelLog(Config_env_log, "3", "CZ   STEP  prepare the JSON response for SUCCESS")

		    //  START 

/*		    fieldDataBytesJsonPayment,err := getJsonResponsePaymentV2(resultadoPayment)					
		        w.Header().Set("Content-Type", "application/json")
		        w.Write(fieldDataBytesJsonPayment)
				utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test handleProcesspayment  4"+"<html><body>"+ result+"</body></html>")
                if err!=nil{
                	utilito.LevelLog(Config_env_log, "3", "Eror en generando response")
                    errorGeneral= err.Error()
                }

*/                
		    //  END
            resultadoPaymentSingle = resultadoPayment
        }

	    if errorGeneral!="" && errorGeneralNbr=="" {//continue next step
	    	utilito.LevelLog(Config_env_log, "3", "CZ   prepare the JSON response for ERROR")

		    //  START 
		    errorGeneral="ERROR_130 -Error preparing the response"	+errorGeneral
			errorGeneralNbr="130"
		    //  END
        }
 utilito.LevelLog(Config_env_log, "3", "CZ  END   handler Listening DB  realizarpago  2")
     return errorGeneral, errorGeneralNbr, resultadoPaymentSingle
}

///---------------------------------------- File validations and file processing


func validateFiles(typeFile string, r *http.Request) ( string, string, []modelito.ExitoDataValidaLine,[]modelito.RequestTokenized,[]modelito.RequestPayment) {
    var errorGeneral string
    var errorGeneralNbr string


    errorGeneral=""

    linesStatus := []modelito.ExitoDataValidaLine{}   //structure to stire the errors in each of the liens of the file

    linesDataTokens := []modelito.RequestTokenized{}   //structure to store the data for all the tokens (all the liens of the file   )     
    linesDataPayments := []modelito.RequestPayment{}   //structure to store the data for all the payment (all the liens of the file   )     

    //start  logic
    
    if errorGeneral=="" {

        utilito.LevelLog(Config_env_log, "3", "CZ   STEP Get the File")
        utilito.LevelLog(Config_env_log, "3", "File Upload Endpoint Hit")

        // Parse our multipart form, 10 << 20 specifies a maximum
        // upload of 10 MB files.
        err:= r.ParseMultipartForm(10 << 20)
        if err != nil {
            
            utilito.LevelLog(Config_env_log, "3", "CZ Error Retrieving the File")
            utilito.LevelLog(Config_env_log, "3",  err.Error())
            errorGeneral="ERROR_110 -Error retriving files ,parameters"	+errorGeneral
            errorGeneralNbr="110"

        } 
        if errorGeneral=="" {  
            utilito.LevelLog(Config_env_log, "3", "CZ Start read the form data")
            formdata := r.MultipartForm // ok, no problem so far, read the Form data

            //get the *fileheaders
            files := formdata.File["file0"] // grab the files, this files was set in the html 

//            midescripcioncita := r.FormValue("description")
            utilito.LevelLog(Config_env_log, "3", "CZ before loop files")

            for i, _ := range files { // loop through the files one by one
               
                utilito.LevelLog(Config_env_log, "3", "CZ loop step 1")
                var elfileindex string
                
                elfileindex = "file0"
                utilito.LevelLog(Config_env_log, "3", "CZ Loop file")
                utilito.LevelLog(Config_env_log, "3", "CZ Loop file:"+elfileindex)
                file, err := files[i].Open()
                utilito.LevelLog(Config_env_log, "3", "CZ open file")
                defer file.Close()
                if err != nil {
                    utilito.LevelLog(Config_env_log, "3",  err.Error())
                    errorGeneral="ERROR_120 -Error file passed not open ,parameters"	+errorGeneral
                    errorGeneralNbr="120"

                }
                //convert multipart file into buffer bytes

                buf := bytes.NewBuffer(nil)
                io.Copy(buf, file)

                micadenita := buf.String()

                utilito.LevelLog(Config_env_log, "4", micadenita)

                utilito.LevelLog(Config_env_log, "3", "MGR Paso linea por linea index")

                lineas := 0

                lineasWithErrors := 0
                for _, line := range strings.Split(strings.TrimSuffix(micadenita, "\n"), "\n") {
                    var u modelito.ExitoDataValidaLine
                    if lineas >= 1{
                        utilito.LevelLog(Config_env_log, "3", "MGR linea de datos")

                        lineas = lineas + 1
                        var respuestaRes string
                        var cualfallo int
                         eachLineaDataToken :=   modelito.RequestTokenized{}
                        eachLineaDataPayment :=   modelito.RequestPayment{}
                        if typeFile =="token" {
                            eachLineaDataToken,respuestaRes,cualfallo =validateAndObtainCampos_token (line, lineas)  //logicrequest.go
                        }

                        if typeFile =="payment" {
                            eachLineaDataPayment, respuestaRes,cualfallo =validateAndObtainCampos_payment (line, lineas)
                        }                        

                         if cualfallo ==0 {  //exito, todos los cmapos de la linea OK, y no errores previos
                            u.Line=strconv.Itoa(lineas)
                            u.Status="OK"
                            u.StatusMessage ="SUCESS"
                            //the dataToken has the data for the line to be tokenized
                            //the dataPaymentn has the data for the line to do the payments

                         }else { //error, al menos un error en la linea
                            u.Line=strconv.Itoa(lineas)
                            u.Status="ERROR540"
                            u.StatusMessage ="ERROR FIELD_"+strconv.Itoa(cualfallo)+" - "+respuestaRes
                            lineasWithErrors =1
                            //the dataToken is set to "" when errors
                         }
                         linesStatus = append(linesStatus,u);
                         linesDataTokens = append(linesDataTokens,eachLineaDataToken);
                         linesDataPayments = append(linesDataPayments,eachLineaDataPayment);
                    }
                    
                    if lineas == 0 {
                        utilito.LevelLog(Config_env_log, "3", "MGR Nombres de campos")
                        lineas = lineas + 1
                    }

                        utilito.LevelLog(Config_env_log, "3", line)
                       
                        

                }//end  -loop through the lines

               if  lineasWithErrors ==1 { //al menos una linea  tuvo un error
                     errorGeneral="ERROR FILE"
                     errorGeneralNbr="540"
                   
               }else{
                    if errorGeneral=="" {  
    
                        utilito.LevelLog(Config_env_log, "3", "Files uploaded successfully : ")
                        utilito.LevelLog(Config_env_log, "3", files[i].Filename+"\n")

                    }    

               }
                //1.count number of lines in the file received
                //2.for each line 
                //      validate the content
                //3.store in the db table AUDIT FILE VALIDATION
                //         seq nbr ,  file name, size, content[all the bytes],commentsParam, validationStatus, validationStatusMessage,validationResponse[for each line,a response OK/Error] ,timestamp
                //3.return result JSON


                

            }//end for - all the files received

        }//end if    
        //errorGeneral,errorGeneralNbr= ProcessGettokenizedcards(w , requestData) //logicbusiness.go
	}

    //end   logic
    
    return errorGeneral,errorGeneralNbr ,linesStatus, linesDataTokens,linesDataPayments
    //return errorGeneral,errorGeneralNbr ,linesStatus
}



func ProcessGetPaymentsForToken(w http.ResponseWriter,  paramInput string) (string,string,[]modelito.Payment) {

    var errorGeneral string
    var errorGeneralNbr string
//    	var result string
   var valoresParaResponder  []modelito.Payment

    errorGeneral=""

		        
    if errorGeneral!="" && errorGeneralNbr=="" {
    	//prepare response with error 502
    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 300. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR_502 -Missing parameter"	+errorGeneral
    	errorGeneralNbr="502"
    }

	////////////////////////////////////////////////DB	
	//	    resultado,errfetchDB:= fetchFromDB ()
	if errorGeneral==""{//continue next step

       	    utilito.LevelLog(Config_env_log, "3", "CZ   STEP Consume DB")
         valoresParaResponder,errorGeneral =logicDBGetPaymentsByToken( errorGeneral,paramInput)  //logicdb.go


    }				    
    if errorGeneral!="" && errorGeneralNbr==""{
    	//prepare response with error 510
    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 510. Error obtaining payments by token:"+errorGeneral)
    	errorGeneral="ERROR_510 -  Error obtaining payments by token -"	+errorGeneral
	    errorGeneralNbr="510"
    }

		/// END

		 

     return errorGeneral, errorGeneralNbr,valoresParaResponder
}



func ProcessGetTokensForCustRef(w http.ResponseWriter,  paramInput string) (string,string,[]modelito.Card) {

    var errorGeneral string
    var errorGeneralNbr string
//    	var result string
   var valoresParaResponder  []modelito.Card

    errorGeneral=""

		        
    if errorGeneral!="" && errorGeneralNbr=="" {
    	//prepare response with error 602
    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 602. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR_602 -Missing parameter"	+errorGeneral
    	errorGeneralNbr="602"
    }

	////////////////////////////////////////////////DB	
	//	    resultado,errfetchDB:= fetchFromDB ()
	if errorGeneral==""{//continue next step

       	    utilito.LevelLog(Config_env_log, "3", "CZ   STEP Consume DB")
         valoresParaResponder,errorGeneral =logicDBGetTokensByCustRef( errorGeneral,paramInput)  //logicdb.go


    }				    
    if errorGeneral!="" && errorGeneralNbr==""{
    	//prepare response with error 510
    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 610. Error obtaining tokens by cust ref:"+errorGeneral)
    	errorGeneral="ERROR_610 -  Error obtaining tokens by cust ref -"	+errorGeneral
	    errorGeneralNbr="610"
    }

		/// END

		 

     return errorGeneral, errorGeneralNbr,valoresParaResponder
}
