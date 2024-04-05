package ctxkey_test

import (
	"context"
	"github.com/simon-engledew/ctxkey"
	"testing"
)

var testKey = ctxkey.New("default")
var otherTestKey = ctxkey.New("other")

func TestKey(t *testing.T) {
	ctx := context.Background()
	if testKey.Value(ctx) != "default" {
		t.Fail()
	}
	ctx = testKey.With(ctx, "override")
	if testKey.Value(ctx) != "override" {
		t.Fail()
	}
	if otherTestKey.Value(ctx) != "other" {
		t.Fail()
	}
	ctx = otherTestKey.With(ctx, "override")
	if testKey.Value(ctx) != "override" {
		t.Fail()
	}
	if otherTestKey.Value(ctx) != "override" {
		t.Fail()
	}
}
