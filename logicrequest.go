package main

import (
	"net/http"
	//"log"
    utilito "banwire/services/file_tokenizer/util"
	"encoding/json"
	modelito "banwire/services/file_tokenizer/model"
	 _ "github.com/lib/pq"   //use go get github.com/lib/pq
	 
)

///////////////7//get


   func obtainParmsGettokenizedcards(r *http.Request, errorGeneral string )(modelito.RequestTokenizedCards, string){
   	var requestData modelito.RequestTokenizedCards
	////////////////////////////////////////////////obtain parms in JSON
   //START    
    utilito.LevelLog(Config_env_log, "3", "cz  handleDBGettokenizedcards")

	    utilito.LevelLog(Config_env_log, "3", "CZ    handlerDB Listening test obtienetarjetastokenizadas")
	    
    	err := r.ParseForm()
		if err != nil {
	    	//prepare response with error 100
	    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 380. Missing parameter:"+errorGeneral)
	    	errorGeneral="ERROR:380 -"	+err.Error()
		}
		v := r.Form
		requestData.Cardreference = v.Get("cardreference")

    //END
   	
   	 return  requestData, errorGeneral
   }



   func obtainParmsProcessPayment(r *http.Request, errorGeneral string) (modelito.RequestPayment,string){
   	 var requestData modelito.RequestPayment
	////////////////////////////////////////////////obtain parms in JSON
   //START    
    utilito.LevelLog(Config_env_log, "3", "cz  handleProcesspayment")
 		 utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test realizarpago")
    		    
    	err := r.ParseForm()
		if err != nil {
	    	//prepare response with error 100
	    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 180. Missing parameter:"+errorGeneral)
	    	errorGeneral="ERROR:180 -"	+err.Error()
		}
		v := r.Form
		requestData.Clientreference = v.Get("clientreference")
		requestData.Paymentreference = v.Get("paymentreference")
		requestData.Token = v.Get("token")
		requestData.Cvv = v.Get("cvv")
		requestData.Amount = v.Get("amount")

   //END
   	 
   	 return requestData,errorGeneral
   }

   func obtainParmsGeneratetokenized(r *http.Request, errorGeneral string) (modelito.RequestTokenized,string) {
	////////////////////////////////////////////////obtain parms in JSON
   //START    
     var requestData modelito.RequestTokenized
    utilito.LevelLog(Config_env_log, "3", "cz  handleGeneratetokenized")
	    utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test handleGeneratetokenized")
    		    
    	err := r.ParseForm()
		if err != nil {
	    	//prepare response with error 280
	    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 280. Error in JSON Request:"+errorGeneral)
	    	errorGeneral="ERROR :280 -Error in JSON Request-"	+err.Error()
		}
		v := r.Form
		requestData.Clientreference = v.Get("clientreference")
		requestData.Paymentreference = v.Get("paymentreference")
		requestData.Card = v.Get("card")
		requestData.Exp = v.Get("exp")
		requestData.Cvv = v.Get("cvv")

   //END
   	  return  requestData, errorGeneral
   }

