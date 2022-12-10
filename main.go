package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Location represents a location on the virtual map.
type Location struct {
	Row int
	Col int
}

// Set represents a set data structure.
type Set[T comparable] struct {
	elements []T
}

// contains checks whether the element is already inside this set.
func (set *Set[T]) contains(element T) bool {
	for _, el := range set.elements {
		if el == element {
			return true
		}
	}

	return false
}

// Add adds the element to this set.
func (set *Set[T]) Add(element T) {
	if !set.contains(element) {
		set.elements = append(set.elements, element)
	}
}

// Size gets the number of items in this set.
func (set *Set[T]) Size() int {
	return len(set.elements)
}

// needToCatchUp checks whether the distance between the head and tail warrants the tail to move to catch-up
// with the head. This function returns 3 values. mustCatchUp indicates whether the tail needs to catch up with the
// head. distX and distY stores the distance of head and tail based on X and Y respectively.
func needToCatchUp(head Location, tail Location) (mustCatchUp bool, distX, distY int) {
	distX = head.Col - tail.Col
	distY = head.Row - tail.Row
	mustCatchUp = distX > 1 || distX < -1 || distY > 1 || distY < -1

	return
}

// move moves each of the elements based on the given direction indicated by dx and dy.
func move(locations []Location, visitedLocations *Set[Location], dx, dy int) {
	// Move the head
	locations[0].Row += dy
	locations[0].Col += dx

	for i := 1; i < len(locations); i++ {
		mustCatchUp, distX, distY := needToCatchUp(locations[i-1], locations[i])
		if mustCatchUp {
			// Check to which destination should we move on the ROW basis.
			dy = 0
			{
				if distY > 0 {
					dy = 1
				} else if distY < 0 {
					dy = -1
				}
			}

			// Check to which destination should we move on the COL basis.
			dx = 0
			{
				if distX > 0 {
					dx = 1
				} else if distX < 0 {
					dx = -1
				}
			}

			locations[i].Row += dy
			locations[i].Col += dx

			// We only add if this is a tail.
			if i == len(locations)-1 {
				visitedLocations.Add(Location{
					Row: locations[i].Row,
					Col: locations[i].Col,
				})
			}
		}
	}
}

func main() {
	// Read the input file.
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %s", err)
	}
	defer f.Close()

	// Create a new buffered reader to read the file content.
	r := bufio.NewReader(f)

	var part1Locations = []Location{{Row: 0, Col: 0}, {Row: 0, Col: 0}}
	var part2Locations = []Location{{Row: 0, Col: 0}, {Row: 0, Col: 0}, {Row: 0, Col: 0}, {Row: 0, Col: 0}, {Row: 0, Col: 0}, {Row: 0, Col: 0}, {Row: 0, Col: 0}, {Row: 0, Col: 0}, {Row: 0, Col: 0}, {Row: 0, Col: 0}}

	var visitedLocations1 = Set[Location]{}
	visitedLocations1.Add(Location{Row: 0, Col: 0})

	var visitedLocations2 = Set[Location]{}
	visitedLocations2.Add(Location{Row: 0, Col: 0})
	for {
		l, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatalf("failed when reading input file: %s", err)
		}

		if err == io.EOF {
			break
		}
		l = strings.TrimSpace(l)

		data := strings.Split(l, " ")
		dir := data[0]
		dist, _ := strconv.Atoi(data[1])

		// We determine the x and y movement based on the given direction.
		dx := 0
		dy := 0
		switch dir {
		case "R":
			dx = 1
			break

		case "L":
			dx = -1
			break

		case "U":
			dy = 1
			break

		case "D":
			dy = -1
			break
		}

		for i := 1; i <= dist; i++ {
			// Part 1 checking
			move(part1Locations, &visitedLocations1, dx, dy)

			// Part 2 checking
			move(part2Locations, &visitedLocations2, dx, dy)
		}
	}

	fmt.Println(visitedLocations1.Size())
	fmt.Println(visitedLocations2.Size())
}
