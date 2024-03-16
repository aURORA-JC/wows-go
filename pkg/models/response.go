package models

type WowsSuccessResponse struct {
	Status string   `json:"status"`
	Meta   WowsMeta `json:"meta"`
}

type WowsMeta struct {
	Count     int `json:"count,omitempty"`
	PageTotal int `json:"page_total,omitempty"`
	Total     int `json:"total,omitempty"`
	Limit     int `json:"limit,omitempty"`
	Page      int `json:"page,omitempty"`
}

type WowsErrorResponse struct {
	Status string `json:"status"`
	Error  struct {
		Field   string `json:"field"`
		Message string `json:"message"`
		Code    int    `json:"code"`
		Value   any    `json:"value"`
	} `json:"error"`
}
