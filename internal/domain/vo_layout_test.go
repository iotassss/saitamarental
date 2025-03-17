package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
)

func TestNewLayout(t *testing.T) {
	_, err := domain.NewLayout("2LDK/3K")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
