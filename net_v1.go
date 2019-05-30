package main

import (
	"net/http"
    "strconv"
    "strings"
//    "regexp"
//	"fmt"
	//"log"
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
        utilito.LevelLog(Config_env_log, "3", "cz  init net_v1")

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

    r.Handle("/v1/validatefiles", netHandle(handlePostVaidateFiles, nil)).Methods("POST")   //logicbusiness.go
    r.Handle("/v1/validatepaymentfiles", netHandle(handlePostVaidatePaymentFiles, nil)).Methods("POST")   //logicbusiness.go
    r.Handle("/v1/consultartokens", netHandle(handlePostConsultaTokens, nil)).Methods("POST")   //logicbusiness.go
    r.Handle("/v1/consultarhistorialtokens", netHandle(handlePostConsultaHistorial, nil)).Methods("POST")   //logicbusiness.go

    r.Handle("/v1/processtokenfile", netHandle(handlePostProcessTokenFile, nil)).Methods("POST")   //logicbusiness.go
    r.Handle("/v1/processpaymentfile", netHandle(handlePostProcessPaymentFile, nil)).Methods("POST")   //logicbusiness.go


	r.Handle("/v1/consultarhistorialClientes", netHandle(handlePostConsultahistorialClientes, nil)).Methods("POST")   //logicbusiness.go
    r.Handle("/v1/consultahistorialToken", netHandle(handlePostConsultaHistorialToken, nil)).Methods("POST")   //logicbusiness.go
    r.Handle("/v1/consultarhistorialPagos", netHandle(handlePostConsultaHistorialPagos, nil)).Methods("POST")   //logicbusiness.go


	r.Handle("/v1/fetchtokenizedcards", netHandle(handleDBPostGettokenizedcards, nil)).Methods("POST")   //logicbusiness.go
	r.Handle("/v1/processpayment", netHandle(v4handleDBPostProcesspayment, nil)).Methods("POST")           //logicbusiness.go    	   
	r.Handle("/v1/generatetokenized", netHandle(handleDBPostGeneratetokenized, nil)).Methods("POST")     //logicbusiness.go


}


//index html angular

func dash01Handler(w http.ResponseWriter, r *http.Request) {
    utilito.LevelLog(Config_env_log, "1", "cz  dash01Handler with param"+Config_env_url)
 
    http.ServeFile(w,r,"index.html")
/*     data := TodoPageData{
			PageTitle: Config_env_server,
     }
     tmpl := template.Must(template.ParseFiles("index.html"))
     tmpl.Execute(w, data)
*/     
    utilito.LevelLog(Config_env_log, "3", "CZ   STEP dash01Handler 01")
}

//index html angular

func dash01payHandler(w http.ResponseWriter, r *http.Request) {
    utilito.LevelLog(Config_env_log, "3", "cz  dash01payHandler with param")
    http.ServeFile(w,r,"indexpay.html")
/*     data := TodoPageData{
			PageTitle: Config_env_server,
     }
     tmpl := template.Must(template.ParseFiles("index.html"))
     tmpl.Execute(w, data)
*/     
    utilito.LevelLog(Config_env_log, "3", "CZ   STEP dash01payHandler 01")
}


//index html angular

func dash01consultaHandler(w http.ResponseWriter, r *http.Request) {
    utilito.LevelLog(Config_env_log, "3", "cz  dash01consultaHandler with param")
//    log.Print("cz  dash01Handler with param"+Config_env_url)
    http.ServeFile(w,r,"indexconsulta.html")
/*     data := TodoPageData{
			PageTitle: Config_env_server,
     }
     tmpl := template.Must(template.ParseFiles("index.html"))
     tmpl.Execute(w, data)
*/     
    utilito.LevelLog(Config_env_log, "3", "CZ   STEP dash01consultaHandler 01")
}


//index html angular

func dash01consultafilesHandler(w http.ResponseWriter, r *http.Request) {
    utilito.LevelLog(Config_env_log, "3", "cz  dash01consultafilesHandler with param")
//    log.Print("cz  dash01Handler with param"+Config_env_url)
    http.ServeFile(w,r,"indexconsultafiles.html")
/*     data := TodoPageData{
			PageTitle: Config_env_server,
     }
     tmpl := template.Must(template.ParseFiles("index.html"))
     tmpl.Execute(w, data)
*/     
    utilito.LevelLog(Config_env_log, "3", "CZ   STEP dash01consultafilesHandler 01")
}


func serveCss01(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/app.min.css")
}

func serveCss02(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/bootstrap-theme.min.css")
}

func serveCss03(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/bootstrap.min.css")
}

func serveCss04(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/font-awesome.min.css")
}

func serveCss05(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/ngToast.min.css")
}

func serveCss06(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/nya-bs-select.min.css")
}

func serveCss07(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/ui-bootstrap-csp.css")
}

func serveJs01(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "js/angular.min.js")
}


   //post

   
  //post

//  handlePostProcessTokenFile

