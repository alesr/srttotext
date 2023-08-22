package srttotext

import (
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
