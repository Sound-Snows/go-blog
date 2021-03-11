package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	//Cfg is File
	Cfg *ini.File

	//RunMode is String
	RunMode string

	//HTTPPort is HttpPort
	HTTPPort int

	//ReadTimeout is read Sql time
	ReadTimeout time.Duration

	//WriteTime is wrte Sql time
	WriteTime time.Duration

	//PageSize is page number
	PageSize int

	//JwtSecret is JWT password
	JwtSecret string

	//Cors domain name
	DomainName string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini':%v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
	LoadDomainName()
}

//LoadBase is function
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}
//LoadDomainName is function
func LoadDomainName(){
	DomainName=Cfg.Section("Cors").Key("DOMAIN_NAME").MustString("*")
}

//LoadServer is function
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTime = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

//LoadApp is function
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)@U#@*!@!")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
