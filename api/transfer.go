package api

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/simplebank/db/sqlc"
)

func (server *Server) transferTransaction(c *gin.Context) {

	var args db.TransferTxParams

	if err := c.ShouldBindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if t := server.checkAccCurrency(args.FromAccountID, args.ToAccountID, c); !t {
		return
	}
	tresult, terr := server.store.TransferTx(context.Background(), args)

	if terr != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(terr))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account created successfully", "account": tresult})

}

func (server *Server) checkAccCurrency(fromId, toId int64, c *gin.Context) bool {

	acc1, aerr1 := server.store.GetAccount(context.Background(), fromId)
	if aerr1 != nil {

		if aerr1 == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(aerr1))
			return false
		}

		c.JSON(http.StatusInternalServerError, errorResponse(aerr1))
		return false
	}

	acc2, aerr2 := server.store.GetAccount(context.Background(), toId)
	if aerr2 != nil {

		if aerr2 == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(aerr2))
			return false
		}

		c.JSON(http.StatusInternalServerError, errorResponse(aerr2))
		return false
	}

	if acc1.Curreny != acc2.Curreny {
		err := fmt.Errorf("the currency of acc id %d and  acc id %d is doesn't match", fromId, toId)
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}

	return true
}
