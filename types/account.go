package types

type Account struct {
	ID         int     `json:"id"`
	Url        string  `json:"url"`
	Bio        string  `json:"bio"`
	Reputation float32 `json:"reputation"`
	Created    int     `json:"created"`
}

type RawAccount struct {
	Account `json:"data"`
}
