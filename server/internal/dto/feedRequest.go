package dto

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jinhanloh2021/pi-climb/internal/models"
)

type FeedCursor struct {
	FollowingCursor string `json:"following_cursor"`
	TrendingCursor  string `json:"trending_cursor"`
}

const cursorDelimiter = "_"

func FormatPostCursor(post *models.Post) string {
	if post == nil {
		return ""
	}
	return fmt.Sprintf("%d%s%d", post.CreatedAt.UnixNano(), cursorDelimiter, post.ID)
}

func ParsePostCursor(cursor string) (int64, int64) {
	if cursor == "" {
		return -1, -1
	}

	parts := strings.Split(cursor, cursorDelimiter)
	if len(parts) != 2 {
		return -1, -1
	}

	timestampNano, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return -1, -1
	}

	postID, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return -1, -1
	}

	return timestampNano, postID
}
