package modelRequest

type Assigner struct {
	Investment int32 `json:"investment" binding:"required"`
}