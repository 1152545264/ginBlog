package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg          *ini.File
	RunMode      string
	HTTPPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
	PageSize     int
	JwtSecrete   string
)

func init() {
	var err error
	confPath := "conf/app.ini"
	Cfg, err = ini.Load(confPath)
	if err != nil {
		log.Fatalf("Fail to parse %s, err:%v", confPath, err)
	}
	LoadBase()
	loadServer()
	LoadApp()
}
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server' , err:%v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeOut = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeOut = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app', err:%v", err)
	}
	JwtSecrete = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
