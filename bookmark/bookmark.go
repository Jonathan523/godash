package bookmark

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	folderCreate "github.com/unjx-de/go-folder"
	"godash/message"
	"io"
	"os"
	"strings"
)

var Categories []Category

const StorageDir = "storage/"
const IconsDir = StorageDir + "icons/"
const bookmarksFile = "bookmarks.json"

func NewBookmarkService() {
	createFolderStructure()
	parseBookmarks()
	go watchBookmarks()
}

func createFolderStructure() {
	folders := []string{StorageDir, IconsDir}
	err := folderCreate.CreateFolders(folders, 0755)
	if err != nil {
		logrus.WithField("error", err).Fatal(message.CannotCreate.String())
	}
	logrus.WithField("folders", folders).Debug("folders created")
}

func copyDefaultBookmarks() {
	source, _ := os.Open("config/" + bookmarksFile)
	defer source.Close()
	destination, err := os.Create(StorageDir + bookmarksFile)
	if err != nil {
		logrus.WithField("file", bookmarksFile).Error(message.CannotCreate.String())
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	if err != nil {
		logrus.WithField("file", bookmarksFile).Error(message.CannotCreate.String())
	}
}

func readBookmarksFile() []byte {
	jsonFile, err := os.Open(StorageDir + bookmarksFile)
	if err != nil {
		copyDefaultBookmarks()
		jsonFile, err = os.Open(StorageDir + bookmarksFile)
		if err != nil {
			logrus.WithField("file", bookmarksFile).Error(message.CannotOpen.String())
			return nil
		}
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		logrus.WithField("file", bookmarksFile).Error(message.CannotParse.String())
		return nil
	}
	return byteValue
}

func replaceIconString() {
	for _, v := range Categories {
		for _, bookmark := range v.Bookmarks {
			if !strings.Contains(bookmark.Icon, "http") {
				bookmark.Icon = "/" + IconsDir + bookmark.Icon
			}
		}
	}
}

func parseBookmarks() {
	byteValue := readBookmarksFile()
	err := json.Unmarshal(byteValue, &Categories)
	if err != nil {
		logrus.WithField("file", bookmarksFile).Error(message.CannotParse.String())
		return
	}
	fmt.Println(Categories)
	replaceIconString()
}

func watchBookmarks() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logrus.WithField("watcher", err).Fatal(message.CannotCreate.String())
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
				logrus.WithField("error", err).Error(message.CannotParse.String())
			case _, ok := <-watcher.Events:
				if !ok {
					return
				}
				parseBookmarks()
				logrus.WithField("bookmarks", len(Categories)).Trace(bookmarksFile + " changed")
			}
		}
	}()

	if err := watcher.Add(StorageDir + bookmarksFile); err != nil {
		logrus.WithField("watcher", err).Fatal(message.CannotCreate.String())
	}
	<-done
}
