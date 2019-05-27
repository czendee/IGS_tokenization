package pgsql

import (
//    "database/sql"
//    "errors"
	"fmt"
	"log"
//	miu "banwire/dash/dashboard_front/util"

//	"banwire/dash/dashboard_front/db"
	 "banwire/dash/dashboard_front/model"
	
)

type Card struct {
    ID    string    `sql:"type:bigserial"`
    Token  string `sql:"type:varchar(30)`
    Last   string    `sql:"type:varchar(4)`
    Bin   string    `sql:"type:varchar(6)`
    Valid   string    `sql:"type:timestamp`
    Score   string    `sql:"type:integer`
    Customer   string    `sql:"type:bigint`
    Brand   string    `sql:"type:varchar(30)`
    Type   string    `sql:"type:varchar(30)`
}

func (_db *Db) GetCardByTokenAndCust(tokencito,customer_id string) (error, model.Card) {

    var miCard model.Card
    
    log.Print("   Step 1 GetCardByTokenAndCust!    \n")
    statement := fmt.Sprintf("SELECT token,last_digits,bin, brand, type_card  reference FROM banwirecard WHERE token='%s' and id_customer='%s' ", tokencito,customer_id)
     log.Print("       Step 2 GetCardByTokenAndCust:\n"+statement)
    return _db.QueryRow(statement).Scan(&miCard.Token,&miCard.Last,&miCard.Bin,&miCard.Brand,&miCard.Type), miCard
}

