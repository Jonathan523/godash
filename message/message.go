package message

type Response struct {
	Message string `json:"message" validate:"required"`
}

type (
	Responses uint
)

const (
	NotFound Responses = iota
	CannotOpen
	CannotParse
	CannotCreate
	CannotStart
)

func (e Responses) String() string {
	switch e {
	case NotFound:
		return "cannot be found"
	case CannotOpen:
		return "cannot be opened"
	case CannotParse:
		return "cannot be parsed"
	case CannotCreate:
		return "cannot be created"
	case CannotStart:
		return "cannot be started"
	default:
		return ""
	}
}
