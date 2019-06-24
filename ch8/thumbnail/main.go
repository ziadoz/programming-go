// photos downloaded from https://picsum.photos/
// https://play.golang.org/p/BipOSeRAr0k
// https://play.golang.org/p/S3E6ftYaGBS
// https://play.golang.org/p/1jlyBQbH7vq
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"gopl.io/ch8/thumbnail/thumbnail"
)

// deleteThumbnails cleans up any existing thumbnails
func deleteThumbnails() {
	thumbs, err := filepath.Glob("./*.thumb.jpg")
	if err != nil {
		log.Fatalf("no thumbnail images in directory: %s", err)
	}

	for _, thumb := range thumbs {
		os.Remove(thumb)
	}
}

// makeThumbnails makes thumbnails of the specified files.
func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Fatal(err)
		}
	}
}

// makeThumbnails2 makes thumbnails of the specified files in parallel, but finishes instantly and does no work.
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f) // NOTE: Ignoring errors.
	}
}

// makeThumbnails3 makes thumbnails of the specified files in parallel and waits for goroutines to finish.
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f) // NOTE: Ignoring errors.
			ch <- struct{}{}
		}(f) // Must pass arg here else goroutines won't have correct param.
	}
	// Wait for goroutines to complete.
	for range filenames {
		<-ch
	}
}

// makeThumbnails3 makes thumbnails of the specified files in parallel.
// Has a subtle bug where the first non nill error is returned, the errors channel is no longer being drained.
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err // NOTE: incorrect: goroutine leak.
		}
	}

	return nil
}

// makeThumbnails5 makes thumbnails for the specified files in parallel.
// It returns the generated file names in an arbitrary order,
// or an error if any step failed.
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))

	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

func makeThumbnails6(filenames []string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines

	for _, f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}

func main() {
	deleteThumbnails()

	dir, _ := os.Getwd()
	images, err := filepath.Glob(filepath.Join(dir, "/*.jpg"))
	if err != nil {
		log.Fatalf("no images in directory: %s", err)
	}

	result := makeThumbnails6(images)
	fmt.Println(result)
}
