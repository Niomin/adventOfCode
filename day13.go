package main

import (
	"strings"
)

func D13() []int {
	input := ".#.##.#.##..###\n...##...#######\n#.####.#.#.###.\n#..##..##..#...\n###..###....###\n.##..##..#.#...\n.#....#..######\n#..##..########\n########.#..#..\n\n.#.###.###..###.#\n..#...##########.\n..######.#..#.###\n..##.#.#......#..\n..#.##....##....#\n.##..#.##.##.##.#\n#.....####..####.\n#.....####..####.\n.##..#.##.##.##.#\n\n####.####.#\n#..#.#..#.#\n#..#.####.#\n####.#..#.#\n.##.##..##.\n######..###\n#####.##.##\n#..########\n......##...\n.##........\n#####....##\n######..###\n..#.#.##.#.\n######..###\n.##...##...\n#..#.####.#\n######..###\n\n.##..#.##\n.##.####.\n####.....\n.##...##.\n.##.#.###\n#####..##\n.##.#.##.\n....#.###\n####.####\n#..###..#\n.#..#####\n#..#.#...\n.....##..\n#####.###\n#####.###\n\n.##..##.#####\n.#.##.#..#..#\n##....###.##.\n#......###..#\n.#.##.#......\n...##.....##.\n#......#...#.\n.#....#......\n#......#.####\n#.#..#.#..##.\n#..##..######\n\n..#..#..#..##..#.\n...##...#...##.##\n...##...#...##.##\n..#..#..#..##..#.\n..#..#...##.#####\n..#..#..###...###\n##.##.##..##...##\n#..##..##.#....##\n#.....###.#.##..#\n...............##\n########.#......#\n\n..#.##....##.#.\n###.########.##\n#...##....##...\n#####..##..####\n.#..########.##\n..##.##..##.##.\n.###.######.###\n#.##...##...##.\n#.##...##...##.\n\n.###.##\n####..#\n####..#\n.###.##\n..#..#.\n##.....\n##..#..\n..#..#.\n.###.##\n\n#..####\n#.#####\n...#...\n..##...\n##....#\n##....#\n..##...\n...#...\n#.#####\n#..####\n.....##\n\n#....#.##\n...##.#.#\n###..#...\n#....#.##\n######...\n..##.##..\n..#...#.#\n..#...#.#\n..##.##..\n######...\n#....#.##\n###..#...\n...##.#.#\n#...##.##\n#...##.##\n\n##.###..#\n##.##.##.\n..##..#..\n..#..##.#\n..#..##.#\n..##..#..\n##.##.##.\n##.###..#\n......##.\n.#######.\n##.#.###.\n\n.#.....\n..#####\n####..#\n...####\n##.#..#\n#.##..#\n.#.#..#\n#..####\n#.##..#\n..##..#\n#..####\n\n....#.#\n#..#..#\n.##..##\n.##..##\n#..#..#\n....#..\n.##.#.#\n#####..\n....##.\n####...\n.....##\n#..##..\n.##...#\n#..#...\n.##...#\n\n..##...#.#.##....\n#.####.###.##..##\n#.####.###.##..##\n..#....#.#.##....\n.#..#..#....#.#.#\n###..#..##.#####.\n##..#..####..#.#.\n..###.#..##.#.##.\n...##.######.....\n..#...##.#.#...#.\n....#...###.#....\n....#...###.#....\n..#...##.#.#...#.\n\n##..##.#.##\n##..##.#.##\n#######..#.\n.####..#...\n#.##.#..#.#\n.#..#.#....\n##..##.#.#.\n.####...#..\n.####.#.###\n..###...###\n######..#.#\n.####.##.#.\n#.##.#.#.##\n.####.#####\n#######..##\n\n#.##...##\n...####.#\n####..###\n###.###.#\n.#.#####.\n.#.#####.\n###.###.#\n####..###\n...####.#\n#.##...##\n#.#.#.###\n#..####..\n##.####..\n\n..###.....##.\n##....##..###\n##....#....##\n####.....#.#.\n...########..\n##...#..##.#.\n...#.#.###..#\n######.##..#.\n##....#...#.#\n...#####.##.#\n#####.#...##.\n..#.######...\n####.##...##.\n###.##..#.#..\n###.##.##.#..\n\n##.#...########\n.#..##.#.......\n.#.#.....#.##.#\n.##..#.###.##.#\n#.#####.##.##.#\n.##.###.#......\n.##.###.#......\n#.#####.##.##.#\n.##..#.###.##.#\n.#.#.....#.##.#\n.#..##.#.......\n##.#...########\n###.....###..##\n.#.#.#.......#.\n####.###..#..#.\n##.##.##..####.\n#.#..#..#..##..\n\n#.#.....#..#...\n.#####.#.##.#.#\n.###.##.#..#.##\n#.#.###.#..#.##\n#.#####.#..#.##\n.###.##.#..#.##\n.#####.#.##.#.#\n#.#.....#..#...\n...##.##....##.\n\n#......#.\n#.####.##\n..####..#\n#......#.\n#......#.\n..####..#\n#.####.##\n#......#.\n########.\n..#.##..#\n##....##.\n.#....#..\n.##..##..\n\n######.##.###\n..##..#####.#\n.#....##.#.##\n.####..##..#.\n#######.#.###\n#######.#.###\n.####..##..#.\n\n.##.#..###..####.\n.##.#.####..####.\n#..#..#..###.#...\n#..#####.###.##.#\n#..###..##.####..\n#####.##...##....\n####.#.#..###..##\n.##.#..#.#.#...#.\n#..#..####.###..#\n.##..##....#####.\n.##.............#\n.....########..##\n#..##.##..#.###.#\n\n####.#..#.###\n...#..##..#..\n..#######.##.\n##.#.#..#.#.#\n...#.#..#.#..\n##.#.#..#.#.#\n#####.##.####\n##...#..#...#\n......##.....\n###.#.##.#.##\n..###....###.\n###...##...##\n....######...\n...#..##..#..\n##.#.####.#.#\n\n.##.#..##.#.#.#..\n.##.#.##.########\n.....#.####.#####\n####....##..###.#\n#..#..##..######.\n#######.###..#..#\n####...#..#.....#\n.....##.#...###..\n.......##.##...#.\n....#....#.#.####\n......#..##.#..#.\n#..###....#...#.#\n#..#...##..##.###\n....#...###....##\n.##....#..#.#...#\n....#....#....#..\n....#....#....#.#\n\n..####..#####\n#########..##\n##.##.####.#.\n########..###\n##....##.####\n...##...#....\n..#..#..#.#.#\n#.####.#.#.##\n##....##.#...\n.........##.#\n..####...#.#.\n...##....#...\n###..####.#..\n\n#..#.#..#.#\n#.###.##.##\n..##.####.#\n.#..##..##.\n#.##..##..#\n.#.########\n..#..#..#..\n..#..#..#..\n.#.########\n#.##..##..#\n....##..##.\n..##.####.#\n#.###.##.##\n#..#.#..#.#\n#.###....##\n\n#...##..##..#\n...##..####..\n.#.##........\n###.#.##..##.\n...##.#.##.#.\n.....#..##..#\n...#.#..##..#\n...##.#.##.#.\n###.#.##..##.\n\n####.######.#\n.#...........\n#.#.###..###.\n.........#...\n#............\n#..#.######.#\n#..#.######.#\n\n####.###.##\n###..#...##\n###..#...#.\n####.###.##\n#..#....#..\n###...##.#.\n.##.##..#..\n#.##.#.##.#\n#..##.##.##\n..#.###..##\n......##.#.\n..#.#.###.#\n#.##.####.#\n..#.#######\n..........#\n..........#\n..#.#######\n\n..#.#..\n##..#.#\n#...##.\n..#.#..\n...#.#.\n....###\n#####..\n#####..\n....###\n\n..#..#..#.#.##.#.\n..##.#..#.#.##.#.\n..######.#......#\n#....##.####..###\n.#....#.##......#\n..#.##.....#..#..\n.#.###..#.######.\n\n#####.#......\n.#.#.....#...\n#..##.#...#..\n#.#....#.....\n#....##..#.#.\n##.#.####..##\n#..#.#.#.....\n#.####..##...\n.#....##.##..\n..##...#.#.##\n##..####.#.##\n..#.#...#.#..\n..#.#...#.#..\n\n#.##...##..\n...#.......\n.###.#....#\n#.#.#.#..#.\n#.##.#....#\n##..##.##.#\n##..##.##.#\n#.##.#....#\n#.#.#.#..#.\n.###.#....#\n...#.......\n#.##...##..\n..#..#....#\n##...#..#.#\n.###..####.\n##..#.#..#.\n##.#.#....#\n\n..##..###\n######.##\n##...##..\n#.##.##..\n##..###..\n.#..#....\n..##..#..\n\n.###.###..##.\n.###.###..##.\n#.##..#...###\n##..#..#...#.\n#..##.##..##.\n.......#..##.\n##.##.##..#.#\n.#.##..###.##\n.#.##..###.##\n##.##.##..#.#\n.......#.###.\n#..##.##..##.\n##..#..#...#.\n#.##..#...###\n.###.###..##.\n\n.##..#..#..\n###.##..##.\n###.##..##.\n.##........\n##.#..##..#\n..##..##..#\n##.#..##..#\n\n....##..#.#.#.#.#\n....##..#.#.#.#.#\n######........###\n####......###....\n.##.##.###..#####\n.......#.#..#..#.\n####...##.#.##.#.\n#####.#..##.##..#\n#..#..##.....####\n....#..#...###.##\n....#....########\n.###.#...#####..#\n.##.#.##.#.....##\n\n.#.###.#.##.#\n........#...#\n.#.##......##\n.#.##......##\n........#...#\n.#.###...##.#\n#.#.#.#......\n####....#.#.#\n.#.....####.#\n.#.#....####.\n.##.#.#...##.\n.#.##.##.....\n..#...###..#.\n###...#..##..\n###...#..##..\n\n###..#...\n.#.##....\n##.##....\n###..#...\n#...##.##\n###......\n#....#.##\n.########\n...####..\n\n...........#..#.#\n#..####..##.#..#.\n#..####..##.#....\n...........#..#.#\n..#....#..#..##..\n#.#....#.##.#..##\n..........#.#....\n#.#....#.#..#.#..\n.########.#.#.##.\n..#.##.#...#...#.\n..........#..#.##\n#..####..#..#.##.\n#.#....#.#.##.#..\n##.#..#.#####.##.\n#..####..#.##..#.\n\n#.#.#.#..#.\n###..######\n##.....##..\n#.##.##..##\n.##..######\n..#.#.#..#.\n..#.#.#..#.\n.##..######\n####.##..##\n##.....##..\n###..######\n#.#.#.#..#.\n..##..####.\n\n#..#.#.#.\n#..#.#.#.\n..#.#..#.\n.#..#....\n..#######\n..#######\n.#..#.#..\n..#.#..#.\n#..#.#.#.\n\n##..###...#\n#....##.#..\n..##..#####\n..##..#.#.#\n#.##.####..\n######....#\n.#..#...#.#\n#....#.##.#\n.#..#....##\n.#..#....##\n#....#..#.#\n\n...##.##.\n..#..##.#\n#.###.###\n#.###.#.#\n.#.#...##\n##...#...\n##..##.##\n##..##.##\n##...#...\n.#.#...##\n#.###.#.#\n#.###.###\n..#..##.#\n....#.##.\n....#.##.\n\n#....##.#\n#....####\n##..##.##\n#.##.#..#\n#....##..\n.#..#..##\n#.##.#.##\n##...####\n#.##.#.##\n#.##.#.##\n##...####\n#.##.#.##\n.#..#..##\n#....##..\n#.##.#..#\n##..##.##\n#....####\n\n#..#....##.##.#\n....##..#..#.#.\n....##..#..#.#.\n#..#....##.##.#\n#.######.#.##.#\n#......##.#..##\n....###...#####\n##...#.#...##..\n.#.......##..##\n.#.......##..##\n##..##.#...##..\n....###...#####\n#......##.#..##\n#.######.#.##.#\n#..#....##.##.#\n\n#...##...#..#\n#.####.#.####\n..#...#..#..#\n#...#.#.##..#\n#..#..###.##.\n...#....#....\n.###..#.##..#\n..##.#.......\n..##.#.......\n.###..#..#..#\n...#....#....\n\n##..##.#.#.\n.#.#.#.###.\n####.##.#.#\n#######.#.#\n.#.#.#.###.\n##..##.#.#.\n#.....#.#.#\n#.....#.#.#\n##..##.#.#.\n.#.#.#.###.\n#######.#.#\n\n......#\n.######\n.#..#..\n.#..#.#\n.####.#\n......#\n##..##.\n#....#.\n#.##.##\n#.##.##\n#....#.\n##..##.\n......#\n\n#.#.######.\n.#.##....##\n####......#\n...#.#..#.#\n##...#..#..\n.#..######.\n#.#..#..#..\n..#....#...\n.##..####..\n#.##.#..#.#\n#.###....##\n..#........\n..#........\n\n.####.####.##\n#.##.#....#.#\n.####......##\n.####..##..##\n#....#....#..\n..##...##...#\n.......##....\n.#..#..#...#.\n..##........#\n.#..#..##..#.\n#....##..##..\n..##..####..#\n.####.#..#.##\n......#..#...\n##..###..###.\n\n.#..##.#.#.#..#.#\n##.#.....#......#\n#.#..#.##.##..##.\n....#...###.##.##\n.......#.##.##.##\n.......#.##.##.##\n....##..###.##.##\n#.#..#.##.##..##.\n##.#.....#......#\n.#..##.#.#.#..#.#\n#...###...#.##.#.\n.#.####..#.#..#.#\n######...##.##.##\n\n####....#####.#\n####....#####.#\n#.###..###.###.\n#..........##..\n#..#.##....#..#\n####.##.#######\n#..#.##.#..##..\n.##......##.#..\n###..##..###..#\n\n......##.##\n####.##.###\n#..#####...\n.##.#.#.##.\n.##.#.#.##.\n#..#####...\n####.##.###\n.....###.##\n#..#...##.#\n.##..##.#..\n####....#..\n\n#..#..#.###.####.\n###.##...##.....#\n##.#........#..#.\n.....##..###.##.#\n..#...#.##..#..#.\n.#....#######..##\n......##.....##..\n#..###.##########\n.#..####.###....#\n..#..##.#####..##\n.##.#.###.#.####.\n#.##..#..#.......\n.....##..##.#..#.\n.###...###..#..#.\n.###...###..#..#.\n\n.####..##..####\n...###....###..\n#.#.##.##.##.#.\n....#.#..#.#...\n#....######....\n#......##......\n....#.#..#.#...\n\n......####..###\n.#..#.#####.###\n#.##.#.........\n#.##.###.#..#.#\n.####..###..###\n######..######.\n..##......##...\n##########..###\n.......#......#\n.#..#..##....##\n......##.####.#\n##..###........\n......###....##\n.####....####..\n#.##.#..#....#.\n\n#####.#\n#.##..#\n##.#.#.\n###..#.\n...#.#.\n###.###\n..#####\n..#####\n###.###\n\n##.###.\n###...#\n##.....\n..#...#\n##.##.#\n..#.###\n..#..##\n......#\n......#\n..#..##\n..#.###\n##..#.#\n..#...#\n##.....\n###...#\n##.###.\n...###.\n\n..##.###.#...##\n#.##.###.....##\n#.##.###.....##\n...#.###.#...##\n.#....#...#.#..\n..####..#####..\n...######.##.##\n.#.#..#.#.#####\n.#...#..#......\n.#...##..######\n.#..#.####..#..\n\n#.##.#..#\n#.##.#..#\n..#.#####\n##.##.##.\n#.#.##.##\n.###.....\n...##....\n\n....#..#....###\n##.#...##.#.###\n###...#....#.##\n..##.#...######\n######.####.###\n.........##....\n#.##..######...\n\n....##.\n#####.#\n....#..\n....##.\n#..##..\n.....#.\n#####.#\n#####.#\n.....#.\n#..##..\n....###\n....#..\n#####.#\n\n.##..#..#\n#.###.###\n.#..##..#\n##.####..\n.#..##...\n......#..\n......#..\n.#..##...\n##.####..\n.##.##..#\n#.###.###\n.##..#..#\n....#..##\n...###...\n.#..###.#\n.#..###.#\n...###...\n\n.#....#.##...\n#......#.#.##\n#.####.###...\n##########...\n.##..##......\n.#....#...###\n.##..##.##.##\n.##.###..##..\n########.####\n..####..##.##\n...##...#..##\n.##..##...#..\n#.####.#..###\n##....###.###\n.#....#.###..\n\n#..#####.#..#\n###.##.######\n#........####\n.#.#..#.#....\n..........##.\n#........#..#\n##..##..#####\n#..#..#..#..#\n#.######.#..#\n\n.#..#.####.#..#\n...#...##...#..\n#...##....##...\n##..##.##.##..#\n##..##.##.##..#\n#...##....##...\n...#...##...#..\n.#..#.####.#..#\n##.#........#.#\n.##.#.####.#.##\n###..........##\n.##.########.##\n.#..##.##.##...\n...#.#.##.#.#..\n.....#.##.#....\n###.#......#.##\n#.###.#..#.###.\n\n#..#.#.##\n....##.#.\n....##.#.\n#..#...##\n.##.#..##\n.....#.##\n.##...#.#\n#..#....#\n....####.\n\n#####..##.#\n#.#.#.#.#..\n###.##.##..\n##.##.#...#\n...#..#.##.\n...#..#.##.\n##.##.....#\n###.##.##..\n#.#.#.#.#..\n#####..##.#\n####..#.#.#\n####..#.#.#\n#####..##.#\n#.#.#.#.#..\n###.##.##..\n\n##.##.######.\n.######.##.##\n##.##.######.\n..#..#......#\n.#..#.#.##.#.\n#.#..#.#..#.#\n###..########\n\n.#.#.#..##..#.#.#\n.###...#..#...###\n#..#.#..##..#.#..\n.#.#...####...#.#\n#.#....#.......#.\n#................\n...#.#......#.#..\n..#.####..####.#.\n...#...####...#..\n#.#.##..##..##.#.\n.....#.####.#....\n#.##.#.####.#.##.\n.###..#....#..###\n.##..########..##\n.##..########..##\n.###..#....#..###\n#.##.#.####.#.##.\n\n#..#.#..#.#\n..###.##..#\n#..#......#\n###.######.\n#.##.####.#\n#.....##...\n#...#....#.\n..#..#..#..\n#.#........\n.###..##..#\n.###..##..#\n#.#........\n..#..#..#..\n#...#....#.\n#.....##...\n\n..####..##.\n#.##.......\n###.....##.\n.###.#..##.\n#...#......\n#.###..####\n##..#.#####\n.#..#.#####\n#.###..####\n\n......#####\n####.##.##.\n####.#.....\n######...#.\n....###.###\n.##.##.####\n.##..#.####\n....###.###\n######...#.\n####.#.....\n####.##.##.\n......#####\n#..#.##.###\n####.#..#..\n#..####.###\n\n..##.##.##..##..#\n..#.####.#..#.###\n#####..#####..###\n.#.##..##.#...#.#\n.#.##..##.#...#.#\n#####..#####..###\n..#.####.#..#.###\n..##.##.##..##..#\n##..####..##.#..#\n#...#..#...#..#..\n##..#..##.##.#...\n\n...##..##...#..\n##.####.##.#.##\n...#.##.#####..\n..#.####..#..##\n..#########..#.\n..#########..#.\n..#.####..#..#.\n...#.##.#####..\n##.####.##.#.##\n...##..##...#..\n########...###.\n\n####.#####.\n...###....#\n...#.....##\n...#.....##\n...###....#\n####.#####.\n#....##.##.\n.....#.#..#\n.##..#.....\n.##..##....\n.....#.#..#\n#....##.##.\n####.#####.\n...###....#\n...#.....##\n\n####....##.####\n.....##.#..###.\n##.##.#####....\n##..###......##\n#.#..####.#.#.#\n##..#.####.#...\n##..#.####.#...\n#.#..####.#.#.#\n##..###......##\n##.##.#####....\n.....##.#..###.\n####....##.####\n.#..####.##.###\n#.#.####..###.#\n#.######..###.#\n\n.#.#..#.#.#.##.#.\n.#.#..#.#.####.#.\n...#..#.....###.#\n####..######.##..\n..........#.#..#.\n..##..##..#...##.\n.###..###.##.#..#\n...####...#.#..#.\n.#.####.#.###.#..\n#.#.##.#.###.#.#.\n.#.#..#.#.#..##..\n\n###..###.\n###...##.\n..##.##.#\n.#..####.\n#..##.#..\n.#.##.#..\n.....#.#.\n.....#.##\n#.#.##.##\n#.#.##.##\n.....#.##\n.....#.#.\n.#.##.#..\n#..##.#..\n.#..####.\n\n..#............#.\n..##.###..###.##.\n..##.##.##.##.##.\n##.#####..#####.#\n..##.##....##.##.\n.....#..##..#....\n###.#.##..##.#.##\n##.#.###..###.#.#\n...#..##..##..#..\n..#.#........#.#.\n#######.#########\n\n###.....#.#.##.\n..#..##.#.##..#\n####.....###.##\n..#.....#...#.#\n..#.#...#.###.#\n##.#..####.#..#\n..#.#.#.###..#.\n##.##..##..#.#.\n##.##.###..#.#.\n\n....#..#.\n####....#\n.##......\n#.##.##.#\n#..##..##\n#..##..##\n#####..##\n\n.#.#.##\n####...\n###....\n..#.#..\n..#....\n###....\n####...\n.#.#.##\n#.##...\n.#.##..\n..#..##\n#....##\n...#.##\n.#...##\n.###...\n\n###..###..#.#..\n.######..##.###\n##########...##\n..####..##.#...\n#.#..#.###.###.\n.##..#..#...##.\n##.##.##.#..#..\n##.##.###..##..\n###..####.#.#.#\n###..####.#.#.#\n##.##.###..##..\n\n#.####.##\n.##..##.#\n#####.###\n..#..#...\n##....##.\n.######.#\n#..##..#.\n#.####.##\n..#..#..#\n...##....\n..####...\n..#..#..#\n..#..#..#\n..####...\n...##....\n..#..#..#\n#.####.##\n\n.#.##..#.#.....\n..#.###.#####..\n..#...#.#.##...\n##.##..#.#..###\n...#..##.###.#.\n..##.#######.##\n..##.#######.##\n...#..##.###.#.\n##.##..#.#..###\n..#...#.#.##...\n..#.###.#####..\n\n#.#########\n.#.#.#..#.#\n#.###.....#\n##..#.##.#.\n.#..#.##.#.\n.###.#..#.#\n#...#.##.#.\n#...##..##.\n.#....##...\n#.#.#....#.\n#.#.#....#.\n\n#..#..#.#\n#..#..#.#\n..#....#.\n##..##.#.\n#.#.##.##\n.##...##.\n##.###...\n#.####...\n######...\n##.###...\n.##...##.\n#.#.##.##\n##..##.#.\n..#....#.\n#..#..#.#\n\n#....#..#\n#.##.#...\n##..##...\n##..##...\n#.##.#..#\n#....#..#\n.####..#.\n.#..#...#\n##..#####\n\n#...##...#....##.\n..#.##.#..###....\n..##..##..##.####\n...####.....##..#\n#.######.#....###\n###.##.####..#..#\n##..##..##.##.##.\n#........##.#####\n...........#.#..#\n\n.#.#..#..#.##\n.##.#...#####\n.##.#...#####\n.#.#..#..####\n#..#.....#...\n.####.##.#...\n..#..###.#...\n#.##.##......\n########.####\n#####..##.###\n#....#..#..##\n..####...#.##\n..#.#.##.#.##\n...#.#.#.....\n#.#..#.##....\n##.#.#.#..#..\n#.#####.#.#..\n\n##.##.##.##.#\n##.###..###.#\n.####.##.####\n#####.##.####\n##.##.##.#..#\n#.#........#.\n#.#........#.\n\n####.#.#.#..#\n.##.#.##.#..#\n.##...###.##.\n.##..##.##..#\n.....###.#..#\n....#.#..####\n######..#.##.\n......##.....\n####.#.#.#..#\n.##..##..####\n.##....#.....\n#####..#.....\n#...##..#....\n.##..########\n.##.#...#####\n#..######....\n.....##..#..#\n\n###....##\n..#.##.#.\n..#..#.#.\n....##...\n##......#\n...#..#..\n###....##\n..######.\n..##..##.\n\n##..##.######\n.####.#######\n.####.##.##.#\n.#..#..#.##.#\n######.#....#\n......#..##..\n######.##..##\n.#..#...###..\n#....#.......\n#######.#..#.\n..##....#..#.\n#....###....#\n########.##.#\n\n#..#.##.#..\n#..#.##.#..\n#..#.#.##..\n#.#####.##.\n####...####\n#..######.#\n.##.###..##\n#####.#.##.\n....###..#.\n\n#..###.#.\n...#...##\n##...##..\n###.###..\n##...#.##\n...###.##\n...#.#.##\n##...#.##\n###.###..\n##...##..\n...#...##\n#..###.#.\n#..###.#.\n\n####.####..##\n#.......#...#\n#.#......##.#\n#..##...#.###\n#..##...#.###\n#.#......##.#\n#.......#...#\n####.####..##\n#.##..#######\n....#..##.##.\n..#...#...###\n##.####.#####\n##..###.#####\n\n###.#.#\n##...##\n#..##.#\n##.#.##\n..##...\n..#.##.\n#####..\n#####..\n..#.##."
	return []int{d13(0, input), d13(1, input)}
}

