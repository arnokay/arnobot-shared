package mbtransformer

import (
	"github.com/arnokay/arnobot-shared/pkg/assert"
	"encoding/json"
)

func TransformRequest(input any) []byte {
  b, err := json.Marshal(input)
  assert.NoError(err, "#transform.Request: provided input cannot be marshaled")

  return b
}

func TransformResponse[T any](resp []byte) T {
  var responce T

  err := json.Unmarshal(resp, &responce)
  assert.NoError(err, "#transform.Response: provided response cannot be unmarshaled")

  return responce
}
