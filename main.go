package main

import (
	"chatbot-backend/bootstrap"
	"chatbot-backend/global"
)

func main() {
	// init
	bootstrap.InitializeConfig()

	// init log
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")

	// init database
	global.App.DB = bootstrap.InitializeDB()
	// release connecting before close project
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

    // init validator
    bootstrap.InitializeValidator()

	// init redis
	global.App.Redis = bootstrap.InitializeRedis()

    // start server
	bootstrap.RunServer()
}
