package central

type HealthResponse struct {
	Code        int     `json:"code"`
	Status      string  `json:"status"`
	Description *string `json:"description"`
	Message     *string `json:"message"`
}

type AddNumberRequest struct {
	First  int `json:"first"`
	Second int `json:"second"`
}

type AddNumberResponse struct {
	Result int `json:"result"`
}
