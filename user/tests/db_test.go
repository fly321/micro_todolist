package tests

import (
	"testing"
	"user/model"
)

func TestDbA(t *testing.T) {
	model.Migration()
}
