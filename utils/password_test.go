package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	pass := RandomString(6)

	hashpwd, herr := HashPassword(pass)

	require.NoError(t, herr)
	require.NotEmpty(t, hashpwd)

	randPwd := RandomString(6)

	cerr := CheckPassword(hashpwd, randPwd)

	require.EqualError(t, cerr, bcrypt.ErrMismatchedHashAndPassword.Error())

	pcerr := CheckPassword(hashpwd, pass)

	require.NoError(t, pcerr)
}
