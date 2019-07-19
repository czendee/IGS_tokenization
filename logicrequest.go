package main

import (
	"net/http"
	"log"
    "strings"
    "strconv"
    "regexp"
    //"io"
    //"bytes"
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

	utilito.LevelLog(Config_env_log, "3", "CZ   handlerDB Listening test obtienetarjetastokenizadas")
	    
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
}//end obtainParmsGettokenizedcards

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
} //end obtainParmsProcessPayment

func obtainParmsProcessDownload(r *http.Request, errorGeneral string) (string, string){
    
    log.Print("Entra funcion obtainParmsProcessDownload")
    var requestData string
	////////////////////////////////////////////////obtain parms in JSON
    //START    
    utilito.LevelLog(Config_env_log, "3", "cz  handleProcesspayment")
 	utilito.LevelLog(Config_env_log, "3", "CZ   handler Listening test realizarpago")
    		    
    err := r.ParseForm()
	if err != nil {
        //prepare response with error 100
	    utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 180. Missing parameter:"+errorGeneral)
	    errorGeneral="ERROR:180 -"	+err.Error()
	}
	
    v := r.Form
	requestData = v.Get("contenidofile")

    var lineaDatos string
    micadenita := requestData

    utilito.LevelLog(Config_env_log, "3", "Respuesta validacion")
    utilito.LevelLog(Config_env_log, "3", micadenita)

    cadenalimpia :=  strings.Replace(micadenita, "{", "", -1)

    for _, linea := range strings.Split(strings.TrimSuffix(cadenalimpia, "}"), "}"){
        utilito.LevelLog(Config_env_log, "3", "linea")
            
        linealimpia :=  strings.Replace(linea, " ", "", -1)
        utilito.LevelLog(Config_env_log, "3", linealimpia)

        for _, campo := range strings.Split(strings.TrimSuffix(linealimpia, ","), ","){
            utilito.LevelLog(Config_env_log, "3", "Campo")
            utilito.LevelLog(Config_env_log, "3", campo)
            lineaDatos = lineaDatos + campo +","
        } // end for campo

        lineaDatos = lineaDatos +"\r\n"
        utilito.LevelLog(Config_env_log, "3", lineaDatos)
    } //end for linea
       
    //END
   	 
   	return lineaDatos,errorGeneral
} //end obtainParmsProcessDownload

