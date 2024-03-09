package main

import (
	"os"
	"runtime/pprof"
	"unicode/utf8"
)

func main() {
	file, _ := os.OpenFile("cpu.pprof", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o600)
	_ = pprof.StartCPUProfile(file)

	a, b := "hello", "hlelo"
	var distance int
	for range 100000000 {
		distance = CheckGuess(a, b)
	}
	_ = distance

	pprof.StopCPUProfile()
}

const (
	EqualGuess   = 0
	CloseGuess   = 1
	DistantGuess = 2
)

// CheckGuess compares the strings with eachother. Possible results:
//   - EqualGuess (0)
//   - CloseGuess (1)
//   - DistantGuess (2)
//
// This works mostly like levensthein distance, but doesn't check further than
// to a distance of 2 and also handles transpositions where the runes are
// directly next to eachother.
func CheckGuess(a, b string) int {
	// Simplifies logic lateron.
	if len(a) < len(b) {
		a, b = b, a
	} else if a == b {
		// Memcompare is an efficient shortcut
		return EqualGuess
	}

	// We only want to indicate a close guess if:
	//   * 1 additional character is found (abc ~ abcd)
	//   * 1 character is missing (abc ~ ab)
	//   * 1 character is wrong (abc ~ adc)
	//   * 2 characters are swapped (abc ~ acb)

	if len(a)-len(b) > CloseGuess {
		return DistantGuess
	}

	var distance int
	aBytes := []byte(a)
	bBytes := []byte(b)
	for {
		aRune, aSize := utf8.DecodeRune(aBytes)
		// If a eaches the end, then so does b, as we make sure a is longer at
		// the top, therefore we can be sure no additional conflict diff occurs.
		if aRune == utf8.RuneError {
			return distance
		}
		bRune, bSize := utf8.DecodeRune(bBytes)

		// Either different runes, or b is empty, returning RuneError (65533).
		if aRune != bRune {
			// Check for transposition (abc ~ acb)
			nextARune, nextASize := utf8.DecodeRune(aBytes[aSize:])
			if nextARune == bRune {
				if nextARune != utf8.RuneError {
					nextBRune, nextBSize := utf8.DecodeRune(bBytes[bSize:])
					if nextBRune == aRune {
						distance++
						aBytes = aBytes[aSize+nextASize:]
						bBytes = bBytes[bSize+nextBSize:]
						continue
					}
				}

				// Make sure to not pop from b, so we can compare the rest, in
				// case we are only missing one character for cases such as:
				//   abc ~ bc
				//   abcde ~ abde
				bSize = 0
			} else if distance == 1 {
				// We'd reach a diff of 2 now. Needs to happen after transposition
				// though, as transposition could still prove us wrong.
				return DistantGuess
			}

			distance++
		}

		aBytes = aBytes[aSize:]
		bBytes = bBytes[bSize:]
	}
}
