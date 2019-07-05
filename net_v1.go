package main

import (
	"net/http"
    "strconv"
    "strings"
    "time"
//    "regexp"
//	"fmt"
	"log"
	"banwire/services/file_tokenizer/db"
	"banwire/services/file_tokenizer/net"
	modelito "banwire/services/file_tokenizer/model"
    utilito "banwire/services/file_tokenizer/util"

//  "banwire/services/file_tokenizer/model"
//	"banwire/services/file_tokenizer/model/pgsql"
//  "encoding/json"
    
//	"time"
//	"encoding/json"
//	 "database/sql"
	 _ "github.com/lib/pq"   //use go get github.com/lib/pq
     
//     	"os"
        "io"
        "bytes"
)

// init loads the routes for version 1
func init() {
//	var _r = net.GetRouter()
//	var r = _r.PathPrefix("/v1").Subrouter()

    var r = net.GetRouter()
	//route for test
	    log.Print("cz  init net_v1")

        r.Handle("/v1/index", netHandle(dash01Handler, nil)).Methods("GET")   
        r.Handle("/v1/indexpay", netHandle(dash01payHandler, nil)).Methods("GET")   
        r.Handle("/v1/indexconsulta", netHandle(dash01consultaHandler, nil)).Methods("GET")   
        r.Handle("/v1/indexconsultafiles", netHandle(dash01consultafilesHandler, nil)).Methods("GET")           
        r.Handle("/app.min.css", netHandle(serveCss01, nil)).Methods("GET")     
        r.Handle("/bootstrap-theme.min.css", netHandle(serveCss02, nil)).Methods("GET")     
        r.Handle("/bootstrap.min.css", netHandle(serveCss03, nil)).Methods("GET")     
        r.Handle("/font-awesome.min.css", netHandle(serveCss04, nil)).Methods("GET")     
        r.Handle("/ngToast.min.css", netHandle(serveCss05, nil)).Methods("GET")             
        r.Handle("/nya-bs-select.min.css", netHandle(serveCss06, nil)).Methods("GET")             
        r.Handle("/ui-bootstrap-csp.css", netHandle(serveCss07, nil)).Methods("GET")             
        r.Handle("/angular.min.js", netHandle(serveJs01, nil)).Methods("GET")     

    //Handle's para index
    r.Handle("/v1/validatefiles", netHandle(handlePostVaidateFiles, nil)).Methods("POST")   //in this net_v1.go
    r.Handle("/v1/processtokenfile", netHandle(handlePostProcessTokenFile, nil)).Methods("POST")   //in this net_v1.go
    r.Handle("/v1/downloadfileValida", netHandle(ForceDownloadValida, nil)).Methods("GET")     //in this net_v1.go
    r.Handle("/v1/downloadfileTokeniza", netHandle(ForceDownloadTokeniza, nil)).Methods("GET")     //in this net_v1.go
    
    //Handle's para index pay
    r.Handle("/v1/processpaymentfile", netHandle(handlePostProcessPaymentFile, nil)).Methods("POST")   //in this net_v1.go
    r.Handle("/v1/validatepaymentfiles", netHandle(handlePostVaidatePaymentFiles, nil)).Methods("POST")   //in this net_v1.go
    //r.Handle("/v1/downloadfile", netHandle(ForceDownload, nil)).Methods("POST")     //in this net_v1.go
    r.Handle("/v1/downloadfile", netHandle(ForceDownload, nil)).Methods("GET")     //in this net_v1.go
    //r.Handle("/v1/downloadfilePago", netHandle(ForceDownloadPago, nil)).Methods("POST")     //in this net_v1.go
    r.Handle("/v1/downloadfilePago", netHandle(ForceDownloadPago, nil)).Methods("GET")     //in this net_v1.go

    //Handle's para indexconsulta
    r.Handle("/v1/consultartokens", netHandle(handlePostConsultaTokens, nil)).Methods("POST")   //in this net_v1.go
    r.Handle("/v1/consultartokens", netHandle(handleGetConsultaTokens, nil)).Methods("GET")   //in this net_v1.go
    //r.Handle("/v1/consultarhistorialtokens", netHandle(handlePostConsultaHistorial, nil)).Methods("POST")   //in this net_v1.go
    r.Handle("/v1/consultarhistorialtokens", netHandle(handleGetConsultaHistorial, nil)).Methods("GET")   //in this net_v1.go

    //Handle's para indexconsultafiles
	//r.Handle("/v1/consultarhistorialClientes", netHandle(handlePostConsultahistorialClientes, nil)).Methods("POST")   //in this net_v1.go
    r.Handle("/v1/consultarhistorialClientes", netHandle(handleGetConsultahistorialClientes, nil)).Methods("GET")   //in this net_v1.go
    //r.Handle("/v1/consultahistorialToken", netHandle(handlePostConsultaHistorialToken, nil)).Methods("POST")   //in this net_v1.go
    //r.Handle("/v1/consultarhistorialPagos", netHandle(handlePostConsultaHistorialPagos, nil)).Methods("POST")   //in this net_v1.go

    //TO DO not needed in this program
	r.Handle("/v1/fetchtokenizedcards", netHandle(handleDBPostGettokenizedcards, nil)).Methods("POST")   //in this net_v1.go
	//r.Handle("/v1/processpayment", netHandle(v4handleDBPostProcesspayment, nil)).Methods("POST")           //in this net_v1.go    	   
	//r.Handle("/v1/generatetokenized", netHandle(handleDBPostGeneratetokenized, nil)).Methods("POST")     //in this net_v1.go

}

//index html angular

func dash01Handler(w http.ResponseWriter, r *http.Request) {
    
    utilito.LevelLog(Config_env_log, "1", "cz  dash01Handler with param")
 
    //log.Print("cz  dash01Handler with param"+Config_env_url)
    http.ServeFile(w,r,"index.html")
    /*data := TodoPageData{
			PageTitle: Config_env_server,
     }
     tmpl := template.Must(template.ParseFiles("index.html"))
     tmpl.Execute(w, data)
    */     
    log.Print("CZ   STEP dash01Handler 01")
}//end dash01Handler

//index html angular

func dash01payHandler(w http.ResponseWriter, r *http.Request) {
    
    log.Print("cz  dash01payHandler with param")
    //log.Print("cz  dash01Handler with param"+Config_env_url)
    http.ServeFile(w,r,"indexpay.html")
    /*data := TodoPageData{
			PageTitle: Config_env_server,
     }
     tmpl := template.Must(template.ParseFiles("index.html"))
     tmpl.Execute(w, data)
    */
    log.Print("CZ   STEP dash01payHandler 01")
}//end dash01payHandler

//index html angular

func dash01consultaHandler(w http.ResponseWriter, r *http.Request) {
    
    log.Print("cz  dash01consultaHandler with param")
    //log.Print("cz  dash01Handler with param"+Config_env_url)
    http.ServeFile(w,r,"indexconsulta.html")
    /*data := TodoPageData{
			PageTitle: Config_env_server,
     }
     tmpl := template.Must(template.ParseFiles("index.html"))
     tmpl.Execute(w, data)
    */
    log.Print("CZ   STEP dash01consultaHandler 01")
}//end dash01consultaHandler

//index html angular

func dash01consultafilesHandler(w http.ResponseWriter, r *http.Request) {
    
    log.Print("cz  dash01consultafilesHandler with param")
    //log.Print("cz  dash01Handler with param"+Config_env_url)
    http.ServeFile(w,r,"indexconsultafiles.html")
    /*data := TodoPageData{
			PageTitle: Config_env_server,
     }
     tmpl := template.Must(template.ParseFiles("index.html"))
     tmpl.Execute(w, data)
    */
    log.Print("CZ   STEP dash01consultafilesHandler 01")
}//end dash01consultafilesHandler

func serveCss01(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/app.min.css")
}//end serveCss01

func serveCss02(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/bootstrap-theme.min.css")
}//end serveCss02

func serveCss03(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/bootstrap.min.css")
}//end serveCss03

func serveCss04(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/font-awesome.min.css")
}//end serveCss04

