package person

import "fmt"

var people = []string{"Andi", "John", "Jane"}

func GetAll() []string {
	return people
}

func getByName(name string) (string, error) {
	for _, personName := range people {
		if name == personName {
			return personName, nil
		}
	}
	return "", fmt.Errorf("no one with %v name", name)
}
