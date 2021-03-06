package binance

import (
	"context"
	"encoding/json"
)

// CreateWithdrawService create withdraw
type CreateWithdrawService struct {
	c       *Client
	asset   string
	address string
	amount  string
	name    *string
}

func (s *CreateWithdrawService) Asset(asset string) *CreateWithdrawService {
	s.asset = asset
	return s
}

func (s *CreateWithdrawService) Address(address string) *CreateWithdrawService {
	s.address = address
	return s
}

func (s *CreateWithdrawService) Amount(amount string) *CreateWithdrawService {
	s.amount = amount
	return s
}

func (s *CreateWithdrawService) Name(name string) *CreateWithdrawService {
	s.name = &name
	return s
}

func (s *CreateWithdrawService) Do(ctx context.Context) (err error) {
	r := &request{
		method:   "POST",
		endpoint: "/wapi/v1/withdraw.html",
		secType:  secTypeSigned,
	}
	m := params{
		"asset":   s.asset,
		"address": s.address,
		"amount":  s.amount,
	}
	if s.name != nil {
		m["name"] = *s.name
	}
	r.SetFormParams(m)
	_, err = s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	return
}

// ListWithdrawsService list withdraws
type ListWithdrawsService struct {
	c         *Client
	asset     *string
	status    *int
	startTime *int64
	endTime   *int64
}

func (s *ListWithdrawsService) Asset(asset string) *ListWithdrawsService {
	s.asset = &asset
	return s
}

func (s *ListWithdrawsService) Status(status int) *ListWithdrawsService {
	s.status = &status
	return s
}

func (s *ListWithdrawsService) StartTime(startTime int64) *ListWithdrawsService {
	s.startTime = &startTime
	return s
}

func (s *ListWithdrawsService) EndTime(endTime int64) *ListWithdrawsService {
	s.endTime = &endTime
	return s
}

func (s *ListWithdrawsService) Do(ctx context.Context) (withdraws []*Withdraw, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/wapi/v1/getWithdrawHistory.html",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.SetParam("asset", *s.asset)
	}
	if s.status != nil {
		r.SetParam("status", *s.status)
	}
	if s.startTime != nil {
		r.SetParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.SetParam("endTime", *s.endTime)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res := new(WithdrawHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}
	return res.Withdraws, nil
}

// WithdrawHistoryResponse define withdraw history response
type WithdrawHistoryResponse struct {
	Withdraws []*Withdraw `json:"withdrawList"`
	Success   bool        `json:"success"`
}

// Withdraw define withdraw info
type Withdraw struct {
	Amount    float64 `json:"amount"`
	Address   string  `json:"address"`
	Asset     string  `json:"asset"`
	TxID      string  `json:"txId"`
	ApplyTime int64   `json:"applyTime"`
	Status    int     `json:"status"`
}
