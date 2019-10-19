// Package twofer implements the game one for you and one for me
package twofer

// ShareWith returns a message with a given name
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
