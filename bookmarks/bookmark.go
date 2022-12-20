package bookmarks

import (
	"github.com/fsnotify/fsnotify"
	folderCreate "github.com/unjx-de/go-folder"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"strings"
)

const StorageDir = "storage/"
const IconsDir = StorageDir + "icons/"
const bookmarksFolder = "bookmarks/"
const bookmarksFile = "config.yaml"

func NewBookmarkService(logging *zap.SugaredLogger) *Config {
	b := Config{log: logging}
	b.createFolderStructure()
	b.parseBookmarks()
	go b.watchBookmarks()
	return &b
}

func (c *Config) createFolderStructure() {
	folders := []string{StorageDir, IconsDir}
	err := folderCreate.CreateFolders(folders, 0755)
	if err != nil {
		c.log.Fatal(err)
	}
	c.log.Debugw("folders created", "folders", folders)
}

func (c *Config) copyDefaultBookmarks() {
	source, _ := os.Open(bookmarksFolder + bookmarksFile)
	defer source.Close()
	destination, err := os.Create(StorageDir + bookmarksFile)
	if err != nil {
		c.log.Error(err)
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	if err != nil {
		c.log.Error(err)
	}
}

func (c *Config) readBookmarksFile() []byte {
	file, err := os.Open(StorageDir + bookmarksFile)
	if err != nil {
		c.copyDefaultBookmarks()
		file, err = os.Open(StorageDir + bookmarksFile)
		if err != nil {
			c.log.Error(err)
			return nil
		}
	}
	defer file.Close()
	byteValue, err := io.ReadAll(file)
	if err != nil {
		c.log.Error(err)
		return nil
	}
	return byteValue
}

func (c *Config) replaceIconString() {
	for _, v := range c.Parsed.Applications {
		for i, bookmark := range v.Entries {
			if !strings.Contains(bookmark.Icon, "http") {
				v.Entries[i].Icon = "/" + IconsDir + bookmark.Icon
			}
		}
	}
}

func (c *Config) parseBookmarks() {
	byteValue := c.readBookmarksFile()
	err := yaml.Unmarshal(byteValue, &c.Parsed)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.replaceIconString()
}

func (c *Config) watchBookmarks() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		c.log.Error(err)
	}
	defer watcher.Close()
	done := make(chan bool)

	go func() {
		for {
			select {
			case err, _ := <-watcher.Errors:
				c.log.Error(err)
			case _, _ = <-watcher.Events:
				c.parseBookmarks()
				c.log.Debug("bookmarks changed", "applications", len(c.Parsed.Applications), "links", len(c.Parsed.Links))
			}
		}
	}()

	if err := watcher.Add(StorageDir + bookmarksFile); err != nil {
		c.log.Fatal()
	}
	<-done
}
