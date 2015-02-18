package bingo

type baseRequest struct {
	Query      string
	Adult      Adult
	accountKey string
}

func (r baseRequest) params() map[string]string {
	params := map[string]string{
		"Query": r.Query,
		"Adult": r.Adult,
	}
	return params
}

func (r baseRequest) AccountKey() string {
	return r.accountKey
}
