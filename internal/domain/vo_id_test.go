package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarent/internal/domain"
)

func TestNewID(t *testing.T) {
	_, err := domain.NewID(1)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