func obtainParmsProcessDownloadPagos(r *http.Request, errorGeneral string) (string, string){
    
    log.Print("Entra funcion obtainParmsProcessDownloadPagos")
    var requestData string
	////////////////////////////////////////////////obtain parms in JSON
    //START    
    
    utilito.LevelLog(Config_env_log, "3", "cz  handleProcesspayment")
 	utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test respuestaPagos")
    		    
    err := r.ParseForm()
	if err != nil {
	    //prepare response with error 100
	    utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 180. Missing parameter:"+errorGeneral)
	    errorGeneral="ERROR:180 -"	+err.Error()
	}
		
    v := r.Form
    log.Print("v: ",v)
	requestData = v.Get("contenidofilePagos")
        
    var lineaDatos string
    micadenita := requestData

    utilito.LevelLog(Config_env_log, "3", "Respuesta pagos")
    utilito.LevelLog(Config_env_log, "3", micadenita)

    cadenalimpia :=  strings.Replace(micadenita, "{", "", -1)

    for _, linea := range strings.Split(strings.TrimSuffix(cadenalimpia, "}"), "}"){
        
        utilito.LevelLog(Config_env_log, "3", "linea")

        linealimpia := strings.Replace(linea, " ", "", -1)
        utilito.LevelLog(Config_env_log, "3", linealimpia)

        for _, campo := range strings.Split(strings.TrimSuffix(linealimpia, ","), ","){
            utilito.LevelLog(Config_env_log, "3", "Campo:")
            utilito.LevelLog(Config_env_log, "3", campo)

            cuenta := 0
            for _, dato := range strings.Split(campo, "\":"){
                //log.Print("dato ",dato)
                cuenta = cuenta + 1
                //log.Print("cuenta_",cuenta)
                if cuenta == 2{
                    datolimpio := strings.Replace(dato, ":", "_", -1)
                    log.Print("Datolimpio: ", datolimpio)
                    lineaDatos = lineaDatos + datolimpio
                }//end if cuenta
     
            }// end for dato
            lineaDatos = lineaDatos + ","
            
        }//end for campo

        log.Print("lineaDatos: ",lineaDatos)
        
        lineaDatos = lineaDatos +"\r\n"
    }//end for linea

    //parte := strings.Split(strings.TrimSuffix(micadenita, "["), "[")
    //mensajes := strings.Split(strings.TrimSuffix(parte[0], ","), ",")
    //utilito.LevelLog(Config_env_log, "3", "mensajes status")
    //utilito.LevelLog(Config_env_log, "3", mensajes[0])
    //limpiar :=  strings.Replace(mensajes[0], "\"", "", -1)
    //limpiar2 :=  strings.Replace(limpiar, " ", "", -1)
    //campoStatus := strings.Split(limpiar2, ":")
    //status_message := campoStatus[1]
    //if status_message == "Success"{
        //utilito.LevelLog(Config_env_log, "3", status_message)
        //limpiar =  strings.Replace(mensajes[2], "\"", "", -1)
        //limpiar2 =  strings.Replace(limpiar, " ", "", -1)
        //payments := strings.Split(limpiar2, ":")
        //log.Print("Payments "+ payments[1])
        //cuenta_i := 0
        //limpiar =  strings.Replace(parte[1], "\n", "", -1)
        //limpiar2 =  strings.Replace(limpiar, " ", "", -1)
        //for _, line := range strings.Split(limpiar2, "},"){
            //utilito.LevelLog(Config_env_log, "3", "For linea")
            //utilito.LevelLog(Config_env_log, "3", line)
            //cuenta_i = cuenta_i + 1
            //log.Print("no vuelta",cuenta_i)
            //for _, campo := range strings.Split(strings.TrimSuffix(line, ","), ","){
                            
                //utilito.LevelLog(Config_env_log, "3", "campo")
                //utilito.LevelLog(Config_env_log, "3", campo)
                //limpia2 := strings.Replace(campo, " ", "", -1)
                //limpia3 := strings.Replace(limpia2, "}", "", -1)
                //limpia4 := strings.Replace(limpia3, "]", "", -1)
                //log.Print("datolimpio4", limpia4)
                //dato := strings.Split(limpia4, ":")
                //lineaDatos = lineaDatos + dato[1] +","
            //}// end for campo

            //log.Print("fuera de for campo")
            //cuenta_i = 0
            //utilito.LevelLog(Config_env_log, "3", lineaDatos)
            //lineaDatos = lineaDatos +"\r\n"
            //}// end for linea
                
            //compara, err := strconv.Atoi(payments[1])
            //if err == nil {

            //}
            //if compara != cuenta_i {
            //    log.Print("ERROR 2048 payments don´t match processed payments")
            //    lineaDatos = "ERROR 2048 payments don´t match processed payments"
            //}else{
            //    log.Print("Archivo Success")
            //}
                
            //log.Print("fuera de for linea")
            //log.Print("Registros procesados: ",cuenta_i)
            //log.Print("Registros correctos: "+payments[1])
            //compara, err := strconv.Atoi(payments[1])
            //if err == nil {
                //log.Print("Registros con error: ",cuenta_i-compara)
            //}//end err
        //}else{
            //log.Print("ERROR 2024 missing parameter")
            //lineaDatos = "ERROR 2024 missing parameter"
        //} //end else-if status_message
        //log.Print("fin if status_message")

        //buscar si en la cadena de caracteres esta status_message y ver si es Success
        //buscar si en la cadena de caracteres cards_tokenized, y ponerle el valor que traiga
        //buscar si esta la cadena de caracteres card data, y buscar los datos entre los corchetes cuadrados y ponerlo en una cadena resto
        //para cada elemento que termine en corche y coma   
            //procesarlo  y dejar el resto en resto  
        //buscar en la cadena resto si hay un corchete y una coma y separas esa parte en otra variable
        //se repite la acción anterior hasta solo entrar un corchete solo que indica el final de los datos
        //cuando solo encuentra un corchete solo se procesa la información
        //el ciclo termina cuando hay un solo corchete
       
   //END
   	 
   	 return lineaDatos,errorGeneral
        
} //end obtainParmsProcessDownloadPagos

func obtainParmsProcessDownloadValida(r *http.Request, errorGeneral string) (string, string){
    
    log.Print("Entra funcion obtainParmsProcessDownloadValida")
    var requestData string
	////////////////////////////////////////////////obtain parms in JSON
    //START    
    utilito.LevelLog(Config_env_log, "3", "cz  handleProcesspayment")
 	utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test respuestaValidacion")
    		    
    err := r.ParseForm()
	if err != nil {
	    //prepare response with error 100
	    utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 180. Missing parameter:"+errorGeneral)
	    errorGeneral="ERROR:180 -"	+err.Error()
	}

	v := r.Form
	requestData = v.Get("contenidofileValida")

    var lineaDatos string
    micadenita := requestData

    utilito.LevelLog(Config_env_log, "3", "Respuesta validacion")
    utilito.LevelLog(Config_env_log, "3", micadenita)

    cadenalimpia :=  strings.Replace(micadenita, "{", "", -1)

    for _, linea := range strings.Split(strings.TrimSuffix(cadenalimpia, "}"), "}"){
        utilito.LevelLog(Config_env_log, "3", "linea")
            
        linealimpia :=  strings.Replace(linea, " ", "", -1)
        utilito.LevelLog(Config_env_log, "3", linealimpia)

        for _, campo := range strings.Split(strings.TrimSuffix(linealimpia, ","), ","){
            utilito.LevelLog(Config_env_log, "3", "Campo")
            utilito.LevelLog(Config_env_log, "3", campo)
            lineaDatos = lineaDatos + campo +","
        } // end for campo

        lineaDatos = lineaDatos +"\r\n"
        utilito.LevelLog(Config_env_log, "3", lineaDatos)
    } //end for linea

    //END
   	 
   	return lineaDatos,errorGeneral

} //end obtainParmsProcessDownloadValida

