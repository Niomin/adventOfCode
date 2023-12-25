package main

import (
	"fmt"
	"strings"
)

func D21() []int {
	q := []string{
		"...........\n.....###.#.\n.###.##..#.\n..#.#...#..\n....#.#....\n.##..S####.\n.##..#...#.\n.......##..\n.##.#.####.\n.##..##.##.\n...........",
		"...................................................................................................................................\n....#...........#.#.#.....#..##.........##............##.#............#.......#.........##........#.#............##..#.....#....##.\n..#.........#....#.#........#.#.........#........##.#...#...........................##..#......#....#....#...#.#..#...#.#.#........\n...###....#......#...........#.##...#.................#....................###.........#.#........#...........#..#.....#........#..\n.##..##...##....#.......#..........#.......#.................................#..##.....##......#........#.....##........#..........\n....#..........##..#...###.#.......#.......#.#.........#....................#..#..#...........................##...................\n......#...#..#...#................................#..#...................................#.#...#..#...........#................#...\n...#........#.......##....#....#..#..#.......#......................#.......#................###.........#............#.........#..\n...........#.#..#....#.#.#..........##....#####.....##.........#..#.#.........#..........#..###..#.#...###.........#........##.#...\n..........#.....#.........#....#.....#..#.#.................##..#................#..................#.##..............##....#..#...\n............##.#.#....#.#...................................#..#...#.##.........#.................##..#........#....#...........#..\n..........#..............##...#............#....##..................#..............................#.##..#........#..##...#........\n......#..#...#..#.....##..#..#...........#..#.#..#............#.......#..#..........#................#...#.#..#..#..............##.\n.......##....#........#.......#....#........##..#........#.....#..#.....#.............................................##.........#.\n.................#...#..........#..#....#...................#..#....##.#........................##...#.......##...........#..#...#.\n.....#......#...............................#..........#.##................#............#....#..........#.....#..#............#....\n................#....#....#.........#.#..................#.....#...........................................#....#...........#......\n....##.....##......#.##....#........##...#................#....................................#.#......#..#.....##............#...\n.#..#.......#.#...........#..........................#........#.........#....#...............#..#..#..#.....#..#...#..#..##...#....\n................##.#.##...............#.#..........##.........#...........#....................#....#..................##..........\n..................#..#.....#...........................................##...#.#...............#..#..#.......###...#......#..#...#..\n.....##..#...#......#......#....#...................##...##....##......##.#.#.............#..............##....##...........#......\n..#......#...#..........#..#......####................#............#..........#.................#...........#....#..##..#...###..#.\n.....#.......#...#......#......#.##...................#...#....#................#.#.#.........#..#...#.....#..##....#.........#....\n.#..#.#..##............#.....#..........................#.#...#.#...#........#....#...........#..#........##...#.#.................\n.#..#.##.#.#..#.......#..#.......#..........#..................#...............##....#..........###......#........##.#.......#...#.\n..#.#....#.........#..#.#.#...#.......................#.#....##.....#.......................................#..#...#............##.\n...#..#...........##........#...............##..#.......#......#.....#........###.#...............#............#.............##.##.\n.....#..#..........#............#.................##...#................#...#..........#..................................#...#....\n..#...................#.#.................................#.....#...............#.#...#..............##......#..#.........##.......\n..#....#...........#..........#..........###...#...##..........#.....#.#..........##.....#.#........#..............................\n..........##....#.#.........#..........#.....#...#.........#..........#...#.............#.###..........#..........#........#....#..\n.....#...#....#..#..#........#.........................#...#............#.#.#.###..#...#..#...........#.#............#.......#.....\n.....##....#..#....#.#....###.................##...........#...##........#.......###....##................#...#.#..#....##.......#.\n.......#.##.........#...................#......................#.........#...#.............................#.###...............#...\n....#..##...##...###......#............#.#........#.#....#..#.....#..##.........#.#............#...........#.#.....##..............\n.............#.....#.....#...........#.#.........#......#.#..#.............#......#.##......#..#...............#...#..#.#...#.#.#..\n......#...#....#....#.....................#.......#...#...........#...#......#..............##...#.................##.#..##...#....\n..#....#.#........#..................#....##.....#......#.......#..#.............#....#...#...#.#................................#.\n....#....#.#.....#.............#.........#.........#....#.#...#...#.................#...............#.......##.....................\n.#....##..#..#.#.................#.........................##...........#....#.#..#...#.....#...#.....................#.....#......\n.#....#..#.....#........................#.......#..#..#......#..#.#...#....#...###...#..........#.#...........#.....##.#...........\n....##......##.....#.......##....#.#......##..#....#.........##.#.##....#........#............##...#.###......................#....\n..#..............#...............#.....##........#....#.#.#........#........#........#.....##....#....##...........#.#.#...........\n.#.......#...#.#.........#.#.......#.....#.#....#....#..............#.#.#..#.#..................#.......#...................#......\n.............................##....#....#..#............##....#......#...#....................#..#.......#........#....#..#..#.....\n..............#.................#...#.#....#..##.......#...........#...###........###....#......##......##.............#.......#...\n.......##...#.................#..#...#.....###..#...#......#.......##....#..#.#..................#....#..........................#.\n........#..#.........###......#.....#...#..#...#.......#....#..#...#...#......##....#.....#...#...#.#.#...#............#.#....#....\n.....#................#....#........#.#...........##.......#.............#...........#........#.......#..#.........................\n.#..................#....#.......#...#................#......#.............#.....#...#.....#...#....#..........#..........##...#...\n..#...................#.....#.#..#.#..#.............#...................#...#...#..#.......#.#.....#.....................##........\n.#.......#.........#...#..............#.#...#.................#.#..#.........##....#...#...#....##.....#.......##........##.....#..\n........#.........#...#....#.#.#.....#...#........#..........#........................#......#...............#..............#.#....\n...##.#.............#..#..##..#..........##.#..#......#.....#.....#...#.........##.#....#.....#..............#.##..........##......\n.##.............#...#...#...#..........#...........##.......#.........#.#...#.............#...#.....##.....##.................#....\n....#........#......#...#..#.#.............#..#.............#.....#.#.....#.......#.#.......#.......#.......#.....##.#.......#.....\n.#..................#..#......#..##........#.......#..................##...#...............#..#...#.........#..#.#..##...........#.\n...........#.#......#....#.#.....#......#........#....#...#.........#.#.#...#..........#...#..#...#..#......##...#.#...#.........#.\n............#.#.#.....##..#.........#..#.#.........#.#...##..#..#.#...................#.##.........###..............#..............\n......................#..#.........#........#.#..#.#.#.#.............#.........##.#..#....#......##........#.#..#..................\n..........##................#......#......#....#..#..#.##...........#..........#.....#....#......#............#.........#..........\n...........#..............#..................#..#.....#.#.#...###.#.....#.........#..............#......#.###...#..................\n........#.#.####.........#........#.....#....#......#.#........#..........#.#........#.....##...##....#.##....##...#....#..........\n.................#....#....#.....###..#...#.##....##......#.......#.......#...##.....###.............#.#...#.....#.................\n.................................................................S.................................................................\n.......#...............#..#..#.........##.............##......#.......##.#......#...#.###......#..#.............#....#....###......\n........##.##.#..###......#...#.##....................#.......#....#....#......#........#...................#......................\n.......#.......#...#....#.........#.##....#.....#...........#.......#....#...#..#.......#..#..#..#........##.........#...#.........\n.........#..........##.#..#.#.#........................#..#.#.......##..#.....##...#.#.........##..................................\n.#...........#.##.#.....#.........#............#....##..#...##.##...............#.#...#..#...#.###.##....#.........#...............\n..................##...#.........#..........#..#.....#.#............#..##....##........#...#............#..##........#..#..........\n.##............#.#...#......#.#..##.....................#..#.##......#.........#..........#............#...........##..............\n....#........##...#...........#...........#...#............#..#.#............#...............##....##.........................#..#.\n.#...................#....#.........#.#...#..........#..........#..#....#.........#...#.#..#....#................#...#..........#..\n...................##...#.....#.....#...#.....#..#......#............####..............#.....#..#..#.#....#.....#..................\n...#...#.............#...#.....#.#.......#.#...#.#........................#...#.....#...#..............#.......##............#.....\n..#..#...........#.....##.....#............#...#.....#.#...#.......###.......#.#............#...................#..................\n........#...........#..#......#..#....................#..#....#.........................#..........................................\n........#.........#..#......#..#.#......#....#..#.........................#...#..#....#..#..#...#......#...#..#..............#.....\n..........#.....................#....#...#...........#......#.........##.......#..............#.#.....#........#................#..\n...#...#.............#...........#...#......##........#......#......##...#..........#.....##...#....#...##...............#.....#...\n.....#..............................#.#...................##...#...#..##..#.................###..............#.........#......##.#.\n....#.............................#.#.....#.....#...#........#..#.............#....#....###......##......#.....................#...\n.#......#...................#..................#....#...........#........#..........###.....#.##...#....#..........##...........##.\n.#.....##....................#...#..........#...#..#..#.........#..##...#....................#........#.............#.##...........\n....#....#.......#...............................#.........#......#####......................#......#.............#.#..........#...\n..#......#...................#..#........#..#.#...#..#.......................##..#.....##..............#..........#....#.#.......#.\n............#...............#....##.#............#...##......#..........##..##...............#.........................#.........#.\n...##......#...#..##............#.#...#....#...#.#.................#.#..#..#.###..#.....#...#.....#..#.............#............#..\n..#.#..#...#.#..#.#..........#.##.#.......#......#......#....#.............#..................................#........#..#....#.#.\n.....#...............##............#..#.##..#.#..#...........#..................................#.............#.....#.###....#.....\n.....#.........................#...............#.............#....#.....##....#.#.........#.#..#............#......................\n..#...#..#...........#..........#.#...#...#....#....................#...#......#.....#.#..#..#....#........#....................#..\n....#..#....#.##.....#...........................##........###....#..#...................#...#.#...............#.#..........#......\n...............#.#.#....................................#..#....#.....#..#...#.......#......#.............#..###..............#....\n....#.#..#..#.......#...#..........#....##.....#..............###..........#...#.........................#....#............#.##....\n.......###...##..........#..........#........#.#.#..#....#......#....#..#...#..........##................##............#...........\n.....#..##..#............#.#.........#.#....#....#......#..#.##.....#..#.#...#.....#.................#...#..#..##..........##......\n...........#.#..#....#..#.................#......#......###...............#.............#....................#..#.#.....##......#..\n...........#..#.........##...............#.#.............#........#..#.........#..........#.............#.......#...........#.##...\n.........#.#...........#................................................##........#.....#.............................#............\n.#.#..#..#....##...#.......#..#..............##......#..............#..............#...#............#..#..#..#...........#.#....#..\n...#.#...#.........#....#.....#...........###....#.#....#.........#.....##......#....#..............##....#...#...#..#.....#.......\n..#..##.......#.........#....................#.#...............#...#....##...#.....#............................#..........#.#.....\n....#....#...........#...#..#.......................##....#...#.....##..........#....##..............#..##...##....#.......#....#..\n......#...#................#...#.##............................#..............................#.#....#............#.#..#.#..##.....\n...............#.##..#.#.#....#....................#.........#.#..#.......#.................#....#.....#.................#...#.#...\n..#....##.#......#..#......#..#................###.....#......#.#.....#...........#.........#.#...........#..#.......#..#.#........\n.......#.......................#...........................#...................#..............#......##............#....#......#...\n..#....#.#..#..#.#.....#...#..#..#.....#........................#....#.#.........#.........#....##..#............#..#......##......\n.........#......##.#..........#..........#............#....#............#....#..........##.....#.#...#...#...##....##....#.......#.\n...#..#.#..................##..#...#....#.#..........#......#.#.....#...................................###.........#........#.....\n.................#......#...#..#........#............#......##.....#....###...........#..#.#.#..................#..#..#....#.....#.\n...#..#..#..#........##..#...........................#......#.##..##....#.#..........#....#......#...#.......#....#....##..........\n..#........##.#.....#..#.#.....#.....#...#..##.....................#....#..#..........#..#.#..#..#..##..........#.......#...#......\n............#..#..#.#.....#.#......#...#.##...............#...#.......#.............#.....##.#....#.##.........#....##........#....\n.##..#......##.................##.#.....#..................#......#..#.#.................##....#.#....#..............#.#.#..#......\n....#....#...#..#..#...#...#.....#............#...........##...........#.........#......#...#.....#...#..#..##..#.#.#....#.........\n..#....#.....#.....#.............#....#..#.#..#................#................#.............#......#...#.#.....#.......##..#.....\n......#.................#...#..........#...#.....#.........#.#.....##.......................#....#....#..#..#.#........#..#........\n......#....#....#.#......#............#...#...................................#.....#....#.......#.##....#...#..................#..\n.....................#...........#....#..#..#....#...#.........#...............#..#.........#........#........#.#...#.#........#...\n.##....#.#...#.....#........................#.............................................#...#.#...#.....##........#...#...#......\n......#......#.....#.......#............#......................................#......#......#.#.##..........#.#..##.#......#......\n.....#..#......#........#....#....#..........##......................................#...........#..........#..#.....###...#.#.....\n....#...................##.........##..#....#.#.#..##.......................###....#.###.#.....#...#.#.......#.....#.#...#....#....\n.##.........#.........#.#......#...##...................................#....##...#...#.......#....##......#..............#....#...\n.............##......#.##.......##.....##...#..#.#...##......................#..#...#...........#...##........#........##.#........\n..........#.......#..................###.##........#......#..........................#......#........#..#..............#...........\n...................................................................................................................................",
	}

	var res []int

	for _, qq := range q {
		input := strings.Split(qq, "\n")
		res = append(res, d21main(input, 64, false))
		res = append(res, d21p2(input, 26501365))
	}

	return res
}

