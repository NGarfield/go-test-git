package main

import (
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func init() {
	runtime.GOMAXPROCS(1)
	initviper()
}
func main() {
	var (
		mysql = newMysqlConn()
		e     = echo.New()
	)
	defer mysql.Close()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":" + viper.GetString("app.port")))
}

func initviper() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot read in viper config:%s", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func newMysqlConn() *gorm.DB {
	conf := mysql.Config{
		DBName:               viper.GetString("mysql.database.test"),
		User:                 viper.GetString("mysql.username"),
		Passwd:               viper.GetString("mysql.password"),
		Net:                  "tcp",
		Addr:                 viper.GetString("mysql.host") + ":" + viper.GetString("mysql.port"),
		AllowNativePasswords: true,
		Timeout:              viper.GetDuration("mysql.timeout"),
		ReadTimeout:          viper.GetDuration("mysql.readtimeout"),
		WriteTimeout:         viper.GetDuration("mysql.writetimeout"),
		ParseTime:            viper.GetBool("mysql.parsetime"),
		Loc:                  time.Local,
	}
	conn, err := gorm.Open("mysql", conf.FormatDSN())
	if err != nil {
		log.Fatalf("cannot open mysql connection:%s", err)
	}

	conn.DB().SetMaxIdleConns(viper.GetInt("mysql.maxidle"))
	conn.DB().SetMaxOpenConns(viper.GetInt("mysql.maxopen"))
	conn.DB().SetConnMaxLifetime(viper.GetDuration("mysql.maxlifetime"))
	conn.LogMode(viper.GetBool("mysql.debug"))

	return conn
}
