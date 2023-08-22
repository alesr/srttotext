# srttotext
Go Utility for Extracting Plain Text from SRT Files

This Go code provides a utility to convert SRT content to plain text by removing timing and sequence number information. The Convert function takes SRT content as a string and returns the extracted textual content.

## Example

```go
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
```
