package balcapi

type (
	PayRequest struct {
		Amt         int    `json:"amt"`         // Төлбөрийн дүн
		Description string `json:"description"` // Төлбөрийн тайлбар
	}
	LimitResponse struct {
		TotalLimit float64 `json:"totalLimit"` // Нийт лимит
		UsedLimit  float64 `json:"usedLimit"`  // Ашигласан дүн
		AvailLimit float64 `json:"availLimit"` // БОломжит дүн
		Status     int     `json:"status"`     // 1 ашиглах боломжтой, 0 боломжгүй
	}
)
