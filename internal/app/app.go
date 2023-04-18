package app

import (
	"DefaultEx2/internal/handlers"
	"DefaultEx2/internal/logger"
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

func RunApp() {

	logger.Init()

	logger.InfoLogger.Println("Starting the application...")

	viper.SetConfigFile("../../internal/cfg/.env")
	viper.ReadInConfig()

	urlExample := "postgres://Chelovek:passpass@localhost:5432/postgres"

	PgClient, err := pgxpool.New(context.Background(), urlExample)
	if err != nil {
		logger.ErrorLogger.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	r := gin.Default()

	r.GET("/up_balance", func(c *gin.Context) {
		handlers.UpBalance(c, PgClient)
	})

	// c.ShouldBindJSON(&person)

	// query := fmt.Sprintf("INSERT INTO user_balance (user_id, name, adress, balance) VALUES ('%d', '%s', '%s', '%d')",
	// 	person.User_id, person.Name, person.Address, person.Balance)

	// PgClient.QueryRow(context.Background(), query).Scan()
	// fmt.Println(person.User_id)

	r.GET("/get_service", func(c *gin.Context) {
		handlers.GetService(c, PgClient)
	})

	r.GET("/get_profit", func(c *gin.Context) {

		handlers.GetProfit(c, PgClient)
	})

	r.Run()

}
