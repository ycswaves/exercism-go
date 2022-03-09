package robotname

import (
	"errors"
	"fmt"
)

// Define the Robot type here.
type Robot struct {
	hasName   bool
	nameIndex int
}

var (
	allPossibleNames      []string
	currNameIndex         int
	recycledNameIndexList []int
)

func createName(letterRemaining int, currStr string) {
	if letterRemaining <= 0 {
		allPossibleNames = append(allPossibleNames, currStr)
		return
	}

	if letterRemaining > 3 {
		for i := 'A'; i <= 'Z'; i++ {
			createName(letterRemaining-1, currStr+string(i))
		}
	} else {
		for i := 0; i <= 9; i++ {
			createName(letterRemaining-1, currStr+fmt.Sprint(i))
		}
	}

	//TODO: randomize allPossibleNames

}

func init() {
	createName(5, "")
}

func (r *Robot) Name() (string, error) {
	maxNameCount := len(allPossibleNames)
	if r.hasName {
		return allPossibleNames[r.nameIndex], nil
	} else if currNameIndex < maxNameCount {
		name := allPossibleNames[currNameIndex]
		r.nameIndex = currNameIndex
		r.hasName = true
		currNameIndex++
		return name, nil
	} else if len(recycledNameIndexList) > 0 {
		name := allPossibleNames[recycledNameIndexList[0]]
		r.nameIndex = recycledNameIndexList[0]
		recycledNameIndexList = recycledNameIndexList[1:]
		r.hasName = true
		return name, nil
	} else {
		return "", errors.New("err")
	}
}

func (r *Robot) Reset() {
	r.hasName = false
	if currNameIndex < len(allPossibleNames)-1 {
		return
	}
	recycledNameIndexList = append(recycledNameIndexList, r.nameIndex)
}
