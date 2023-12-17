package enc

const (
	CodeSuccess    = 0
	MessageSuccess = "success"
)

var (
	StatusSuccess = Status{Code: CodeSuccess, Message: MessageSuccess}
)

// Status 响应状态信息
type Status struct {
	// Code 状态码
	Code int `json:"code"`
	// Message 状态信息
	Message string `json:"message"`
}

// Enc 响应数据包装
type Enc struct {
	Status
	// Data 数据
	Data interface{} `json:"data"`
	// Extra 额外数据
	Extra map[string]interface{} `json:"extra"`
}

// NewSuccessEnc creates a success status Enc.
func NewSuccessEnc(data interface{}) *Enc {
	return &Enc{
		Status: StatusSuccess,
		Data:   data,
	}
}

// NewErrorEnc creates an error status Enc.
func NewErrorEnc(status Status) *Enc {
	return &Enc{
		Status: status,
		Data:   nil,
	}
}

// WithMessage puts message to Enc.
func (e *Enc) WithMessage(message string) *Enc {
	e.Message = message
	return e
}

// WithExtra puts extra data to Enc.
func (e *Enc) WithExtra(extra map[string]interface{}) *Enc {
	e.Extra = extra
	return e
}
