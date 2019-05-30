package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	//"log"
    utilito "banwire/services/file_tokenizer/util"

	"banwire/services/file_tokenizer/db"
	"banwire/services/file_tokenizer/path"
)

// Initializes the Config variables
var Config config

var HTTPListen string
var configFile string

var RunMode string

var Config_DB_pass string
var Config_DB_user string
var Config_DB_name string
var Config_DB_server string
var Config_DB_port int
var Config_WS_crbanwire_pass string
var Config_WS_crbanwire_url string

var Config_env_server string
var Config_env_url string
var Config_env_port string
var Config_env_log string


/*    const (
        DB_USER     = "lerepagr"        
        DB_PASSWORD = "Ag8q2utgSsVy2tyR7_M9cNYbzsqSvwma"
        DB_NAME     = "lerepagr"
        DB_SERVER     = "stampy.db.elephantsql.com" //"54.163.207.112"
        DB_PORT      = 5432

        CR_BANWIRE_USER      = "pruebasbw" //this was mentioned by charly Jan 2019
    )
*/

func init() {
	flag.StringVar(&RunMode, "mode", "api", "Service mode run (Options: api, batch)")
//to be set later	flag.StringVar(&HTTPListen, "http", ":8095", "Path where HTTP Server will listening")
	flag.StringVar(&configFile, "config", "./conf/config.json", "Path of configuration file")
//	flag.StringVar(&configFile, "config", "config.json", "Path of configuration file")

	configFile = path.RelativePath(configFile)
}


// loadConfiguration loads the configuration file
func LoadConfiguration() {
    utilito.LevelLog(Config_env_log, "3", "Loading configuration...v3.5")

	if d, e := ioutil.ReadFile(configFile); e == nil {
		e := json.Unmarshal(d, &Config)
		if e != nil {
//			log.Panicf("Error in unmarshalling the configuration file: %s", e.Error())
            utilito.LevelLog(Config_env_log, "3", "Configuration was not was loaded!")
		}else{
			utilito.LevelLog(Config_env_log, "3", "Configuration was YESS was loaded!")
		}

        
		utilito.LevelLog(Config_env_log, "3", "Configuration was loaded!(check previo")
	} else {
		//log.Panicf("Error in opening the configuration file: %s", e)
			utilito.LevelLog(Config_env_log, "3", "Error in opening the configuration file %s")
            utilito.LevelLog(Config_env_log, "3", e.Error())
	}
    //reset the port
    	flag.StringVar(&HTTPListen, "http", Config_env_port, "Path where HTTP Server will listening")

}

// config is the configuration structure object
type config struct {
//	Database configDatabase `json:"database"`
	Database configDatabase `json:"othersources"`
}

// configDatabase is the database structure object for configuration
type configDatabase struct{}

// UnmarshalJSON handles desearialization of configDatabase
// and loads the othersources: database connections and webservices connections
func (c *configDatabase) UnmarshalJSON(data []byte) error {
 utilito.LevelLog(Config_env_log, "3", "UnmarshalJSON 00!:othersources")
	var cc = []struct {
		Type string                   `json:"type"`
		Nodes []map[string]interface{} `json:"nodes"`
	}{}
			utilito.LevelLog(Config_env_log, "3", "UnmarshalJSON 01!")
	err := json.Unmarshal(data, &cc)
	if err != nil {
		return err
	}
			utilito.LevelLog(Config_env_log, "3", "UnmarshalJSON 02!")
	for _, d := range cc {
					utilito.LevelLog(Config_env_log, "3", "UnmarshalJSON 03.!"+d.Type)
		switch d.Type {
		case "crbanwire":
			utilito.LevelLog(Config_env_log, "3", "UnmarshalJSON 2.04.2!")
			for _, n := range d.Nodes {
							utilito.LevelLog(Config_env_log, "3", "UnmarshalJSON 2.05!")
				if active, _ := n["active"].(bool); active {
                    passcrban,_:=n["passwordcrbanwire"].(string)
                    urlcrban,_:=n["urlcrbanwire"].(string)
                    
                    utilito.LevelLog(Config_env_log, "3", "---- The value  was loaded"+passcrban)

                     Config_WS_crbanwire_pass = passcrban
                     Config_WS_crbanwire_url = urlcrban
                    utilito.LevelLog(Config_env_log, "3", "---- The crbanwire value  was assigned es "+Config_WS_crbanwire_pass)
				}
				utilito.LevelLog(Config_env_log, "3", "UnmarshalJSON 2.06!")
			}
		case "envserver":
			utilito.LevelLog(Config_env_log, "3", "UnmarshalJSON 2.04.2! envserver")
			for _, n := range d.Nodes {
							utilito.LevelLog(Config_env_log, "3", "UnmarshalJSON 2.05!")
				if active, _ := n["active"].(bool); active {
                    valenv,_:=n["envlevel"].(string)
                    valurl,_:=n["envurl"].(string)
                    valport,_:=n["envport"].(string)
                    vallog,_:=n["envlog"].(string)
                    utilito.LevelLog(Config_env_log, "3", "---- The value  was loaded"+valenv)

                     Config_env_server = valenv
                     Config_env_url = valurl
                     Config_env_port = valport
                     Config_env_log = vallog
                    utilito.LevelLog(Config_env_log, "3", "---- The env level value  was assigned es "+Config_env_server)
                    utilito.LevelLog(Config_env_log, "3", "---- The env level port  was assigned es "+Config_env_port)
				}
				utilito.LevelLog(Config_env_log, "3", "UnmarshalJSON 2.06! envserver")
			}

		case "postgresql":
			utilito.LevelLog(Config_env_log, "3", "UnmarshalJSON 04!")
			for _, n := range d.Nodes {
							utilito.LevelLog(Config_env_log, "3", "UnmarshalJSON 05!")
				if active, _ := n["active"].(bool); active {
					host, _ := n["host"].(string)
					port, _ := n["port"].(float64)
					_db, _ := n["db"].(string)
					user, _ := n["user"].(string)
					pass, _ := n["password"].(string)
                    
                    

                    Config_DB_pass =pass
                    Config_DB_user =user
                    Config_DB_name =_db
                    Config_DB_server =host
                    Config_DB_port =int(port)            
                    
                    utilito.LevelLog(Config_env_log, "3", "---- The DB values  was assigned "+Config_DB_server)                    
                    utilito.LevelLog(Config_env_log, "3", "---- The DB values  was assigned "+Config_DB_user)
                    utilito.LevelLog(Config_env_log, "3", "---- The DB values  was assigned "+Config_DB_pass)
                    utilito.LevelLog(Config_env_log, "3", "---- The DB values  was assigned "+Config_DB_name)
					if e := db.Connection.Set(db.NewPgDb(host, int(port), _db, user, pass)); e == nil {
						utilito.LevelLog(Config_env_log, "3", "---- The postgresql database was loaded"+host)
						utilito.LevelLog(Config_env_log, "3", "---- The postgresql database was loaded"+_db)
					} else {
						return e
					}
				}
							utilito.LevelLog(Config_env_log, "3", "UnmarshalJSON 06!")
			}

			break
		}
			utilito.LevelLog(Config_env_log, "3", "UnmarshalJSON 07!")
	}

	return nil
}