func d13(smudges int, input string) int {
	patterns := strings.Split(input, "\n\n")
	total := 0
	for _, p := range patterns {
		pattern := strings.Split(p, "\n")

		if colNum := getReflectColNum(smudges, pattern); colNum != -1 {
			total += (colNum + 1) * 1
		} else {
			total += (getReflectRowNum(smudges, pattern) + 1) * 100
		}

	}
	return total
}

func getReflectColNum(smudges int, pattern []string) int {
	for col := 0; col < len(pattern[0])-1; col++ {
		if reflectsAfterCol(pattern, col, smudges) {
			return col
		}
	}
	return -1
}

func getReflectRowNum(smudges int, pattern []string) int {
	for row := 0; row < len(pattern)-1; row++ {
		if reflectsAfterRow(pattern, row, smudges) {
			return row
		}
	}
	return -1
}

func reflectsAfterCol(pattern []string, col int, smudges int) bool {
	smudgesFound := 0
	for a := 0; a <= col; a++ {
		b := col + 1 + col - a
		if b >= len(pattern[0]) {
			continue
		}
		for y := range pattern {
			if pattern[y][a] != pattern[y][b] {
				if smudgesFound < smudges {
					smudgesFound++
					continue
				}
				return false
			}
		}
	}
	return smudges == smudgesFound
}

func reflectsAfterRow(pattern []string, row int, smudges int) bool {
	smudgesFound := 0
	for a := 0; a <= row; a++ {
		b := row + 1 + row - a
		if b >= len(pattern) {
			continue
		}
		for x := range pattern[0] {
			if pattern[a][x] != pattern[b][x] {
				if smudgesFound < smudges {
					smudgesFound++
					continue
				}
				return false
			}
		}
	}
	return smudges == smudgesFound
}
