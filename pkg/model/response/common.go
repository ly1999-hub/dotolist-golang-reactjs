package responsemodel

type ResponseCreate struct {
	ID string `json:"id"`
}

type ResponseUpdate struct {
	ID string `json:"id"`
}
type ResponseDelete struct {
	ID string `json:"id"`
}

type ResponseList struct {
	List interface{} `json:"list"`
}
