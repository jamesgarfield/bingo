package bingogen

type Request struct {
	Query      string
	Adult      string
	accountKey string
}

func (r Request) params() map[string]string {
	params := map[string]string{
		"Query": r.Query,
		"Adult": r.Adult,
	}
	return params
}

func (r Request) AccountKey() string {
	return r.accountKey
}
