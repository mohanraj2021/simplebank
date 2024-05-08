package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/simplebank/db/sqlc"
)

func (server *Server) addaccount(c *gin.Context) {

	type createAccountRequest struct {
		Owner   string `json:"name" binding:"required"`
		Curreny string `json:"currency" binding:"required,oneof=INR USD"`
	}
	var crAcc createAccountRequest

	if err := c.ShouldBindJSON(&crAcc); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	createAccountParam := db.CreateAccountParams{
		Owner:   crAcc.Owner,
		Curreny: crAcc.Curreny,
		Balance: 0,
	}

	acc, caerr := server.store.CreateAccount(context.Background(), createAccountParam)

	if caerr != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(caerr))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account created successfully", "account": acc})
}

func (server *Server) getAccount(c *gin.Context) {
	type getAccounrById struct {
		Id int64 `uri:"id" binding:"required,min=1"`
	}
	var getAcc getAccounrById
	if err := c.ShouldBindUri(&getAcc); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	acc, aerr := server.store.GetAccount(context.Background(), getAcc.Id)
	if aerr != nil {

		if aerr == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(aerr))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(aerr))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account retrieved successfully", "account": acc})

}

func (server *Server) getAccountList(c *gin.Context) {
	type accListLimits struct {
		Start int32 `form:"start" binding:"required,min=1"`
		End   int32 `form:"end" binding:"required,min=5,max=10"`
	}
	var aListParam accListLimits

	if err := c.ShouldBindQuery(&aListParam); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// args := db.ListAccountsParams{
	// 	Limit:  aListParam.Start,
	// 	Offset: aListParam.End,
	// }

	args := db.ListAccountsParams{
		Limit:  5,
		Offset: 0,
	}
	accList, alErr := server.store.ListAccounts(context.Background(), args)

	if alErr != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(alErr))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "accounts": accList})
}

func (server *Server) updateAccountBalance(c *gin.Context) {
	type updateAccountParam struct {
		Id      int64 `json:"id" binding:"required,gte=1"`
		Balance int64 `json:"amount" binding:"required,gte=1"`
	}

	var upAcc updateAccountParam

	if err := c.ShouldBindJSON(&upAcc); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := db.AddAccountBalanceParams{
		ID:     upAcc.Id,
		Amount: upAcc.Balance,
	}

	acc, upaccerr := server.store.AddAccountBalance(context.Background(), args)

	if upaccerr != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(upaccerr))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "accounts": acc})

}

func (server *Server) deleteAccount(c *gin.Context) {
	type getAccounrById struct {
		Id int64 `uri:"id" binding:"required,min=1"`
	}
	var delAcc getAccounrById
	if err := c.ShouldBindUri(&delAcc); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if _, delaccerr := server.store.DeleteAccount(context.Background(), delAcc.Id); delaccerr != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(delaccerr))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "acount deleted successfully"})
}
