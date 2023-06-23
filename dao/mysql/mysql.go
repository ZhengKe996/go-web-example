package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var db *sqlx.DB

func InitDB() (err error) {
	// DSN:Data Source Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", viper.GetString("mysql.user"), viper.GetString("mysql.password"), viper.GetString("mysql.host"), viper.GetString("mysql.port"), viper.GetString("mysql.dbName"))
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.MaxOpenConns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.MaxIdleConns"))
	return
}

func Close() {
	db.Close()
}
