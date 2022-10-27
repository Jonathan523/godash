package message

type Response struct {
	Message string `json:"message" validate:"required"`
}

type (
	Responses uint
)
