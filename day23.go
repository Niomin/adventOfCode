package main

import (
	"strings"
)

func D23() []int {
	inputs := []string{
		"#.#####################\n#.......#########...###\n#######.#########.#.###\n###.....#.>.>.###.#.###\n###v#####.#v#.###.#.###\n###.>...#.#.#.....#...#\n###v###.#.#.#########.#\n###...#.#.#.......#...#\n#####.#.#.#######.#.###\n#.....#.#.#.......#...#\n#.#####.#.#.#########v#\n#.#...#...#...###...>.#\n#.#.#v#######v###.###v#\n#...#.>.#...>.>.#.###.#\n#####v#.#.###v#.#.###.#\n#.....#...#...#.#.#...#\n#.#########.###.#.#.###\n#...###...#...#...#.###\n###.###.#.###v#####v###\n#...#...#.#.>.>.#.>.###\n#.###.###.#.###.#.#v###\n#.....###...###...#...#\n#####################.#",
		"#.###########################################################################################################################################\n#.#...#####...###...#...#...#.......#...#...###...#...#...###...#.......###...#...#.......#...###.....#...#...###...#...#...#...........#...#\n#.#.#.#####.#.###.#.#.#.#.#.#.#####.#.#.#.#.###.#.#.#.#.#.###.#.#.#####.###.#.#.#.#.#####.#.#.###.###.#.#.#.#.###.#.#.#.#.#.#.#########.#.#.#\n#.#.#...###.#...#.#.#.#.#.#.#.....#.#.#.#.#...#.#.#.#.#.#...#.#...#.....#...#.#.#...#.....#.#.....#...#.#.#.#.#...#.#.#.#.#.#.#.........#.#.#\n#.#.###.###.###.#.#.#.#.#.#.#####.#.#.#.#.###.#.#.#.#.#.###.#.#####.#####.###.#.#####.#####.#######.###.#.#.#.#.###.#.#.#.#.#.#.#########.#.#\n#.#.###...#...#...#.#.#...#...>.>.#...#.#...#...#...#.#...#.#.....#.#...#...#.#.#.....#####...#.....#...#.#.#.#...#.#.#.#.#.#.#.###.....#.#.#\n#.#.#####.###.#####.#.#########v#######.###.#########.###.#.#####.#.#.#.###.#.#.#.###########.#.#####.###.#.#.###.#.#.#.#.#.#.#.###.###.#.#.#\n#.#.#.....#...#.....#.........#...#.....#...#.........#...#.#.....#.#.#.#...#.#.#...#...>.>.#.#.#...#.#...#.#.#...#.#.#.#.#.#.#.....#...#.#.#\n#.#.#.#####.###.#############.###.#.#####.###.#########.###.#.#####.#.#.#.###.#.###.#.###v#.#.#.#.#.#.#.###.#.#.###.#.#.#.#.#.#######.###.#.#\n#...#.....#...#.....#...#.....###.#.#...#...#.#...#...#.#...#.....#.#.#...###.#.###...#...#...#.#.#.#.#.#...#...#...#.#.#.#...###.....#...#.#\n#########.###.#####v#.#.#.#######.#.#.#.###.#.#.#.#.#.#.#.#######.#.#.#######.#.#######.#######.#.#.#.#.#.#######.###.#.#.#######.#####.###.#\n#.....#...#...#...#.>.#.#.#.......#.#.#...#.#.#.#.#.#.#.#.#.>.>...#...#...###...#.......#...#...#.#.#.#.#.....###...#.#...###.....#...#...#.#\n#.###.#.###.###.#.#v###.#.#.#######.#.###.#.#.#.#.#.#.#.#.#.#v#########.#.#######.#######.#.#.###.#.#.#.#####.#####.#.#######.#####.#.###.#.#\n#.#...#.....###.#.#...#...#.......#.#...#.#.#...#...#...#...#.#####.....#.#...###.........#.#...#.#.#.#.>.>.#.....#.#...#.....#...#.#.....#.#\n#.#.###########.#.###.###########.#.###.#.#.#################.#####.#####.#.#.#############.###.#.#.#.###v#.#####.#.###.#.#####.#.#.#######.#\n#.#.............#.....#.....#.....#.....#...#...###...###...#.......#.....#.#.###...........#...#.#.#.#...#.......#...#.#.#...#.#...#.......#\n#.#####################.###.#.###############.#.###.#.###.#.#########.#####.#.###.###########.###.#.#.#.#############.#.#.#.#.#.#####.#######\n#.....#...#.............#...#.....#...###...#.#.....#.....#...#...#...#.....#...#.........###.....#...#.........#...#...#...#...#...#.......#\n#####.#.#.#.#############.#######.#.#.###.#.#.###############.#.#.#.###.#######.#########.#####################.#.#.#############.#.#######.#\n#####...#...#...#...#...#.#.....#.#.#.....#.#...#...#...#####.#.#.#.###...#...#...#...#...#...#.....#...###...#.#.#.#.......#.....#.........#\n#############.#.#.#.#.#.#.#.###.#.#.#######.###.#.#.#.#.#####.#.#.#.#####.#.#.###.#.#.#.###.#.#.###.#.#.###.#.#.#.#.#.#####.#.###############\n#.......#...#.#...#.#.#.#.#...#...#.#.......#...#.#...#.....#...#.#.....#...#...#...#...#...#.#...#...#.....#...#.#...#.....#.....#.........#\n#.#####.#.#.#.#####.#.#.#.###.#####.#.#######.###.#########.#####.#####.#######.#########.###.###.###############.#####.#########.#.#######.#\n#.....#...#.#...#...#.#.#.#...#...#.#.#.....#.....#...#...#.....#.#.....###...#.......###...#.....#.....#...#...#.#.....#...#.....#.#.......#\n#####.#####.###.#.###.#.#.#.###.#.#.#.#.###.#######.#.#.#.#####.#.#.#######.#.#######.#####.#######.###.#.#.#.#.#.#.#####.#.#.#####.#.#######\n#.....#...#.....#.....#...#...#.#...#...#...#.......#...#.......#...#...#...#.........#...#.........###...#...#...#.......#.#.......#.......#\n#.#####.#.###################.#.#########.###.#######################.#.#.#############.#.#################################.###############.#\n#.......#.................###...#.........###.....#.......#.....#...#.#.#.......#...###.#.....#.........###...#...#...#.....#...#.......#...#\n#########################.#######.###############.#.#####.#.###.#.#.#.#.#######.#.#.###.#####.#.#######.###.#.#.#.#.#.#.#####.#.#.#####.#.###\n#.......................#.#.....#.....#.......#...#.#.....#.#...#.#.#.#.#...#...#.#...#.#.....#.......#.....#...#...#...#.....#.#.....#.#...#\n#.#####################.#.#.###.#####.#.#####.#.###.#.#####.#.###.#.#.#.#.#.#v###.###.#.#.###########.###################.#####.#####.#.###.#\n#.#.............#.....#...#...#.#...#...#.....#.....#.....#.#...#.#...#.#.#.>.>.#.#...#.#.#.....#...#.......#.....#...###.....#.......#.#...#\n#.#.###########.#.###.#######.#.#.#.#####.###############.#.###.#.#####.#.###v#.#.#.###.#.#.###.#.#.#######.#.###.#.#.#######.#########.#.###\n#...#...#...#...#.#...#.......#...#.....#.......#.....#...#...#.#.#...#...#...#...#...#.#.#...#.#.#.###...#.#...#...#.....###.........#...###\n#####.#.#.#.#.###.#.###.###############.#######.#.###.#.#####.#.#.#.#.#####.#########.#.#.###.#.#.#.###.#.#.###.#########.###########.#######\n#.....#...#.#.....#...#...........#...#.........#...#...#.....#.#...#.#.....#####.....#.#...#.#.#.#.#...#.#.....#...#.....#...#.....#.......#\n#.#########.#########.###########.#.#.#############.#####.#####.#####.#.#########.#####.###.#.#.#.#.#.###.#######.#.#.#####.#.#.###.#######.#\n#.....#...#...........#...#.....#...#...#...#.....#.#...#.....#.#.....#.......#...#...#.###.#.#.#.#.#...#.#.......#...#...#.#.#...#.#.......#\n#####.#.#.#############.#.#.###.#######.#.#.#.###.#.#.#.#####.#.#.###########.#.###.#.#.###.#.#.#.#.###.#.#v###########.#.#.#.###.#.#.#######\n#.....#.#.#.......#...#.#.#.#...#.......#.#.#...#.#...#.#...#.#.#.#.....###...#.#...#.#.#...#.#.#.#.#...#.>.>.#.......#.#.#.#.#...#.#.......#\n#.#####.#.#.#####.#.#.#.#.#.#.###.#######.#.###.#.#####v#.#.#.#.#.#.###.###.###.#.###.#.#.###.#.#.#.#.#####v#.#.#####.#.#.#.#.#.###.#######.#\n#.#...#.#...#.....#.#.#.#.#.#.###.....#...#...#.#.....>.>.#...#.#.#...#.....###...#...#.#...#.#.#.#...###...#.#.#.....#.#.#.#.#.#...#.......#\n#.#.#.#.#####v#####.#.#.#.#.#.#######v#.#####.#.#######v#######.#.###.#############.###.###.#.#.#.#######.###.#.#.#####.#.#.#.#.#.###v#######\n#...#...#...#.>.....#.#.#...#...#...>.>.###...#...#.....###...#...###.............#...#.#...#.#.#.###...#...#...#...#...#.#.#.#.#...>.#.....#\n#########.#.#v#######.#.#######.#.###v#####.#####.#.#######.#.###################.###.#.#.###.#.#.###.#.###.#######.#.###.#.#.#.#####v#.###.#\n#.........#...#.....#.#.......#.#.###.....#.......#.........#.......#.............###.#.#...#.#.#.#...#.#...###...#.#...#.#.#.#.#.....#.#...#\n#.#############.###.#.#######.#.#.#######.#########################.#.###############.#.###.#.#.#.#.###.#.#####.#.#.###.#.#.#.#.#.#####.#.###\n#.......#...#...#...#.#...#...#.#.#.......#.......###...............#.#.....#.......#.#...#.#.#.#.#...#...#...#.#.#.#...#...#...#.......#...#\n#######.#.#.#.###.###.#.#.#.###.#.#.#######.#####.###.###############.#.###.#.#####.#.###.#.#.#.#.###.#####.#.#.#.#.#.#####################.#\n#.....#...#...###...#.#.#...###...#.......#.#...#.....#...#...#.....#...###...#...#.#.....#...#...#...#...#.#...#.#.#.#...#...#...#.........#\n#.###.#############.#.#.#################.#.#.#.#######.#.#.#.#.###.###########.#.#.###############.###.#.#.#####.#.#.#.#.#.#.#.#.#.#########\n#...#.#.............#...#.....#...###.....#.#.#...#...#.#.#.#...#...#...........#...#.....#.....###.....#.#.#.....#.#...#.#.#...#...#.......#\n###.#.#.#################.###.#.#.###.#####.#.###.#.#.#.#.#.#####.###.###############.###.#.###.#########.#.#.#####.#####.#.#########.#####.#\n#...#...#.............###...#...#.....#####...###...#.#.#.#.#.....###.............#...#...#...#.#.......#.#.#.....#.....#.#...........#.....#\n#.#######.###########.#####.#########################.#.#.#.#.###################.#.###.#####.#.#.#####.#.#.#####.#####.#.#############.#####\n#...#...#.#...........#.....#...........#...###...###...#...#.....#...#...........#...#.....#.#.#.....#.#...###...#...#...#.............#...#\n###.#.#.#.#.###########.#####.#########.#.#.###.#.###############.#.#.#.#############.#####.#.#.#####.#.#######.###.#.#####.#############.#.#\n###.#.#...#...........#.......#...#.....#.#...#.#...#.....#.....#.#.#.#.........#...#.#.....#.#.###...#.....#...#...#.....#...........#...#.#\n###.#.###############.#########.#.#.#####.###.#.###.#.###.#.###.#.#.#.#########.#.#.#.#.#####.#.###.#######.#.###.#######.###########.#.###.#\n###...#...............#.......#.#.#...###.#...#.#...#...#.#...#...#.#.#...#.....#.#...#.#...#.#...#.....#...#...#.#.......#...........#.#...#\n#######.###############.#####.#.#.###v###.#.###.#.#####.#.###v#####.#.#.#.#.#####.#####.#.#.#.###.#####.#.#####.#.#.#######.###########.#.###\n###...#...........#...#.....#.#.#...>.>...#.#...#...#...#.#.>.>.#...#.#.#.#.#...#...#...#.#.#...#.#...#.#.#...#.#.#...#.....#...........#.###\n###.#.###########v#.#.#####.#.#.#####v#####.#.#####.#.###.#.#v#.#.###.#.#.#v#.#.###.#.###.#.###.#.#.#.#.#.#.#.#v#.###.#.#####.###########.###\n#...#.....###...#.>.#.#...#.#.#.....#.....#.#...#...#...#.#.#.#...#...#.#.>.>.#.....#...#.#.....#.#.#.#.#.#.#.>.>.#...#.......###...#...#...#\n#.#######.###.#.#v###.#.#.#.#.#####.#####.#.###.#.#####.#.#.#.#####.###.###v###########.#.#######.#.#.#.#.#.###v###.#############.#.#.#.###.#\n#...#.....#...#.#.###.#.#.#.#.#...#.#.....#.....#...#...#...#.....#.....#...#.........#.#...#.....#.#.#.#.#.###...#...#...###...#.#.#.#.#...#\n###.#.#####.###.#.###.#.#.#.#.#.#.#.#.#############.#.###########.#######.###.#######.#.###.#.#####.#.#.#.#.#####.###.#.#.###.#.#v#.#.#.#.###\n#...#.......#...#...#.#.#...#.#.#.#.#...........#...#.###.........#.....#...#.#.......#...#.#.#...#.#...#...#...#.#...#.#...#.#.>.#...#...###\n#.###########.#####.#.#.#####.#.#.#.###########.#.###.###.#########.###.###.#.#.#########.#.#.#.#.#.#########.#.#.#.###.###.#.###v###########\n#.....#.....#.#...#.#.#...###...#.#.#...........#...#...#...........#...###...#.......#...#.#...#.#.#...#.....#...#...#...#...###...........#\n#####.#.###.#.#.#.#.#.###.#######.#.#.#############.###.#############.###############.#.###.#####.#.#.#.#.###########.###.#################.#\n#####...#...#...#.#.#.#...#.....#...#...........###.....#.........#...#...#.........#.#.....#...#...#.#.#.......#####.....#.................#\n#########.#######.#.#.#.###.###.###############.#########.#######.#.###.#.#.#######.#.#######.#.#####.#.#######.###########.#################\n#.......#.......#.#.#.#.###...#...#.............#...#...#.......#...#...#.#...#...#...###.....#.#...#.#.......#...........#.................#\n#.#####.#######.#.#.#.#.#####.###.#.#############.#.#.#.#######.#####.###.###.#.#.#######.#####.#.#.#.#######.###########.#################.#\n#.....#.....#...#...#...#...#.#...#.......###...#.#...#.....#...#.....#...#...#.#...#...#...#...#.#.#...#...#...........#.#####.......#.....#\n#####.#####.#.###########.#.#.#.#########.###.#.#.#########.#.###.#####.###.###.###.#.#.###.#.###.#.###v#.#.###########.#.#####.#####.#.#####\n#...#.....#...#...#...###.#.#.#...#.....#...#.#.#...#.......#...#.....#.###...#.#...#.#.#...#...#.#.#.>.>.#.....#...###...#.....#...#...#...#\n#.#.#####.#####.#.#.#.###.#.#.###.#.###.###.#.#.###.#.#########v#####.#.#####.#.#.###.#.#.#####.#.#.#.#v#######.#.#.#######.#####.#.#####.#.#\n#.#...#...#...#.#.#.#.###.#.#.#...#...#.....#.#.....#.........>.>.....#.....#...#.#...#...#####.#.#.#.#.#...#...#.#.#...###.......#.......#.#\n#.###.#v###.#.#.#.#.#.###.#.#.#.#####.#######.#################v###########.#####.#.###########.#.#.#.#.#.#.#.###.#.#.#.###################.#\n#...#.#.>.#.#.#.#...#...#.#.#.#...###.......#.........###...###...#.........###...#...........#.#.#.#.#.#.#.#...#.#.#.#.#...#...###...#...#.#\n###.#.#v#.#.#.#.#######.#.#.#.###.#########.#########.###.#.#####.#.###########.#############.#.#.#.#.#.#.#.###.#.#.#.#.#.#.#.#.###v#.#.#.#.#\n#...#...#.#.#.#.....#...#.#.#.#...#...###...#.....#...#...#...#...#...#...#...#.....#.........#.#.#.#.#...#...#.#.#.#.#...#.#.#...>.#...#...#\n#.#######.#.#.#####.#.###.#.#.#.###.#.###v###.###.#.###.#####.#.#####.#.#.#.#.#####.#.#########.#.#.#.#######.#.#.#.#.#####.#.#####v#########\n#...#...#...#.#...#.#...#.#.#.#...#.#...>.>.#.#...#...#.....#.#...###...#.#.#...#...#.........#...#.#.#.......#.#.#...#...#...#...#.........#\n###.#.#.#####.#.#.#.###.#.#.#.###.#.#####v#.#.#.#####.#####.#.###.#######.#.###.#.###########.#####.#.#.#######.#.#####.#.#####.#.#########.#\n#...#.#.###...#.#.#.#...#.#.#...#.#.#.....#...#.......###...#.#...#.......#.#...#.#...#.....#.#...#...#.......#...###...#.......#...#.....#.#\n#.###.#.###.###.#.#.#.###.#.###.#.#.#.###################.###.#.###.#######.#.###v#.#.#.###.#.#.#.###########.#######.#############.#.###.#.#\n#.#...#...#.....#...#...#.#...#.#.#.#.....#...#.....###...###.#...#.....#...#.#.>.>.#...###...#.#...#.........###...#.....#.......#.#...#...#\n#.#.#####.#############.#.###.#.#.#.#####.#.#.#.###.###.#####.###.#####.#.###.#.#v#############.###.#.###########.#.#####.#.#####.#.###.#####\n#...#...#.............#.#.#...#.#.#...###...#.#...#.....#####.....###...#...#.#.#.#...........#.#...#.............#...###...#.....#.#...#...#\n#####.#.#############.#.#.#.###.#.###.#######.###.###################.#####.#.#.#.#.#########.#.#.###################.#######.#####.#.###.#.#\n#.....#...............#...#.....#.....#.......###.........#.....#.....#.....#...#...#.........#.#.#...#...............#...#...#...#...###.#.#\n#.#####################################.#################.#.###.#.#####.#############.#########.#.#.#.#.###############.#.#.###.#.#######.#.#\n#...........#...###...#...#.........###...........#.......#...#.#.#.....#...#.........###.....#.#.#.#.#...#.....#...#...#.#.....#.........#.#\n###########.#.#.###.#.#.#.#.#######.#############.#.#########.#.#.#.#####.#.#.###########.###.#.#.#.#.###.#.###.#.#.#.###.#################.#\n###...###...#.#...#.#.#.#.#.#.......#.............#...#####...#.#...#...#.#.#...........#.#...#.#.#.#.###...#...#.#.#.#...#.....#.........#.#\n###.#.###v###.###.#.#.#.#.#.#.#######.###############.#####.###.#####.#.#.#.###########.#.#.###.#.#.#.#######v###.#.#.#.###.###.#.#######.#.#\n#...#...#.>.#.#...#.#...#...#.......#.............###...#...#...#...#.#.#.#.....#.......#.#.....#...#.......>.>.#.#.#.#...#...#...#.....#...#\n#.#####.#v#.#.#.###.###############.#############.#####.#.###.###.#.#.#.#.#####.#.#######.###################v#.#.#.#.###.###.#####.###.#####\n#.....#...#.#.#...#.#.......#.......#...#.........#.....#...#...#.#.#.#.#.....#.#...#...#...........###...#...#.#.#.#.#...###.......#...#...#\n#####.#####.#.###.#.#.#####.#.#######.#.#.#########.#######.###.#.#.#.#.#####.#.###.#.#.###########.###.#.#.###.#.#.#.#.#############.###.#.#\n#...#.....#...###.#.#.#...#...###...#.#.#...###...#...#...#...#.#.#.#.#.#...#.#.###...#.#.....#.....#...#.#...#.#.#...#...###...#...#...#.#.#\n#.#.#####.#######.#.#.#.#.#######.#.#.#.###.###.#.###v#.#.###.#.#.#.#.#.#.#.#.#.#######v#.###.#.#####.###.###.#.#.#######.###v#.#.#.###.#.#.#\n#.#.#...#.......#...#.#.#.###...#.#.#.#.###.#...#.#.>.>.#...#.#...#.#.#.#.#.#.#...#...>.>.###.#.#.....###.....#...#.......#.>.#...#.#...#.#.#\n#.#.#.#.#######.#####.#.#.###.#.#.#.#.#.###v#.###.#.#v#####.#.#####.#.#.#.#.#.###.#.###v#####.#.#.#################.#######.#v#####.#.###.#.#\n#.#.#.#.........#.....#.#.#...#...#.#.#...>.>.###.#.#.#...#...#.....#.#.#.#.#.#...#.#...#####...#.....#...###...###.........#.....#...#...#.#\n#.#.#.###########.#####.#.#.#######.#.#####v#####.#.#.#.#.#####.#####.#.#.#.#.#.###.#.###############.#.#.###.#.#################.#####.###.#\n#.#.#.........###...#...#.#.#.....#...#...#...###.#.#...#.....#.#...#.#.#.#.#.#...#.#.....#.......#...#.#.....#...#...............#.....#...#\n#.#.#########.#####.#.###.#.#.###.#####.#.###.###.#.#########.#.#.#.#.#.#.#.#.###.#.#####.#.#####.#.###.#########.#.###############.#####.###\n#.#.....#...#...###.#.#...#.#...#...#...#.....#...#.#.........#.#.#...#...#...###.#.#.....#.....#...###.#.........#...#.............#.....###\n#.#####.#.#.###.###.#.#.###.###.###.#.#########.###.#.#########.#.###############.#.#.#########.#######.#.###########.#.#############.#######\n#.#...#...#.....#...#.#...#...#...#.#.......###.....#.........#...#.....#...#...#...#.....#.....#.......#.........###...#.............#.....#\n#.#.#.###########.###.###.###.###.#.#######.#################.#####.###.#.#.#.#.#########.#.#####.###############.#######.#############.###.#\n#...#...........#...#...#...#.#...#.........#...#...#...#...#...###.#...#.#...#.....#...#.#.......#...#...#...#...#.....#...............#...#\n###############.###.###.###.#.#.#############.#.#.#.#.#.#.#.###.###.#.###.#########.#.#.#.#########.#.#.#.#.#.#.###.###.#################.###\n#.............#...#...#.###...#...#...#.....#.#...#...#...#.#...#...#...#.#.....###...#.#.#...#.....#.#.#...#.#.###...#.....#...#...#...#...#\n#.###########.###.###.#.#########.#.#.#.###.#.#############.#.###.#####.#.#.###.#######.#.#.#.#.#####.#.#####.#.#####.#####.#.#.#.#.#.#.###.#\n#...........#.#...###...#####...#...#...#...#.............#...###...#...#.#.###...#...#...#.#.#.....#...#...#...#...#.....#...#...#.#.#...#.#\n###########.#.#.#############.#.#########.###############.#########.#.###.#.#####.#.#.#####.#.#####.#####.#.#####.#.#####.#########.#.###.#.#\n#...........#...###.....#.....#.....#.....###...#.....#...#.....###.#.###...#.....#.#...###.#.......#...#.#.###...#...###.......###...###...#\n#.#################.###.#.#########.#.#######.#.#.###.#.###.###.###.#.#######.#####.###.###.#########.#.#.#.###.#####.#########.#############\n#.#.......#...#...#...#.#.........#.#.#...###.#.#...#.#...#...#.....#...#...#...###...#.#...#...#.....#...#...#...#...#...#.....#...#.......#\n#.#.#####.#.#.#.#.###.#.#########.#.#.#.#.###.#.###.#.###v###.#########.#.#.###.#####.#.#.###.#.#.###########.###.#.###.#.#.#####.#.#.#####.#\n#...#.....#.#...#.###.#.###.......#.#.#.#...#.#.#...#.#.>.>.#.#.........#.#.###...#...#.#.#...#...#...........#...#...#.#.#.....#.#.#.#.....#\n#####.#####.#####.###.#.###.#######.#.#.###.#.#.#.###.#.###.#.#.#########.#.#####.#.###.#.#.#######.###########.#####.#.#.#####.#.#.#.#.#####\n#.....#...#.#.....#...#.#...#...###...#...#.#.#.#.###...###.#.#...###...#.#.###...#.#...#...###...#.........###...#...#.#.###...#.#.#.#.....#\n#.#####.#.#.#.#####.###.#.###.#.#########.#.#.#.#.#########.#.###.###.#.#.#.###v###.#.#########.#.#########.#####.#.###.#.###v###.#.#.#####.#\n#.......#.#.#.###...#...#.....#...###...#.#.#.#.#.........#.#.###...#.#.#.#.#.>.>.#.#...###...#.#...........#...#.#.#...#.#.>.###.#.#.#.....#\n#########.#.#.###.###.###########.###.#.#.#.#.#.#########.#.#.#####.#.#.#.#.#.###.#.###.###.#.#.#############.#.#.#.#.###.#.#v###.#.#.#.#####\n###...#...#.#...#.#...#.....#.....#...#...#...#.......#...#.#.#.....#.#.#.#.#.###.#.###...#.#.#.....#.....#...#.#.#.#.#...#.#...#.#.#.#.#...#\n###.#.#.###.###.#.#.###.###.#.#####.#################.#.###.#.#.#####.#.#.#.#.###.#.#####.#.#.#####.#.###.#.###.#.#.#.#.###.###.#.#.#.#.#.#.#\n#...#...#...#...#.#.###.#...#...###...............#...#...#...#...#...#...#.#...#.#.#.....#.#...#...#.###.#.###.#.#...#...#...#...#.#.#.#.#.#\n#.#######.###.###.#.###.#.#####v#################.#.#####.#######.#.#######.###.#.#.#.#####.###.#.###.###.#.###.#.#######.###.#####.#.#.#.#.#\n#.#.....#...#...#.#...#.#...#.>.>.#...#...........#...#...#.......#.#.......#...#.#.#.....#.#...#...#...#...#...#.....#...#...###...#.#...#.#\n#.#.###.###.###.#.###.#.###.#.###.#.#.#.#############.#.###.#######.#.#######.###.#.#####.#.#.#####v###.#####.#######.#.###.#####.###.#####.#\n#.#.#...#...#...#...#.#.#...#...#.#.#.#...........#...#...#...#...#.#...#...#...#.#.#.....#.#.#...>.>.#...#...#.....#.#...#.#.....#...#.....#\n#.#.#.###.###.#####.#.#.#.#####.#.#.#.###########.#.#####.###.#.#.#.###.#.#.###.#.#.#.#####.#.#.#####.###.#.###.###.#.###.#.#.#####.###.#####\n#...#.....###.......#...#.......#...#.............#.......###...#...###...#.....#...#.......#...#####.....#.....###...###...#.......###.....#\n###########################################################################################################################################.#",
	}

	var res []int
	for _, input := range inputs {
		data := strings.Split(input, "\n")
		res = append(res, d23p1(data, 1, 0, len(data[0])-2, len(data)-1, false))
		res = append(res, d23p2(data))
	}

	return res
}

