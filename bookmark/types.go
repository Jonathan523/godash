package bookmark

type Category struct {
	Description string     `json:"description"`
	Bookmarks   []Bookmark `json:"bookmarks"`
}

type Bookmark struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	Url  string `json:"url"`
}
