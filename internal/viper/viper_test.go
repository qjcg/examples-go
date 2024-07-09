package viper

import (
	"testing"

	"github.com/spf13/viper"
)

func TestSetDefaultThenGet(t *testing.T) {
	conf := viper.New()
	conf.SetDefault("meaningOfLife", 42)

	want := 42
	got := conf.GetInt("meaningOfLife")
	if want != got {
		t.Fatalf("want %v got %v", want, got)
	}
}