type d23node struct {
	x, y    int
	visited map[[2]int]bool
}

func d23p2(input []string) int {
	f := d23getForks(input)
	m := make(map[[2]int]bool, len(f))
	for _, pt := range f {
		m[pt] = true
	}

	graph := make(map[[2]int]map[[2]int]int)

	for _, pt := range f {
		graph[pt] = d23bfs(input, pt, m, true)
	}

	res := d23dfs(graph, f[0], f[len(f)-1], map[[2]int]bool{})

	return res
}

func d23dfs(graph map[[2]int]map[[2]int]int, from [2]int, to [2]int, visited map[[2]int]bool) int {
	if from == to {
		return 0
	}
	res := 0
	for pt, price := range graph[from] {
		if visited[pt] {
			continue
		}
		visited[pt] = true
		res = max(price+d23dfs(graph, pt, to, visited), res)
		visited[pt] = false
	}
	return res
}

func d23getForks(input []string) [][2]int {
	var res [][2]int
	res = append(res, [2]int{1, 0})
	for i := range input {
		for j := range input[i] {
			if input[i][j] == '#' {
				continue
			}
			n := [][2]int{{j - 1, i}, {j + 1, i}, {j, i + 1}, {j, i - 1}}
			cnt := 0
			for _, p := range n {
				if !d23existsNotSharp(input, p) {
					continue
				}
				cnt++
			}
			if cnt > 2 {
				res = append(res, [2]int{j, i})
			}
		}
	}
	res = append(res, [2]int{len(input[0]) - 2, len(input) - 1})

	return res
}

