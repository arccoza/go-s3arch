package main

import (
	"github.com/minio/minio-go/v6"
	"github.com/kljensen/snowball"
	"github.com/davecgh/go-spew/spew"
	"log"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	endpoint := "localhost:9000"
    accessKeyID := "minioadmin"
    secretAccessKey := "minioadmin"
    useSSL := false

    minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
    if err != nil {
        log.Fatalln(err)
    }

    log.Printf("%#v\n", minioClient)

    bucketName := "emails"
    location := ""

    err = minioClient.MakeBucket(bucketName, location)
    if err != nil {
        // Check to see if we already own this bucket (which happens if you run this twice)
        exists, errBucketExists := minioClient.BucketExists(bucketName)
        if errBucketExists == nil && exists {
            log.Printf("We already own %s\n", bucketName)
        } else {
            log.Fatalln(err)
        }
    } else {
        log.Printf("Successfully created %s\n", bucketName)
    }

    msgs := []string{
`The greatest glory in living lies not in never falling, but in rising every time we fall.`,
`The way to get started is to quit talking and begin doing.`,
`Be yourself; everyone else is already taken.`,
`Two things are infinite: the universe and human stupidity; and I'm not sure about the universe.`,
`So many books, so little time.`,
`Be who you are and say what you feel, because those who mind don't matter, and those who matter don't mind.`,
`If you tell the truth, you don't have to remember anything.`,
	}

	index(minioClient, msgs)

	// for i, text := range msgs {
	// 	id := strconv.Itoa(i)
	// 	file := strings.NewReader(text)

	// 	_, err := minioClient.PutObject("emails", "msgs/" + id, file, int64(file.Len()),
	// 		minio.PutObjectOptions{
	// 			ContentType: "application/octet-stream",
	// 			UserTags: map[string]string{"id": id},
	// 		})

	// 	if err != nil {
	// 	    fmt.Println(err)
	// 	}

	// 	file = strings.NewReader("")
	// 	for _, word := range stem(text) {
	// 		_, err := minioClient.PutObject("emails", "index/" + word, file, int64(file.Len()),
	// 			minio.PutObjectOptions{
	// 				ContentType: "application/octet-stream",
	// 				UserTags: map[string]string{"id": id},
	// 			})
	// 		if err != nil {
	// 		    fmt.Println(err)
	// 		}
	// 	}
	// }

	find(minioClient, "lies not")
}

func index(minioClient *minio.Client, msgs []string) {
	for i, text := range msgs {
		id := strconv.Itoa(i)
		file := strings.NewReader(text)

		_, err := minioClient.PutObject("emails", "msgs/" + id, file, int64(file.Len()),
			minio.PutObjectOptions{
				ContentType: "application/octet-stream",
				UserMetadata: map[string]string{"id": id},
			})

		if err != nil {
		    fmt.Println(err)
		}

		file = strings.NewReader(id)
		for _, word := range stem(text) {
			_, err := minioClient.PutObject("emails", "index/" + word + "/" + id, file, int64(file.Len()),
				minio.PutObjectOptions{
					ContentType: "application/octet-stream",
					UserTags: map[string]string{"id": id},
				})
			if err != nil {
			    fmt.Println(err)
			}
		}
	}
}

func find(minioClient *minio.Client, query string) {
	// Create a done channel to control 'ListObjectsV2' go routine.
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	isRecursive := true

	// var ret []string
	for _, w := range stem(query) {
		fmt.Println(w)
		objectCh := minioClient.ListObjectsV2("emails", "index/" + w, isRecursive, doneCh)
		for object := range objectCh {
		    if object.Err != nil {
		        fmt.Println(object.Err)
		        return
		    }
		    fmt.Println()
		    spew.Dump(object)
		}
	}
}

func stem(text string) []string {
	var ret []string

	for _, w := range strings.Split(text, " ") {
		stemmed, err := snowball.Stem(w, "english", true)
		if err == nil{
			// fmt.Println(stemmed)
			ret = append(ret, stemmed)
		}
	}

	return ret
}