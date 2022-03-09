package wallet

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-kit/log"
	"github.com/stretchr/testify/assert"

	"coins-wallet/infrastructure/inmem"
)

func TestCreateAccountHandler_Success(t *testing.T) {
	accountRepo := inmem.NewAccountRepository()
	service := NewService(accountRepo, inmem.NewPaymentRepository())
	h := MakeHandler(service, log.NewNopLogger())

	req, err := http.NewRequest(http.MethodPost, "/v1/accounts", bytes.NewBuffer([]byte(`{"id":"fake999", "balance":10, "currency":"AUD"}`)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	resp := rr.Result()
	defer resp.Body.Close()

	var response createAccountResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "", response.Err)

	accounts, err := accountRepo.GetAll()
	assert.Nil(t, err)
	assert.Len(t, accounts, 4)
}

func TestGetAllAccountsHandler_Success(t *testing.T) {
	service := NewService(inmem.NewAccountRepository(), inmem.NewPaymentRepository())
	h := MakeHandler(service, log.NewNopLogger())

	req, err := http.NewRequest(http.MethodGet, "/v1/accounts", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	resp := rr.Result()
	defer resp.Body.Close()

	var response getAllAccountResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}

	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, response.Res, 3)
}

func TestGetAllPaymentsHandler_Success(t *testing.T) {
	service := NewService(inmem.NewAccountRepository(), inmem.NewPaymentRepository())
	h := MakeHandler(service, log.NewNopLogger())

	req, err := http.NewRequest(http.MethodGet, "/v1/payments", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	resp := rr.Result()
	defer resp.Body.Close()

	var response getAllPaymentResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}

	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, response.Res, 0)
}

func TestTransferHandler_Success(t *testing.T) {
	paymentRepo := inmem.NewPaymentRepository()
	service := NewService(inmem.NewAccountRepository(), paymentRepo)
	h := MakeHandler(service, log.NewNopLogger())

	req, err := http.NewRequest(http.MethodPost, "/v1/transfer", bytes.NewBuffer([]byte(`{"senderId":"bob123", "receiverId":"alice456", "amount":1}`)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	resp := rr.Result()
	defer resp.Body.Close()

	var response createAccountResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "", response.Err)

	payments, err := paymentRepo.GetAll()
	assert.Nil(t, err)
	assert.Len(t, payments, 2)
}
