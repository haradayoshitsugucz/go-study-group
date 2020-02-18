package chapter3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKadai4(t *testing.T) {
	k := Face{}
	k.Watch()

	assert.True(t, k.Eye.isOpen)

	k.Eat()
	assert.True(t, k.Mouth.hasFood)

	k.Breathe()
	assert.True(t, k.Nose.isOpen)
	assert.True(t, k.Mouth.isOpen)
}
