package main

import (
	"io"
	"bufio"
	"bytes"
	"path"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/minio/minio-go/v6"
	// "golang.org/x/text/unicode/norm"
)

type Index struct {
	Bucket string
	Prefix string
	Store minio.Client
}

func SaveIndex(store *minio.Client, bucket, prefix, id string, scanner *bufio.Scanner) error {
	object := bytes.NewReader([]byte{})
	size := int64(object.Len())
	opts := minio.PutObjectOptions{
		ContentType: "application/octet-stream",
		// UserTags: map[string]string{"id": id},
	}

	for scanner.Scan() {
		token := scanner.Text()
		spew.Dump(token)

		n, err := store.PutObject(bucket, path.Join(prefix, token, id), object, size, opts)
		spew.Dump(n)
		if err != nil {
		    spew.Dump(err)
		}
	}
	
	return scanner.Err()
}

func SearchIndex(store *minio.Client, bucket, prefix string, query io.Reader) error {
	scanner := TokenizeText(query, [5]bool{true, false, false, false, false})
	isRecursive := true
	done := make(chan struct{})
	defer close(done)
	
	for scanner.Scan() {
		token := scanner.Text()
		// spew.Dump(token)

		rs := []rune(token)
		len := len(rs)
		min := 2
		if len < 4 {
			min = 0
		}

		for i := len; i > min; i-- {
			ngram := string(rs[:i])
			weight := float64(i) / float64(len)
			spew.Dump(ngram)
			ch := store.ListObjectsV2(bucket, path.Join(prefix, ngram), isRecursive, done)

			for object := range ch {
			    if object.Err != nil {
			        fmt.Println(object.Err)
			        return object.Err
			    }
			    
			    spew.Dump(object.Key, weight)
			    fmt.Println()
			}
		}

		// for i, j := 0, norm.NFC.FirstBoundary(token); j > -1; i, j = j, norm.NFC.FirstBoundary(token[j:]) {
		// 	spew.Dump(string(token[i:j]))
		// }
		// spew.Dump(norm.NFC.QuickSpan([]byte("hello world")))
	}
	
	return scanner.Err()
}
