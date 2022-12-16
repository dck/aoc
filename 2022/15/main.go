package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

type Diamond struct {
	sensor Point
	beacon Point
	r      int
}

func (d *Diamond) Radius() int {
	if d.r == 0 {
		d.r = abs(d.sensor.x-d.beacon.x) + abs(d.sensor.y-d.beacon.y)
	}
	return d.r
}

func (d *Diamond) Inside(p Point) bool {
	r := d.Radius()

	distX := abs(d.sensor.x - p.x)
	distY := abs(d.sensor.y - p.y)

	return distX+distY <= r
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	diamonds := []Diamond{}
	for scanner.Scan() {
		line := scanner.Text()
		sensor, beacon := parseInput(line)

		diamonds = append(diamonds, Diamond{
			sensor: sensor,
			beacon: beacon,
		})
	}

	//fmt.Println(part1(diamonds, 2000000))

	borders := map[Point]int{}
	for _, d := range diamonds {
		r := d.Radius() + 1

		for x := d.sensor.x - r; x <= d.sensor.x+r; x++ {
			shift := r - abs(d.sensor.x-x)
			y1 := d.sensor.y - shift
			y2 := d.sensor.y + shift

			borders[Point{x, y1}] += 1

			if y2 != y1 {
				borders[Point{x, y2}] += 1
			}
		}
	}

	for p, i := range borders {
		if i >= 4 {
			found := false
			for _, d := range diamonds {
				if d.Inside(p) {
					found = true
					break
				}
			}

			if !found {
				fmt.Println(p.x*4000000 + p.y)
			}
		}
	}

}

func part1(diamonds []Diamond, y int) int {
	beacons := map[Point]bool{}
	for _, d := range diamonds {
		beacons[d.beacon] = true
	}

	filled := map[int]bool{}
	for _, d := range diamonds {
		x1, x2, err := overlapping(d, y)
		if err != nil {
			continue
		}

		for i := x1; i <= x2; i++ {
			p := Point{i, y}
			if _, ok := beacons[p]; ok {
				continue
			}

			filled[i] = true
		}
	}
	return len(filled)
}

func overlapping(d Diamond, y int) (int, int, error) {
	if y < d.sensor.y-d.Radius() || y > d.sensor.y+d.Radius() {
		return 0, 0, fmt.Errorf("there is no overlapping")
	}

	dist := abs(d.sensor.y - y)
	shiftOnY := d.Radius() - dist

	return d.sensor.x - shiftOnY, d.sensor.x + shiftOnY, nil
}

func parseInput(input string) (Point, Point) {
	var sx, sy, bx, by int
	_, _ = fmt.Sscanf(input, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)

	return Point{sx, sy}, Point{bx, by}
}
