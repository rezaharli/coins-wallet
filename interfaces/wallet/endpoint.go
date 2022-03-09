package wallet

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"coins-wallet/domain/entity/account"
	"coins-wallet/domain/entity/payment"
)

type createAccountResponse struct {
	Err string `json:"err,omitempty"`
}

func makeCreateAccountEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(account.Account)
		err := s.CreateAccount(req)
		res := createAccountResponse{}
		if err != nil {
			res.Err = err.Error()
		}
		return res, nil
	}
}

type getAllAccountRequest struct{}

type getAllAccountResponse struct {
	Res []account.Account `json:"res,omitempty"`
	Err string            `json:"err,omitempty"`
}

func makeGetAllAccountEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		accounts, err := s.GetAllAccount()
		res := getAllAccountResponse{Res: accounts}
		if err != nil {
			res.Err = err.Error()
		}
		return res, nil
	}
}

type getAllPaymentRequest struct{}

type getAllPaymentResponse struct {
	Res []payment.Payment `json:"res,omitempty"`
	Err string            `json:"err,omitempty"`
}

func makeGetAllPaymentEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		payments, err := s.GetAllPayment()
		res := getAllPaymentResponse{Res: payments}
		if err != nil {
			res.Err = err.Error()
		}
		return res, nil
	}
}

type transferRequest struct {
	SenderId   account.AccountID `json:"senderId"`
	ReceiverId account.AccountID `json:"receiverId"`
	Amount     float32           `json:"amount"`
}

type transferResponse struct {
	Res []payment.Payment `json:"res,omitempty"`
	Err string            `json:"err,omitempty"`
}

func makeTransferEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transferRequest)
		payments, err := s.Transfer(req.SenderId, req.ReceiverId, req.Amount)
		res := transferResponse{Res: payments}
		if err != nil {
			res.Err = err.Error()
		}
		return res, nil
	}
}