func obtainParmsProcessDownloadTokeniza(r *http.Request, errorGeneral string) (string, string){
    
    log.Print("Entra funcion obtainParmsProcessDownloadTokeniza")
    var requestData string
	////////////////////////////////////////////////obtain parms in JSON
    //START    
    utilito.LevelLog(Config_env_log, "3", "cz  handleProcesspayment")
 	utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test respuestaTokenizacion")
    	    
    err := r.ParseForm()
	if err != nil {
	    //prepare response with error 100
	    utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 180. Missing parameter:"+errorGeneral)
	    errorGeneral="ERROR:180 -"	+err.Error()
	}
		
    v := r.Form
	requestData = v.Get("contenidofileTokeniza")
        
    var lineaDatos string
    micadenita := requestData

    utilito.LevelLog(Config_env_log, "3", "Respuesta pagos")
    utilito.LevelLog(Config_env_log, "3", micadenita)

    cadenalimpia :=  strings.Replace(micadenita, "{", "", -1)

    for _, linea := range strings.Split(strings.TrimSuffix(cadenalimpia, "}"), "}"){
        
        utilito.LevelLog(Config_env_log, "3", "linea")

        linealimpia := strings.Replace(linea, " ", "", -1)
        utilito.LevelLog(Config_env_log, "3", linealimpia)

        for _, campo := range strings.Split(strings.TrimSuffix(linealimpia, ","), ","){
            utilito.LevelLog(Config_env_log, "3", "Campo:")
            utilito.LevelLog(Config_env_log, "3", campo)

            cuenta := 0
            for _, dato := range strings.Split(campo, "\":"){
                //log.Print("dato ",dato)
                cuenta = cuenta + 1
                //log.Print("cuenta_",cuenta)
                if cuenta == 2{
                    datolimpio := strings.Replace(dato, ":", "_", -1)
                    log.Print("Datolimpio: ", datolimpio)
                    lineaDatos = lineaDatos + datolimpio
                }//end if cuenta
     
            }// end for dato
            lineaDatos = lineaDatos + ","
            
        }//end for campo

        log.Print("lineaDatos: ",lineaDatos)
        
        lineaDatos = lineaDatos +"\r\n"
    }//end for linea

       
    //END
   	 
   	return lineaDatos,errorGeneral
} //end obtainParmsProcessDownloadTokeniza


func obtainParmsConsultarTokens(r *http.Request, errorGeneral string) (string, string){
    
    log.Print("Entra a obtainParmsConsultarTokens GET")
    var requestData string
	////////////////////////////////////////////////obtain parms in JSON
    //START    
    utilito.LevelLog(Config_env_log, "3", "cz  handleConsultarTokensGET ")
 	utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test ConsultarTokens GET")
    	    
    err := r.ParseForm()
	if err != nil {
	    //prepare response with error 100
	    utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 180. Missing parameter:"+errorGeneral)
	    errorGeneral="ERROR:180 -"	+err.Error()
	}
		
    v := r.Form
    
	requestData = v.Get("contenidofileconsultaTok")
    

    var lineaDatos string
    micadenita := requestData

    utilito.LevelLog(Config_env_log, "3", "Respuesta ConsultarToken GET:")
    utilito.LevelLog(Config_env_log, "3", micadenita)

    //END
   	 lineaDatos= micadenita

   	return lineaDatos,errorGeneral
} //end obtainParmsConsultarTokens

