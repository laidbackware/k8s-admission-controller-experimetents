package exampleadmissioncontroller

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeserialize(t *testing.T) {
	var body []byte
	writer := httptest.NewRecorder()
	_, err := unmarshalService(body, writer)

	assert.Nil(t, err)
}