func d23existsNotSharp(input []string, p [2]int) bool {
	if p[0] < 0 || p[0] == len(input) || p[1] < 0 || p[1] == len(input[0]) {
		return false
	}
	if input[p[1]][p[0]] == '#' {
		return false
	}
	return true
}

func d23bfs(input []string, from [2]int, to map[[2]int]bool, ignoreSlopes bool) map[[2]int]int {
	nodes := make([]*d23node, 1, 10000)
	nodes[0] = &d23node{x: from[0], y: from[1], visited: make(map[[2]int]bool)}
	slopes := map[byte][2]int{'^': {0, -1}, '>': {1, 0}, '<': {-1, 0}, 'v': {0, 1}}

	steps := -1
	res := make(map[[2]int]int)
	for len(nodes) > 0 {
		steps++
		l := len(nodes)

		for i := 0; i < l; i++ {
			x, y, visited := nodes[i].x, nodes[i].y, nodes[i].visited
			if !d23existsNotSharp(input, [2]int{x, y}) {
				continue
			}
			if visited[[2]int{x, y}] {
				continue
			}
			pt := [2]int{x, y}
			if pt != from && to[pt] {
				res[pt] = max(res[pt], steps)
				continue
			}
			visited[[2]int{x, y}] = true
			var difs [][2]int
			if !ignoreSlopes && input[y][x] != '.' {
				difs = [][2]int{slopes[input[y][x]]}
			} else {
				for _, s := range slopes {
					x, y := nodes[i].x+s[0], nodes[i].y+s[1]
					if !d23existsNotSharp(input, [2]int{x, y}) {
						continue
					}
					if visited[[2]int{x, y}] || input[y][x] == '#' {
						continue
					}
					difs = append(difs, s)
				}
			}
			for _, d := range difs {
				if len(difs) > 1 {
					v := d23cv(visited)
					nodes = append(nodes, &d23node{x: x + d[0], y: y + d[1], visited: v})
				} else {
					nodes = append(nodes, &d23node{x: x + d[0], y: y + d[1], visited: visited})
				}
			}
		}

		nodes = nodes[l:]
	}
	return res
}

func d23p1(input []string, x, y int, x1, y1 int, ignoreSlopes bool) int {
	bfs := d23bfs(input, [2]int{x, y}, map[[2]int]bool{{x1, y1}: true}, ignoreSlopes)
	return bfs[[2]int{x1, y1}]
}

func d23cv(visited map[[2]int]bool) map[[2]int]bool {
	clone := make(map[[2]int]bool)
	for i, v := range visited {
		clone[i] = v
	}
	return clone
}