package main

import (
	"strings"
)

func D14() []int {
	inputs := []string{
		"O....#....\nO.OO#....#\n.....##...\nOO.#O....O\n.O.....O#.\nO.#..O.#.#\n..O..#O..O\n.......O..\n#....###..\n#OO..#....",
		"#.###...OO.#.O.O..O#.....#O#..O....#.#OO....O..#.....O##..#....#..#O...OOO#....##...OO#..O....O....O\n...O..........#.OO#....O....#..O.#.....##.##O..OOOO#....#OO..O...O#O.O###O..#..O.O..O..OO...OO#.O..O\nO............O....#....O....OO.#...#O..#.O......OO..#.........#O#..#.#......#..O#...###O...O.....OO.\n..#...#..O.#O#..OO#O..#O.....O...#.............OO.O..O...O...O....O....#.#.#O........#...O..OO#..#.#\n#.........#..#....O##.O#...O#..O#....O.##O.#.#..#....#..OO..#..##..O.##.O#..O#.....O......O....#O...\n...O#O....OO.....#....O#.......O#O..............O....OO..O#.........O#.O.OO..O...O#.O..O...O...#....\n..O.O#.....O#....OO#..O.#.#O...#..##.OO.O.O..#O....OO#......#...#...##.#..#O#OO#..O.OO..O.#.#.#...O#\n....O.O.OO.#.O#.#.#....O##.....##..#OO....O#O..O.O.O#.O.#.O...#..O..O......#..O.O##.#.#O........#O.#\n.O.#O#O.#.O.#..###...O.#.#O#....#O#....O..O.OO...O....##.......OO.......#........#....#..O...O.O....\n#O.##.O....#O..O.....#...OO..O#..O..#...#.....O..#.OO.........#O#.O.#O#...O.##.O.....O.........#....\n....OO#..O.O#.#.O..........O#........#.#O..O.OO.O..#O....#.O#O#..#.#.....O#...O.#.OOOOOO.#O.##.O...O\n..#.OO.O.....#.....#O..##..OO....O#.O.#....##.OOO..#O....O..O.#...O.#...O##O....#..O....OO.#O..O..#.\n....#O.O#.....O.....O.....O.O.##...O..O#OOOO..O..OO.O#.O#.O...O.OO.#O...#..O..O..O.#...O#......O.OO.\n...O....O.....OO....O..OO#.O#...O....O.##..OOO.#....#OO...#.O...........OOO....OO..........#...#...#\n......O.....OO#.##.O....#.O.....O.......O...O......#.O..O#.#..#.O#.#.O.##......#...##.#OO#O.O...##O.\n...#O......##..O.OO......O.##OO..O..#.O......O.##.#O....O...........#O..O#.O##.#.O...##....#.OOO.OO.\n.O#..OO...O.OO#O...#....O..O#OO....OO##.#.##O###O.##...O.O............OOO....#OOO..#.#..OO.#...#####\nO.##O.OO.#......##.O..O...#...OO.....O..#........O....O#O#..O#OO##....O..O..#O.#.........#..O.....#.\n.O....#.##...#O.O.OO...O#O#......#..O..#O#OO#.#.#O...#O##O#O.#...#..#...#.O...#.OO..OO....#O.#OO....\n....#...#OO....OOO.........O.O....#O....#......O..O.OO.O....#..#..#.#...O.O.OO....O#.......#.OO.O#O#\n..#O.#........OO.#...#...#.....#OO.......O.OO#.........##.........##.........#O..#.O#....O##......O.\n#O#O.#.O..#O#..O..O...#....#.###..#O.O..O#.#........#....#O.O.OO.#.###O#.#.O#..#O...#......O..O#....\nOO..#...............O.OOO...#.......O#.OO......O..O..O.#...OO.##.O........#..O#........##.#.#.O#..##\nO...#O.#.....O...##.O..O..#.O..O..O#......OO.OO#OO..O..O...OO#.#....O.O...#.#.#.#O.#..O..O.O#....O.#\nO...OO..O...O.O.O.O...#...#.....O.#O...........O...O......O.....#..O...OOO##..#.OO..OO...O.OO.##...#\n..O....##O#..#..#..O#.O......##.....##O.....#OO##..OO...........O#......#....#.###.O..OO#.#.......#.\n.O..#..O.....##....O...#.OO..#.OO..OOO#O..#..O....##O....OO....O.....#O.O..##O....#....OO..O...#O.O.\n..#..#.O#.O...#O..O...#..O.............O#.###.O..O..#..........O.#..#..O#.#..OO....#.........O......\n...OO..##O..#O#OO#O...##O#.....O...O.O....#...O..OOO..O..#.....O.O#.O.......O..O.#O...OO.#...O.##...\n#O..O#....O..O...O..O....#..........O.O...#.O..O.O.....O...O.#O##.O..O##......#..O......#....O.....O\nOO#OOO..O.#...OO.......OO......#.O......OO.O.O#.#....O#OOO#O.#.#O#.O..O..#.....##O.#....O..#.OOO.#.O\n.#.O.......O...O...#O....O.OO.#.O#......#OO##.#.##..........O.O#...OO#.OO#..O...#...#.O...O.....#...\n.O....OOO#..#..O#......O#..#O.#..#O.O#.O...##O..O..#..O...#O.OO...........O...#O.O.......O..#O......\n..#....O...O.....O.O.O...O#.#.....#...#O.O.O#...#...O.O#O....O.OOO#OO...O.#O..O##....O.O.###...OO.OO\n..#...#..O#..#OO#..O....O..O.#.O....O#.#......O..O.O###.O#.#.....O.#.##..OO#..O....#.O...#..#..#O...\n.....#.O.....#.##...#O........#...O....O......O#...OO#.......OO#...O.#..##O.O...O#..O.#..O....O..##.\n#.##.O#.......O.O......OO....O.#.#.O.O....#.O.O...O#..#...O..O##.O.O...#..O#.O..O.O...O..#...O.O.#..\n.......O.#.O...#O.O.....O.###..#OO..OO.....OO#O.....O..O.O..##..#........O..#O.#..##..O..O..O..O....\n.#...O....O#..O..O........O.O.O........OOO.#....O#...#.O....O...OOOO....O.....#..#.....#...#O.O.#.O.\n#O#..O.#O........O#....#...OO#.......O....#O....#.#.....O##...O.O...OO..OO#O#.#....#.O#.......O.OO..\n.###..#...#..#...#OOO#..#O...........#...#...O....O#..O#..O..O.#..#.#..O.#...OO.#.........OO.##.O.#.\n.OO..#..O.OO.O.#O..O..#O##..O.####..O..OO.#.OO.......O.O#OO....OO#OOO..#.#OO........#..#.#.O.O.O..OO\n.O.OO#.O.#..O..#O...#.#..O....##.O#.#O.O........O..#.OO....O...OOO#O.#...O................O.....O.O.\n.O......#...O.......O.##.#...O...O..OOO....OO..O....#..O.O.O.O..OO..O......O...O#.##..O.....O.OO...O\n..OO...O..O#.O.O.O.O..O.#.##.O..##.O......OO....O..O.O..O.....#..#..............OOO.#...#OO.O..O..##\n#.....O...OOO#........#O#.#....#....OOOOO..#.#OOO..O.O#O#O.....#..O.O..............O.O#O..#O.OO#.O..\n....O.O..O#..O.....#O#...O.#O.##.O..O.#.#.O.....#O...OOO....O....O......#..O.O.#......O...#O.O...O.#\n.O.#.O...O.O#....O..#...O....OO.......O.OO.O##..O#..........OO#.O#O#.O.O...#...O.O......#...........\n.#O.....O..#....#O............#..O..OO..O.O.OO.#O.....#....#O.#.....O....O..O..O..#...#...O......#..\n....OOO.O......#.O..#O..O..#..O.O...O...#O#.O...O.O......OO...O...#..O#.O..#.#O.#O.#O.##OO..O.......\n.O..O...#....O.O...O.OO.OO......OOOO...O#....#.O..#.#..#...O........#..#...###...#..........O....#.#\n##.#O...O......O#OO....O.....O..OO.O#.OO...##.OO#..O..OO.......#.#.OO.#....O#O.O.O...#.O##.#.O...#..\n..#O#..#.OO.O..#O.O...............#.#O...#....#...#..#..#O...........#......##.......O##..O.#...#..O\nO.....#.OO##.O..O..###OO....O...............OO..#.....OOO#.O...#O...#.#..#..##...#O#.OO...O......#..\n.##....#.O...O#O..O.##.##.O.#.O...#..#.##...#.......O.OOOOO#....#..O...............O.#O...O##.#....#\n#.##.O..#.#...O.....OO#.....O.O.OO.OO#...O..O#.....O...#.......##......#.O.....O#.#O.....O..#.#.....\n..#..O...O.O#...#O..O...O.O#.#OO...O...#......O.O.....#.O...O..OO....#...OO..O...#....#O..O.O#.....O\n.....O.O.....O.OO#O.O....OO#......#.O.O....#.O..OO.O.......O..O.......###.O....#O..#....O.O##....OO#\n#.#.O.##.....O....#.#.O.....O..O.O...##O#.O.#O.....#.#........O..O#...#......O#.#O#...#O........O.#.\n..O..#..........#..O.O.#..OO.O.O..O..O.OO.#O..#.O..OO..OO.O#..O#OO#O.#.#.O...O.............O....O..O\n.#...#O.#........O#..#.....O.##...#..O#..O..O..O.....O#.....OO..##..O..##O.#....O..#O.##.....O..OO..\n....#..O.#..#.O...#.....O.O.O.......#.O..O..........O.#O.#.O#..#.....O.#.....O..#O...OO...O..#......\nO...#...O........O.O.#O##O.....O....#...O....#....##...#.O.OO.....O.O##.OO#.#.O.OO.O.#O..O......O...\n..OO.....O...O.O...O..#..O##OO...#..#.........#O.O#O.......O#OO.O..O..O..O.##O#O.#O..OOO..#.O.O.O.#.\nO..O##...OO..#.#..O.#.O.O...#..O.....#....O#..O.O..O..OO.OO.........O..#.#.#...##..#.....O....O.#O..\n.#.#......O....O#..O.##.#O.O..OO...#O#O.#..OO.O#....O.##O#....O......#...#O....O..........O.#O....O.\n###...O....O...OO....O...O.O....#..#...#.O....O......O.O....#..##...O.#........OO.....#.O....O......\n...#.#.O..O#...#.O#.O#.#O#...OO.O#.#O...#O.......O#.#O..#..##.##O##.....##.O.....O...#....#.....##..\nOO.O..#.O...#....O#...#.##.#....OO..#..#..O....O#...O#..#..OO...#O.#....##O....#O.O....O.#..#.....O#\n.#.O.O..O..#...O...O......O#O....##....O......#......O..#.#...........OOO.#.....O.#.#.............O.\nO.....O...#OO.OOO..#....#....#O....O.O..........O...O...#.OO..O.O.##...O.OOOO..O...O....O..O.#......\n#O........O..#O#.O..O#.......#OO.....#....#O##.O.O#.O.#O.O....O.O#O....O.O.OO....#.......O...#O..O..\n#...O.#.#...#.O.O.#....O.OO...O..O....#...........O....#.OO.OO#.#OO.....OO..O.##....O..#O.....#O..#.\n##..#.OO.#OO..OO..##.........O....#...OOO..##O#..O.O.#.##O.#.##O.O...O..#O.O..O.O......#.#OO...O#O..\n.O....O#...O.....O#..O...#....O....O#......O..#.#O........#.O#..#..O.#.##..#..........#........O.#.O\n.OO#.O.O.###.........O.O..#...#.O..O...OO..#.OO....#..#.........O#.....OO..O......#.......O....O.O#.\nOO.OO.O.#.OO..#O..O.......##....O.O.#.##..##..#..O..O#...O...#...#...O.OOO.#O..O...#.O.O....OO......\n#.#..O.O.......OO.#.#..O..OO#O.....O#.#......O....O..O##.O.O##..O#.........#OO..O....#.#....O.O..O..\n#......#.O.#.....OO...#O.....#.O.........O.#..##O..OOO.##O....O...##...#..O#.O..#...O.OO...OO#....##\nO......OO#..O#.#.#...OO.#O....#O#.O#..#..OO...#..#..O.#....#.O.#O.##O...#.....O....OO..........O.##.\n.O...O..#O.OO...OOO.#..#..#.#..OO.O......#O#..O.........#........O.OO.O.###....#O#.#.#.#O..OOO....O.\nO.....#O#....#....OO..OO...O.....##O.O#.##.O.#......O.#...O#.#O.OO##.........O#...#O##.O...#..O.#..#\n.#.##...#O.#O..O#..#.#.#...#...O.O.O#..#.O..O.O#.#......OOO....#.O.O..O.O##..OOO..OO......O....OO#O.\nO..##......O...O...O#OO#.......O..O..#....#.#..O....O..OO..O##O#O..#..O...O.O.O..##...#...###.....O.\n.O..O...O...#...O...##...O#....O.....O.O.O#..#.....#....O..#.#...OO...OO.O..O..O##....O#O.#O..#....O\n.#..O..O..#.#.O.O.#..#.O.#.#.#.#..O..O..O..O...#O....##..#.#.O...OO..#O....#..#.#..O.....O.O.....O.#\n#...OO.#.....#.OOO#....O..O...O#O..#.OO..O.O.#...........O.....O.OO.#OO.O##O.......O....O...........\n...#.......##.O..#O..O#.#O.....O.#.O.....#O#.#.......#....OO#.....#.#..O.OO#O..O...#..#.O#.#OOO..O..\n.#.O............#OO.##O.O.O..#...O....O#O....O....OO..O..O.OO.#...OO#.#..##.OO..##..O.##..#...O.#O..\nO#O...#...O#.#..OO.....O.OO..O#...O..#O#OO#..#O##..O.....#.O.#.O.OOOO...O..#.#..O..O...#O..#O#.O....\n.O..##.#..O.O.O...O...O.O..#..O.#.O.O.O#.#.O...O.#.#.O##.#.O.O.##...O...O.O...#O#OOOO#...#.OO....#O.\nO..O.O....O.O....O.O.O.O.......O.O...O..#O#.OOO#............#.........O..O#.OO#..O#..#.#....#.#.....\n.##.O...O..........OO...#O..#.......#..#..##.#.O..O.OO.#..#...O..O#.#.#O..#.OOO#..O..O.OO.#..#..#...\n..O.......O#.O...OO.....O.OO..OO#.........#...O#.#.O......O.#...#...........OO..O....##....#.O...O..\n...#..O..#....#....##O...O.O....#....#O.....##.#.....OOOO..##.O##........#.O..O.....O...#....O....#O\n...O#O...O....O...O...O..#......#..OO..O...#..#...#O......O.O..O.O.O...O..O..#...##.O#.O..O###..OOO#\n...O.O.#OOOO.OO...O.......##.#O.#.O...#.....#..O........#.#.O#.#O#...O##.......O..O.....##..O....#.O\n..O.O#.O.OO..#...O....O....O...#O..##.....OO.#.#.#.OO.#..OO.......O....##..O#.O.O.......#O#O#O..OO.#\nOO#......OO.......O..O..#.OOOO#O......#.###.#O...#O#O...#..O..O....O#.#...#O.###O.#.OO....OO....#...\n#.#.O..#....##....#..#................O.......#O.#O#O...##O.OO...O...#...O.##....O.O......##.O..#..O",
	}
	var res []int
	for _, input := range inputs {
		lines := strings.Split(input, "\n")
		matrix := make([][]byte, len(lines))
		for i := range matrix {
			matrix[i] = []byte(lines[i])
		}
		res = append(res, d14p1(matrix, true))
		res = append(res, d14p2(matrix))
	}

	return res
}

