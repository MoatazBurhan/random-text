package main

import (
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	Count     = 2
	Separator = "-"
)

type Chatterer struct {
	Words      []string
	randSource *rand.Rand
}

func NewChatterer() (c Chatterer) {
	c.Words = readDictionary()
	source := rand.NewSource(time.Now().UnixNano())
	c.randSource = rand.New(source)
	return
}

func (c Chatterer) Chatter() string {
	pieces := []string{}
	for i := 0; i < Count; i++ {
		randomInt := c.randomInt()
		pieces = append(pieces, c.Words[randomInt%len(c.Words)])
	}

	return strings.Join(pieces, Separator)
}

func (c Chatterer) randomInt() int {
	return c.randSource.Int()
}

func readDictionary() (words []string) {
	file, err := os.Open("./dict/words")
	if err != nil {
		panic(err)
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	words = strings.Split(string(bytes), "\n")
	return
}
