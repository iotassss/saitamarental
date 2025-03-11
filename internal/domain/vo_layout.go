package domain

import "fmt"

const (
	LayoutOneRoom = "ワンルーム"
	Layout1K1DK   = "1K/1DK"
	Layout1LDK2K  = "1LDK/2K"
	Layout2LDK3K  = "2LDK/3K"
	Layout3LDK4K  = "3LDK/4K~"
)

var layouts = []string{
	LayoutOneRoom,
	Layout1K1DK,
	Layout1LDK2K,
	Layout2LDK3K,
	Layout3LDK4K,
}

type Layout string

func NewLayout(layout string) (Layout, error) {
	if layout == "" {
		return "", fmt.Errorf("layout cannot be empty")
	}

	found := false
	for _, l := range layouts {
		if l == layout {
			found = true
			break
		}
	}
	if !found {
		return "", fmt.Errorf("layout '%s' is not recognized", layout)
	}

	return Layout(layout), nil
}

func (l Layout) Value() string {
	return string(l)
}
