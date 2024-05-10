package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	db "github.com/simplebank/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Message": "Server is running"})
	})

	router.POST("/addaccount", server.addaccount)
	router.GET("/getaccounts", server.getAccountList)
	router.GET("/getaccount/:id", server.getAccount)
	router.POST("/updateaccount", server.updateAccountBalance)
	router.DELETE("/deleteaccount/:id", server.deleteAccount)
	router.POST("/transfer", server.transferTransaction)
	server.router = router

	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(addr string) {
	server.router.Run(addr)
}
