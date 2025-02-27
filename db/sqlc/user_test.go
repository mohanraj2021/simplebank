package db

import (
	"context"
	"testing"
	"time"

	"github.com/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func createRandomUser(t *testing.T) User {
	cuserArgs := CreateUserParams{
		Username:       utils.RandomAccountName(6),
		Fullname:       utils.RandomAccountName(6),
		Hashedpassword: "secret",
		Email:          utils.RandomAccountEmail(6),
	}

	user, err := testQueries.CreateUser(context.Background(), cuserArgs)

	require.NoError(t, err)
	require.Equal(t, cuserArgs.Username, user.Username)
	require.Equal(t, cuserArgs.Fullname, user.Fullname)
	require.Equal(t, cuserArgs.Hashedpassword, user.Hashedpassword)
	require.Equal(t, cuserArgs.Email, user.Email)

	require.NotEmpty(t, user)
	require.Greater(t, time.Now(), user.CreatedAt)

	return user
}