func d21p2(input []string, totalSteps int) int {
	steps := []int{len(input) / 2, len(input)/2 + len(input), len(input)/2 + 2*len(input)}
	vals := []int{d21main(input, steps[0], true), d21main(input, steps[1], true), d21main(input, steps[2], true)}
	emptyVals := []int{sqr(steps[0] + 1), sqr(steps[1] + 1), sqr(steps[2] + 1)}

	diff := []int{emptyVals[0] - vals[0], emptyVals[1] - vals[1], emptyVals[2] - vals[2]}
	magicConst1 := diff[2] + diff[0] - 2*diff[1]
	magicConst2 := diff[1] - diff[0] - magicConst1

	cycles := totalSteps / len(input)

	step := diff[0]
	for i := 0; i < cycles; i++ {
		step += magicConst2 + magicConst1*(i+1)
	}
	empty := sqr(totalSteps + 1)
	return empty - step
}

func d21main(input []string, cnt int, isInfinite bool) int {
	steps := make([]int, cnt+1)
	n_ := d21gs(input)
	nodes := map[[2]int]int{n_[0]: 1}

	for i := 0; i < cnt; i++ {
		visited := make(map[[2]int]bool)
		steps[cnt] = len(nodes)
		newNodes := make(map[[2]int]int)
		for node, cnt := range nodes {
			for _, c := range [][2]int{{node[0] - 1, node[1]}, {node[0] + 1, node[1]}, {node[0], node[1] - 1}, {node[0], node[1] + 1}} {
				i, j := c[0], c[1]

				if isInfinite {
					i %= len(input)
					i = (i + len(input)) % len(input)
					j %= len(input[0])
					j = (j + len(input[0])) % len(input[0])
				} else {
					if i < 0 || i >= len(input) || j < 0 || j >= len(input[0]) {
						continue
					}
				}

				if input[i][j] == '#' {
					continue
				}
				if _, ok := visited[c]; ok {
					continue
				}
				newNodes[c] += cnt
				visited[c] = true
			}
		}
		nodes = newNodes
	}
	res := 0
	for _, cnt := range nodes {
		res += cnt
	}

	return res
}
func d21gs(input []string) [][2]int {
	nodes := make([][2]int, 0)

	for i := range input {
		for j := range input[i] {
			if input[i][j] == 'S' {
				nodes = append(nodes, [2]int{i, j})
			}
		}
	}
	return nodes
}

func d21pr(input []string, nodes [][2]int) {
	output := make([][]byte, len(input))
	for i := range output {
		output[i] = []byte(input[i])
	}
	for i := range nodes {
		output[nodes[i][0]][nodes[i][1]] = 'O'
	}
	for _, out := range output {
		fmt.Println(string(out))
	}
}