package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/getaccount/%d", account.ID)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	require.NoError(t, err)

	server.router.ServeHTTP(recorder, req)

	var acc map[string]db.Account

	data, recBodyErr := ioutil.ReadAll(recorder.Body)

	require.NoError(t, recBodyErr)

	err = json.Unmarshal(data, &acc)
	require.NoError(t, err)

	require.Equal(t, acc["account"], account)

}

func TestGetAccountLists(t *testing.T) {
	acc1 := randomAccount()
	acc2 := randomAccount()
	acc3 := randomAccount()
	acc4 := randomAccount()
	acc5 := randomAccount()
	accList := []db.Account{acc1, acc2, acc3, acc4, acc5}

	controller := gomock.NewController(t)

	defer controller.Finish()

	store := mockdb.NewMockStore(controller)

	args := db.ListAccountsParams{
		Limit:  5,
		Offset: 0,
	}

	store.EXPECT().ListAccounts(gomock.Any(), args).Times(1).Return(accList, nil)

	server := NewServer(store)

	recorder := httptest.NewRecorder()

	url := "/getaccounts?start=1&end=5"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, req)

	var accs map[string][]db.Account

	bAccs, errs := ioutil.ReadAll(recorder.Body)
	require.NoError(t, errs)

	unmarshalErr := json.Unmarshal(bAccs, &accs)

	require.NoError(t, unmarshalErr)

	require.Equal(t, accList, accs["accounts"])

}

func TestAddAccount(t *testing.T) {

	account := randomAccount()
	createAccParams := db.CreateAccountParams{
		Owner:   account.Owner,
		Curreny: account.Curreny,
	}

	controller := gomock.NewController(t)
	store := mockdb.NewMockStore(controller)

	store.EXPECT().CreateAccount(gomock.Any(), createAccParams).Times(1).Return(account, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()

	data := fmt.Sprintf(`{"name":"%s","currency":"%s"}`, account.Owner, account.Curreny)

	req, err := http.NewRequest(http.MethodPost, "/addaccount", bytes.NewBuffer([]byte(data)))

	require.NoError(t, err)

	server.router.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusOK, recorder.Code)

	var retAcc map[string]db.Account

	unMarErr := json.Unmarshal(recorder.Body.Bytes(), &retAcc)

	require.NoError(t, unMarErr)

	require.Equal(t, retAcc["account"], account)

}

func TestUpdateAccountBalance(t *testing.T) {

	account := randomAccount()
	args := db.AddAccountBalanceParams{
		ID:     account.ID,
		Amount: 200,
	}

	controller := gomock.NewController(t)
	store := mockdb.NewMockStore(controller)

	store.EXPECT().AddAccountBalance(gomock.Any(), args).Times(1).Return(account, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()

	data := fmt.Sprintf(`{"id":%d,"amount":%d}`, account.ID, 200)

	req, err := http.NewRequest(http.MethodPost, "/updateaccount", bytes.NewBuffer([]byte(data)))

	require.NoError(t, err)

	server.router.ServeHTTP(recorder, req)

	var retAcc map[string]db.Account

	unMarErr := json.Unmarshal(recorder.Body.Bytes(), &retAcc)

	require.NoError(t, unMarErr)

	require.Equal(t, retAcc["account"], account)

}

func TestDeleteAcount(t *testing.T) {

	account := randomAccount()

	controller := gomock.NewController(t)
	store := mockdb.NewMockStore(controller)

	store.EXPECT().DeleteAccount(gomock.Any(), account.ID).Times(1)

	server := NewServer(store)

	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/deleteaccount/%d", account.ID)

	req, err := http.NewRequest(http.MethodDelete, url, nil)

	require.NoError(t, err)

	server.router.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusOK, recorder.Code)

}

func randomAccount() db.Account {

	return db.Account{
		ID:      utils.RandomINT(1, 1000),
		Owner:   utils.RandomAccountName(10),
		Balance: utils.RandomINT(100, 1000),
		Curreny: utils.RandomCurrency(),
	}
}
