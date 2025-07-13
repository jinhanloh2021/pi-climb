package dto

type FeedCursor struct {
	FollowingCursor string `json:"follwing_cursor"`
	TrendingCursor  string `json:"trending_cursor`
}
