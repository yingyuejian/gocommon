package enc

var (
	// StatusUnknownError 未知错误
	StatusUnknownError = Status{Code: -1, Message: "unknown error"}

	// StatusMissingRequestParameter 丢失必要的请求参数
	StatusMissingRequestParameter = Status{Code: -2, Message: "missing request parameter"}
	// StatusMissingRequestHeader 丢失必要的请求头
	StatusMissingRequestHeader = Status{Code: -3, Message: "missing request header"}

	// StatusDatetimeFormatError 日期时间格式错误
	StatusDatetimeFormatError = Status{Code: -4, Message: "datetime format error"}
	// StatusOperationNotAllowed 操作不允许
	StatusOperationNotAllowed = Status{Code: -5, Message: "operation not allowed"}
	// StatusIllegalArgument 参数非法
	StatusIllegalArgument = Status{Code: -6, Message: "illegal argument"}
	// StatusArgumentTypeMismatch 参数类型不匹配
	StatusArgumentTypeMismatch = Status{Code: -7, Message: "argument type mismatch"}
	// StatusArgumentValidationFailed 参数验证失败
	StatusArgumentValidationFailed = Status{Code: -8, Message: "argument validation failed"}

	// StatusRequestMethodNotSupported 请求方法不支持
	StatusRequestMethodNotSupported = Status{Code: -9, Message: "request method not supported"}
	// StatusHTTPMediaTypeNotSupported HTTP 媒体类型不支持
	StatusHTTPMediaTypeNotSupported = Status{Code: -10, Message: "http media type not supported"}
	// StatusContentTypeNotSupported Content-Type 不支持
	StatusContentTypeNotSupported = Status{Code: -11, Message: "content type not supported"}
	// StatusImageFormatNotSupported 图片格式不支持
	StatusImageFormatNotSupported = Status{Code: -12, Message: "image format not supported"}

	// StatusHTTPMessageNotReadable HTTP 消息不可读
	StatusHTTPMessageNotReadable = Status{Code: -13, Message: "http message not readable"}
)

// ToEnc converts Status to Enc.
func (s Status) ToEnc() *Enc {
	return NewErrorEnc(s)
}
