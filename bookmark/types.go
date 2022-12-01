package bookmark

type Entry struct {
	Category  string     `json:"category"`
	Bookmarks []Bookmark `json:"bookmarks"`
}

type Bookmark struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	Url  string `json:"url"`
}