func obtainParmsConsultarHistPagoTokens(r *http.Request, errorGeneral string) (string, string){
    
    log.Print("Entra a obtainParmsConsultarHistPagoTokens")
    var requestData string
	////////////////////////////////////////////////obtain parms in JSON
    //START    
    utilito.LevelLog(Config_env_log, "3", "cz  handleConsultarHistPagoTokens")
 	utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test ConsultarHistPagoTokens")
    	    
    err := r.ParseForm()
	if err != nil {
	    //prepare response with error 100
	    utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 180. Missing parameter:"+errorGeneral)
	    errorGeneral="ERROR:180 -"	+err.Error()
	}
		
    v := r.Form
    
	requestData = v.Get("contenidohitorialTokenk")
    
        
    var lineaDatos string
    micadenita := requestData

    utilito.LevelLog(Config_env_log, "3", "Respuesta Consultar historial Token")
    utilito.LevelLog(Config_env_log, "3", micadenita)
/*
    cadenalimpia :=  strings.Replace(micadenita, "{", "", -1)

    for _, linea := range strings.Split(strings.TrimSuffix(cadenalimpia, "}"), "}"){
        utilito.LevelLog(Config_env_log, "3", "linea")
            
        linealimpia :=  strings.Replace(linea, " ", "", -1)
            
        utilito.LevelLog(Config_env_log, "3", linealimpia)

        for _, campo := range strings.Split(strings.TrimSuffix(linealimpia, ","), ","){
            utilito.LevelLog(Config_env_log, "3", "Campo")
            utilito.LevelLog(Config_env_log, "3", campo)
            dato := strings.Split(campo, ":")
            lineaDatos = lineaDatos + dato[1] +","
        } // end for campo

        lineaDatos = lineaDatos +"\r\n"
        utilito.LevelLog(Config_env_log, "3", lineaDatos)
    } //end for linea
    
  */

//   	 lineaDatos= "ESTO LLEGO: "+micadenita
       lineaDatos= micadenita     
    //END
   	 
   	return lineaDatos,errorGeneral
} //end obtainParmsConsultarHistPagoTokens
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
 }//end obtainParmsGeneratetokenized

////////////////////////Post
func obtainPostParmsGettokenizedcards(r *http.Request, errorGeneral string )(modelito.RequestTokenizedCards, string){
    log.Print("entra a obtainPostParmsGettokenizedcards")
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
}//end obtainPostParmsGettokenizedcards

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
}//end obtainPostParmsProcessPayment

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
}//end obtainPostParmsGeneratetokenized

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
}//end obtainPostParmsGeneratetokenized