func d14p2(matrix [][]byte) int {
	hashes := make(map[int]int)
	var from, loop int
	for i := 1; i < 1000; i++ {
		d14makeCycle(matrix)
		hash := d14getHash(matrix)
		if _, ok := hashes[hash]; ok {
			from = hashes[hash]
			loop = i - from
			break
		}
		hashes[hash] = i
	}
	num := (1_000_000_000 - from) % loop
	for i := 0; i < num; i++ {
		d14makeCycle(matrix)
	}

	return d14p1(matrix, false)
}

func d14getHash(matrix [][]byte) int {
	res := 0
	for i, line := range matrix {
		for j, c := range line {
			res += int(c) * (i + 3) * (j * 37 % 143)
		}
	}
	return res
}

func d14makeCycle(matrix [][]byte) {
	d14rot1(matrix, 1)
	d14rot2(matrix, 1)
	d14rot1(matrix, -1)
	d14rot2(matrix, -1)
}

func d14rot1(matrix [][]byte, n int) {
	for i := range matrix[0] {
		cnt, top := 0, -1
		if n == -1 {
			top = len(matrix[0])
		}
		for j := range matrix {
			l := j
			if n == -1 {
				l = len(matrix) - 1 - j
			}
			if matrix[l][i] == 'O' {
				cnt++
				matrix[l][i] = '.'
			}

			if matrix[l][i] == '#' || j == len(matrix)-1 {
				for k := top + n; k != top+(cnt+1)*n; k += n {
					matrix[k][i] = 'O'
				}
				top = l
				cnt = 0
			}
		}
	}
}