////////////////////////Post



   func obtainPostParmsGettokenizedcards(r *http.Request, errorGeneral string )(modelito.RequestTokenizedCards, string){
   	var requestData modelito.RequestTokenizedCards
	////////////////////////////////////////////////obtain parms in JSON
   //START    
    utilito.LevelLog(Config_env_log, "3", "cz  handleDBGettokenizedcards")

	    utilito.LevelLog(Config_env_log, "3", "CZ    handlerDB Listening test obtienetarjetastokenizadas")
	 
 			decoder := json.NewDecoder(r.Body)
		
			err := decoder.Decode(&requestData)
			if err != nil {
		    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 380. JSON format/Missing parameter:"+errorGeneral)
		    	errorGeneral="ERROR:380 -Input JSON format/Missing parameter"	+err.Error()

			}
		
			//post   cardreference := requestData.Cardreference

    //END
   	
   	 return  requestData, errorGeneral
   }



   func obtainPostParmsProcessPayment(r *http.Request, errorGeneral string) (modelito.RequestPayment,string){
   	 var requestData modelito.RequestPayment
	////////////////////////////////////////////////obtain parms in JSON
   //START    
    utilito.LevelLog(Config_env_log, "3", "cz  handleProcesspayment")
 		utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test realizarpago")
 			decoder := json.NewDecoder(r.Body)
		
			err := decoder.Decode(&requestData)
			if err != nil {
		    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 180. JSON format/Missing parameter:"+errorGeneral)
		    	errorGeneral="ERROR:180 -Input JSON format/Missing parameter"	+err.Error()

			}

   //END
   	 
   	 return requestData,errorGeneral
   }

   func obtainPostParmsGeneratetokenized(r *http.Request, errorGeneral string) (modelito.RequestTokenized,string) {
	////////////////////////////////////////////////obtain parms in JSON
   //START    
     var requestData modelito.RequestTokenized
    utilito.LevelLog(Config_env_log, "3", "cz  handleGeneratetokenized")
	    utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test handleGeneratetokenized")
 			decoder := json.NewDecoder(r.Body)
		
			err := decoder.Decode(&requestData)
			if err != nil {
		    	utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 280. JSON format/Missing parameter:"+errorGeneral)
		    	errorGeneral="ERROR:280 -Input JSON format/Missing parameter"	+err.Error()

			}

   //END
   	  return  requestData, errorGeneral
   }


////////////////////////validate input params

	    func validaReqProcessPayment( parRequestData modelito.RequestPayment) string {
            var resultado string
            
            	if parRequestData.Clientreference != "" {
	            	if len(parRequestData.Clientreference)>100 {
	
						resultado="Client reference is required"
			        }
				}else{
					resultado="Client reference is required"
		        }

				if parRequestData.Paymentreference != "" {
					if len(parRequestData.Paymentreference) >100 {
	
						resultado="Payment reference max lenght is 100"
			        }

				}else{
					resultado="Payment reference is required"
		        }

				if parRequestData.Token != "" {

				}else{
					resultado="Token is required"
		        }

				if parRequestData.Cvv != "" {
					if len(parRequestData.Cvv)==3 ||  len(parRequestData.Cvv)==4 {
	
					}else{
						resultado="Cvv must be 3 or 4 digits"
			        }

				}else{
					resultado="Cvv is required"
		        }
				if parRequestData.Amount != "" {

				}else{
					resultado="Amount is required"
		        }
            //lenght




		        
			/// END

            return resultado
	    }


 
 	    func validaReqGenerateTokenized( parRequestData modelito.RequestTokenized) string {
            var resultado string
            
            	if parRequestData.Clientreference != "" {
					if len(parRequestData.Paymentreference) >100 {
	
						resultado="Customer reference max lenght is 100"
			        }
				}else{
					resultado="Client reference is required"
		        }

				if parRequestData.Paymentreference != "" {
					if len(parRequestData.Paymentreference) >100 {
	
						resultado="Payment reference max lenght is 100"
			        }
				}else{
					resultado="Payment reference is required"
		        }

				if parRequestData.Card != "" {
					if len(parRequestData.Card)==16 || len(parRequestData.Card)==15{
	
					}else{
						resultado="Card Number must be 16 digits"
			        }
				}else{
					resultado="Card is required"
		        }

				if parRequestData.Exp != "" {
					if  len(parRequestData.Exp)==4 {
	
					}else{
						resultado="Valid Thru  4 digits"
			        }
				}else{
					resultado="Valid Thru is required"
		        }
/*				if parRequestData.Cvv != "" {
					if len(parRequestData.Cvv)==3 ||  len(parRequestData.Cvv)==4 {
	
					}else{
						resultado="Cvv must be 3 or 4 digits"
			        }
				}else{
					resultado="Cvv is required"
		        }
*/
			/// END

            return resultado
	    }

 	    func validaReqFetchCards( parRequestData modelito.RequestTokenizedCards) string {
            var resultado string
            
            	if parRequestData.Cardreference != "" {

				}else{
					resultado="Card reference is required"
		        }

			/// END

            return resultado
	    }
