package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	db "github.com/simplebank/db/sqlc"
	"github.com/simplebank/token"
	"github.com/simplebank/types"
)

type Server struct {
	store  db.Store
	maker  token.Maker
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{
		store: store,
	}

	tokenMaker, terr := token.NewJWTMaker(types.SecreteKey)
	if terr != nil {
		log.Fatal("Unable to create JWT maker")
		os.Exit(1)
	}

	server.maker = tokenMaker

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Message": "Server is running"})
	})

	router.POST("/addaccount", server.addaccount)

	secured := router.Group("/secured").Use(authMiddleware(tokenMaker))
	{
		secured.GET("/getaccounts", server.getAccountList)
	}

	router.GET("/getaccounts", server.getAccountList)
	router.GET("/getaccount/:id", server.getAccount)
	router.POST("/updateaccount", server.updateAccountBalance)
	router.DELETE("/deleteaccount/:id", server.deleteAccount)
	router.POST("/transfer", server.transferTransaction)
	router.POST("/users", server.addUser)
	router.GET("/users/:username", server.getUser)
	router.POST("/users/login", server.userLogin)

	server.router = router

	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(addr string) {
	server.router.Run(addr)
}
