package lib

import (
	"strings"
)

//Function to build name id
func BuildID(name string) (string, error) {
	//Lets make lowercase and remove spaces
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")
	return name, nil
}

func CleanNameToKey(name string) string {
	//Lets make lowercase and remove spaces
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "_")
	return name
}
