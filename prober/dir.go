package prober

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

type DirListen struct {
	dir string
}

func Listen() {
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watch.Close()
	//添加要监控的对象，文件或文件夹
	err = watch.Add()
	if err != nil {
		log.Fatal(err)
	}
}
