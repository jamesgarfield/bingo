package bingo

func (r ImageRequest) params() map[string]string {
	params := map[string]string{"Query": r.Query, "Adult": r.Adult}
	return params
}
func (r ImageRequest) AccountKey() string {
	return r.accountKey
}
