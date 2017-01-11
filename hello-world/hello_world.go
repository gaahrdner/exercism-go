// Package greeting greets a person given their name, or simply responds with
// "Hello World!"
package greeting

const testVersion = 3

// HelloWorld returns a friendly greeting
func HelloWorld(name string) string {
	// Write some code here to pass the test suite.
	if name == "" {
		return "Hello, World!"
	}
	return "Hello, " + name + "!"
}
