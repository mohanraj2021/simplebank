package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/simplebank/db/mock"
	db "github.com/simplebank/db/sqlc"
	"github.com/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func TestGetAccount(t *testing.T) {
	account := randomAccount()

	controller := gomock.NewController(t)
	defer controller.Finish()
	store := mockdb.NewMockStore(controller)

	store.EXPECT().GetAccount(gomock.Any(), account.ID).Times(1).Return(account, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/getaccount/%d", account.ID)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	require.NoError(t, err)

	server.router.ServeHTTP(recorder, req)

}

func randomAccount() db.Account {

	return db.Account{
		ID:      utils.RandomINT(1, 1000),
		Owner:   utils.RandomAccountName(10),
		Balance: utils.RandomINT(100, 1000),
		Curreny: utils.RandomCurrency(),
	}
}
