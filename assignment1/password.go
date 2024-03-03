package assignment1

import (
	"crypto/md5"
	"encoding/hex"
	"log"
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
							candidate := string([]byte{byte(char1), byte(char2), byte(char3), byte(char4), byte(char5), byte(char6)})
							cansidates <- candidate
						}
					}
				}
			}
		}
	}

	close(cansidates)
}

func combinations(d data) {
	defer d.wg.Done()
	for candidate := range d.Candidates {
		if newMD5(candidate) == d.TargetMD5 {
			d.Found <- candidate
			return
		}
	}
}

func Run(numGo int) {

	log.Println("Prgram started...")
	MD5 := "a74277500228f7b4cfa8694098443fc5"

	candidates := make(chan string, 100)
	found := make(chan string)

	var wg sync.WaitGroup

	data := data{
		TargetMD5:  MD5,
		Candidates: candidates,
		Found:      found,
		wg:         &wg,
	}

	log.Println("Generating candidates.....")
	go genCandidates(candidates)

	log.Println("checking for all combinations started.....")
	for i := 0; i < numGo; i++ {
		wg.Add(1)
		go combinations(data)
	}

	go func() {
		wg.Wait()
		close(found)
	}()

	password := <-found
	log.Print("password found.. ", password)
}
