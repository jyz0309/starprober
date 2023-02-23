package prober

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

type DirListener struct {
	dir string
}

func (l *DirListener) Listen() {
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watch.Close()
	err = watch.Add(l.dir)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case ev := <-watch.Events:
				{
					if ev.Op&fsnotify.Create == fsnotify.Create {
						log.Println("Create File:", ev.Name)
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						log.Println("Write File:", ev.Name)
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						log.Println("Delete File:", ev.Name)
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						log.Println("Rename File:", ev.Name)
					}
				}
			case err := <-watch.Errors:
				{
					log.Println("error: ", err)
					return
				}
			}
		}
	}()
}
