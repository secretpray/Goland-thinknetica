package thing

import (
	"io"
)

type Employee struct {
	Id   uint
	Rank string
	Name string
	Age  uint
}

func FindMaxAge(creatures ...interface{}) interface{} {
	if len(creatures) == 0 {
		return nil
	}

	maxAge := uint(0) // Initialize maxAge as uint
	var maxCreature interface{}

	for _, creature := range creatures {
		switch c := creature.(type) {
		case Customer:
			if c.Age > maxAge {
				maxAge = c.Age
				maxCreature = c
			}
		case Employee:
			if c.Age > maxAge {
				maxAge = c.Age
				maxCreature = c
			}
		}
	}

	return maxCreature
}

func PassOnlyStrings(writer io.Writer, args ...interface{}) []string {
	var result []string

	for _, val := range args {
		if str, ok := val.(string); ok {
			_, _ = writer.Write([]byte(str))
			result = append(result, str)
		}
	}

	return result
}
