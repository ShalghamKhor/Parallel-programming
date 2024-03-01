package assignment1

import (
	"crypto/md5"
	"encoding/hex"
	"sync"
)

type data struct {
	TargetMD5  string
	Candidates <-chan string
	Found      chan<- string
	wg         *sync.WaitGroup
}

func newMD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func genCandidates(cansidates chan<- string) {
	chars := "abcdefghijklmnopqrstuvwxyz0123456789"
	for _, char1 := range chars {
		for _, char2 := range chars {
			for _, char3 := range chars {
				for _, char4 := range chars {
					for _, char5 := range chars {
						for _, char6 := range chars {
							candidate := string([]byte{char1, char2, char3, char4, char5, char6})
							cansidates <- candidate
						}
					}
				}
			}
		}
	}
}
