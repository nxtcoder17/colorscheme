package colors

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseHex converts hex color to RGB
func ParseHex(hex string) (r, g, b int) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) == 6 {
		r64, _ := strconv.ParseInt(hex[0:2], 16, 64)
		g64, _ := strconv.ParseInt(hex[2:4], 16, 64)
		b64, _ := strconv.ParseInt(hex[4:6], 16, 64)
		return int(r64), int(g64), int(b64)
	}
	return 0, 0, 0
}

// ToHex converts RGB to hex string (without #)
func ToHex(r, g, b int) string {
	return fmt.Sprintf("%02x%02x%02x", clamp(r), clamp(g), clamp(b))
}

func clamp(v int) int {
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}
	return v
}

// Darken reduces brightness by percent (0-100)
func Darken(hex string, percent int) string {
	r, g, b := ParseHex(hex)
	factor := float64(100-percent) / 100
	return ToHex(int(float64(r)*factor), int(float64(g)*factor), int(float64(b)*factor))
}

// Lighten increases brightness by percent (0-100)
func Lighten(hex string, percent int) string {
	r, g, b := ParseHex(hex)
	factor := float64(percent) / 100
	return ToHex(r+int(float64(255-r)*factor), g+int(float64(255-g)*factor), b+int(float64(255-b)*factor))
}
