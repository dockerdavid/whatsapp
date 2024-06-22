package queue

import (
	"compi-whatsapp/pkg/meow"
	"sync"
	"time"
)

type FileQueue struct {
	To      string `json:"to" validate:"required"`
	URL     string `json:"url" validate:"required"`
	Caption string `json:"caption" validate:"required"`
}

var (
	fileQueue                 []*FileQueue
	queueMutex                sync.Mutex
	maxAmountOfFilesPerMinute = 10
)

func AddFileToQueue(file *FileQueue) {
	queueMutex.Lock()
	defer queueMutex.Unlock()
	fileQueue = append(fileQueue, file)
}

func SendPendingFiles() {
	queueMutex.Lock()
	defer queueMutex.Unlock()

	if len(fileQueue) == 0 {
		return
	}

	var filesToSend []*FileQueue

	if len(fileQueue) < maxAmountOfFilesPerMinute {
		filesToSend = fileQueue
		fileQueue = nil
	} else {
		filesToSend = fileQueue[:maxAmountOfFilesPerMinute]
		fileQueue = fileQueue[maxAmountOfFilesPerMinute:]
	}

	for _, file := range filesToSend {
		go func(file *FileQueue) {
			meow.SendFile(&meow.File{
				To:      file.To,
				URL:     file.URL,
				Caption: file.Caption,
			})
		}(file)
		time.Sleep(1 * time.Second)
	}
}
