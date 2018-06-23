package model

type Counter struct {
	BaseModel
	Key      string `json:"key"`
	TargetID uint   `json:"target_id"`
	Count    uint   `json:"count"`
}
