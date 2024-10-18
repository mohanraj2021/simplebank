package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/simplebank/db/sqlc"
	"github.com/simplebank/types"
	"github.com/simplebank/utils"
)

type createUserResponse struct {
	Username  string    `json:"username"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (server *Server) addUser(c *gin.Context) {

	type CreateUserStruct struct {
		Username string `json:"username" binding:"required,alphanum"`
		Fullname string `json:"fullname" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	var createUser CreateUserStruct

	if err := c.ShouldBindJSON(&createUser); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashPwd, herr := utils.HashPassword(createUser.Password)

	if herr != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(herr))
		return
	}

	args := db.CreateUserParams{
		Username:       createUser.Username,
		Fullname:       createUser.Fullname,
		Email:          createUser.Email,
		Hashedpassword: hashPwd,
	}

	user, cuErr := server.store.CreateUser(context.Background(), args)
	if cuErr != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(cuErr))
		return
	}
	resp := createUserResponse{
		Username:  user.Username,
		Fullname:  user.Fullname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{"user": resp})
}

func (server *Server) getUser(c *gin.Context) {
	username := c.Param("username")
	user, err := server.store.GetUser(context.Background(), username)
	if err != nil {
		if rerr, ok := err.(*pq.Error); ok {
			switch rerr.Code.Name() {
			case "unique_violation":
				c.JSON(http.StatusBadRequest, errorResponse(err))
				return
			}
		}
	}
	resp := createUserResponse{
		Username:  user.Username,
		Fullname:  user.Fullname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	c.JSON(http.StatusOK, gin.H{"user": resp})
}

func (server *Server) userLogin(c *gin.Context) {
	type userLogin struct {
		Username string `json:"username" binding:"required,alphanum,gte=5"`
		Password string `json:"password" binding:"required,gte=6"`
	}
	var reqUser userLogin
	if err := c.BindJSON(&reqUser); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userObj, derr := server.store.GetUser(context.Background(), reqUser.Username)

	if derr != nil {
		if rerr, ok := derr.(*pq.Error); ok {
			// switch {

			// }
			fmt.Println(rerr.Code.Name())
		}
		c.JSON(http.StatusForbidden, gin.H{"message": "username or password is incorrect"})
		return
	}
	if perr := utils.CheckPassword(userObj.Hashedpassword, reqUser.Password); perr != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "username or password is incorrect"})
		return
	}

	token, terr := server.maker.CreateToken(userObj.Username, time.Duration(types.Token_Duration*int(time.Hour)))
	if terr != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(terr))
		return
	}

	type createUserResponse struct {
		Username  string
		Token     string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	resp := createUserResponse{
		Username: userObj.Username,
		Token:    token,
	}

	c.JSON(http.StatusOK, gin.H{"user": resp})
}