func handlePostProcessTokenFile(w http.ResponseWriter, r *http.Request) {
       	defer func() {
		db.Connection.Close(nil)
	}()
    var errorGeneral string
    var errorGeneralNbr string



    
//   	var requestData modelito.RequestTokenizedCards


    errorGeneral=""

    linesStatus := []modelito.ExitoDataTokenLine{}   //structure to stire the errors in each of the liens of the file

//    requestData, errorGeneral=obtainPostParmsGettokenizedcards(r,errorGeneral) //logicrequest_post.go
    if errorGeneral!="" {
        utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 100. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:100 -Missing parameter"	+errorGeneral
    	errorGeneralNbr="100"
    }
	////////////////////////////////////////////////process business rules
	/// START
    if errorGeneral=="" {

        utilito.LevelLog(Config_env_log, "3", "CZ  ProcessTokenFile  STEP Get the File")
        utilito.LevelLog(Config_env_log, "3", " ProcessTokenFile File Upload Endpoint Hit")

        // Parse our multipart form, 10 << 20 specifies a maximum
        // upload of 10 MB files.
        err:= r.ParseMultipartForm(10 << 20)
        if err != nil {
            
            utilito.LevelLog(Config_env_log, "3", "CZ ProcessTokenFile Error Retrieving the File")
            utilito.LevelLog(Config_env_log, "3", err.Error())
            errorGeneral="ERROR:110 -Error retriving files ,parameters"	+errorGeneral
            errorGeneralNbr="110"

        } 
        if errorGeneral=="" {  
            utilito.LevelLog(Config_env_log, "3", "CZ ProcessTokenFile Start read the form data")
            formdata := r.MultipartForm // ok, no problem so far, read the Form data

            //get the *fileheaders
            files := formdata.File["file0"] // grab the files, this files was set in the html 

            utilito.LevelLog(Config_env_log, "3", "CZ ProcessTokenFile before loop files")

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
                    errorGeneral="ERROR:120 -Error file passed not open ,parameters"	+errorGeneral
                    errorGeneralNbr="120"

                }
                //convert multipart file into buffer bytes

                buf := bytes.NewBuffer(nil)
                io.Copy(buf, file)

                micadenita := buf.String()

                utilito.LevelLog(Config_env_log, "3", micadenita)

                utilito.LevelLog(Config_env_log, "3", "ProcessTokenFile MGR Paso linea por linea index")

                lineas := 0

                lineasWithErrors := 0
                for _, line := range strings.Split(strings.TrimSuffix(micadenita, "\n"), "\n") {
                    var u modelito.ExitoDataTokenLine
                    if lineas >= 1{
                        utilito.LevelLog(Config_env_log, "3", "MGR linea de datos")

                        lineas = lineas + 1
                         respuestaRes,cualfallo :=campos_token (line, lineas)
                         if cualfallo ==0 {  //exito, todos los cmapos de la linea OK, y no errores previos
                            u.Line=strconv.Itoa(lineas)
                            u.Status="OK"
                            u.StatusMessage ="SUCESS"

                         }else { //error, al menos un error en la linea
                            u.Line=strconv.Itoa(lineas)
                            u.Status="ERROR540"
                            u.StatusMessage ="ERROR LINEA:"+strconv.Itoa(cualfallo)+" - "+respuestaRes
                            lineasWithErrors =1
                         }
                         linesStatus = append(linesStatus,u);
                    }
                    
                    if lineas == 0 {
                        utilito.LevelLog(Config_env_log, "3","MGR Nombres de campos")
                        lineas = lineas + 1
                    }

                        utilito.LevelLog(Config_env_log, "3", line)
                       
                        

                }//end  -loop through the lines

               if  lineasWithErrors ==1 { //al menos una linea  tuvo un error
                     errorGeneral="ERROR FILE"
                     errorGeneralNbr="540"
                   
               }else{
                    if errorGeneral=="" {  
    
                        utilito.LevelLog(Config_env_log, "3", " ProcessTokenFile Files uploaded successfully : ")
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
//		errorGeneral,errorGeneralNbr= ProcessGettokenizedcards(w , requestData) //logicbusiness.go
	}


	/// END
    if errorGeneral!=""{
    	//send error response if any
    	//prepare an error JSON Response, if any
		utilito.LevelLog(Config_env_log, "3", "CZ ProcessTokenFile  STEP Get the ERROR response JSON ready")
		
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

}


func handlePostVaidateFiles(w http.ResponseWriter, r *http.Request) {
	defer func() {
		db.Connection.Close(nil)
	}()
    var errorGeneral string
    var errorGeneralNbr string



    
//   	var requestData modelito.RequestTokenizedCards


    errorGeneral=""

    linesStatus := []modelito.ExitoDataTokenLine{}   //structure to stire the errors in each of the liens of the file

//    requestData, errorGeneral=obtainPostParmsGettokenizedcards(r,errorGeneral) //logicrequest_post.go
    if errorGeneral!="" {
        utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 100. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:100 -Missing parameter"	+errorGeneral
    	errorGeneralNbr="100"
    }
	////////////////////////////////////////////////process business rules
	/// START
    if errorGeneral=="" {

        utilito.LevelLog(Config_env_log, "3", "CZ   STEP Get the File")
        utilito.LevelLog(Config_env_log, "3", "File Upload Endpoint Hit")

        // Parse our multipart form, 10 << 20 specifies a maximum
        // upload of 10 MB files.
        err:= r.ParseMultipartForm(10 << 20)
        if err != nil {
            
            utilito.LevelLog(Config_env_log, "3", "CZ Error Retrieving the File")
            utilito.LevelLog(Config_env_log, "3",  err.Error())
            errorGeneral="ERROR:110 -Error retriving files ,parameters"	+errorGeneral
            errorGeneralNbr="110"

        } 
        if errorGeneral=="" {  
            utilito.LevelLog(Config_env_log, "3", "CZ Start read the form data")
            formdata := r.MultipartForm // ok, no problem so far, read the Form data

            //get the *fileheaders
            files := formdata.File["file0"] // grab the files, this files was set in the html 

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
                    errorGeneral="ERROR:120 -Error file passed not open ,parameters"	+errorGeneral
                    errorGeneralNbr="120"

                }
                //convert multipart file into buffer bytes

                buf := bytes.NewBuffer(nil)
                io.Copy(buf, file)

                micadenita := buf.String()

                utilito.LevelLog(Config_env_log, "3", micadenita)

                utilito.LevelLog(Config_env_log, "3", "MGR Paso linea por linea index")

                lineas := 0

                lineasWithErrors := 0
                for _, line := range strings.Split(strings.TrimSuffix(micadenita, "\n"), "\n") {
                    var u modelito.ExitoDataTokenLine
                    if lineas >= 1{
                        utilito.LevelLog(Config_env_log, "3", "MGR linea de datos")

                        lineas = lineas + 1
                         respuestaRes,cualfallo :=campos_token (line, lineas)
                         if cualfallo ==0 {  //exito, todos los cmapos de la linea OK, y no errores previos
                            u.Line=strconv.Itoa(lineas)
                            u.Status="OK"
                            u.StatusMessage ="SUCESS"

                         }else { //error, al menos un error en la linea
                            u.Line=strconv.Itoa(lineas)
                            u.Status="ERROR540"
                            u.StatusMessage ="ERROR FIELD:"+strconv.Itoa(cualfallo)+" - "+respuestaRes
                            lineasWithErrors =1
                         }
                         linesStatus = append(linesStatus,u);
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
//		errorGeneral,errorGeneralNbr= ProcessGettokenizedcards(w , requestData) //logicbusiness.go
	}


	/// END
    if errorGeneral!=""{
    	//send error response if any
    	//prepare an error JSON Response, if any
		utilito.LevelLog(Config_env_log, "3", "CZ   STEP Get the ERROR response JSON ready")
		
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
					
}//en function validate

//  handlePostProcessTokenFile

func handlePostProcessPaymentFile(w http.ResponseWriter, r *http.Request) {
       	defer func() {
		db.Connection.Close(nil)
	}()
    var errorGeneral string
    var errorGeneralNbr string



    
//   	var requestData modelito.RequestTokenizedCards


    errorGeneral=""

    linesStatus := []modelito.ExitoDataTokenLine{}   //structure to stire the errors in each of the liens of the file

//    requestData, errorGeneral=obtainPostParmsGettokenizedcards(r,errorGeneral) //logicrequest_post.go
    if errorGeneral!="" {
        utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 100. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:100 -Missing parameter"	+errorGeneral
    	errorGeneralNbr="100"
    }
	////////////////////////////////////////////////process business rules
	/// START
    if errorGeneral=="" {

        utilito.LevelLog(Config_env_log, "3", "CZ  ProcessTokenFile  STEP Get the File")
        utilito.LevelLog(Config_env_log, "3", " ProcessTokenFile File Upload Endpoint Hit")

        // Parse our multipart form, 10 << 20 specifies a maximum
        // upload of 10 MB files.
        err:= r.ParseMultipartForm(10 << 20)
        if err != nil {
            
            utilito.LevelLog(Config_env_log, "3", "CZ ProcessTokenFile Error Retrieving the File")
            utilito.LevelLog(Config_env_log, "3",  err.Error())
            errorGeneral="ERROR:110 -Error retriving files ,parameters"	+errorGeneral
            errorGeneralNbr="110"

        } 
        if errorGeneral=="" {  
            utilito.LevelLog(Config_env_log, "3", "CZ ProcessTokenFile Start read the form data")
            formdata := r.MultipartForm // ok, no problem so far, read the Form data

            //get the *fileheaders
            files := formdata.File["file0"] // grab the files, this files was set in the html 

            utilito.LevelLog(Config_env_log, "3", "CZ ProcessTokenFile before loop files")

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
                    errorGeneral="ERROR:120 -Error file passed not open ,parameters"	+errorGeneral
                    errorGeneralNbr="120"

                }
                //convert multipart file into buffer bytes

                buf := bytes.NewBuffer(nil)
                io.Copy(buf, file)

                micadenita := buf.String()

                utilito.LevelLog(Config_env_log, "3", micadenita)

                utilito.LevelLog(Config_env_log, "3", "ProcessTokenFile MGR Paso linea por linea index")

                lineas := 0

                lineasWithErrors := 0
                for _, line := range strings.Split(strings.TrimSuffix(micadenita, "\n"), "\n") {
                    var u modelito.ExitoDataTokenLine
                    if lineas >= 1{
                        utilito.LevelLog(Config_env_log, "3", "MGR linea datos")

                        lineas = lineas + 1
                         respuestaRes,cualfallo :=campos_payment (line, lineas)
                         if cualfallo ==0 {  //exito, todos los cmapos de la linea OK, y no errores previos
                            u.Line=strconv.Itoa(lineas)
                            u.Status="OK"
                            u.StatusMessage ="SUCESS"

                         }else { //error, al menos un error en la linea
                            u.Line=strconv.Itoa(lineas)
                            u.Status="ERROR540"
                            u.StatusMessage ="ERROR LINEA:"+strconv.Itoa(cualfallo)+" - "+respuestaRes
                            lineasWithErrors =1
                         }
                         linesStatus = append(linesStatus,u);
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
    
                        utilito.LevelLog(Config_env_log, "3", " ProcessTokenFile Files uploaded successfully : ")
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
//		errorGeneral,errorGeneralNbr= ProcessGettokenizedcards(w , requestData) //logicbusiness.go
	}


	/// END
    if errorGeneral!=""{
    	//send error response if any
    	//prepare an error JSON Response, if any
		utilito.LevelLog(Config_env_log, "3", "CZ ProcessTokenFile  STEP Get the ERROR response JSON ready")
		
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

}


func handlePostVaidatePaymentFiles(w http.ResponseWriter, r *http.Request) {
	defer func() {
		db.Connection.Close(nil)
	}()
    var errorGeneral string
    var errorGeneralNbr string



    
//   	var requestData modelito.RequestTokenizedCards


    errorGeneral=""

    linesStatus := []modelito.ExitoDataTokenLine{}   //structure to stire the errors in each of the liens of the file

//    requestData, errorGeneral=obtainPostParmsGettokenizedcards(r,errorGeneral) //logicrequest_post.go
    if errorGeneral!="" {
        utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 100. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:100 -Missing parameter"	+errorGeneral
    	errorGeneralNbr="100"
    }
	////////////////////////////////////////////////process business rules
	/// START
    if errorGeneral=="" {

        utilito.LevelLog(Config_env_log, "3", "CZ   STEP Get the File")
        utilito.LevelLog(Config_env_log, "3", "File Upload Endpoint Hit")

        // Parse our multipart form, 10 << 20 specifies a maximum
        // upload of 10 MB files.
        err:= r.ParseMultipartForm(10 << 20)
        if err != nil {
            
            utilito.LevelLog(Config_env_log, "3", "CZ Error Retrieving the File")
            utilito.LevelLog(Config_env_log, "3",  err.Error())
            errorGeneral="ERROR:110 -Error retriving files ,parameters"	+errorGeneral
            errorGeneralNbr="110"

        } 
        if errorGeneral=="" {  
            utilito.LevelLog(Config_env_log, "3", "CZ Start read the form data")
            formdata := r.MultipartForm // ok, no problem so far, read the Form data

            //get the *fileheaders
            files := formdata.File["file0"] // grab the files, this files was set in the html 

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
                    errorGeneral="ERROR:120 -Error file passed not open ,parameters"	+errorGeneral
                    errorGeneralNbr="120"

                }
                //convert multipart file into buffer bytes

                buf := bytes.NewBuffer(nil)
                io.Copy(buf, file)

                micadenita := buf.String()

                utilito.LevelLog(Config_env_log, "3", micadenita)

                utilito.LevelLog(Config_env_log, "3", "MGR Paso linea por linea index")

                lineas := 0

                lineasWithErrors := 0
                for _, line := range strings.Split(strings.TrimSuffix(micadenita, "\n"), "\n") {
                    var u modelito.ExitoDataTokenLine
                    if lineas >= 1{
                        utilito.LevelLog(Config_env_log, "3", "MGR Linea de datos")

                        lineas = lineas + 1
                         respuestaRes,cualfallo :=campos_payment (line, lineas)
                         if cualfallo ==0 {  //exito, todos los cmapos de la linea OK, y no errores previos
                            u.Line=strconv.Itoa(lineas)
                            u.Status="OK"
                            u.StatusMessage ="SUCESS"

                         }else { //error, al menos un error en la linea
                            u.Line=strconv.Itoa(lineas)
                            u.Status="ERROR540"
                            u.StatusMessage ="ERROR FIELD:"+strconv.Itoa(cualfallo)+" - "+respuestaRes
                            lineasWithErrors =1
                         }
                         linesStatus = append(linesStatus,u);
                    }
                    
                    if lineas == 0 {
                        utilito.LevelLog(Config_env_log, "3", "MGR Linea nombre de campos")
                        lineas = lineas + 1
                    }

                        utilito.LevelLog(Config_env_log, "3", line)
                       
                        

                }//end  -loop through the lines

               if  lineasWithErrors ==1 { //al menos una linea  tuvo un error
                     errorGeneral="ERROR FILE"
                     errorGeneralNbr="540"
                   
               }else{
                    if errorGeneral=="" {  
    
                        utilito.LevelLog(Config_env_log, "3",  "Files uploaded successfully : ")
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
//		errorGeneral,errorGeneralNbr= ProcessGettokenizedcards(w , requestData) //logicbusiness.go
	}


	/// END
    if errorGeneral!=""{
    	//send error response if any
    	//prepare an error JSON Response, if any
		utilito.LevelLog(Config_env_log, "3", "CZ   STEP Get the ERROR response JSON ready")
		
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
					
}//en function handlePostVaidatePaymentFiles



   
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
		utilito.LevelLog(Config_env_log, "3", "CZ   STEP Get the ERROR response JSON ready")
		
			/// START
		fieldDataBytesJson,err := getJsonResponseError(errorGeneral, errorGeneralNbr)
		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		if(err!=nil){
			
		}
	
    } 
					
}




// handleGeneratetokenized for receive and handle the request from client
func handleDBPostGeneratetokenized(w http.ResponseWriter, r *http.Request) {
	defer func() {
		db.Connection.Close(nil)
	}()
     var requestData modelito.RequestTokenized
     var errorGeneral string
     var errorGeneralNbr string
     
    errorGeneral=""


    requestData,errorGeneral =obtainPostParmsGeneratetokenized(r,errorGeneral)   //logicrequest_post.go



	////////////////////////////////////////////////validate parms
	/// START
    
    if errorGeneral=="" {

		errorGeneral,errorGeneralNbr= ProcessGeneratetokenized(w , requestData)
	}

    if errorGeneral!=""{
    	//send error response if any
    	//prepare an error JSON Response, if any
		utilito.LevelLog(Config_env_log, "3", "CZ   STEP Get the ERROR response JSON ready")
		
			/// START
		fieldDataBytesJson,err := getJsonResponseError(errorGeneral, errorGeneralNbr)
		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		if(err!=nil){
			
		}
	
    } 
					
}



///////////////////////////////v4
///////////////////////////////v4



// v4handleDBProcesspayment  receive and handle the request from client, access DB
func v4handleDBPostProcesspayment(w http.ResponseWriter, r *http.Request) {
	defer func() {
		db.Connection.Close(nil)
	}()
    var errorGeneral string
    var errorGeneralNbr string
    var requestData modelito.RequestPayment
    
    errorGeneral=""
requestData,errorGeneral =obtainPostParmsProcessPayment(r,errorGeneral)  //logicrequest_post.go

	////////////////////////////////////////////////validate parms
	/// START
	////////////////////////////////////////////////validate parms
	/// START
    
    if errorGeneral=="" {

		errorGeneral,errorGeneralNbr= v4ProcessProcessPayment(w , requestData)    //logicbusiness.go 
	}
 
    if errorGeneral!=""{
    	//send error response if any
    	//prepare an error JSON Response, if any
		utilito.LevelLog(Config_env_log, "3", "CZ   STEP Get the ERROR response JSON ready")
		
			/// START
		fieldDataBytesJson,err := getJsonResponseError(errorGeneral, errorGeneralNbr)
		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		if(err!=nil){
			
		}
	
    } 
					
}

   // handlePostConsultaTokens  receive and handle the request from client, access DB, and web
func handlePostConsultaTokens(w http.ResponseWriter, r *http.Request) {
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
		utilito.LevelLog(Config_env_log, "3", "CZ   STEP Get the ERROR response JSON ready")
		
			/// START
		fieldDataBytesJson,err := getJsonResponseError(errorGeneral, errorGeneralNbr)
		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		if(err!=nil){
			
		}
	
    } 
					
}//end handlePostConsultaTokens

   // handlePostConsultaHistorial  receive and handle the request from client, access DB, and web
func handlePostConsultaHistorial(w http.ResponseWriter, r *http.Request) {
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
		utilito.LevelLog(Config_env_log, "3", "CZ   STEP Get the ERROR response JSON ready")
		
			/// START
		fieldDataBytesJson,err := getJsonResponseError(errorGeneral, errorGeneralNbr)
		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		if(err!=nil){
			
		}
	
    } 
					
}//end handlePostConsultaHistorial

//func handlePostConsultahistorialClientes

func handlePostConsultahistorialClientes(w http.ResponseWriter, r *http.Request) {
	defer func() {
		db.Connection.Close(nil)
	}()
    var errorGeneral string
    var errorGeneralNbr string

    
//   	var requestData modelito.RequestTokenizedCards


    errorGeneral=""
    linesStatus := []modelito.ExitoDataTokenLine{}   //structure to stire the errors in each of the liens of the file

	//    requestData, errorGeneral=obtainPostParmsGettokenizedcards(r,errorGeneral) //logicrequest_post.go
    if errorGeneral!="" {
        utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 100. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:100 -Missing parameter"	+errorGeneral
    	errorGeneralNbr="100"
    }


	////////////////////////////////////////////////process business rules
	/// START
    if errorGeneral=="" {//if errorGeneral1

        utilito.LevelLog(Config_env_log, "3", "CZ   STEP Get the File")
        utilito.LevelLog(Config_env_log, "3", "File Upload Endpoint Hit")

        // Parse our multipart form, 10 << 20 specifies a maximum
        // upload of 10 MB files.
        err:= r.ParseMultipartForm(10 << 20)
        if err != nil {
            
            utilito.LevelLog(Config_env_log, "3", "CZ Error Retrieving the File")
            utilito.LevelLog(Config_env_log, "3",  err.Error())
            errorGeneral="ERROR:110 -Error retriving files ,parameters"	+errorGeneral
            errorGeneralNbr="110"

        }

        if errorGeneral=="" {//if error general2
            utilito.LevelLog(Config_env_log, "3", "CZ Start read the form data")
            formdata := r.MultipartForm // ok, no problem so far, read the Form data
            
            //get the *fileheaders
            files := formdata.File["file0"] // grab the files, this files was set in the html 

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

                utilito.LevelLog(Config_env_log, "3", str)// print the content as a 'string'

                utilito.LevelLog(Config_env_log, "3", "MGR paso linea por linea")

                lineas := 0

                lineasWithErrors := 0

                for _, line := range strings.Split(strings.TrimSuffix(str, "\n"), "\n") {//inicio for
                    var u modelito.ExitoDataTokenLine
                    if lineas >= 1{ //if -data
                        utilito.LevelLog(Config_env_log, "3", "MGR Linea de datos")

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
                        utilito.LevelLog(Config_env_log, "3", "MGR Linea 0 de textos")
                        lineas = lineas + 1
                    } //end if -line name fields

                    utilito.LevelLog(Config_env_log, "3", line)

                }//end for

                if  lineasWithErrors ==1 { //al menos una linea  tuvo un error
                    errorGeneral="ERROR FILE"
                }else{
                    if errorGeneral=="" {  

                        utilito.LevelLog(Config_env_log, "3", "Files uploaded successfully : ")
                        utilito.LevelLog(Config_env_log, "3", files[i].Filename+"\n")
                    }
                }//end else

                //1.count number of lines in the file received
                //2.for each line 
                //      validate the content
                //3.store in the db table AUDIT FILE VALIDATION
                //         seq nbr ,  file name, size, content[all the bytes],commentsParam, validationStatus, validationStatusMessage,validationResponse[for each line,a response OK/Error] ,timestamp
                //3.return result JSON

                //errorGeneral,errorGeneralNbr= ProcessGettokenizedcards(w , requestData) //logicbusiness.go
                utilito.LevelLogint(Config_env_log, "3", lineasWithErrors)

            }//end -loop through the files one by one
        
        } //end if error general
    
    }//end if error general1





	/// END
    if errorGeneral!=""{
    	//send error response if any
    	//prepare an error JSON Response, if any
		utilito.LevelLog(Config_env_log, "3", "CZ   STEP Get the ERROR response JSON ready")
		
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


}//end function handlePostConsultahistorialClientes

//func handlePostConsultaHistorialToken

func handlePostConsultaHistorialToken(w http.ResponseWriter, r *http.Request) {
	defer func() {
		db.Connection.Close(nil)
	}()
    var errorGeneral string
    var errorGeneralNbr string

    
//   	var requestData modelito.RequestTokenizedCards


    errorGeneral=""
    linesStatus := []modelito.ExitoDataTokenLine{}   //structure to stire the errors in each of the liens of the file

	//    requestData, errorGeneral=obtainPostParmsGettokenizedcards(r,errorGeneral) //logicrequest_post.go
    if errorGeneral!="" {
        utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 100. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:100 -Missing parameter"	+errorGeneral
    	errorGeneralNbr="100"
    }


	////////////////////////////////////////////////process business rules
	/// START
    if errorGeneral=="" {//if errorGeneral1

        utilito.LevelLog(Config_env_log, "3", "CZ   STEP Get the File")
        utilito.LevelLog(Config_env_log, "3", "File Upload Endpoint Hit")

        // Parse our multipart form, 10 << 20 specifies a maximum
        // upload of 10 MB files.
        err:= r.ParseMultipartForm(10 << 20)
        if err != nil {
            
            utilito.LevelLog(Config_env_log, "3", "CZ Error Retrieving the File")
            utilito.LevelLog(Config_env_log, "3", err.Error())
            errorGeneral="ERROR:110 -Error retriving files ,parameters"	+errorGeneral
            errorGeneralNbr="110"

        }

        if errorGeneral=="" {//if error general2
            utilito.LevelLog(Config_env_log, "3", "CZ Start read the form data")
            formdata := r.MultipartForm // ok, no problem so far, read the Form data
            
            //get the *fileheaders
            files := formdata.File["file0"] // grab the files, this files was set in the html 

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

                utilito.LevelLog(Config_env_log, "3", str)// print the content as a 'string'            

                utilito.LevelLog(Config_env_log, "3", "MGR paso linea por linea")

                lineas := 0

                lineasWithErrors := 0

                for _, line := range strings.Split(strings.TrimSuffix(str, "\n"), "\n") {//inicio for
                    var u modelito.ExitoDataTokenLine
                    if lineas >= 1{ //if -data
                        utilito.LevelLog(Config_env_log, "3", "MGR Linea de datos")

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
                        utilito.LevelLog(Config_env_log, "3", "MGR Linea nombres de campos")
                        lineas = lineas + 1
                    } //end if -line name fields

                    utilito.LevelLog(Config_env_log, "3", line)

                }//end for

                if  lineasWithErrors ==1 { //al menos una linea  tuvo un error
                    errorGeneral="ERROR FILE"
                }else{
                    if errorGeneral=="" {  

                        utilito.LevelLog(Config_env_log, "3",  "Files uploaded successfully : ")
                        utilito.LevelLog(Config_env_log, "3",  files[i].Filename+"\n")
                    }
                }//end else

                //1.count number of lines in the file received
                //2.for each line 
                //      validate the content
                //3.store in the db table AUDIT FILE VALIDATION
                //         seq nbr ,  file name, size, content[all the bytes],commentsParam, validationStatus, validationStatusMessage,validationResponse[for each line,a response OK/Error] ,timestamp
                //3.return result JSON

                //errorGeneral,errorGeneralNbr= ProcessGettokenizedcards(w , requestData) //logicbusiness.go
                utilito.LevelLogint(Config_env_log, "3", lineasWithErrors)

            }//end -loop through the files one by one
        
        } //end if error general
    
    }//end if error general1





	/// END
    if errorGeneral!=""{
    	//send error response if any
    	//prepare an error JSON Response, if any
		utilito.LevelLog(Config_env_log, "3", "CZ   STEP Get the ERROR response JSON ready")
		
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

    
//   	var requestData modelito.RequestTokenizedCards


    errorGeneral=""
    linesStatus := []modelito.ExitoDataTokenLine{}   //structure to stire the errors in each of the liens of the file

	//    requestData, errorGeneral=obtainPostParmsGettokenizedcards(r,errorGeneral) //logicrequest_post.go
    if errorGeneral!="" {
        utilito.LevelLog(Config_env_log, "3", "CZ    Prepare Response with 100. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:100 -Missing parameter"	+errorGeneral
    	errorGeneralNbr="100"
    }


	////////////////////////////////////////////////process business rules
	/// START
    if errorGeneral=="" {//if errorGeneral1

        utilito.LevelLog(Config_env_log, "3", "CZ   STEP Get the File")
        utilito.LevelLog(Config_env_log, "3", "File Upload Endpoint Hit")

        // Parse our multipart form, 10 << 20 specifies a maximum
        // upload of 10 MB files.
        err:= r.ParseMultipartForm(10 << 20)
        if err != nil {
            
            utilito.LevelLog(Config_env_log, "3", "CZ Error Retrieving the File")
            utilito.LevelLog(Config_env_log, "3", err.Error())
            errorGeneral="ERROR:110 -Error retriving files ,parameters"	+errorGeneral
            errorGeneralNbr="110"

        }

        if errorGeneral=="" {//if error general2
            utilito.LevelLog(Config_env_log, "3", "CZ Start read the form data")
            formdata := r.MultipartForm // ok, no problem so far, read the Form data
            
            //get the *fileheaders
            files := formdata.File["file0"] // grab the files, this files was set in the html 

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
                    utilito.LevelLog(Config_env_log, "3", err.Error())
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

                utilito.LevelLog(Config_env_log, "3", str)// print the content as a 'string'             

                utilito.LevelLog(Config_env_log, "3", "MGR paso linea por linea")

                lineas := 0

                lineasWithErrors := 0

                for _, line := range strings.Split(strings.TrimSuffix(str, "\n"), "\n") {//inicio for
                    var u modelito.ExitoDataTokenLine
                    if lineas >= 1{ //if -data
                        utilito.LevelLog(Config_env_log, "3", "MGR Linea de datos")

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
                        utilito.LevelLog(Config_env_log, "3", "MGR Linea nombres de campos")
                        lineas = lineas + 1
                    } //end if -line name fields

                    utilito.LevelLog(Config_env_log, "3", line)

                }//end for

                if  lineasWithErrors ==1 { //al menos una linea  tuvo un error
                    errorGeneral="ERROR FILE"
                }else{
                    if errorGeneral=="" {  

                        utilito.LevelLog(Config_env_log, "3", "Files uploaded successfully : ")
                        utilito.LevelLog(Config_env_log, "3", files[i].Filename+"\n")
                    }
                }//end else

                //1.count number of lines in the file received
                //2.for each line 
                //      validate the content
                //3.store in the db table AUDIT FILE VALIDATION
                //         seq nbr ,  file name, size, content[all the bytes],commentsParam, validationStatus, validationStatusMessage,validationResponse[for each line,a response OK/Error] ,timestamp
                //3.return result JSON

                //errorGeneral,errorGeneralNbr= ProcessGettokenizedcards(w , requestData) //logicbusiness.go
                utilito.LevelLogint(Config_env_log, "3", lineasWithErrors)

            }//end -loop through the files one by one
        
        } //end if error general
    
    }//end if error general1





	/// END
    if errorGeneral!=""{
    	//send error response if any
    	//prepare an error JSON Response, if any
		utilito.LevelLog(Config_env_log, "3", "CZ   STEP Get the ERROR response JSON ready")
		
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
//              campoValue = strings.Replace(campo, "\n", "", -1) // only works with a single character
//              re := regexp.MustCompile(`\r?\n`)
//              campoValue = re.ReplaceAllString(campoValue, "y")

              campoValue = strings.Replace(campo, "\"", "", -1) // only works with a single character
              var largo string
              largo = strconv.Itoa ( len(campoValue))
              utilito.LevelLog(Config_env_log, "3", "largo del campo es:"+largo+":valor del campo es:"+campoValue)
              
              resultado, cualfallo = valida_campo_token(campoValue, numcampos)

              if cualfallo >0 {
                  
                  utilito.LevelLog(Config_env_log, "3", "fallo es valor en :"+campo)
                  break
              }
        }
        
        return resultado, cualfallo
}

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
                    }

                }else{
                    resultado = "External Identifier is required"
                    cualfallo = 1
                }
            }
            
            if numcampos == 2{
            	if campo != "" {
					if len(campo) >100 {
	
						resultado="Customer reference max lenght is 100"
                        cualfallo = 2
			        }
				}else{
					resultado="Client reference is required"
                    cualfallo = 2
		        }
                
            }

            if numcampos == 3{
                if campo != "" {
					if len(campo) >100 {
	
						resultado="Payment reference max lenght is 100"
                        cualfallo = 3
			        }
				}else{
					resultado="Payment reference is required"
                    cualfallo = 3
		        }
            }

            if numcampos == 4{
                if campo != "" {
					if len(campo)==16 || len(campo)==15{
	
					}else{
						resultado="Card Number must be 16 digits:"+campo
                        cualfallo = 4
			        }
				}else{
					resultado="Card is required"
                    cualfallo = 4
		        }
            }

            if numcampos == 5{
                utilito.LevelLog(Config_env_log, "3", "\n")
                if campo != "" {
					if  len(campo)==4 || len(campo)==5 { // 2 for the double quotes and 1 for the end of line
	
					}else{
						resultado="Valid Thru  4 digits"
                        cualfallo = 5
			        }
				}else{
					resultado="Valid Thru is required"
                    cualfallo = 5
		        }
            }
    return resultado, cualfallo
}

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
}

func valida_campo_pay (campo string, numcampos int)(string, int){
    utilito.LevelLog(Config_env_log, "3", "MGR valida campo payment")

    var resultado string
    var cualfallo int 
    cualfallo = 0
            if numcampos == 1{
                if campo != "" {
                    if len(campo) > 30 {
                        resultado = "External identifier max leng is 30"
                        cualfallo = 1
                    }

                }else{
                    resultado = "External Identifier is required"
                    cualfallo = 1
                }
            }
            
            if numcampos == 2{
            	if campo != "" {
					if len(campo) >100 {
	
						resultado="Client reference is required"
                        cualfallo = 2
			        }
				}else{
					resultado="Client reference is required"
                    cualfallo = 2
		        }

            }

            if numcampos == 3{
                if campo != "" {
					if len(campo) >100 {
	
						resultado="Payment reference max lenght is 100"
                        cualfallo = 3
			        }
				}else{
					resultado="Payment reference is required"
                    cualfallo = 3
		        }

            }

            if numcampos == 4{
                if campo != "" {

				}else{
					resultado="Token is required"
                    cualfallo = 4
		        }
            }

            if numcampos == 5{
                           
                if campo != "" {
					if len(campo)==3 ||  len(campo)==4 {
	
					}else{
						resultado="Cvv must be 3 or 4 digits"
                        cualfallo = 5
			        }
				}else{
					resultado="Cvv is required"
                    cualfallo = 5
		        }
            } 
            
			if numcampos == 6{
                if campo != "" {

				}else{
					resultado="Amount is required"
                    cualfallo = 6
		        }
            }	

    return resultado, cualfallo
}