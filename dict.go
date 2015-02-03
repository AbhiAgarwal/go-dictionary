package dict

import (
	"fmt"
	"os/exec"
	"strings"
)

type Dictionary struct {
	Words    []string
	wordsmap map[string]bool
}

func (d *Dictionary) PopulateByUNIXCommand() {
	cmd := "cat /usr/share/dict/words"
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err == nil {
		words := fmt.Sprintf("%s", out)
		d.Words = strings.Split(words, "\n")
		d.wordsmap = make(map[string]bool, len(d.Words))
		for i := 0; i < len(d.Words); i++ {
			d.wordsmap[d.Words[i]] = true
		}
	}
}

func (d *Dictionary) InDictionary(word string) bool {
	if d.wordsmap[word] {
		return true
	}
	return false
}
