package token

import (
	"fmt"
	"testing"
	"time"

	"github.com/simplebank/types"
	"github.com/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func TestJWTToken(t *testing.T) {
	username := utils.RandomAccountName(5)

	tokenmaker, tmerr := NewJWTMaker(types.SecreteKey)
	require.NoError(t, tmerr)
	require.NotEmpty(t, tokenmaker)

	token, terr := tokenmaker.CreateToken(username, time.Duration(1*time.Hour))
	fmt.Println(token)

	require.NoError(t, terr)
	require.NotEmpty(t, token)

	payload, verr := tokenmaker.VerifyToken(token)
	require.NoError(t, verr)
	require.NotEmpty(t, payload)
}
