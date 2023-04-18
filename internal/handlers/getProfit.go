package handlers

import (
	"DefaultEx2/internal/logger"
	"DefaultEx2/internal/structs"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetProfit(c *gin.Context, p *pgxpool.Pool) {

	var profit structs.Profit
	c.ShouldBindJSON(&profit)

	var reserve structs.Reserve

	aprove := "aproved"

	rows, err := p.Query(context.Background(), "select * from reserve where user_id=$1 and service_id=$2 and order_id=$3", profit.User_id, profit.Service_id, profit.Order_id)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	for rows.Next() {
		rows.Scan(&reserve.User_id, &reserve.Name, &reserve.Service_id, &reserve.Order_id, &reserve.Cost, &reserve.Сonfirmation)
	}

	if profit.Sum == reserve.Cost || reserve.Сonfirmation == "wating" {

		_, err = p.Exec(context.Background(), "insert into profit (user_id, name, service_id, order_id, sum) VALUES($1,$2,$3,$4,$5)",
			profit.User_id, profit.Name, profit.Service_id, profit.Order_id, profit.Sum)
		if err != nil {
			logger.ErrorLogger.Println(err)
		}

		_, err = p.Exec(context.Background(), "update reserve set сonfirmation=$1", aprove)
		if err != nil {
			logger.ErrorLogger.Println(err)
		}

	}

}
