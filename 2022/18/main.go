package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
	z int
}

var modifiers = [...]Point{
	{1, 0, 0},
	{-1, 0, 0},
	{0, 1, 0},
	{0, -1, 0},
	{0, 0, 1},
	{0, 0, -1},
}

func (p Point) Modify(a Point) Point {
	return Point{
		p.x + a.x,
		p.y + a.y,
		p.z + a.z,
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cubes := []Point{}
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ",")

		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		cubes = append(cubes, Point{x, y, z})
	}

	fmt.Println(Part1(cubes))
	fmt.Println(Part2(cubes))
}

func Part1(cubes []Point) int {
	total := len(cubes) * 6
	visited := map[Point]bool{}
	for _, c := range cubes {
		visited[c] = true

		for _, m := range modifiers {
			overlapping := c.Modify(m)

			if _, ok := visited[overlapping]; ok {
				total -= 2
			}
		}
	}

	return total
}

func Part2(cubes []Point) int {
	lava := map[Point]bool{}
	minPoint := Point{math.MaxInt, math.MaxInt, math.MaxInt}
	maxPoint := Point{math.MinInt, math.MinInt, math.MinInt}

	for _, c := range cubes {
		lava[c] = true
		minPoint.x = Min(minPoint.x, c.x)
		minPoint.y = Min(minPoint.y, c.y)
		minPoint.z = Min(minPoint.z, c.z)

		maxPoint.x = Max(maxPoint.x, c.x)
		maxPoint.y = Max(maxPoint.y, c.y)
		maxPoint.z = Max(maxPoint.z, c.z)
	}

	minPoint = minPoint.Modify(Point{-1, -1, -1})
	maxPoint = maxPoint.Modify(Point{1, 1, 1})

	water := map[Point]bool{}
	queue := []Point{minPoint}
	total := 0

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		if water[point] {
			continue
		}

		water[point] = true

		for _, m := range modifiers {
			adj := point.Modify(m)

			if adj.x < minPoint.x || adj.y < minPoint.y || adj.z < minPoint.z || adj.x > maxPoint.x || adj.y > maxPoint.y || adj.z > maxPoint.z {
				continue
			}

			if water[adj] {
				continue
			}

			if lava[adj] {
				total += 1
				continue
			}

			queue = append(queue, adj)
		}
	}

	return total
}
