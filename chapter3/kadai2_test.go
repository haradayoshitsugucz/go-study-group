package chapter3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKadai2(t *testing.T) {
	k := NewKadai2(1, "hoge")

	assert.Equal(t, 1, k.id)
	assert.Equal(t, "hoge", k.name)
}
