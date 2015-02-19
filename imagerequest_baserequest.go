package bingo

func (r ImageRequest) AccountKey() string {
	return r.accountKey
}
func (r ImageRequest) params() map[string]string {
	params := map[string]string{"Query": r.Query, "Adult": string(r.Adult)}
	return params
}