func serveCss05(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/ngToast.min.css")
}//end serveCss05

func serveCss06(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/nya-bs-select.min.css")
}//end serveCss06

func serveCss07(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/ui-bootstrap-csp.css")
}//end serveCss07

func serveJs01(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "js/angular.min.js")
}//end serveJs01

    //post

   
    //post

//  handlePostProcessTokenFile

func handlePostProcessTokenFile(w http.ResponseWriter, r *http.Request) {
    
    defer func() {
		db.Connection.Close(nil)
	}()
    var errorGeneral string
    var errorGeneralNbr string
    var resultadoTokenJson modelito.Card

    //this is used to store the file procession status and data
    var u modelito.Filetrans
    u.Transtype ="process tokenizer"
    /*u.Transtype =
    u.Filename =

    u.Trans_status =
    u.Trans_statusmssg =
    u.Trans_processstatus =
    u.TransCreated_at =
    u.Trans_user =
    u.Trans_data_received =
    u.Trans_val_response =
    u.Trans_process_responser =
    u.Trans_process_qty =
    */
     tokensCreados := 0

    //var requestData modelito.RequestTokenizedCards
    errorGeneral=""

    linesStatus := []modelito.ExitoDataValidaLine{}   //structure to stire the errors in each of the liens of the file
    inputDataToken:= []modelito.RequestTokenized{} //this is for processing not for just validating
    inputDataPayment:= []modelito.RequestPayment{} //this is for processing not for just validating

    errorGeneral,errorGeneralNbr ,linesStatus,inputDataToken,inputDataPayment =  validateFiles("token", r)  //logicbusiness.go

    if errorGeneral!="" {
        utilito.LevelLog(Config_env_log, "3","CZ    Prepare Response with 100. Validation File failed-Tokens:"+errorGeneral)
    	errorGeneral="ERROR:100 -Validation File failed-Tokens: "	+errorGeneral
    	errorGeneralNbr="100"
    }

    var lineasWithErrors =0   // this will help to identify if Tokenization wwas done SUCESSfor each and all the lines,
                               //or some had errors


	////////////////////////////////////////////////process business rules
	/// START
    processLinesStatus := []modelito.ExitoDataTokenLine{}  //an array for the status of the process (tokenization)

    if errorGeneral=="" {   //process business Tokenization
        // use this structuire inputDataToken to call methods for the tokenization
     
        //the results of each of the tokanizations, will be returned here

        utilito.LevelLog(Config_env_log, "1","CZ  ProcessTokenFile  STEP Get the File")
        utilito.LevelLog(Config_env_log, "3"," ProcessTokenFile File Upload Endpoint Hit")
        //for each token in the array, call this method

        lineaProcess := 1

        var howmany int
        howmany = len(inputDataPayment)
        howmany = howmany+1

        log.Print("info lienaProcess: ",lineaProcess)
     
        for _, d := range inputDataToken {
            lineaProcess =lineaProcess +1
             log.Print("info lienaProcess: ",lineaProcess)
             var responseGeneral string

            responseGeneral,errorGeneralNbr,resultadoTokenJson= ProcessGeneratetokenized(w , d) //logicbusiness.go

             var u modelito.ExitoDataTokenLine

            if responseGeneral !=""{
                      if strings.Contains(responseGeneral, "ERROR") {
                            u.Line=strconv.Itoa(lineaProcess)
                            u.Status="ERROR540"
                            u.StatusMessage ="ERROR FIELD_"+strconv.Itoa(lineaProcess)+" - "+responseGeneral
                            now := time.Now()
                            u.Date = now
	                        u.Token = "0"
                            u.LastDigits = "0"
	                        u.Marca = "0"
	                        u.Vigencia = "0"
	                        u.Bin = "0"
	                        u.Score = "0"
	                        u.Type = "0"
                            lineasWithErrors =1
                        utilito.LevelLog(Config_env_log, "3"," ProcessTokenFile File -Process tokenizer ERROR line:"+responseGeneral)
                        
                           errorGeneral="ERROR555"
                      }else{
                          utilito.LevelLog(Config_env_log, "3"," ProcessTokenFile File -Process tokenizer OK line:")
                         //sucess for this line/tokenizer
                        u.Line=strconv.Itoa(lineaProcess)
                        u.Status="OK"
                        u.StatusMessage =responseGeneral
                        now := time.Now()
                        u.Date = now
	                    u.Token = resultadoTokenJson.Token
                        u.LastDigits = resultadoTokenJson.Last
	                    u.Marca = resultadoTokenJson.Brand
	                    u.Vigencia = resultadoTokenJson.Valid
	                    u.Bin = resultadoTokenJson.Bin
	                    u.Score = resultadoTokenJson.Score
	                    u.Type = resultadoTokenJson.Type
                        tokensCreados =    tokensCreados +1
                      }
            }else{
                utilito.LevelLog(Config_env_log, "3"," ProcessTokenFile File -Process tokenizer OK line:")
                //sucess for this line/tokenizer
                u.Line=strconv.Itoa(lineaProcess)
                u.Status="OK"
                u.StatusMessage ="OK"
                now := time.Now()
                u.Date = now
                u.Token = resultadoTokenJson.Token
                u.LastDigits = resultadoTokenJson.Last
                u.Marca = resultadoTokenJson.Brand
                u.Vigencia = resultadoTokenJson.Valid
                u.Bin = resultadoTokenJson.Bin
                u.Score = resultadoTokenJson.Score
                u.Type = resultadoTokenJson.Type
                tokensCreados =    tokensCreados +1
            }
            if responseGeneral==""{
                //this is not expected, as the result will be returned 
                utilito.LevelLog(Config_env_log, "3"," ProcessTokenFile File -Process tokenizer NOT expected:")
            }
            // add this tokenization into the sattus for all the lines
           utilito.LevelLog(Config_env_log, "3"," ProcessTokenFile File -before line:"+u.StatusMessage )

           processLinesStatus = append(processLinesStatus,u)

 		}//end for



	}//end if - process business Tokenization


	/// END
    if errorGeneral!=""{

        if   lineasWithErrors ==0 { //all the lines were tokenized
                errorGeneral ="SUCCESS"
        }
        if   lineasWithErrors ==1 { //not all were tokenized
                errorGeneral ="PARTIAL SUCCESS -SOME LINES WERE NOT TOKENIZED"
        }


    	//send error response if any
    	//prepare an error JSON Response, if any
		utilito.LevelLog(Config_env_log, "3","CZ ProcessTokenFile  STEP Get the ERROR response JSON ready")
		// START
		 //old  getJsonResponseError(errorGeneral, errorGeneralNbr)

        fieldDataBytesJson,err := getJsonResponseTokenFile(errorGeneral, errorGeneralNbr, linesStatus,processLinesStatus  )  //logicresponse.go 
		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		if(err!=nil){
			
		}
	
    }else{
        utilito.LevelLog(Config_env_log, "3","CZ ProcessTokenFile  STEP SUCESS, prepare response JSON ready")

       if   lineasWithErrors ==0 { //all the lines were tokenized
            errorGeneral ="SUCCESS"
       }
       if   lineasWithErrors ==1 { //not all were tokenized
            errorGeneral ="PARTIAL SUCCESS -SOME LINES WERE NOT TOKENIZED"
       }

        errorGeneralNbr ="OK"
        fieldDataBytesJson,err := getJsonResponseTokenFile(errorGeneral, errorGeneralNbr, linesStatus,processLinesStatus  )  //logicresponse.go 

		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		if(err!=nil){
			
		}

    } 
    //call the function to store the filetransaction info
    //type tokenizer

    u.Trans_status =errorGeneralNbr
    u.Trans_statusmssg =errorGeneral
    u.Trans_processstatus =errorGeneral
    u.Trans_process_qty ="3"

     resultCreateFiletransRecord :=logicProcessCreateFileTrans(u, errorGeneral) 
     if resultCreateFiletransRecord!=""{
              //error inserting the record for file trans
     }else{
         //success inserting the record for file trans
     }
}//end handlePostProcessTokenFile

