package bingo

type baseRequest struct {
	Query   string
	Adult   Adult
	AcctKey string
}

func (r baseRequest) AccountKey() string {
	return r.AcctKey
}

func (r baseRequest) params() map[string]string {
	params := map[string]string{
		"Query": r.Query,
		"Adult": string(r.Adult),
	}
	return params
}
