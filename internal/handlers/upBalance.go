package handlers

import (
	"DefaultEx2/internal/logger"
	"DefaultEx2/internal/structs"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func UpBalance(c *gin.Context, p *pgxpool.Pool) {

	var person structs.Person

	c.ShouldBindJSON(&person)

	_, err := p.Exec(context.Background(),
		"insert into user_balance (user_id, name, adress, balance) values($1,$2,$3,$4) on conflict (user_id) do update set balance = user_balance.balance + $5",
		person.User_id, person.Name, person.Adress, person.Balance, person.Balance)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
}
