package schema

type ErrorShowType int

const (
	SILENT        ErrorShowType = 0
	WARN_MESSAGE  ErrorShowType = 1
	ERROR_MESSAGE ErrorShowType = 2
	NOTIFICATION  ErrorShowType = 3
	REDIRECT      ErrorShowType = 9
)

type Response struct {
	Success      bool          `json:"success"`
	Data         any           `json:"data,omitempty"`
	ErrorMessage string        `json:"errorMessage,omitempty"`
	ErrorCode    int           `json:"errorCode,omitempty"`
	ShowType     ErrorShowType `json:"showType,omitempty"`
}
