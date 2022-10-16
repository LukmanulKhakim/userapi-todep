package config

import (
	"os"

	"sync"

	"github.com/labstack/gommon/log"
)

type AppConfig struct {
	Driver   string
	Name     string
	Address  string
	Port     int
	Username string
	Password string
}

func NewConfig() *AppConfig {
	cfg := initConfig()
	if cfg == nil {
		log.Fatal("Cannot run configuration setup")
		return nil
	}

	return cfg
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func initConfig() *AppConfig {
	var defaultConfig AppConfig

	// err := godotenv.Load("config.env")
	// if err != nil {
	// 	log.Error("config error :", err.Error())
	// 	return nil
	// }
	// SECRET = os.Getenv("SECRET")
	// cnv, err := strconv.Atoi(os.Getenv("SERVERPORT"))
	// if err != nil {
	// 	log.Fatal("Cannot parse port variable")
	// 	return nil
	// }
	// SERVERPORT = int16(cnv)
	defaultConfig.Name = os.Getenv("Name")
	defaultConfig.Username = os.Getenv("Username")
	defaultConfig.Password = os.Getenv("Password")
	defaultConfig.Address = os.Getenv("Address")

	// cnv, err := strconv.Atoi(os.Getenv("Port"))
	// if err != nil {
	// 	log.Fatal("Cannot parse DB Port variable")
	// 	return nil
	// }
	// defaultConfig.Port = cnv

	return &defaultConfig
	// app.DBUser = os.Getenv("DB_USER")
	// app.DBPwd = os.Getenv("DB_PWD")
	// app.DBHost = os.Getenv("DB_HOST")
	// port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	// if err != nil {
	// 	log.Error("config error :", err.Error())
	// 	return nil
	// }
	// app.DBPort = uint(port)
	// app.DBName = os.Getenv("DB_NAME")
	// app.JWTSecret = os.Getenv("JWT_SECRET")

	// return &app
}
