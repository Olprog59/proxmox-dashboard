package commons

import (
	"fmt"
)

func TransformBytesForHuman(b int64, nb int) string {
	if b < 1024 {
		return fmt.Sprintf("%d B", b)
	}

	size := float64(b)
	suffixes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}

	var i int
	for i = 0; size >= 1024 && i < len(suffixes)-1; i++ {
		size /= 1024
	}

	var nbStr string
	switch nb {
	case 1:
		nbStr = fmt.Sprintf("%.1f", size)
	case 2:
		nbStr = fmt.Sprintf("%.2f", size)
	case 3:
		nbStr = fmt.Sprintf("%.3f", size)
	case 4:
		nbStr = fmt.Sprintf("%.4f", size)
	default:
		nbStr = fmt.Sprintf("%.0f", size)
	}

	return fmt.Sprintf("%s %s", nbStr, suffixes[i])
}
