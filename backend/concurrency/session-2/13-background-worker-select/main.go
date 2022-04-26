package main

import "fmt"

func main() {

	imageQueue := make(chan string, 100)
	videoQueue := make(chan string, 100)
	quit := make(chan bool)
	resizer := func() {
		for {
			select {
			case image := <-imageQueue:
				imageResizer(image)
			case video := <-videoQueue:
				videoResizer(video)
			case <-quit:
				break
			}
		}
	}
	go resizer()

	imageQueue <- "image.jpg"
	videoQueue <- "video.mp4"
	quit <- true
}

func imageResizer(image string) {
	fmt.Printf("%s resized\n", image)
}

func videoResizer(video string) {
	fmt.Printf("%s resized\n", video)
}
