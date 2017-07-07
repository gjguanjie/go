package configs

//"io"
//"log"
//"os"

import (
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

/*
var (
Info  *log.Logger
Error *log.Logger
)
*/

func init() {
	/*
		file, err := os.OpenFile("goweb.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("Failed to open error log file", err)
		}
		Info = log.New(io.MultiWriter(file, os.Stdout), "INFO:", log.Ldate|log.Ltime|log.Llongfile)
		Error = log.New(io.MultiWriter(file, os.Stdout), "ERROR:", log.Ldate|log.Ltime|log.Llongfile)
	*/
	//var logfilename = beego.AppConfig.String("logfilename")
	logs.SetLogger(logs.AdapterFile, `{"file":"goweb.log","level":6}`)

}
