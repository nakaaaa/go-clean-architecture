package model

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestUser_TableName(t *testing.T) {
	var rp User
	got := rp.TableName()
	assert.Equal(t, got, "users")
}
