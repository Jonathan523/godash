package bookmark

import (
	"encoding/json"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	folderCreate "github.com/unjx-de/go-folder"
	"io"
	"launchpad/message"
	"os"
)

type Service struct {
	Bookmarks []Bookmark
}

var Bookmarks []Bookmark

const StorageDir = "storage/"
const folder = "bookmark/"
const bookmarksFile = "bookmarks.json"

func init() {
	createFolderStructure()
	parseBookmarks()
	go watchBookmarks()
}

func createFolderStructure() {
	folders := []string{StorageDir}
	err := folderCreate.CreateFolders(folders, 0755)
	if err != nil {
		logrus.WithField("error", err).Fatal(message.CannotCreate.String())
	}
	logrus.WithField("folders", folders).Debug("folders created")
}

func copyBookmarks() {
	source, _ := os.Open(folder + bookmarksFile)
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
		copyBookmarks()
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

func parseBookmarks() {
	byteValue := readBookmarksFile()
	err := json.Unmarshal(byteValue, &Bookmarks)
	if err != nil {
		logrus.WithField("file", bookmarksFile).Error(message.CannotParse.String())
		return
	}
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
				logrus.WithField("bookmarks", len(Bookmarks)).Trace(bookmarksFile + " changed")
			}
		}
	}()

	if err := watcher.Add(StorageDir + bookmarksFile); err != nil {
		logrus.WithField("watcher", err).Fatal(message.CannotCreate.String())
	}
	<-done
}
