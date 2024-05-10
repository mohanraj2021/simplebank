package api

import (
	"context"
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
	tresult, terr := server.store.TransferTx(context.Background(), args)

	if terr != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(terr))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account created successfully", "account": tresult})

}
