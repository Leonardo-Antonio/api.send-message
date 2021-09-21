package entity

type Message struct {
	Name  string `json:"name,omitempty" xml:"name,omitempty"`
	Text  string `json:"text,omitempty" xml:"text,omitempty"`
	Email string `json:"email,omitempty" xml:"email,omitempty"`
}
