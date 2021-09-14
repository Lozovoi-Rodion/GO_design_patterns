package state

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type state int

const (
	Locked state = iota
	Failed
	Unlocked
)

func OpenLock() {
	code := "1234"
	currentState := Locked
	entry := strings.Builder{}

	for {
		switch currentState {
		case Locked:
			r, _, _ := bufio.NewReader(os.Stdin).ReadRune()
			entry.WriteRune(r)

			if entry.String() == code {
				currentState = Unlocked
				break
			}

			if strings.Index(code, entry.String()) != 0 {
				currentState = Failed
			}
		case Failed:
			fmt.Println("FAILED")
			entry.Reset()
			currentState = Locked
		case Unlocked:
			fmt.Println("UNLOCKED")
			return
		}
	}
}
