package main

import "testing"

func TestNoSlashCorrectUrl(t *testing.T) {
	withNoSlash := "/rungo/unit-testing-made-easy-in-go-25077669318"
	output := noTwoSlash(withNoSlash)
	if output != withNoSlash {
		t.Errorf("Output is not equal to correct url string")
	}
}

func TestNoSlashDoubleSlashUrl(t *testing.T) {
	withSlash := "/rungo//unit-testing-made-easy-in-go-25077669318"
	correctURL := "/rungo/unit-testing-made-easy-in-go-25077669318"
	output := noTwoSlash(withSlash)
	for i := 0; i < len(correctURL)-1; i++ {
		if correctURL[i] != output[i] {
			t.Errorf("Found defference in index %d, out: %x, correct: %x", i, output[i], correctURL[i])
		}
	}
	if output != correctURL {
		t.Errorf(
			"Output is not equal to expected url string without slashes\nExp: %s, len: %d\nGot: %s, len: %d",
			correctURL,
			len(correctURL),
			output,
			len(correctURL),
		)
	}
}
