package jsonplaceholder

// Album - album
type Album struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id`
	Title  string `json:"title"`
}

// Photo - photo object
type Photo struct {
	AlbumID      int    `json:"albumId`
	ID           int    `json:"id`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}
