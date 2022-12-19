package bookmarks

import (
	"encoding/json"
	"github.com/fsnotify/fsnotify"
	folderCreate "github.com/unjx-de/go-folder"
	"go.uber.org/zap"
	"io"
	"os"
	"strings"
)

const StorageDir = "storage/"
const IconsDir = StorageDir + "icons/"
const bookmarksFolder = "bookmarks/"
const bookmarksFile = "bookmarks.json"

func NewBookmarkService(logging *zap.SugaredLogger) *Bookmarks {
	b := Bookmarks{log: logging}
	b.createFolderStructure()
	b.parseBookmarks()
	go b.watchBookmarks()
	return &b
}

func (b *Bookmarks) createFolderStructure() {
	folders := []string{StorageDir, IconsDir}
	err := folderCreate.CreateFolders(folders, 0755)
	if err != nil {
		b.log.Fatal(err)
	}
	b.log.Debug("folders created")
}

func (b *Bookmarks) copyDefaultBookmarks() {
	source, _ := os.Open(bookmarksFolder + bookmarksFile)
	defer source.Close()
	destination, err := os.Create(StorageDir + bookmarksFile)
	if err != nil {
		b.log.Error(err)
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	if err != nil {
		b.log.Error(err)
	}
}

func (b *Bookmarks) readBookmarksFile() []byte {
	jsonFile, err := os.Open(StorageDir + bookmarksFile)
	if err != nil {
		b.copyDefaultBookmarks()
		jsonFile, err = os.Open(StorageDir + bookmarksFile)
		if err != nil {
			b.log.Error(err)
			return nil
		}
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		b.log.Error(err)
		return nil
	}
	return byteValue
}

func (b *Bookmarks) replaceIconString() {
	for _, v := range b.Categories {
		for i, bookmark := range v.Entries {
			if !strings.Contains(bookmark.Icon, "http") {
				v.Entries[i].Icon = "/" + IconsDir + bookmark.Icon
			}
		}
	}
}

func (b *Bookmarks) parseBookmarks() {
	byteValue := b.readBookmarksFile()
	err := json.Unmarshal(byteValue, &b.Categories)
	if err != nil {
		b.log.Error(err)
		return
	}
	b.replaceIconString()
}

func (b *Bookmarks) watchBookmarks() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		b.log.Error(err)
	}
	defer watcher.Close()
	done := make(chan bool)

	go func() {
		for {
			select {
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				b.log.Error(err)
			case _, ok := <-watcher.Events:
				if !ok {
					return
				}
				b.parseBookmarks()
				b.log.Debug("bookmarks changed", "categories", len(b.Categories))
			}
		}
	}()

	if err := watcher.Add(StorageDir + bookmarksFile); err != nil {
		b.log.Fatal()
	}
	<-done
}