func d14rot2(matrix [][]byte, r int) {
	for i := range matrix {
		cnt, top := 0, -1
		if r == -1 {
			top = len(matrix)
		}
		for j := range matrix[0] {
			l := j
			if r == -1 {
				l = len(matrix[0]) - 1 - j
			}
			if matrix[i][l] == 'O' {
				cnt++
				matrix[i][l] = '.'
			}
			if matrix[i][l] == '#' || j == len(matrix[0])-1 {
				for k := top + r; k != top+(cnt+1)*r; k += r {
					matrix[i][k] = 'O'
				}
				top = l
				cnt = 0
			}
		}
	}
}

func d14p1(matrix [][]byte, withRotation bool) int {
	res := 0
	for i := range matrix[0] {
		col := make([]byte, len(matrix))
		for j := range matrix {
			col[j] = matrix[j][i]
		}
		if withRotation {
			res += d14col(col)
		} else {
			res += d14simple(col)
		}
	}

	return res
}

func d14simple(col []byte) int {
	res := 0
	for i := range col {
		if col[i] == 'O' {
			res += len(col) - i
		}
	}
	return res
}

func d14col(col []byte) int {
	res := 0
	m := len(col)
	subs := strings.Split(string(col), "#")
	for _, s := range subs {
		c := strings.Count(s, "O")

		res += (m + (m - c + 1)) * c / 2

		m -= len(s) + 1
	}

	return res
}
