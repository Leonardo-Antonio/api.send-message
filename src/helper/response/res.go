package response

type response struct {
	MessageType string `json:"message_type,omitempty" xml:"message_type,omitempty"`
	Message     string `json:"message,omitempty" xml:"message,omitempty"`
}

const MSJ = "message"
const ERR = "error"

func New(messageType, message string) *response {
	return &response{messageType, message}
}
