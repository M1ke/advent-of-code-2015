/*

Santa is delivering presents to an infinite two-dimensional grid of houses.

He begins by delivering a present to the house at his starting location, and then an elf at the North Pole calls him via radio and tells him where to move next. Moves are always exactly one house to the north (^), south (v), east (>), or west (<). After each move, he delivers another present to the house at his new location.

However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off, and Santa ends up visiting some houses more than once. How many houses receive at least one present?

For example:

> delivers presents to 2 houses: one at the starting location, and one to the east.
^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.

*/

package main

import (
    "fmt"
    "log"
    "strconv"
    "io/ioutil"
)

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func main() {
    moves, err := ioutil.ReadFile("3-input")
    check(err)
    var coords [2]int
    coords[0] = 0
    coords[1] = 0
    history := make([]string, 0)
    history = append(history, "0,0")

    var coord string

    for _,move := range moves {
    	fmt.Println(string(move))
    	switch rune(move) {
    	case 'v':
    		coords[1] -= 1
    	case '<':
    		coords[0] -= 1
    	case '>':
    		coords[0] += 1
    	case '^':
    		coords[1] += 1
    	}
    	coord = strconv.Itoa(coords[0]) + "," + strconv.Itoa(coords[1])
    	fmt.Println(coord)
    	if (!stringInSlice(coord, history)){
    		history = append(history, coord)
    	}
    }

    fmt.Println(len(history))
}