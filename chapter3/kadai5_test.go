package chapter3

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKadai5(t *testing.T) {
	m := Master{
		id:   1,
		name: "hoge",
	}

	b, err := json.Marshal(m)
	assert.NoError(t, err)
	assert.Equal(t, `{"id":1,"name":"hoge"}`, string(b))
}
