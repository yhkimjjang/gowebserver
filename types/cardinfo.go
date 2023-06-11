package types

type CardInfo struct {
	Code    string        `json:"code"`
	Name    string        `json:"name"`
	Benefit []interface{} `json:"benfit"`
}
