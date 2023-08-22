package srttotext

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertSRTToText(t *testing.T) {
	srtContent := `1
00:00:00,498 --> 00:00:02,827
Here's what I love most about food and diet.

2
00:00:02,827 --> 00:00:06,383
We all eat several times a day, and we're totally in charge

3
00:00:06,383 --> 00:00:09,427
of what goes on our plate and what stays off.
`

	expected := "Here's what I love most about food and diet. We all eat several times a day, and we're totally in charge of what goes on our plate and what stays off."

	observed := Convert(srtContent)

	assert.Equal(t, expected, observed)
}

func ExampleConvert() {
	// Sample SRT content
	input := `
1
00:00:01,000 --> 00:00:04,000
Hello, world!

2
00:00:05,000 --> 00:00:08,000
This is a test.
`

	output := Convert(input)
	fmt.Println(output)

	// Output:
	// Hello, world! This is a test.
}
