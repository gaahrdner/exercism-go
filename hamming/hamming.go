package hamming

import "errors"

const testVersion = 5

// Distance calculates the Hamming distance between two DNA strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("DNA sequences are not the same length")
	}
	hamming := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hamming++
		}
	}
	return hamming, nil
}
