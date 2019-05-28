CREATE USER banwire;

CREATE DATABASE gs_ivr_tokenization OWNER banwire;


CREATE TABLE banwirecard(
id_card bigserial,
token    varchar(100),   --constrain unique
last_digits   varchar(4),        
bin         varchar(6),
valid_thru    timestamp, 
score     integer,
is_banned          boolean,
banned_at         timestamp, 
created_at         timestamp,
last_update_at timestamp ,
id_customer      bigint,
brand   varchar(20) ,
type_card varchar(15)
);


CREATE TABLE banwirecustomer(
id_customer bigserial,
reference    varchar(100),   --constrain unique
created_at         timestamp,
last_update_at timestamp 
);

CREATE TABLE banwirepayment(
id_payment bigserial,
token    varchar(100),   --constrain unique
created_at         timestamp,
amount      bigint,
);

CREATE TABLE banwirefiletransaction(
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

);

--banwirefiletransaction(

CREATE TABLE banwirecr_audit(
id_audit bigserial,
transaction_type  varchar (20),  --type is tokenizer, process payment
transaction_process_status    varchar(100),   --OK, ERROR
transaction_process_status_mssg    varchar(100),   --OK, ERROR
transaction_date             timestamp, 
transaction_source         varchar(19),     --file or ivr or service
transaction_user           varchar (50),
transaction_data_receiced   varchar (2500),   --data received with the information for the token, or to do a payment [do not store the TDC number/Valid Throug]
transaction_data_sent_cr   varchar (2500),    --data sent to CR_banwire to process a token or a payment  [do not store the TDC number/Valid Throug]
transaction_data_received_cr  varchar(2500),   --data sent back by CR as response
transaction_data_response  varchar(2500)     --data sent back as response
);


--select * from  banwirecr_audit;