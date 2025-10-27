package main

import (
	"fmt"
	"log"
	"time"
)

type Document struct {
	Key   string
	Title string
}

type Feed struct{}

func (f *Feed) Count() int {
	return 42
}

func (f *Feed) Fetch(key string) (Document, error) {
	time.Sleep(time.Second)

	doc := Document{
		Key:   key,
		Title: "Title for " + key,
	}
	return doc, nil
}

type FetchCounter interface {
	Fetch(key string) (Document, error)
	Count() int
}

func process(fc FetchCounter) {
	fmt.Printf("There are %d documents\n", fc.Count())

	keys := []string{"a", "a", "a", "b", "b", "b"}

	for _, key := range keys {
		doc, err := fc.Fetch(key)
		if err != nil {
			log.Printf("Could not fetch %s : %v", key, err)
			return
		}
		fmt.Printf("%s : %v\n", key, doc)
	}
}

type CachingFeed struct {
	*Feed
	cache map[string]Document
}

func NewCachingFeed(f *Feed) *CachingFeed {
	return &CachingFeed{
		Feed:  f,
		cache: make(map[string]Document),
	}
}

func (cf *CachingFeed) Fetch(key string) (Document, error) {
	if doc, exists := cf.cache[key]; exists {
		fmt.Printf("[CACHE MISS] key: %s\n", key)
		return doc, nil
	}

	fmt.Printf("[CACHE MISS] Key: %s\n", key)
	doc, err := cf.Feed.Fetch(key)
	if err != nil {
		return Document{}, err
	}

	cf.cache[key] = doc
	return doc, nil
}

func main() {
	process(&Feed{})
	cachingFeed := NewCachingFeed(&Feed{})
	process(cachingFeed)
	process(cachingFeed)
}
