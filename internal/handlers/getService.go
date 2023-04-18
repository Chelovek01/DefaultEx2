package handlers

import (
	"DefaultEx2/internal/logger"
	"DefaultEx2/internal/structs"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetService(c *gin.Context, p *pgxpool.Pool) {

	var reserve structs.Reserve
	c.ShouldBindJSON(&reserve)

	var userBalance int32

	err := p.QueryRow(context.Background(), "SELECT balance FROM user_balance where user_id=$1", reserve.User_id).Scan(&userBalance)
	if err != nil {
		logger.ErrorLogger.Printf("Ошибка получения данных о балансе пользователя err=%s", err)
	}

	if reserve.Cost > int32(userBalance) && int32(userBalance) != 0 {
		data := fmt.Sprintf("Недостаточно сдредств для получения услуги у пользовтеля, %d", reserve.Order_id)
		logger.InfoLogger.Printf(data)
		c.String(http.StatusOK, data)
	}

	if (reserve.Cost < int32(userBalance) || reserve.Cost == int32(userBalance)) && int32(userBalance) != 0 {

		updateBalanse := int32(userBalance) - reserve.Cost
		p.QueryRow(context.Background(), "update user_balance set balance=$1 where user_id=$2", updateBalanse, reserve.User_id)

		_, err = p.Exec(context.Background(), "INSERT INTO reserve (user_id, name, service_id, order_id, cost, сonfirmation) VALUES($1,$2,$3,$4,$5,$6) returning user_id",
			reserve.User_id, reserve.Name, reserve.Service_id, reserve.Order_id, reserve.Cost, reserve.Сonfirmation)

		if err != nil {
			logger.ErrorLogger.Println(err)
		} else {
			logger.InfoLogger.Printf("Средства пользователя %d списаны с баланса в резерв, для дальнейшего получения услуги %d", reserve.User_id, reserve.Service_id)
			c.String(http.StatusOK, fmt.Sprintf("Средства зарезервированны пользователем, %d для услуги %d", reserve.Order_id, reserve.Service_id))
		}
	}

}
