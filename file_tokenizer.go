package main

import (
	"banwire/services/file_tokenizer/net"
	"flag"
	//"log"
    utilito "banwire/services/file_tokenizer/util"
	"sync"
)

var loaded sync.WaitGroup

const (
	DefaultRunMode = ""
	ApiRunMode     = "api"
	BatchRunMode   = "batch"
)

func init() {
	loaded.Add(1)
	
}

func main() {
	flag.Parse()
	LoadConfiguration()

	loaded.Done()
	utilito.LevelLog(Config_env_log, "3", "Service is ready for run...")
    
	switch RunMode {
	case BatchRunMode:
		BatchTest()
		break
	case DefaultRunMode, ApiRunMode:
		startServer()
		break
	default:
		utilito.LevelLog(Config_env_log, "3", "Run mode is unknown")
		break
	}

}



func startServer() {
	utilito.LevelLog(Config_env_log, "3", "HTTP Listening: "+ HTTPListen)
	net.HTTPListener(HTTPListen)
}
