package model
import (
    "database/sql"
//    "errors"
	"fmt"
	"log"

	
)

type Filetrans struct {
    ID    string    `sql:"type:bigserial"`
    Transtype  string `sql:"type:varchar(20)`
    Filename  string `sql:"type:varchar(100)`
    Trans_status  string `sql:"type:varchar(100)`
    Trans_statusmssg  string `sql:"type:varchar(100)`
    Trans_processstatus string `sql:"type:varchar(100)`
    TransCreated_at   string    `sql:"type:timestamp`
    Trans_user string `sql:"type:varchar(50)`
    Trans_data_received string `sql:"type:varchar(2500)`
    Trans_val_response string `sql:"type:varchar(2500)`
    Trans_process_responser string `sql:"type:varchar(2500)`
    Trans_process_qty string `sql:"type:varchar(2500)`
    
/*
    id_file bigserial,
transaction_type  varchar (20),  --type is validate token, validate payment,  process tokenizer, process payment
name_file    varchar(100),   
transaction_validation_status    varchar(100),   --OK, ERROR, 
transaction_validation_status_mssg    varchar(100),   --OK, ERROR, 
transaction_process_status    varchar(100),   --PENDING, PROCCESED 
transaction_date             timestamp, 
transaction_user           varchar (50),
transaction_data_received   varchar (2500),     --data received in the file [do not store the TDC number/Valid Throug]
transaction_validation_response  varchar(2500), --data send bakc as response for the validation
transaction_process_response  varchar(2500),    --data send bakc as response for the process, idication each line if Ok, or ERROR
transaction_processed_qty  varchar(2500)      ---nuumber of cards tokenized with SUCCESS or number of payments  SUCCESSFULLY
*/
}
func (u *Filetrans) getFiletrans(db *sql.DB) error {
    statement := fmt.Sprintf("SELECT name_file, transaction_type FROM banwirefiletransaction WHERE id_file=%d", u.ID)
    return db.QueryRow(statement).Scan(&u.Filename, &u.Transtype)
}



func (u *Filetrans) CreateFiletrans(db *sql.DB) error {

	    statement := fmt.Sprintf("INSERT INTO banwirefiletransaction( "+
        "transaction_type, "+
        "name_file , "+
        "transaction_validation_status ,   "+
        "transaction_validation_status_mssg ,"+
        "transaction_process_status ,"+
        "transaction_date , "+
        "transaction_user , "+
        "transaction_data_received , "+
        "transaction_validation_response,  "+
        "transaction_process_response  ,  "+
        "transaction_processed_qty  ) "+
        " VALUES( "+
        "'%s',   "+
        "'%s',   "+
        "'%s',   "+
        "'%s',   "+        
        "'%s',   "+           
        "current_timestamp, "+
        "'%s',   "+
        "'%s',   "+
        "'%s',   "+
        "'%s',   "+
        " %s  )", 
    u.Transtype,
    u.Filename,
    u.Trans_status ,
    u.Trans_statusmssg ,
    u.Trans_processstatus ,
//    u.TransCreated_at   ,
    u.Trans_user ,
    u.Trans_data_received ,
    u.Trans_val_response ,
    u.Trans_process_responser ,
    u.Trans_process_qty )

      
      	   log.Print("exec ejecutado:sql "+statement)
	    _, err := db.Exec(statement)
     	log.Print("exec ejecutado")
	    if err != nil {
     	   log.Print("exec ejecutado:error "+err.Error())
	        return err
	    }
     	   log.Print("exec ejecutado:todo ok ")
    return nil
}


func GetTodayFilesByType(db *sql.DB, eltoken string ) ([]Payment, error) {
     	log.Print("procesando GetTodayPaymentsByTokenCard")
        statement := fmt.Sprintf("SELECT token,created_at,amount FROM banwirefiletransaction WHERE token='%s' and created_at- interval '6h' >= (now()- interval '6h')::date + interval '1minutes' ",eltoken)
        
//        return db.QueryRow(statement).Scan(&u.Token, &u.Created_at,&u.Amount),nil

log.Print("procesando GetTodayPaymentsByTokenCard"+"SELECT token,created_at,amount FROM banwirefiletransaction WHERE token='%s' and created_at- interval '6h' >= (now()- interval '6h')::date + interval '1minutes' ",eltoken)

    rows, err := db.Query(statement)
    log.Print("GetTodayPaymentsByTokenCard 02.1!\n")
    if err != nil {
        return nil, err
    }
    log.Print("GetTodayPaymentsByTokenCard 02.5!\n")
    defer rows.Close()
    dailypayments := []Payment{}
    for rows.Next() {
    	 log.Print("GetTodayPaymentsByTokenCard 03!\n")
        var u Payment
        if err := rows.Scan( &u.Token, &u.Created_at,&u.Amount); err != nil {

        	log.Print("GetTodayPaymentsByTokenCard err!\n"+err.Error())
            return nil, err
        }
    	 log.Print("GetTodayPaymentsByTokenCard 03.5!\n")
        dailypayments = append(dailypayments, u)
    }
    log.Print("GetTodayPaymentsByTokenCard 04!\n")
    return dailypayments, nil

  //   }else{
  //       return errorGeneral
  //   }
}




