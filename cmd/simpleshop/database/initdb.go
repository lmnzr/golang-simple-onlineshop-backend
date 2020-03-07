package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/lmnzr/simpleshop/cmd/simpleshop/helper/config"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/helper/env"
	logutil "github.com/lmnzr/simpleshop/cmd/simpleshop/helper/log"
)

func setup() map[string]int {
	configmap := make(map[string]int)

	config, conferr := config.GetConfig()

	if conferr != nil {
		configmap["maxconn"] = 25
		configmap["maxiddle"] = 25
		configmap["maxlifetime"] = 5
	} else {
		configmap["maxconn"] = config.GetInt("dbMaxConn")
		configmap["maxiddle"] = config.GetInt("dbMaxIddle")
		configmap["maxlifetime"] = config.GetInt("dbMaxLifeTime")
	}
	return configmap

}

//OpenDbConnection :
func OpenDbConnection() (*sql.DB, error) {
	configmap := setup()

	dbMaxConns := configmap["maxconn"]
	dbMaxIdleConns := configmap["maxiddle"]
	dbMaxLifeTime := configmap["maxlifetime"]

	logutil.LoggerDB().Info("open db connection")

	dbUser := env.Getenv("DB_USER", "lmnzr")
	dbPass := env.Getenv("DB_PASS", "root")
	dbHost := env.Getenv("DB_HOST", "localhost")
	dbPort := env.Getenv("DB_PORT", "3306")
	dbSchema := env.Getenv("DB_SCHEMA", "simpleshop")
	dbConnString := fmt.Sprintf("%[1]s:%[2]s@tcp(%[3]s:%[4]s)/%[5]s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbSchema)

	db, err := sql.Open("mysql", dbConnString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(dbMaxConns)
	db.SetMaxIdleConns(dbMaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(dbMaxLifeTime) * time.Minute)

	return db, nil
}
