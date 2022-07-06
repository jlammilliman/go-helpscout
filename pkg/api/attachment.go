package api

// Attatchements are a subsidiary to Conversations, and must have an associated conversationId
type Attachment struct {
	Id       int
	Hash     string
	MimeType string
	Filename string
	Size     int
	Width    int
	Height   int
	Url      string
}
