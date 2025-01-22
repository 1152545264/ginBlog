package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	PrefixUrl      string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

var ServerSetting = &Server{}

type DataBase struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DataBaseSetting = &DataBase{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeOut time.Duration
}

var RedisSetting = &Redis{}

var cfg *ini.File

func SetUp() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini' : %v", err)
	}
	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DataBaseSetting)
	mapTo("redis", RedisSetting)

	AppSetting.ImageMaxSize *= 1 << 20
	ServerSetting.ReadTimeOut *= time.Second
	ServerSetting.WriteTimeOut *= time.Second
	RedisSetting.IdleTimeOut *= time.Second
}

func mapTo(section string, v any) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("cfg.Mapto %s err : 5v\n", section, err)
	}
}
