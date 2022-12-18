package bookmarks

import "go.uber.org/zap"

type Bookmarks struct {
	log        *zap.SugaredLogger
	Categories []Category
}

type Category struct {
	Description string  `json:"description"`
	Entries     []Entry `json:"entries"`
}

type Entry struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	Url  string `json:"url"`
}