func validaReqGenerateTokenized( parRequestData modelito.RequestTokenized) string {
    
    var resultado string
            
            	if parRequestData.Clientreference != "" {
					if len(parRequestData.Clientreference) > 30 {
	
						resultado="Client reference max lenght is 30"
			        }
				}else{
					resultado="Client reference is required"
		        }

				if parRequestData.Paymentreference != "" {
					if len(parRequestData.Paymentreference) > 50 {
	
						resultado="Payment reference max lenght is 50"
			        }
				}else{
					resultado="Payment reference is required"
		        }

				if parRequestData.Card != "" {
					if len(parRequestData.Card)==16 {
	
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
}//end validaReqGenerateTokenized

func validaReqFetchCards( parRequestData modelito.RequestTokenizedCards) string {
    
    var resultado string
            
        if parRequestData.Cardreference != "" {

		}else{
		    resultado="Card reference is required"
		}

		/// END

        return resultado

}//end validaReqFetchCards


//// File_tokenizer   solution needs these to validate the input of the lines received.
//// File_tokenizer   solution needs these to validate the input of the lines received.
///START

//Función campos_token
func campos_token (line string, lineas int)(string, int){
        
    utilito.LevelLog(Config_env_log, "3", "MGR campo por campo")
    numcampos := 0
    var resultado string
    var cualfallo int
    resultado ="OK"
    cualfallo =0
    for _, campo := range strings.Split(strings.TrimSuffix(line, ","), ","){
        
        numcampos = numcampos + 1

        var campoValue string
        
        //campoValue = strings.Replace(campo, "\n", "", -1) // only works with a single character
        //re := regexp.MustCompile(`\r?\n`)
        //campoValue = re.ReplaceAllString(campoValue, "y")

        campoValue = strings.Replace(campo, "\"", "", -1) // only works with a single character
        var largo string
        largo = strconv.Itoa ( len(campoValue))
        utilito.LevelLog(Config_env_log, "3", "largo del campo es:"+largo+":valor del campo es:"+campoValue)
              
        resultado, cualfallo = valida_campo_token(campoValue, numcampos)//logicrequest.go

        if cualfallo >0 {
            
            utilito.LevelLog(Config_env_log, "3", "fallo es valor en :"+campo)
            break
        }
    }//end for campo
        
    return resultado, cualfallo

} //end campos_token

//Función valida_campo_token
func valida_campo_token (campo string, numcampos int)(string, int){
    
    utilito.LevelLog(Config_env_log, "3", "MGR valida campo token nbr"+strconv.Itoa(numcampos)+" with value:"+campo+"*")
    var resultado string
    var cualfallo int 
    cualfallo = 0
        
        if numcampos == 1{
            if campo != "" {
                if len(campo) > 30 {
                    resultado = "External identifier max leng is 30"
                        
                    cualfallo = 1
                }else{
                    matched, err := regexp.MatchString(`\W`, campo)
                    //log.Print("match1: ",matched)
                    if matched == true {
                        log.Print("External identifier characters don´t match")
                        resultado = "External identifier characters don´t match"
                        
                        cualfallo = 1
                    }
                    if err != nil{
                        log.Print("err: ",err)
                    }
                }

            }else{
                resultado = "External Identifier is required"
                cualfallo = 1
            }
        }//end if numcapos1
            
        if numcampos == 2{
            if campo != "" {
				if len(campo) >30 {
	
					resultado="Customer reference max lenght is 30"
                    cualfallo = 2
			    }else{
                    matched, err := regexp.MatchString(`\W`, campo)
                    //log.Print("match2: ",matched)
                    if matched == true {
                        log.Print("Customer reference characters don´t match")
                        resultado = "Customer reference characters don´t match"
                        
                        cualfallo = 2
                    }
                    if err != nil{
                        log.Print("err: ",err)
                    }
                }

			}else{
				resultado="Client reference is required"
                cualfallo = 2
		    }
                
        }//end if numcampos2

        if numcampos == 3{
            if campo != "" {
				if len(campo) >50 {
	
					resultado="Payment reference max lenght is 50"
                    cualfallo = 3
			    }else{
                    matched, err := regexp.MatchString(`\D`, campo)
                    //log.Print("match3: ",matched)
                    if matched == true{
                        log.Print("Payment reference characters don´t match")
                        resultado = "Payment reference characters don´t match"
                        
                        cualfallo = 3
                    }
                    if err != nil{
                        log.Print("err: ",err)
                    }
                }
			}else{
				resultado="Payment reference is required"
                cualfallo = 3
		    }
        }//end if numcampos3

        if numcampos == 4{
            if campo != "" {
				if len(campo)==16{
                    matched, err := regexp.MatchString(`\D`, campo)
                    
                    //log.Print("match4: ",matched,)
                    if matched == true {
                        log.Print("Card Number characters don´t match")
                        resultado = "Card Number characters don´t match"
                        
                        cualfallo = 4
                    }
                    if err != nil{
                        log.Print("err: ",err)
                    }
				}else{
					resultado="Card Number must be 16 digits:"+campo
                    cualfallo = 4
			    }
			}else{
				resultado="Card is required"
                cualfallo = 4
		    }
        }//end if numcampos4

        if numcampos == 5{
            utilito.LevelLog(Config_env_log, "3", "\n")
            if campo != "" {
				if  len(campo)==4 || len(campo)==5 { // 2 for the double quotes and 1 for the end of line
                    matched, err := regexp.MatchString(`\D`, campo)
                    
                    //log.Print("match5: ",matched)
                    if matched == true {
                        log.Print("Valid Thru characters don´t match")
                        resultado = "Valid Thru characters don´t match"
                        
                        cualfallo = 5
                    }
                    if err != nil{
                        log.Print("err: ",err)
                    }
				}else{
					resultado="Valid Thru  4 digits"
                    cualfallo = 5
			    }
			}else{
				resultado="Valid Thru is required"
                cualfallo = 5
		    }
        }//end if numcampos5

    return resultado, cualfallo

} //end func valida_campo_token

//Función campos_payment
func campos_payment (line string, lineas int)(string, int){
    
    utilito.LevelLog(Config_env_log, "3", "MGR campo por campo")
    numcampos := 0
    var resultado string
    var cualfallo int
    for _, campo := range strings.Split(strings.TrimSuffix(line, ","), ","){
        utilito.LevelLog(Config_env_log, "3", campo)
        numcampos = numcampos + 1
        resultado, cualfallo = valida_campo_pay(campo, numcampos)
    }
        
    return resultado, cualfallo

} //end func campos_payment

//Función valida_campo_pay
func valida_campo_pay (campo string, numcampos int)(string, int){
    
    utilito.LevelLog(Config_env_log, "3", "MGR valida campo payment nbr"+strconv.Itoa(numcampos)+" with value: "+campo+"*")
    log.Print("leng:",len(campo))
    var resultado string
    var cualfallo int 
    cualfallo = 0
        
        if numcampos == 1{
            if campo != "" {
                if len(campo) > 30 {
                    resultado = "External identifier max leng is 30"
                    cualfallo = 1
                }else{
                    matched, err := regexp.MatchString(`\W`, campo)
                    //log.Print(matched)
                    if matched == true {
                        log.Print("External identifier characters don´t match")
                        resultado = "External identifier characters don´t match"
                        
                        cualfallo = 1
                    }
                    if err != nil{
                        log.Print("err: ",err)
                    }
                }//end if len(campo)

            }else{
                resultado = "External Identifier is required"
                cualfallo = 1
            }
        }//end if numcampos1
            
        if numcampos == 2{
            if campo != "" {
				if len(campo) >30 {
	
					resultado="Client reference is required"
                    cualfallo = 2
			    }else{
                    matched, err := regexp.MatchString(`\W`, campo)
                    //log.Print(matched)
                    if matched == true {
                        log.Print("Client reference characters don´t match")
                        resultado = "Client reference characters don´t match"
                        
                        cualfallo = 2
                    }
                    if err != nil{
                        log.Print("err: ",err)
                    }
                }//end if len(campo)
			}else{
				resultado="Client reference is required"
                cualfallo = 2
		    }

        }//end if numcapos2

        if numcampos == 3{
            if campo != "" {
				if len(campo) >30 {
	
					resultado="Payment reference max lenght is 30"
                    cualfallo = 3
			    }else{
                    matched, err := regexp.MatchString(`\W`, campo)
                    //log.Print(matched)
                    if matched == true {
                        log.Print("Payment reference characters don´t match")
                        resultado = "Payment reference characters don´t match"
                        
                        cualfallo = 3
                    }
                    if err != nil{
                        log.Print("err: ",err)
                    }
                }//end if len(campo)
			}else{
				resultado="Payment reference is required"
                cualfallo = 3
		    }

        }//end if numcampos3

        if numcampos == 4{
            if campo != "" {
                if len(campo) ==31 {
	
					matched, err := regexp.MatchString(`\W`, campo)
                    //log.Print(matched)
                    if matched == true {
                        log.Print("Token characters don´t match")
                        resultado = "Token characters don´t match"
                        
                        cualfallo = 4
                    }
                    if err != nil{
                        log.Print("err: ",err)
                    }
                    
			    }else{

                    resultado="Token lenght must be 31: "+campo
                    cualfallo = 4
                }//end if len(campo)

			}else{
				resultado="Token is required"
                cualfallo = 4
		    }
        }//end if numcampos4

        if numcampos == 5{
                           
            if campo != "" {
				if len(campo)==3 ||  len(campo)==4 {

                    matched, err := regexp.MatchString(`\D`, campo)
                    //log.Print(matched)
                    if matched == true {
                        log.Print("Cvv characters don´t match")
                        resultado = "Cvv characters don´t match"
                        
                        cualfallo = 5
                    }
                    if err != nil{
                        log.Print("err: ", err)
                    }
				}else{
					resultado="Cvv must be 3 or 4 digits"
                    cualfallo = 5
			    }
			}else{
				resultado="Cvv is required"
                cualfallo = 5
		    }
        }//end if numcampos5 
            
		if numcampos == 6{
            if campo != "" {

                    matched, err := regexp.MatchString(`\D`, campo)
                    //matched2, err := regexp.MatchString('\.',campo)
                    //log.Print("match2: ",matched2)
                    if matched == true {
                        log.Print("Amount characters don´t match")
                        resultado = "Amount characters don´t match"
                        
                        cualfallo = 6
                    }
                    if err != nil {
                        log.Print("err: ", err)
                    }
			}else{
				resultado="Amount is required"
                cualfallo = 6
		    }
        }//end if numcampos6	

    return resultado, cualfallo

}//end func valida_campo_pay

///END
//// File_tokenizer   solution needs these to validate the input of the lines received.

//// File_tokenizer   solution needs these to validate the input of the lines received.
//// File_tokenizer   solution needs these to validate the input of the lines received.
///START

//Función validateAndObtainCampos_token
func validateAndObtainCampos_token (line string, lineas int)(modelito.RequestTokenized,string, int){
    
    ////////////////////////////////////////////////validate and if OK, then set the values in the Tokenized data structure 
    //START    
    requestData := modelito.RequestTokenized{}


     
    /*
        Clientreference string            `json:"clientreference"`
        Paymentreference string            `json:"paymentreference"`
        Card  string      `json:"card"`
        Exp  string      `json:"exp"`
        Cvv  string      `json:"Cvv"`
    */
    utilito.LevelLog(Config_env_log, "3", "MGR campo por campo")
    numcampos := 0
    var resultado string
    var cualfallo int
    resultado ="OK"
    cualfallo =0
    
    for _, campo := range strings.Split(strings.TrimSuffix(line, ","), ","){
        
        numcampos = numcampos + 1

        var campoValue string

        limpia := strings.Replace(campo, ":", "", -1) // para eliminar cualquier caracter de ":"
        limpia2 := strings.Replace(limpia, "\r", "", -1) //elimina cualquier retorno de carro
        campoValue = strings.Replace(limpia2, "\"", "", -1) // only works with a single character
        log.Print("Prueba: "+campoValue)
        var largo string
        largo = strconv.Itoa ( len(campoValue))
        utilito.LevelLog(Config_env_log, "1", "largo del campo es: "+largo+" :valor del campo es: "+campoValue)
              
        resultado, cualfallo = valida_campo_token(campoValue, numcampos)

        if cualfallo >0 {
                  
            utilito.LevelLog(Config_env_log, "3", "fallo es valor en :"+campo)
            //set empty values for dataToken to be returned
            requestData.Clientreference = ""
            requestData.Paymentreference = ""
            requestData.Card = ""
            requestData.Exp = ""
            requestData.Cvv = ""
            break
        }else{
            //set the value 
            if numcampos==1 {//IGS data id external identifier
            }

            if numcampos==2 {
                requestData.Clientreference = campoValue
            }
            
            if numcampos==3 {
                requestData.Paymentreference = campoValue
            }
            
            if numcampos==4 {
                requestData.Card = campoValue
            }
            
            if numcampos==5 {
                requestData.Exp = campoValue
            }
            
            if numcampos==6 {//cvv not required
                requestData.Cvv = campoValue
            }                       

        }//end else cualfallo
    }//end for campo
        
    return requestData,resultado, cualfallo
    
} //end validateAndObtainCampos_token

//Funcion validateAndObtainCampos_payment
func validateAndObtainCampos_payment (line string, lineas int)(modelito.RequestPayment,string, int){
    
    ////////////////////////////////////////////////validate and if OK, then set the values in the Payment data structure 
    //START    
    requestData := modelito.RequestPayment{}  //model/request.go

    /*
	    Clientreference string            `json:"clientreference"`
	    Paymentreference  string      `json:"paymentreference"`
	    Token  string      `json:"Token"`
	    Cvv  string      `json:"Cvv"`
	    Amount  string      `json:"Amount"`
    */

    utilito.LevelLog(Config_env_log, "3", "MGR campo por campo")
    numcampos := 0
    var resultado string
    var cualfallo int
    resultado = "OK"
    cualfallo = 0

    for _, campo := range strings.Split(strings.TrimSuffix(line, ","), ","){
        
        //utilito.LevelLog(Config_env_log, "3", campo)
        numcampos = numcampos + 1
        //resultado, cualfallo = valida_campo_pay(campo, numcampos)

        var campoValue string

        limpia := strings.Replace(campo, ".", "", -1) // para eliminar cualquier caracter de "."
        limpia2 := strings.Replace(limpia, "\r", "", -1) // quita cualquier retorno de carro que encuentre
        campoValue = strings.Replace(limpia2, "\"", "", -1) // only works with a single character

        log.Print("Prueba: "+campoValue)
        var largo string
        largo = strconv.Itoa ( len(campoValue))
        utilito.LevelLog(Config_env_log, "1", "largo del campo es:"+largo+":valor del campo es:"+campoValue)
              
        resultado, cualfallo = valida_campo_pay(campoValue, numcampos)//logicrequest.go


        if cualfallo >0 {
                  
            utilito.LevelLog(Config_env_log, "3", "fallo es valor en :"+campo)
            //set empty values for dataToken to be returned
            requestData.Clientreference = ""
            requestData.Paymentreference = ""
            requestData.Token = ""
            requestData.Cvv = ""
            requestData.Amount = ""                  
            break
        }else{
            //set the value 
            if numcampos==1 {
            }

            if numcampos==2 {
                requestData.Clientreference = campoValue
            }

            if numcampos==3 {
                requestData.Paymentreference = campoValue
            }
            
            if numcampos==4 {
                requestData.Token = campoValue
            }
            
            if numcampos==5 {
                requestData.Cvv = campoValue
            }
            
            if numcampos==6 {
                requestData.Amount = campoValue
            }                       

        }//end else cualfallo

    }//end for campo
        
    return requestData,resultado, cualfallo

} //end func validateAndObtainCampos_payment

///END
//// File_tokenizer   solution needs these to validate the input of the lines received.


func obtainParmsConsultarTokensPOST(r *http.Request, errorGeneral string) (string, string){
    
    log.Print("Entra a obtainParmsConsultarTokens :)")
    var requestData string
	////////////////////////////////////////////////obtain parms in JSON
    //START    
    utilito.LevelLog(Config_env_log, "3", "cz  handleConsultarTokens")
 	utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test ConsultarTokens")
    	    
    err := r.ParseForm()
	if err != nil {
	    //prepare response with error 100
	    utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 180. Missing parameter:"+errorGeneral)
	    errorGeneral="ERROR:180 -"	+err.Error()
	}
		
    v := r.Form
    
	requestData = v.Get("contenidofileconsultaTok")
    

    var lineaDatos string
    micadenita := requestData

    utilito.LevelLog(Config_env_log, "3", "Respuesta ConsultarToken")
    utilito.LevelLog(Config_env_log, "3", micadenita)

    cadenalimpia :=  strings.Replace(micadenita, "{", "", -1)

    for _, linea := range strings.Split(strings.TrimSuffix(cadenalimpia, "}"), "}"){
        utilito.LevelLog(Config_env_log, "3", "linea")
            
        linealimpia :=  strings.Replace(linea, " ", "", -1)
            
        utilito.LevelLog(Config_env_log, "3", linealimpia)

        for _, campo := range strings.Split(strings.TrimSuffix(linealimpia, ","), ","){
            utilito.LevelLog(Config_env_log, "3", "Campo")
            utilito.LevelLog(Config_env_log, "3", campo)
            dato := strings.Split(campo, ":")
            lineaDatos = lineaDatos + dato[1] +","
        } // end for campo

        lineaDatos = lineaDatos +"\r\n"
        utilito.LevelLog(Config_env_log, "3", lineaDatos)
    } //end for linea

       
    //END
   	 
   	return lineaDatos,errorGeneral
} //end obtainParmsConsultarTokensPOST

func obtainParmsConsultarHistorialClientes(r *http.Request, errorGeneral string) (string, string){
    
    log.Print("Entra a obtainParmsConsultarHistorialClientes GET")
    var requestData string
	////////////////////////////////////////////////obtain parms in JSON
    //START    
    utilito.LevelLog(Config_env_log, "3", "cz  handleConsultarHistorialClientesGET ")
 	utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test ConsultarHistorialClientes GET")
    	    
    err := r.ParseForm()
	if err != nil {
	    //prepare response with error 100
	    utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 180. Missing parameter:"+errorGeneral)
	    errorGeneral="ERROR:180 -"	+err.Error()
	}
		
    v := r.Form
    
	requestData = v.Get("contenidofileconsultaHistorialC")
    

    var lineaDatos string
    micadenita := requestData

    utilito.LevelLog(Config_env_log, "3", "Respuesta ConsultarHistorialClientes GET:")
    utilito.LevelLog(Config_env_log, "3", micadenita)

/*
    cadenalimpia :=  strings.Replace(micadenita, "{", "", -1)

    for _, linea := range strings.Split(strings.TrimSuffix(cadenalimpia, "}"), "}"){
        utilito.LevelLog(Config_env_log, "3", "linea")
            
        linealimpia :=  strings.Replace(linea, " ", "", -1)
            
        utilito.LevelLog(Config_env_log, "3", linealimpia)

        for _, campo := range strings.Split(strings.TrimSuffix(linealimpia, ","), ","){
            utilito.LevelLog(Config_env_log, "3", "Campo")
            utilito.LevelLog(Config_env_log, "3", campo)
            dato := strings.Split(campo, ":")
            lineaDatos = lineaDatos + dato[1] +","
        } // end for campo

        lineaDatos = lineaDatos +"\r\n"
        utilito.LevelLog(Config_env_log, "3", lineaDatos)
    } //end for linea

  */     
    
   	// lineaDatos= "ESTO LLEGO: "+micadenita
    lineaDatos= micadenita

    //END

   	return lineaDatos,errorGeneral
} //end obtainParmsConsultarHistorialClientes
func obtainParmsConsultarHistorialCPOST(r *http.Request, errorGeneral string) (string, string){
    
    log.Print("Entra a obtainParmsConsultarHistorialCPOST :)")
    var requestData string
	////////////////////////////////////////////////obtain parms in JSON
    //START    
    utilito.LevelLog(Config_env_log, "3", "cz  handleConsultarHistorilaC")
 	utilito.LevelLog(Config_env_log, "3", "CZ    handler Listening test HistorialClientes")
    	    
    err := r.ParseForm()
	if err != nil {
	    //prepare response with error 100
	    utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 180. Missing parameter:"+errorGeneral)
	    errorGeneral="ERROR:180 -"	+err.Error()
	}
		
    v := r.Form
    
	requestData = v.Get("contenidofileconsultaHistorialC")
    

    var lineaDatos string
    micadenita := requestData

    utilito.LevelLog(Config_env_log, "3", "Respuesta ConsultarHistorialC")
    utilito.LevelLog(Config_env_log, "3", micadenita)

    cadenalimpia :=  strings.Replace(micadenita, "{", "", -1)

    for _, linea := range strings.Split(strings.TrimSuffix(cadenalimpia, "}"), "}"){
        utilito.LevelLog(Config_env_log, "3", "linea")
            
        linealimpia :=  strings.Replace(linea, " ", "", -1)
            
        utilito.LevelLog(Config_env_log, "3", linealimpia)

        for _, campo := range strings.Split(strings.TrimSuffix(linealimpia, ","), ","){
            utilito.LevelLog(Config_env_log, "3", "Campo")
            utilito.LevelLog(Config_env_log, "3", campo)
            dato := strings.Split(campo, ":")
            lineaDatos = lineaDatos + dato[1] +","
        } // end for campo

        lineaDatos = lineaDatos +"\r\n"
        utilito.LevelLog(Config_env_log, "3", lineaDatos)
    } //end for linea

       
    //END
   	 
   	return lineaDatos,errorGeneral
} //end obtainParmsConsultarHistorialCPOST