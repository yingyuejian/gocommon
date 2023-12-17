package enc

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSuccessEnc(t *testing.T) {
	data := "My Data"
	extra := map[string]interface{}{"a": 1}
	enc := NewSuccessEnc(data).WithExtra(extra)

	t.Logf("%+v\n", enc)
	bytes, _ := json.Marshal(enc)
	t.Logf("JSON: %v\n", string(bytes))

	assert.Equal(t, CodeSuccess, enc.Code)
	assert.Equal(t, MessageSuccess, enc.Message)
	assert.Equal(t, data, enc.Data)
	assert.Equal(t, extra, enc.Extra)
}

func TestNewErrorEnc(t *testing.T) {
	errStatus := Status{Code: -1000, Message: "error 1000"}
	enc := NewErrorEnc(errStatus)

	t.Logf("%+v\n", enc)
	bytes, _ := json.Marshal(enc)
	t.Logf("JSON: %v\n", string(bytes))

	assert.Equal(t, errStatus.Code, enc.Code)
	assert.Equal(t, errStatus.Message, enc.Message)
	assert.Equal(t, nil, enc.Data)
}

func TestNewErrorEnc2(t *testing.T) {
	enc := NewErrorEnc(StatusContentTypeNotSupported)

	t.Logf("%+v\n", enc)
	bytes, _ := json.Marshal(enc)
	t.Logf("JSON: %v\n", string(bytes))

	assert.Equal(t, StatusContentTypeNotSupported.Code, enc.Code)
	assert.Equal(t, StatusContentTypeNotSupported.Message, enc.Message)
	assert.Equal(t, nil, enc.Data)
}

func TestEnc_WithMessage(t *testing.T) {
	enc := NewSuccessEnc("My Data").WithMessage("My Message")
	t.Logf("%+v\n", enc)
	assert.Equal(t, "My Message", enc.Message)
	assert.Equal(t, "My Data", enc.Data)
}