func handlePostVaidateFiles(w http.ResponseWriter, r *http.Request) {
	defer func() {
		db.Connection.Close(nil)
	}()
    var errorGeneral string
    var errorGeneralNbr string

    //this is used to store the file procession status and data
    var u modelito.Filetrans
    u.Transtype ="validate token"
    /*u.Transtype =
    u.Filename =

    u.Trans_status =
    u.Trans_statusmssg =
    u.Trans_processstatus =
    u.TransCreated_at =
    u.Trans_user =
    u.Trans_data_received =
    u.Trans_val_response =
    u.Trans_process_responser =
    u.Trans_process_qty =
    */

    //var requestData modelito.RequestTokenizedCards

    errorGeneral=""

    linesStatus := []modelito.ExitoDataValidaLine{}   //structure to stire the errors in each of the liens of the file
    notToBeUsedDataToken:= []modelito.RequestTokenized{} //this is for processing not for just validating
    notToBeUsedDataPayment:= []modelito.RequestPayment{} //this is for processing not for just validating

    var howmany int
    howmany = len(notToBeUsedDataToken)
    howmany = howmany+1
                
    howmany = len(notToBeUsedDataPayment)
    howmany = howmany+1

    errorGeneral,errorGeneralNbr ,linesStatus,notToBeUsedDataToken,notToBeUsedDataPayment =  validateFiles("token", r)  //logicbusiness.go

    
    if errorGeneral!="" {
        log.Print("CZ    Prepare Response with 100. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:100 -Missing parameter"	+errorGeneral
    	errorGeneralNbr="100"
    }
    if errorGeneral!=""{
    	//send error response if any
    	//prepare an error JSON Response, if any
		log.Print("CZ   STEP Get the ERROR response JSON ready")
		
		// START
		 //old  getJsonResponseError(errorGeneral, errorGeneralNbr)

        fieldDataBytesJson,err := getJsonResponseErrorValidateFile(errorGeneral, errorGeneralNbr, linesStatus  )  //logicresponse.go 
		//////////    write the response (ERROR)
         w.Header().Set("Content-Disposition", "attachment; filename=foo.pdf")
         //w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

         //w.Write(pdfData)

		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	



		if(err!=nil){
			
		}
	
    }else{

/*        var  cardTokenized modelito.Card
        fieldDataBytesJson,err := getJsonResponseValidateFileV2(cardTokenized)
        w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		if(err!=nil){
			
		}//end if
*/
        errorGeneral ="SUCCESS"
        errorGeneralNbr ="OK"
        fieldDataBytesJson,err := getJsonResponseErrorValidateFile(errorGeneral, errorGeneralNbr, linesStatus  )  //logicresponse.go 
		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		if(err!=nil){
			
		}

    } 

    //call the function to store the filetransaction info
    //type tokenizer

    u.Trans_status =errorGeneralNbr
    u.Trans_statusmssg =errorGeneral
    u.Trans_processstatus =errorGeneral
    u.Trans_process_qty ="0"

     resultCreateFiletransRecord :=logicProcessCreateFileTrans(u, errorGeneral) 
     if resultCreateFiletransRecord!=""{
              //error inserting the record for file trans
     }else{
         //success inserting the record for file trans
     }


}//end handlePostVaidateFiles

//  handlePostProcessTokenFile

func handlePostProcessPaymentFile(w http.ResponseWriter, r *http.Request) {
    
    defer func() {
		db.Connection.Close(nil)
	}()
    var errorGeneral string
    var errorGeneralNbr string
    var resultadoPaymentJson modelito.ExitoData

    //this is used to store the file procession status and data
    var u modelito.Filetrans
    u.Transtype ="process payment"
    /*u.Transtype =
    u.Filename =

    u.Trans_status =
    u.Trans_statusmssg =
    u.Trans_processstatus =
    u.TransCreated_at =
    u.Trans_user =
    u.Trans_data_received =
    u.Trans_val_response =
    u.Trans_process_responser =
    u.Trans_process_qty =
    */

    //var requestData modelito.RequestTokenizedCards

    errorGeneral=""

    linesStatus := []modelito.ExitoDataValidaLine{}   //structure to stire the errors in each of the liens of the file
    inputDataToken:= []modelito.RequestTokenized{} //this is for processing not for just validating
    inputDataPayment:= []modelito.RequestPayment{} //this is for processing not for just validating

    var howmany int
    howmany = len(inputDataToken)
    howmany = howmany+1

    howmany = len(inputDataPayment)
    howmany = howmany+1

    errorGeneral,errorGeneralNbr ,linesStatus,inputDataToken,inputDataPayment =  validateFiles("payment", r)  //logicbusiness.go
    
    if errorGeneral!="" {
        log.Print("CZ    Prepare Response with 300. Validation File failed-Payments:"+errorGeneral)
    	errorGeneral="ERROR:300 -Validation File failed-Payments"	+errorGeneral
    	errorGeneralNbr="300"
    }

    var lineasWithErrors =0   // this will help to identify if Tokenization wwas done SUCESSfor each and all the lines,
                               //or some had errors


	////////////////////////////////////////////////process business rules
	/// START
    processLinesStatus := []modelito.ExitoDataPayLine{}  //an array for the status of the process (tokenization)

    if errorGeneral=="" {   //process business Tokenization
        // use this structuire inputDataToken to call methods for the tokenization
     
         //the results of each of the tokanizations, will be returned here



    utilito.LevelLog(Config_env_log, "1","CZ  ProcessPaymentFile  STEP Get the File")
    utilito.LevelLog(Config_env_log, "3"," ProcessPaymentFile File Upload Endpoint Hit")
    //for each token in the array, call this method


    lineaProcess := 1

    var howmany int
    howmany = len(inputDataToken)
    howmany = howmany+1

    for _, d := range inputDataPayment {
     		lineaProcess =lineaProcess +1
             
             var responseGeneral string
                               //d is supposed to be of type   requestData modelito.RequestPayment
            responseGeneral,errorGeneralNbr,resultadoPaymentJson= v4ProcessProcessPayment(w , d) //logicbusiness.go

            log.Print("info responseGeneral: "+responseGeneral)
             var u modelito.ExitoDataPayLine

            if responseGeneral !=""{
                      if strings.Contains(responseGeneral, "ERROR") {
                            u.Line=strconv.Itoa(lineaProcess)
                            u.StatusMessage ="ERROR FIELD_"+strconv.Itoa(lineaProcess)+" - "+responseGeneral
                            u.Status="ERROR640"
                            u.Token = "0"
                            u.PaymentReference = "0"
                            u.Authcode = "0"
                            u.Idtransaction = "0"
                            u.Marca = "0"
                            u.Bin = "0"
                            u.LastDigits = "0"
                            u.Type = "0"
                            lineasWithErrors =1
                        utilito.LevelLog(Config_env_log, "3"," ProcessPaymentFile File -Process  ERROR line:"+responseGeneral)
                        
                           errorGeneral="ERROR640"
                      }else{
                          utilito.LevelLog(Config_env_log, "3"," ProcessPaymentFile File -Process  OK line:")
                         //sucess for this line/tokenizer
                        u.Line=strconv.Itoa(lineaProcess)
                        u.StatusMessage ="OK"
                        u.Status="OK"
                        u.Token = resultadoPaymentJson.Token
                        u.PaymentReference = resultadoPaymentJson.PaymentReference
                        u.Authcode = resultadoPaymentJson.Authcode
                        u.Idtransaction = resultadoPaymentJson.Idtransaction
                        u.Marca = resultadoPaymentJson.Marca
                        u.Bin = resultadoPaymentJson.Bin
                        u.LastDigits = resultadoPaymentJson.LastDigits
                        u.Type = resultadoPaymentJson.Type
                      }
            }else{
                    utilito.LevelLog(Config_env_log, "3"," ProcessPaymentFile File -Process  OK line:")
                    //sucess for this line/tokenizer
                u.Line=strconv.Itoa(lineaProcess)
                u.StatusMessage ="OK"
                u.Status="OK"
                u.Token = resultadoPaymentJson.Token
                u.PaymentReference = resultadoPaymentJson.PaymentReference
                u.Authcode = resultadoPaymentJson.Authcode
                u.Idtransaction = resultadoPaymentJson.Idtransaction
                u.Marca = resultadoPaymentJson.Marca
                u.Bin = resultadoPaymentJson.Bin
                u.LastDigits = resultadoPaymentJson.LastDigits
                u.Type = resultadoPaymentJson.Type
            }
            if responseGeneral==""{
                //this is not expected, as the result will be returned 
                utilito.LevelLog(Config_env_log, "3"," ProcessPaymentFile File -Process - NOT expected:")
            }
            // add this tokenization into the sattus for all the lines
           utilito.LevelLog(Config_env_log, "3"," ProcessPaymentFile File -before line:"+u.StatusMessage )

           processLinesStatus = append(processLinesStatus,u)

 		}//end for



	}//end if - process business payment


	/// END
    if errorGeneral!=""{

        if   lineasWithErrors ==0 { //all the lines were tokenized
                errorGeneral ="SUCCESS"
        }
        if   lineasWithErrors ==1 { //not all were tokenized
                errorGeneral ="PARTIAL SUCCESS -SOME LINES WITH PAYMENTS  WERE NOT PROCESSED"
        }


    	//send error response if any
    	//prepare an error JSON Response, if any
		utilito.LevelLog(Config_env_log, "3","CZ ProcessPaymentFile  STEP Get the ERROR response JSON ready")
		// START
		 //old  getJsonResponseError(errorGeneral, errorGeneralNbr)


        fieldDataBytesJson,err := getJsonResponsePaymentFile(errorGeneral, errorGeneralNbr, linesStatus,processLinesStatus  )  //logicresponse.go 
		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		if(err!=nil){
			
		}
	
    }else{
        utilito.LevelLog(Config_env_log, "3","CZ ProcessPaymentFile  STEP SUCESS, prepare response JSON ready")

       if   lineasWithErrors ==0 { //all the lines were tokenized
            errorGeneral ="SUCCESS"
       }
       if   lineasWithErrors ==1 { //not all were tokenized
            errorGeneral ="PARTIAL SUCCESS -SOME LINES WITH PAYMENTS  WERE NOT PROCESSED"
       }

        errorGeneralNbr ="OK"
        fieldDataBytesJson,err := getJsonResponsePaymentFile(errorGeneral, errorGeneralNbr, linesStatus,processLinesStatus  )  //logicresponse.go 
		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		if(err!=nil){
			
		}

    } 

    //call the function to store the filetransaction info
    //type tokenizer

    u.Trans_status =errorGeneralNbr
    u.Trans_statusmssg =errorGeneral
    u.Trans_processstatus =errorGeneral
    u.Trans_process_qty ="0"

     resultCreateFiletransRecord :=logicProcessCreateFileTrans(u, errorGeneral) 
     if resultCreateFiletransRecord!=""{
              //error inserting the record for file trans
     }else{
         //success inserting the record for file trans
     }

} //end handlePostProcessPaymentFile

func handlePostVaidatePaymentFiles(w http.ResponseWriter, r *http.Request) {
	
    defer func() {
		db.Connection.Close(nil)
	}()
    var errorGeneral string
    var errorGeneralNbr string


    //this is used to store the file procession status and data
    var u modelito.Filetrans
    u.Transtype ="validate payment"
    /*u.Transtype =
    u.Filename =

    u.Trans_status =
    u.Trans_statusmssg =
    u.Trans_processstatus =
    u.TransCreated_at =
    u.Trans_user =
    u.Trans_data_received =
    u.Trans_val_response =
    u.Trans_process_responser =
    u.Trans_process_qty =
    */
    
    //var requestData modelito.RequestTokenizedCards

    errorGeneral=""

    linesStatus := []modelito.ExitoDataValidaLine{}   //structure to stire the errors in each of the liens of the file

    notToBeUsedDataToken:= []modelito.RequestTokenized{} //this is for processing not for just validating
    notToBeUsedDataPayment:= []modelito.RequestPayment{} //this is for processing not for just validating

    var howmany int
    howmany = len(notToBeUsedDataToken)
    howmany = howmany+1
            
    howmany = len(notToBeUsedDataPayment)
    howmany = howmany+1


    errorGeneral,errorGeneralNbr ,linesStatus,notToBeUsedDataToken,notToBeUsedDataPayment =  validateFiles("payment", r)  //logicbusiness.go

    if errorGeneral !=""{//error validation
 

    }
    if errorGeneral==""{ //validation OK
        //now get the 
 
    }

    if errorGeneral!=""{
    	//send error response if any
    	//prepare an error JSON Response, if any
		log.Print("CZ   STEP Get the ERROR response JSON ready")
		
		// START
		 //old  getJsonResponseError(errorGeneral, errorGeneralNbr)

        fieldDataBytesJson,err := getJsonResponseErrorValidateFile(errorGeneral, errorGeneralNbr, linesStatus  )  //logicresponse.go 
		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		
        if(err!=nil){
			
		}
	
    }else{

/*        var  cardTokenized modelito.Card
        fieldDataBytesJson,err := getJsonResponseValidateFileV2(cardTokenized)
        w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		if(err!=nil){
			
		}//end if
*/
        errorGeneral ="SUCCESS"
        errorGeneralNbr ="OK"
        fieldDataBytesJson,err := getJsonResponseErrorValidateFile(errorGeneral, errorGeneralNbr, linesStatus  )  //logicresponse.go 
		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		
        if(err!=nil){
			
		}

    } 

    //call the function to store the filetransaction info
    //type tokenizer

    u.Trans_status =errorGeneralNbr
    u.Trans_statusmssg =errorGeneral
    u.Trans_processstatus =errorGeneral
    u.Trans_process_qty ="0"

    resultCreateFiletransRecord :=logicProcessCreateFileTrans(u, errorGeneral) 
    if resultCreateFiletransRecord!=""{
              //error inserting the record for file trans
    }else{
         //success inserting the record for file trans
    }

}//en function handlePostVaidatePaymentFiles

////-------------------------------------------------------Not used for these in the file logic
   
// handleDBGettokenizedcards  receive and handle the request from client, access DB, and web
func handleDBPostGettokenizedcards(w http.ResponseWriter, r *http.Request) {
    
    defer func() {
		db.Connection.Close(nil)
	}()
    var errorGeneral string
    var errorGeneralNbr string

   	var requestData modelito.RequestTokenizedCards


    errorGeneral=""
    requestData, errorGeneral=obtainPostParmsGettokenizedcards(r,errorGeneral) //logicrequest_post.go

	////////////////////////////////////////////////process business rules
	/// START
    
    if errorGeneral=="" {

		errorGeneral,errorGeneralNbr= ProcessGettokenizedcards(w , requestData) //logicbusiness.go
	}
	/// END
    
    if errorGeneral!=""{
    	//send error response if any
    	//prepare an error JSON Response, if any
		log.Print("CZ   STEP Get the ERROR response JSON ready")
		
			/// START
		fieldDataBytesJson,err := getJsonResponseError(errorGeneral, errorGeneralNbr)
		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		
        if(err!=nil){
			
		}
	
    } 
					
}//end handleDBPostGettokenizedcards

///////////////////////////////v4
///////////////////////////////v4

// handleGetConsultaTokens  receive and handle the request from client, access DB, and web
func handleGetConsultaTokens(w http.ResponseWriter, r *http.Request) {
    defer func() {
		db.Connection.Close(nil)
	}()
    
    var errorGeneral string
    var errorGeneralNbr string
	var  recordsFound []modelito.Card

    log.Print("Entra a handleGetConsultaTokens funcion boton \"Consultar tokens\" indexconsulta")
    //file := "ConsultarToken.txt"
    
    
    paramsReceived, err:= obtainParmsConsultarTokens(r , errorGeneral) //logisrequest.go
    
    if(err!=""){
        log.Print("CZ    Prepare Response with 601. Get Tokens done with thisToken failed:"+errorGeneral)
    	errorGeneral="ERROR:601 -Get tokens  by cust reference- parameter:"	+errorGeneral
    	errorGeneralNbr="601"

	}else{

		errorGeneral,errorGeneralNbr,recordsFound= ProcessGetTokensForCustRef(w , paramsReceived) //logicbusiness.go

    }

    log.Print("Regreso de función obtainParmsConsultarTokens")
    
    // downloadBytes:= []byte(paramsReceived)

    // Generate the server headers

    log.Print("Generador de cabezeras")


    log.Print("Fin obtainParmsConsultarTokens")

    if errorGeneral != ""{
    	//send error response if any
    	//prepare an error JSON Response, if any
		log.Print("CZ   STEP Get the ERROR response JSON ready")
		
			/// START
		fieldDataBytesJson,err := getJsonResponseError(errorGeneral, errorGeneralNbr)
		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		
        		
        if(err!=nil){
			
	    }
	
    }else{
    	//send error response if any
    	//prepare an error JSON Response, if any
		log.Print("CZ   STEP Get the OK response JSON ready")
		errorGeneralNbr ="500"
			/// START
		fieldDataBytesJson,err := getJsonResponseConsultarTokens( errorGeneral, errorGeneralNbr, recordsFound)
		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		
        		
        if(err!=nil){
			
	    }

    }



}//end handleGetConsultaTokens



///////////////////////////////v4


// handleGetConsultaHistorial  receive and handle the request from client, access DB, and web
func handleGetConsultaHistorial(w http.ResponseWriter, r *http.Request) {
    defer func() {
		db.Connection.Close(nil)
	}()
    
    log.Print("Entra a handleGetConsultaHistorial funcion boton \"Consultar historial pagos del token\" indexconsulta")
    
    var errorGeneral string
    var errorGeneralNbr string

   var  recordsFound []modelito.Payment
    
    paramsReceived, errorGeneral:= obtainParmsConsultarHistPagoTokens(r , errorGeneral) //logisrequest.go
    
    if(errorGeneral!=""){
        log.Print("CZ    Prepare Response with 501. Get Payments done with thisToken failed:"+errorGeneral)
    	errorGeneral="ERROR:501 -Get payments by token- parameter:"	+errorGeneral
    	errorGeneralNbr="501"

	}else{

		errorGeneral,errorGeneralNbr,recordsFound= ProcessGetPaymentsForToken(w , paramsReceived) //logicbusiness.go

    }
    log.Print("Regreso de función handleGetConsultaHistorial")
    
//    downloadBytes:= []byte(paramsReceived)

    // Generate the server headers


    log.Print("Entra a handlePostConsultaHistorial funcion boton \"Consultar historial pagos del token\" indexconsulta")


    if errorGeneral != ""{
    	//send error response if any
    	//prepare an error JSON Response, if any
		log.Print("CZ   STEP Get the ERROR response JSON ready")
		
			/// START
		fieldDataBytesJson,err := getJsonResponseError(errorGeneral, errorGeneralNbr)
		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		
        		
        if(err!=nil){
			
	    }
	
    }else{
    	//send error response if any
    	//prepare an error JSON Response, if any
		log.Print("CZ   STEP Get the OK response JSON ready")
		errorGeneralNbr ="500"
			/// START
		fieldDataBytesJson,err := getJsonResponseConsultarPayments(paramsReceived, errorGeneralNbr,recordsFound) //logicresponse.go
		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		
        		
        if(err!=nil){
			
	    }

    }

    
    log.Print("Fin obtainParmsConsultarHistPagoTokens")


}//end handleGetConsultaHistorial

//func handlePostConsultahistorialClientes
func handlePostConsultahistorialClientes(w http.ResponseWriter, r *http.Request) {
	
    log.Print("Entra a obtainParmsConsultarHistorialCPOST funcion boton \"Consultar\" indexconsultafiles")
    file := "ConsultarToken.txt"
    
    var errorGeneral string
    
    htmlStrDownloadJson, err:= obtainParmsConsultarHistorialCPOST(r , errorGeneral) //logisrequest.go
    
    if(err!=""){

	}//end if
    log.Print("Regreso de función obtainParmsConsultarHistorialCPOST")
    
    downloadBytes:= []byte(htmlStrDownloadJson)

    // Generate the server headers

    log.Print("Generador de cabezeras")
		
        w.Header().Set("Content-Type", "text/plain;charset: uft-8")
        w.Header().Set("Content-Disposition", "attachment; filename="+file+"")
		w.Write(downloadBytes)	

    log.Print("Fin obtainParmsConsultarHistorialC")

    /*defer func() {
		db.Connection.Close(nil)
	}()
    var errorGeneral string
    var errorGeneralNbr string

    //var requestData modelito.RequestTokenizedCards

    errorGeneral=""
    linesStatus := []modelito.ExitoDataValidaLine{}   //structure to stire the errors in each of the liens of the file

	//    requestData, errorGeneral=obtainPostParmsGettokenizedcards(r,errorGeneral) //logicrequest_post.go
    
    if errorGeneral!="" {
        log.Print("CZ    Prepare Response with 100. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:100 -Missing parameter"	+errorGeneral
    	errorGeneralNbr="100"
    }


	////////////////////////////////////////////////process business rules
	/// START
    
    if errorGeneral=="" {//if errorGeneral1

        log.Print("CZ   STEP Get the File")
        log.Print("File Upload Endpoint Hit")

        // Parse our multipart form, 10 << 20 specifies a maximum
        // upload of 10 MB files.
        err:= r.ParseMultipartForm(10 << 20)
        
        if err != nil {
            
            log.Print("CZ Error Retrieving the File")
            log.Print(err)
            errorGeneral="ERROR:110 -Error retriving files ,parameters"	+errorGeneral
            errorGeneralNbr="110"

        }

        if errorGeneral=="" {//if error general2
            log.Print("CZ Start read the form data")
            formdata := r.MultipartForm // ok, no problem so far, read the Form data
            
            //get the *fileheaders
            files := formdata.File["file0"] // grab the files, this files was set in the html 

            log.Print("CZ before loop files")

            for i, _ := range files { // loop through the files one by one
                log.Print("CZ loop step 1")
                var elfileindex string

                elfileindex = "file0"
                log.Print("CZ Loop file")
                log.Print("CZ Loop file:"+elfileindex)
                file, err := files[i].Open()
                log.Print("CZ open file")
                defer file.Close()
                if err != nil {
                    log.Print(w, err)
                    errorGeneral="ERROR:120 -Error file passed not open ,parameters"	+errorGeneral
                    errorGeneralNbr="120"

                }

                //convert multipart file into buffer bytes

                buf := bytes.NewBuffer(nil)
                io.Copy(buf, file)
                //if _, err := io.Copy(buf, file); err != nil {
                //      return nil, err
                //}                
                //log.Print(buf) // print the content as 'bytes'

                // convert content to a 'string'
                str := buf.String()

                log.Print(str) // print the content as a 'string'                

                log.Print("MGR paso linea por linea")

                lineas := 0

                lineasWithErrors := 0

                for _, line := range strings.Split(strings.TrimSuffix(str, "\n"), "\n") {//inicio for
                    var u modelito.ExitoDataValidaLine
                    if lineas >= 1{ //if -data
                        log.Printf("MGR Linea %d de datos", lineas)

                        lineas = lineas + 1

                        respuestaRes,cualfallo :=campos_payment (line, lineas)

                        if cualfallo ==0 {  //exito, todos los cmapos de la linea OK
                            u.Line=strconv.Itoa(lineas)
                            u.Status="OK"
                            u.StatusMessage ="SUCESS"
                        }else { //error, al menos un error en la linea
                            u.Line=strconv.Itoa(lineas)
                            u.Status="ERROR550"
                            u.StatusMessage ="ERROR LINEA:"+strconv.Itoa(cualfallo)+" - "+respuestaRes
                            lineasWithErrors = 1
                        }//fin else

                        linesStatus = append(linesStatus,u);
                    }//end -data

                    if lineas == 0 { //if -line name fields
                        log.Print("MGR Linea 0 de textos")
                        lineas = lineas + 1
                    } //end if -line name fields

                    log.Println(line)

                }//end for

                if  lineasWithErrors ==1 { //al menos una linea  tuvo un error
                    errorGeneral="ERROR FILE"
                }else{
                    if errorGeneral=="" {  

                        log.Print(w, "Files uploaded successfully : ")
                        log.Print(w, files[i].Filename+"\n")
                    }
                }//end else

                //1.count number of lines in the file received
                //2.for each line 
                //      validate the content
                //3.store in the db table AUDIT FILE VALIDATION
                //         seq nbr ,  file name, size, content[all the bytes],commentsParam, validationStatus, validationStatusMessage,validationResponse[for each line,a response OK/Error] ,timestamp
                //3.return result JSON

                //errorGeneral,errorGeneralNbr= ProcessGettokenizedcards(w , requestData) //logicbusiness.go
                log.Println(lineasWithErrors)

            }//end -loop through the files one by one
        
        } //end if error general
    
    }//end if error general1


	/// END
    if errorGeneral!=""{
    	//send error response if any
    	//prepare an error JSON Response, if any
		log.Print("CZ   STEP Get the ERROR response JSON ready")
		
		// START
        //old  getJsonResponseError(errorGeneral, errorGeneralNbr)

        fieldDataBytesJson,err := getJsonResponseErrorValidateFile(errorGeneral, errorGeneralNbr, linesStatus  )  //logicresponse.go 
        
        //////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		
        if(err!=nil){
			
		}
	
    }else{
        var  cardTokenized modelito.Card
        fieldDataBytesJson,err := getJsonResponseValidateFileV2(cardTokenized)
        w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		
        if(err!=nil){

	    }//end if

    }*/


}//end function handlePostConsultahistorialClientes

//func handlePostConsultaHistorialToken

func handlePostConsultaHistorialToken(w http.ResponseWriter, r *http.Request) {
	
    defer func() {
		db.Connection.Close(nil)
	}()
    var errorGeneral string
    var errorGeneralNbr string

    //var requestData modelito.RequestTokenizedCards


    errorGeneral=""
    linesStatus := []modelito.ExitoDataValidaLine{}   //structure to stire the errors in each of the liens of the file

	//    requestData, errorGeneral=obtainPostParmsGettokenizedcards(r,errorGeneral) //logicrequest_post.go
    if errorGeneral!="" {
        log.Print("CZ    Prepare Response with 100. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:100 -Missing parameter"	+errorGeneral
    	errorGeneralNbr="100"
    }


	////////////////////////////////////////////////process business rules
	/// START
    
    if errorGeneral=="" {//if errorGeneral1

        log.Print("CZ   STEP Get the File")
        log.Print("File Upload Endpoint Hit")

        // Parse our multipart form, 10 << 20 specifies a maximum
        // upload of 10 MB files.
        err:= r.ParseMultipartForm(10 << 20)
        if err != nil {
            
            log.Print("CZ Error Retrieving the File")
            log.Print(err)
            errorGeneral="ERROR:110 -Error retriving files ,parameters"	+errorGeneral
            errorGeneralNbr="110"

        }

        if errorGeneral=="" {//if error general2
            log.Print("CZ Start read the form data")
            formdata := r.MultipartForm // ok, no problem so far, read the Form data
            
            //get the *fileheaders
            files := formdata.File["file0"] // grab the files, this files was set in the html 

            log.Print("CZ before loop files")

            for i, _ := range files { // loop through the files one by one
                log.Print("CZ loop step 1")
                var elfileindex string

                elfileindex = "file0"
                log.Print("CZ Loop file")
                log.Print("CZ Loop file:"+elfileindex)
                file, err := files[i].Open()
                log.Print("CZ open file")
                defer file.Close()
                if err != nil {
                    log.Print(w, err)
                    errorGeneral="ERROR:120 -Error file passed not open ,parameters"	+errorGeneral
                    errorGeneralNbr="120"

                }

                //convert multipart file into buffer bytes

                buf := bytes.NewBuffer(nil)
                io.Copy(buf, file)
                //if _, err := io.Copy(buf, file); err != nil {
                //      return nil, err
                //}                
                //log.Print(buf) // print the content as 'bytes'

                // convert content to a 'string'
                str := buf.String()

                log.Print(str) // print the content as a 'string'                

                log.Print("MGR paso linea por linea")

                lineas := 0

                lineasWithErrors := 0

                for _, line := range strings.Split(strings.TrimSuffix(str, "\n"), "\n") {//inicio for
                    var u modelito.ExitoDataValidaLine
                    if lineas >= 1{ //if -data
                        log.Printf("MGR Linea %d de datos", lineas)

                        lineas = lineas + 1

                        respuestaRes,cualfallo :=campos_payment (line, lineas)

                        if cualfallo ==0 {  //exito, todos los cmapos de la linea OK
                            u.Line=strconv.Itoa(lineas)
                            u.Status="OK"
                            u.StatusMessage ="SUCESS"
                        }else { //error, al menos un error en la linea
                            u.Line=strconv.Itoa(lineas)
                            u.Status="ERROR550"
                            u.StatusMessage ="ERROR LINEA:"+strconv.Itoa(cualfallo)+" - "+respuestaRes
                            lineasWithErrors = 1
                        }//fin else

                        linesStatus = append(linesStatus,u);
                    }//end -data

                    if lineas == 0 { //if -line name fields
                        log.Print("MGR Linea 0 de textos")
                        lineas = lineas + 1
                    } //end if -line name fields

                    log.Println(line)

                }//end for

                if  lineasWithErrors ==1 { //al menos una linea  tuvo un error
                    errorGeneral="ERROR FILE"
                }else{
                    if errorGeneral=="" {  

                        log.Print(w, "Files uploaded successfully : ")
                        log.Print(w, files[i].Filename+"\n")
                    }
                }//end else

                //1.count number of lines in the file received
                //2.for each line 
                //      validate the content
                //3.store in the db table AUDIT FILE VALIDATION
                //         seq nbr ,  file name, size, content[all the bytes],commentsParam, validationStatus, validationStatusMessage,validationResponse[for each line,a response OK/Error] ,timestamp
                //3.return result JSON

                //errorGeneral,errorGeneralNbr= ProcessGettokenizedcards(w , requestData) //logicbusiness.go
                log.Println(lineasWithErrors)

            }//end -loop through the files one by one
        
        } //end if error general
    
    }//end if error general1


	/// END
    if errorGeneral!=""{
    	//send error response if any
    	//prepare an error JSON Response, if any
		log.Print("CZ   STEP Get the ERROR response JSON ready")
		
		// START
        //old  getJsonResponseError(errorGeneral, errorGeneralNbr)

        fieldDataBytesJson,err := getJsonResponseErrorValidateFile(errorGeneral, errorGeneralNbr, linesStatus  )  //logicresponse.go 
        
        //////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		if(err!=nil){
			
		}
	
    }else{
        var  cardTokenized modelito.Card
        fieldDataBytesJson,err := getJsonResponseValidateFileV2(cardTokenized)
        w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		if(err!=nil){

	    }//end if

    }


}//end function handlePostConsultaHistorialToken

//func handlePostConsultaHistorialPagos

func handlePostConsultaHistorialPagos(w http.ResponseWriter, r *http.Request) {
	
    defer func() {
		db.Connection.Close(nil)
	}()
    var errorGeneral string
    var errorGeneralNbr string

    //var requestData modelito.RequestTokenizedCards


    errorGeneral=""
    linesStatus := []modelito.ExitoDataValidaLine{}   //structure to stire the errors in each of the liens of the file

	//    requestData, errorGeneral=obtainPostParmsGettokenizedcards(r,errorGeneral) //logicrequest_post.go
    if errorGeneral!="" {
        log.Print("CZ    Prepare Response with 100. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:100 -Missing parameter"	+errorGeneral
    	errorGeneralNbr="100"
    }


	////////////////////////////////////////////////process business rules
	/// START
    
    if errorGeneral=="" {//if errorGeneral1

        log.Print("CZ   STEP Get the File")
        log.Print("File Upload Endpoint Hit")

        // Parse our multipart form, 10 << 20 specifies a maximum
        // upload of 10 MB files.
        err:= r.ParseMultipartForm(10 << 20)
        if err != nil {
            
            log.Print("CZ Error Retrieving the File")
            log.Print(err)
            errorGeneral="ERROR:110 -Error retriving files ,parameters"	+errorGeneral
            errorGeneralNbr="110"

        }

        if errorGeneral=="" {//if error general2
            log.Print("CZ Start read the form data")
            formdata := r.MultipartForm // ok, no problem so far, read the Form data
            
            //get the *fileheaders
            files := formdata.File["file0"] // grab the files, this files was set in the html 

            log.Print("CZ before loop files")

            for i, _ := range files { // loop through the files one by one
                log.Print("CZ loop step 1")
                var elfileindex string

                elfileindex = "file0"
                log.Print("CZ Loop file")
                log.Print("CZ Loop file:"+elfileindex)
                file, err := files[i].Open()
                log.Print("CZ open file")
                defer file.Close()
                if err != nil {
                    log.Print(w, err)
                    errorGeneral="ERROR:120 -Error file passed not open ,parameters"	+errorGeneral
                    errorGeneralNbr="120"

                }

                //convert multipart file into buffer bytes

                buf := bytes.NewBuffer(nil)
                io.Copy(buf, file)
                //if _, err := io.Copy(buf, file); err != nil {
                //      return nil, err
                //}                
                //log.Print(buf) // print the content as 'bytes'

                // convert content to a 'string'
                str := buf.String()

                log.Print(str) // print the content as a 'string'                

                log.Print("MGR paso linea por linea")

                lineas := 0

                lineasWithErrors := 0

                for _, line := range strings.Split(strings.TrimSuffix(str, "\n"), "\n") {//inicio for
                    var u modelito.ExitoDataValidaLine
                    if lineas >= 1{ //if -data
                        log.Printf("MGR Linea %d de datos", lineas)

                        lineas = lineas + 1

                        respuestaRes,cualfallo :=campos_payment (line, lineas)

                        if cualfallo ==0 {  //exito, todos los cmapos de la linea OK
                            u.Line=strconv.Itoa(lineas)
                            u.Status="OK"
                            u.StatusMessage ="SUCESS"
                        }else { //error, al menos un error en la linea
                            u.Line=strconv.Itoa(lineas)
                            u.Status="ERROR550"
                            u.StatusMessage ="ERROR LINEA:"+strconv.Itoa(cualfallo)+" - "+respuestaRes
                            lineasWithErrors = 1
                        }//fin else

                        linesStatus = append(linesStatus,u);
                    }//end -data

                    if lineas == 0 { //if -line name fields
                        log.Print("MGR Linea 0 de textos")
                        lineas = lineas + 1
                    } //end if -line name fields

                    log.Println(line)

                }//end for

                if  lineasWithErrors ==1 { //al menos una linea  tuvo un error
                    errorGeneral="ERROR FILE"
                }else{
                    if errorGeneral=="" {  

                        log.Print(w, "Files uploaded successfully : ")
                        log.Print(w, files[i].Filename+"\n")
                    }
                }//end else

                //1.count number of lines in the file received
                //2.for each line 
                //      validate the content
                //3.store in the db table AUDIT FILE VALIDATION
                //         seq nbr ,  file name, size, content[all the bytes],commentsParam, validationStatus, validationStatusMessage,validationResponse[for each line,a response OK/Error] ,timestamp
                //3.return result JSON

                //errorGeneral,errorGeneralNbr= ProcessGettokenizedcards(w , requestData) //logicbusiness.go
                log.Println(lineasWithErrors)

            }//end -loop through the files one by one
        
        } //end if error general
    
    }//end if error general1


	/// END
    if errorGeneral!=""{
    	//send error response if any
    	//prepare an error JSON Response, if any
		log.Print("CZ   STEP Get the ERROR response JSON ready")
		
		// START
        //old  getJsonResponseError(errorGeneral, errorGeneralNbr)

        fieldDataBytesJson,err := getJsonResponseErrorValidateFile(errorGeneral, errorGeneralNbr, linesStatus  )  //logicresponse.go 
        
        //////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		if(err!=nil){
			   
		}
	
    }else{
        var  cardTokenized modelito.Card
        fieldDataBytesJson,err := getJsonResponseValidateFileV2(cardTokenized)
        w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		if(err!=nil){

	    }//end if

    }
    

}//end function handlePostConsultaHistorialPagos


func ForceDownload(w http.ResponseWriter, r *http.Request) {
    
    log.Print("Empieza a función ForceDownload")

    file := "banwireResponse.txt"
    //downloadBytes, err := ioutil.ReadFile(file)
    
    var errorGeneral string
    
    htmlStrDownloadJson, err:= obtainParmsProcessDownload(r , errorGeneral) //logisrequest.go
    //hacer una func similar a esta func obtainParmsProcessPayment(r *http.Request, errorGeneral string) (modelito.RequestPayment,string){
    // que reciba lo misoomo, y solo busque dos parametros: cualArchivo y lo que viaja en respuestaGeneral que mando el index.html
    //y el indexpay.html 
    
    if(err!=""){

	}//end if
    log.Print("Regreso de función obtainParmsProcessDownload")
        
    //if err != nil {
    //  utilito.LevelLog(Config_env_log, "3",err.tost)
    //}
    downloadBytes:= []byte(htmlStrDownloadJson)

    // set the default MIME type to send
    //mime := http.DetectContentType(downloadBytes)

    //fileSize := len(string(downloadBytes))
    // Generate the server headers

    log.Print("Generador de cabezeras")
		
        w.Header().Set("Content-Type", "text/plain;charset: uft-8")
        w.Header().Set("Content-Disposition", "attachment; filename="+file+"")
		w.Write(downloadBytes)	

                
    //    w.Header().Set("Expires", "0")
    //    w.Header().Set("Content-Transfer-Encoding", "binary")
    //    w.Header().Set("Content-Length", strconv.Itoa(fileSize))
    //    w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

         //b := bytes.NewBuffer(downloadBytes)
         //if _, err := b.WriteTo(w); err != nil {
         //              fmt.Fprintf(w, "%s", err)
         //      }

        //force it down the client's.....
        //http.ServeContent(w, r, file, time.Now(), bytes.NewReader(downloadBytes))

        //http.ServeFile(w, r, "css/app.min.css")
    log.Print("Fin ForceDownload")
 } //end ForceDownload

func ForceDownloadPago(w http.ResponseWriter, r *http.Request) {
    
    log.Print("Empieza funcion ForceDownloadPago")
    file := "banwireResponsePagos.txt"
    //downloadBytes, err := ioutil.ReadFile(file)
    
    var errorGeneral string
    htmlStrDownloadJson, err:= obtainParmsProcessDownloadPagos(r , errorGeneral) //logisrequest.go
    ///hacer una func similar a esta func obtainParmsProcessPayment(r *http.Request, errorGeneral string) (modelito.RequestPayment,string){
    // que reciba lo misoomo, y solo busque dos parametros: cualArchivo y lo que viaja en respuestaGeneral que mando el index.html
    //y el indexpay.html 

    if(err!=""){

	}//end if

    log.Print("Función obtainParmsProcessDownloadPagos")
        
    
        //if err != nil {
        //  utilito.LevelLog(Config_env_log, "3",err.tost)
        //}
    downloadBytes:= []byte(htmlStrDownloadJson)
        
    // set the default MIME type to send
    //mime := http.DetectContentType(downloadBytes)

    //fileSize := len(string(downloadBytes))
    // Generate the server headers

	log.Print("Paso Generador de cabezeras")
        w.Header().Set("Content-Type", "text/plain;charset: uft-8")
        w.Header().Set("Content-Disposition", "attachment; filename="+file+"")
		w.Write(downloadBytes)	

                
    //    w.Header().Set("Expires", "0")
    //    w.Header().Set("Content-Transfer-Encoding", "binary")
    //    w.Header().Set("Content-Length", strconv.Itoa(fileSize))
    //    w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

         //b := bytes.NewBuffer(downloadBytes)
         //if _, err := b.WriteTo(w); err != nil {
         //              fmt.Fprintf(w, "%s", err)
         //      }

    //force it down the client's.....
    //http.ServeContent(w, r, file, time.Now(), bytes.NewReader(downloadBytes))

    //http.ServeFile(w, r, "css/app.min.css")
    log.Print("Fin ForceDownloadPago")

} //end ForceDownloadPago

//Función ForceDownloadValida para index
func ForceDownloadValida(w http.ResponseWriter, r *http.Request) {
    
    log.Print("Empieza funcion ForceDownloadValida")
    file := "banwireResponseValidacion.txt"
    //downloadBytes, err := ioutil.ReadFile(file)
    
    var errorGeneral string
    htmlStrDownloadJson, err:= obtainParmsProcessDownloadValida(r , errorGeneral) //logisrequest.go
    ///hacer una func similar a esta func obtainParmsProcessPayment(r *http.Request, errorGeneral string) (modelito.RequestPayment,string){
    // que reciba lo misoomo, y solo busque dos parametros: cualArchivo y lo que viaja en respuestaGeneral que mando el index.html
    //y el indexpay.html 

    if(err!=""){

	}//end if

    log.Print("Termina función obtainParmsProcessDownloadValida")
        
    
        //if err != nil {
        //  utilito.LevelLog(Config_env_log, "3",err.tost)
        //}
    downloadBytes:= []byte(htmlStrDownloadJson)
        
         // set the default MIME type to send
         //mime := http.DetectContentType(downloadBytes)

         //fileSize := len(string(downloadBytes))
         // Generate the server headers

	log.Print("Paso Generador de cabezeras")
        w.Header().Set("Content-Type", "text/plain;charset: uft-8")
        w.Header().Set("Content-Disposition", "attachment; filename="+file+"")
		w.Write(downloadBytes)	

                
    //    w.Header().Set("Expires", "0")
    //    w.Header().Set("Content-Transfer-Encoding", "binary")
    //    w.Header().Set("Content-Length", strconv.Itoa(fileSize))
    //    w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

         //b := bytes.NewBuffer(downloadBytes)
         //if _, err := b.WriteTo(w); err != nil {
         //              fmt.Fprintf(w, "%s", err)
         //      }

         //force it down the client's.....
         //http.ServeContent(w, r, file, time.Now(), bytes.NewReader(downloadBytes))

    //http.ServeFile(w, r, "css/app.min.css")
    log.Print("Fin ForceDownloadValida")
} //end ForceDownloadValida

//Función ForceDownloadTokeniza para index
func ForceDownloadTokeniza(w http.ResponseWriter, r *http.Request) {
    
    log.Print("Empieza funcion ForceDownloadTokeniza")
    file := "banwireResponseTokenizacion.txt"
    //downloadBytes, err := ioutil.ReadFile(file)
    
    var errorGeneral string
    
    htmlStrDownloadJson, err:= obtainParmsProcessDownloadTokeniza(r , errorGeneral) //logisrequest.go
    ///hacer una func similar a esta func obtainParmsProcessPayment(r *http.Request, errorGeneral string) (modelito.RequestPayment,string){
    // que reciba lo misoomo, y solo busque dos parametros: cualArchivo y lo que viaja en respuestaGeneral que mando el index.html
    //y el indexpay.html 
       
    log.Print("Termina función obtainParmsProcessDownloadTokeniza")
        
    if(err!=""){

	}//end if
        
        //if err != nil {
            //  utilito.LevelLog(Config_env_log, "3",err.tost)
        //}
    downloadBytes:= []byte(htmlStrDownloadJson)
        
         // set the default MIME type to send
         //mime := http.DetectContentType(downloadBytes)

         //fileSize := len(string(downloadBytes))
         // Generate the server headers

		log.Print("Paso Generador de cabezeras")
        w.Header().Set("Content-Type", "text/plain;charset: uft-8")
        w.Header().Set("Content-Disposition", "attachment; filename="+file+"")
		w.Write(downloadBytes)	

                
    //    w.Header().Set("Expires", "0")
    //    w.Header().Set("Content-Transfer-Encoding", "binary")
    //    w.Header().Set("Content-Length", strconv.Itoa(fileSize))
    //    w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

         //b := bytes.NewBuffer(downloadBytes)
         //if _, err := b.WriteTo(w); err != nil {
         //              fmt.Fprintf(w, "%s", err)
         //      }

         // force it down the client's.....
    //     http.ServeContent(w, r, file, time.Now(), bytes.NewReader(downloadBytes))

    //      http.ServeFile(w, r, "css/app.min.css")
    log.Print("Fin ForceDownloadTokeniza")
} //end ForceDownloadTokeniza



// handlePostConsultaTokens  receive and handle the request from client, access DB, and web
func handlePostConsultaTokens(w http.ResponseWriter, r *http.Request) {
	
    log.Print("Entra a handlePostConsultaTokens funcion boton \"Consultar token\" indexconsulta")
    file := "ConsultarToken.txt"
    
    var errorGeneral string
    
    htmlStrDownloadJson, err:= obtainParmsConsultarTokensPOST(r , errorGeneral) //logisrequest.go
    
    if(err!=""){

	}//end if
    log.Print("Regreso de función obtainParmsConsultarTokens")
    
    downloadBytes:= []byte(htmlStrDownloadJson)

    // Generate the server headers

    log.Print("Generador de cabezeras")
		
        w.Header().Set("Content-Type", "text/plain;charset: uft-8")
        w.Header().Set("Content-Disposition", "attachment; filename="+file+"")
		w.Write(downloadBytes)	

    log.Print("Fin obtainParmsConsultarTokens")

    				
}//end handlePostConsultaTokens

func handleGetConsultahistorialClientes(w http.ResponseWriter, r *http.Request) {
    defer func() {
		db.Connection.Close(nil)
	}()
    
    var errorGeneral string
    var errorGeneralNbr string
	
    log.Print("Entra a handleGetConsultahistorialClientes funcion boton \"Consultar\" indexconsultafiles")
    //file := "ConsultarToken.txt"
    
    
    paramsReceived, err:= obtainParmsConsultarHistorialClientes(r , errorGeneral) //logisrequest.go
    
    if(err!=""){

	}//end if
    log.Print("Regreso de función obtainParmsConsultarHistorialClientes")
    
    // downloadBytes:= []byte(paramsReceived)

    // Generate the server headers

    log.Print("Generador de cabezeras")


    log.Print("Fin obtainParmsConsultarHistorial")

    /*log.Print("Entra a handlePostConsultaTokens funcion boton \"Consultar token\" indexconsulta")

   	var requestData modelito.RequestTokenizedCards

    errorGeneral = ""
    requestData, errorGeneral = obtainParmsConsultarTokens(r,errorGeneral) //logicrequest.go
    
    
    log.Print("Regresa de obtainParmsConsultarTokens")


	////////////////////////////////////////////////process business rules
	/// START
    if errorGeneral=="" {

		errorGeneral,errorGeneralNbr= ProcessGettokenizedcards(w , requestData) //logicbusiness.go
	}
	/// END
*/
    if errorGeneral != ""{
    	//send error response if any
    	//prepare an error JSON Response, if any
		log.Print("CZ   STEP Get the ERROR response JSON ready")
		
			/// START
		fieldDataBytesJson,err := getJsonResponseError(errorGeneral, errorGeneralNbr)
		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		
        		
        if(err!=nil){
			
	    }
	
    }else{
    	//send error response if any
    	//prepare an error JSON Response, if any
		log.Print("CZ   STEP Get the OK response JSON ready")
		errorGeneralNbr ="500"
			/// START
		fieldDataBytesJson,err := getJsonResponseConsultarHistorailClientes(paramsReceived, errorGeneralNbr) //logicresponse.go
		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		
        		
        if(err!=nil){
			
	    }

    }


}//end handleGetConsultahistorialClientes


