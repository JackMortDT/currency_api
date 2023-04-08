package command

type Rates struct {
	Meta Meta `json:"meta"`
	Data Data `json:"data"`
}

type Meta struct {
	LastUpdated string `json:"last_updated_at"`
}

type Data map[string]struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}
