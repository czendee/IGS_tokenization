CREATE USER banwire;

CREATE DATABASE dashboard_banwire OWNER banwire;





---mysql



CREATE TABLE banwiredash02graph01(
id_record BIGINT AUTO_INCREMENT,
numsecuecial_grupoddatos    bigint,  
numsecuecial    bigint,     
nombrecolumna    varchar(100),   
created_at         timestamp,
valorcolumna    varchar(100),       
valoramount      bigint,
graphnbr    varchar(10),
    PRIMARY KEY (id_record)
);


CREATE TABLE banwiredash02graph02(
id_record BIGINT AUTO_INCREMENT,
numsecuecial_grupoddatos    bigint, 
numsecuecial    bigint,      
nombrecolumna    varchar(100),   
created_at         timestamp,
valorcolumna    varchar(100),       
valoramount      bigint,
graphnbr    varchar(10),
    PRIMARY KEY (id_record)
);


-----postgres

CREATE TABLE banwiredash02graph01(
id_record bigserial,
numsecuecial_grupoddatos    bigint,  
numsecuecial    bigint,     
nombrecolumna    varchar(100),   
created_at         timestamp,
valorcolumna    varchar(100),       
valoramount      bigint,
graphnbr    varchar(10),
);


CREATE TABLE banwiredash02graph02(
id_record  bigserial,
numsecuecial_grupoddatos    bigint, 
numsecuecial    bigint,      
nombrecolumna    varchar(100),   
created_at         timestamp,
valorcolumna    varchar(100),       
valoramount      bigint,
graphnbr    varchar(10),
);

