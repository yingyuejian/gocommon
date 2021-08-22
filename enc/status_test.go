package enc

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatus_ToEnc(t *testing.T) {
	enc := StatusUnknownError.ToEnc()

	t.Logf("%+v\n", enc)
	bytes, _ := json.Marshal(enc)
	t.Logf("JSON: %v\n", string(bytes))

	assert.Equal(t, StatusUnknownError.Code, enc.Code)
	assert.Equal(t, StatusUnknownError.Message, enc.Message)
	assert.Equal(t, nil, enc.Data)
}

func TestPrintAllPresetStatus(t *testing.T) {
	t.Log(StatusUnknownError)
	t.Log(StatusMissingRequestParameter)
	t.Log(StatusMissingRequestHeader)
	t.Log(StatusDatetimeFormatError)
	t.Log(StatusOperationNotAllowed)
	t.Log(StatusIllegalArgument)
	t.Log(StatusArgumentTypeMismatch)
	t.Log(StatusArgumentValidationFailed)
	t.Log(StatusRequestMethodNotSupported)
	t.Log(StatusHTTPMediaTypeNotSupported)
	t.Log(StatusContentTypeNotSupported)
	t.Log(StatusImageFormatNotSupported)
	t.Log(StatusHTTPMessageNotReadable)
}
