package utils

import "fmt"

// Warning create a Warning Error structure
type Warning struct {
	Message string
}

// Warnings define an array of Warnings
type Warnings []Warning

// Error create a Error structure
type Error struct {
	Message string
}

func (warning *Warning) Error() string {
	return fmt.Sprintf("⚠️  %s\n", warning.Message)
}

func (warnings Warnings) Error() string {
	var output string
	for i := 0; i < len(warnings); i++ {
		output = fmt.Sprintf("%s%s", output, warnings[i].Error())
	}
	return output
}

func (error *Error) Error() string {
	return fmt.Sprintf("❌ %s\n", error.Message)
}
