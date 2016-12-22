package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetMessageMock() GetRedisFunc {
	return func(key string) string {
		return "this is mock message"
	}
}

func TestGetRedisMessage(t *testing.T) {
	c = Config{
		GetRedis: GetMessageMock(),
	}

	assert.Equal(t, "Here The Message this is mock message", GetMessage(), "TestGetRedisMessage")
}
