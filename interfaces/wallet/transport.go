package wallet

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
	"github.com/gorilla/mux"

	"coins-wallet/domain/entity/account"
)

func MakeHandler(service Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	createAccountHandler := kithttp.NewServer(
		makeCreateAccountEndpoint(service),
		decodeCreateAccountRequest,
		encodeResponse,
		opts...,
	)

	getAllAccountHandler := kithttp.NewServer(
		makeGetAllAccountEndpoint(service),
		decodeGetAllAccountRequest,
		encodeResponse,
		opts...,
	)

	getAllPaymentHandler := kithttp.NewServer(
		makeGetAllPaymentEndpoint(service),
		decodeGetAllPaymentRequest,
		encodeResponse,
		opts...,
	)

	transferHandler := kithttp.NewServer(
		makeTransferEndpoint(service),
		decodeTransferRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle("/v1/accounts", createAccountHandler).Methods("POST")
	r.Handle("/v1/accounts", getAllAccountHandler).Methods("GET")
	r.Handle("/v1/payments", getAllPaymentHandler).Methods("GET")
	r.Handle("/v1/transfer", transferHandler).Methods("POST")
	return r
}

func decodeCreateAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		ID       string  `json:"id"`
		Balance  float32 `json:"balance"`
		Currency string  `json:"currency"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return account.Account{
		ID:       account.AccountID(body.ID),
		Balance:  body.Balance,
		Currency: body.Currency,
	}, nil
}

func decodeGetAllAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return getAllAccountRequest{}, nil
}

func decodeGetAllPaymentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return getAllPaymentRequest{}, nil
}

func decodeTransferRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		SenderId   string  `json:"senderId"`
		ReceiverId string  `json:"receiverId"`
		Amount     float32 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return transferRequest{
		SenderId:   account.AccountID(body.SenderId),
		ReceiverId: account.AccountID(body.ReceiverId),
		Amount:     body.Amount,
	}, nil
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
