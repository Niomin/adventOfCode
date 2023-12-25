package main

import (
	"strconv"
	"strings"
)

func D18() []int {
	data := []string{
		"R 6 (#70c710)\nD 5 (#0dc571)\nL 2 (#5713f0)\nD 2 (#d2c081)\nR 2 (#59c680)\nD 2 (#411b91)\nL 5 (#8ceee2)\nU 2 (#caa173)\nL 1 (#1b58a2)\nU 2 (#caa171)\nR 2 (#7807d2)\nU 3 (#a77fa3)\nL 2 (#015232)\nU 2 (#7a21e3)",
		"L 7 (#2ac8e2)\nD 5 (#14b771)\nL 14 (#302602)\nU 5 (#14b773)\nL 4 (#55b782)\nD 6 (#535223)\nL 3 (#3bf082)\nD 8 (#7be433)\nL 9 (#4c1060)\nD 12 (#3a69b3)\nL 4 (#5f1760)\nD 6 (#6832c3)\nL 6 (#2f0d60)\nU 3 (#0f6c51)\nL 5 (#2a1042)\nU 7 (#464ee1)\nR 5 (#2a1040)\nU 5 (#4ce141)\nL 2 (#1261c0)\nU 11 (#1a05f3)\nL 3 (#22a220)\nU 6 (#6f3d13)\nL 6 (#677110)\nD 6 (#38ccd3)\nL 6 (#1a3370)\nD 3 (#4f3b43)\nL 12 (#0086e0)\nD 3 (#68f4b1)\nL 4 (#5c97f0)\nD 5 (#4bee81)\nR 8 (#556760)\nU 5 (#447f53)\nR 3 (#2a7580)\nD 5 (#7063e3)\nR 7 (#2c84b0)\nD 9 (#193141)\nL 7 (#1ae260)\nD 10 (#4e32c1)\nL 3 (#4f4a70)\nU 10 (#356251)\nL 8 (#7ea0b2)\nD 4 (#580771)\nR 4 (#7ea0b0)\nD 10 (#002e51)\nL 12 (#0f93b0)\nD 5 (#024911)\nL 9 (#16ca20)\nU 2 (#14e323)\nL 3 (#77c5d0)\nU 7 (#4db283)\nL 8 (#7b4722)\nU 5 (#1f8db3)\nR 8 (#0a5fd2)\nU 3 (#3dd443)\nL 11 (#5ad430)\nU 11 (#33f903)\nR 5 (#5ad432)\nD 7 (#1f0493)\nR 14 (#4e8c52)\nU 7 (#28bcc3)\nR 3 (#4f1792)\nU 5 (#31ad71)\nL 11 (#196172)\nU 6 (#6707d1)\nL 11 (#556db2)\nU 3 (#6707d3)\nL 6 (#457dc2)\nU 3 (#31ad73)\nL 8 (#01d4d2)\nU 4 (#525cb3)\nL 6 (#225382)\nU 12 (#0fc663)\nL 4 (#701292)\nU 5 (#0fc661)\nL 7 (#0398a2)\nU 8 (#525cb1)\nL 2 (#46ac42)\nU 5 (#6c89c3)\nL 8 (#320b70)\nD 3 (#319c33)\nL 5 (#7b0110)\nD 11 (#319c31)\nR 5 (#4286b0)\nD 4 (#2db113)\nL 8 (#806100)\nU 14 (#5c7523)\nL 4 (#1a96c0)\nU 8 (#40b8d3)\nR 8 (#5bab20)\nU 6 (#43c6f3)\nL 8 (#3a4b90)\nU 5 (#3e3133)\nL 3 (#679780)\nU 9 (#584113)\nL 8 (#559c02)\nU 7 (#60bf53)\nL 4 (#5aca62)\nU 8 (#4d0163)\nR 9 (#2237e2)\nD 5 (#106f13)\nR 12 (#5c2362)\nD 8 (#80b7d3)\nR 7 (#26b6b2)\nD 6 (#08d3b3)\nR 14 (#59afa2)\nU 4 (#241333)\nR 9 (#0c0eb2)\nU 3 (#28a8c3)\nR 4 (#47b022)\nU 12 (#279703)\nR 7 (#47b020)\nU 5 (#4e01a3)\nL 5 (#6c4722)\nU 5 (#607023)\nR 5 (#609b62)\nU 5 (#6ec7a3)\nL 9 (#152722)\nU 6 (#022cf1)\nL 14 (#481a02)\nU 5 (#7ce951)\nL 5 (#4316f2)\nU 6 (#0f1061)\nL 2 (#2b14e2)\nU 3 (#3e10c1)\nL 12 (#435d72)\nU 2 (#6d75e3)\nL 8 (#4a4142)\nU 8 (#5ec183)\nL 12 (#5a9f02)\nU 8 (#390f33)\nR 8 (#655e52)\nU 11 (#757e23)\nR 10 (#031912)\nU 7 (#10ea33)\nR 14 (#4a55c0)\nD 6 (#1d2fc3)\nR 9 (#7aa260)\nD 12 (#07c993)\nR 9 (#315c00)\nU 8 (#88b083)\nR 5 (#52fb50)\nU 5 (#2b4ed3)\nL 9 (#428cf0)\nU 6 (#0225f3)\nR 9 (#36f490)\nU 6 (#0225f1)\nR 7 (#595110)\nD 4 (#374603)\nL 4 (#44efe0)\nD 6 (#2b1b03)\nR 4 (#664590)\nD 7 (#2b1b01)\nR 5 (#1bfe40)\nU 6 (#2c7a81)\nR 2 (#3bc8f0)\nU 9 (#416101)\nR 10 (#003f10)\nD 7 (#3a4a51)\nR 2 (#003f12)\nD 5 (#4ba511)\nR 3 (#5190e0)\nD 3 (#400191)\nR 12 (#1a09a0)\nD 8 (#3aef23)\nR 11 (#0cb780)\nD 10 (#6a6393)\nR 9 (#561c80)\nU 10 (#124da3)\nR 6 (#2b0af0)\nU 8 (#563841)\nR 12 (#73b520)\nU 3 (#199441)\nR 5 (#73b522)\nU 4 (#47d3d1)\nR 4 (#70a9a0)\nU 8 (#18f861)\nL 4 (#18de92)\nU 3 (#5bc8a1)\nL 12 (#59e082)\nU 2 (#4aacb1)\nR 12 (#0b5a82)\nU 6 (#708aa1)\nL 5 (#3902b2)\nU 5 (#171153)\nR 5 (#3745f2)\nU 5 (#081ad3)\nR 5 (#45f1c0)\nD 8 (#527f53)\nR 6 (#45f1c2)\nD 6 (#47e013)\nL 6 (#3745f0)\nD 4 (#01abd3)\nR 8 (#51f1e2)\nD 8 (#6aa421)\nL 5 (#6114f2)\nD 7 (#786041)\nR 5 (#31b050)\nD 3 (#4f62d1)\nR 10 (#5ba7a0)\nU 8 (#14c261)\nR 4 (#2d27d0)\nU 7 (#2cd791)\nR 4 (#2db660)\nU 5 (#7a4801)\nR 9 (#19e300)\nU 5 (#17adf1)\nL 10 (#19e302)\nU 10 (#63c5a1)\nL 2 (#6881e0)\nU 6 (#2ba711)\nR 12 (#148ef0)\nU 6 (#1a9ba3)\nR 9 (#0597f0)\nD 8 (#4f1ee3)\nR 10 (#4dd080)\nD 4 (#33a753)\nL 4 (#3fb6d0)\nD 3 (#1ebdb3)\nR 10 (#0b0872)\nD 8 (#0e1853)\nL 10 (#3547f2)\nD 5 (#21e0e3)\nR 4 (#52cee2)\nD 5 (#3b66c3)\nL 10 (#5b91a0)\nD 6 (#26bac3)\nR 11 (#222710)\nD 2 (#2d5211)\nR 7 (#44f020)\nU 12 (#505ed1)\nR 3 (#358270)\nD 12 (#505ed3)\nR 5 (#16c910)\nU 7 (#36a931)\nR 12 (#849ea0)\nU 6 (#36a933)\nR 10 (#1e8c20)\nU 10 (#55b611)\nL 11 (#2cded0)\nD 5 (#5f5281)\nL 12 (#00e5b2)\nU 5 (#4a2981)\nL 3 (#379182)\nU 8 (#3e19a1)\nR 6 (#644222)\nU 9 (#0f4fe1)\nR 6 (#570b02)\nU 7 (#551e33)\nR 7 (#07ea32)\nD 5 (#35bb93)\nR 4 (#0a12d2)\nD 4 (#0cb943)\nL 4 (#3b7242)\nD 7 (#28d521)\nR 7 (#2d9d60)\nU 10 (#7332e1)\nR 3 (#44d090)\nD 4 (#7baf51)\nR 7 (#01b8e0)\nU 4 (#1b4e01)\nR 6 (#0b27b2)\nU 3 (#642a61)\nL 4 (#0b27b0)\nU 6 (#21beb1)\nR 4 (#64df80)\nU 7 (#08c4b3)\nR 10 (#626c00)\nD 6 (#34e6b3)\nR 5 (#06b262)\nU 4 (#6e38c3)\nR 5 (#06b260)\nD 4 (#710243)\nR 5 (#05c140)\nD 4 (#1e5571)\nL 15 (#54bee0)\nD 6 (#0a4df1)\nR 6 (#1f4bd0)\nU 3 (#0a4df3)\nR 13 (#4057c0)\nU 11 (#701d41)\nR 6 (#16f352)\nU 4 (#449eb3)\nL 15 (#69f132)\nU 6 (#449eb1)\nR 15 (#337df2)\nU 6 (#2a2871)\nL 6 (#3946a0)\nU 6 (#112121)\nR 15 (#107a70)\nD 5 (#66a2e1)\nR 2 (#4b4f50)\nD 12 (#277781)\nR 8 (#32cc40)\nD 3 (#1184e1)\nR 5 (#439160)\nD 8 (#43e2d1)\nR 3 (#496410)\nD 2 (#0fb6d1)\nR 9 (#12a250)\nD 12 (#2abc23)\nR 4 (#123e22)\nD 15 (#834853)\nR 4 (#123e20)\nU 4 (#453473)\nR 8 (#0281c0)\nU 12 (#27c921)\nR 7 (#543410)\nU 3 (#2cd3a3)\nR 5 (#3bdd30)\nU 3 (#73d693)\nR 13 (#2bfb40)\nU 6 (#3f93d1)\nL 13 (#1f66f0)\nU 8 (#3e3911)\nR 6 (#1f66f2)\nU 10 (#137901)\nL 8 (#3ea9b0)\nU 3 (#0f6451)\nL 10 (#2ea050)\nU 8 (#450f13)\nR 14 (#1758c2)\nU 4 (#21cd73)\nR 2 (#1758c0)\nU 9 (#39b8e3)\nR 14 (#0255c0)\nD 4 (#290d23)\nR 4 (#0ad710)\nD 9 (#0a59f3)\nR 8 (#3b9950)\nD 4 (#7bae53)\nR 11 (#3e54a0)\nD 8 (#753163)\nR 9 (#4dc1b0)\nD 6 (#2c5541)\nL 2 (#564600)\nD 13 (#292593)\nL 8 (#5b16b0)\nD 4 (#292591)\nL 4 (#21f3f0)\nU 4 (#2c5543)\nL 8 (#0bafd0)\nD 4 (#4cda43)\nL 10 (#6241a0)\nD 5 (#5e31a3)\nR 10 (#1c5f90)\nD 3 (#4e18a1)\nR 7 (#2445d0)\nD 9 (#24fc71)\nR 6 (#299cc0)\nD 6 (#37a491)\nR 7 (#299cc2)\nU 6 (#43e8d1)\nR 10 (#2445d2)\nU 9 (#410ad1)\nR 10 (#67e350)\nU 8 (#5d9801)\nR 9 (#12bd80)\nD 8 (#30b691)\nR 15 (#2ebe00)\nD 2 (#6062d1)\nR 5 (#113860)\nD 13 (#55fa81)\nR 8 (#657e80)\nU 11 (#6be1c1)\nR 9 (#3d83c0)\nU 2 (#83eee3)\nR 11 (#0755a0)\nU 10 (#389d53)\nR 6 (#669170)\nU 10 (#246853)\nR 13 (#26ef30)\nD 10 (#362743)\nR 6 (#0a7380)\nD 8 (#0b2353)\nR 8 (#370cd0)\nD 10 (#2d3c43)\nR 14 (#72d862)\nD 8 (#5dc733)\nL 5 (#72d860)\nD 8 (#034b23)\nL 6 (#1d5e70)\nD 7 (#163871)\nL 4 (#085d00)\nD 9 (#5c3901)\nR 3 (#5c2d00)\nD 8 (#1eead1)\nR 12 (#189f80)\nD 4 (#12f0c1)\nR 5 (#612270)\nD 3 (#24db21)\nR 9 (#120470)\nD 12 (#645371)\nR 3 (#4c8030)\nD 11 (#28ed73)\nR 3 (#491000)\nD 10 (#3e0563)\nR 8 (#181d10)\nD 6 (#139513)\nR 6 (#171f80)\nD 3 (#467203)\nR 7 (#171f82)\nD 12 (#3f62a3)\nR 5 (#181d12)\nD 4 (#2d1f13)\nL 9 (#333730)\nD 2 (#2eef11)\nL 10 (#84f660)\nD 4 (#386021)\nL 3 (#491a90)\nD 7 (#849781)\nR 6 (#491a92)\nD 3 (#2e2731)\nL 6 (#3403a0)\nD 6 (#588431)\nL 10 (#6dde22)\nD 11 (#4b8081)\nL 4 (#0770b2)\nD 4 (#525a71)\nR 7 (#106a92)\nD 10 (#12b753)\nR 4 (#06ce12)\nU 10 (#2da5b3)\nR 5 (#7eb4d2)\nD 6 (#2da5b1)\nR 3 (#21deb2)\nD 8 (#12b751)\nR 12 (#333f42)\nD 12 (#4df5d1)\nL 7 (#1e5502)\nD 3 (#3b11d1)\nL 8 (#2add42)\nD 5 (#2a2461)\nL 6 (#2add40)\nD 7 (#3332c1)\nL 3 (#0e2cd2)\nU 7 (#756e81)\nL 7 (#4756b2)\nD 5 (#4d0ca1)\nL 3 (#56ef42)\nD 11 (#3d05a1)\nL 4 (#46a2c2)\nU 6 (#5bdcb3)\nL 5 (#2e3fe2)\nU 11 (#781b83)\nR 5 (#4f8f62)\nU 3 (#21dae3)\nL 3 (#3e9a92)\nU 2 (#1add33)\nL 9 (#3732e2)\nD 12 (#07f113)\nL 2 (#0c88b2)\nD 10 (#853513)\nL 8 (#0c88b0)\nD 8 (#091c53)\nL 4 (#326e32)\nD 3 (#2a2593)\nR 4 (#642632)\nD 13 (#6a4c83)\nL 4 (#6a1472)\nD 6 (#082b33)\nL 3 (#47d362)\nD 8 (#48ba61)\nL 7 (#300792)\nD 5 (#6148c1)\nL 8 (#0100d0)\nU 11 (#20efd3)\nL 6 (#377920)\nD 11 (#7895f3)\nL 4 (#3a8d50)\nD 3 (#7895f1)\nL 12 (#26a8b0)\nD 7 (#20efd1)\nL 7 (#0e6f20)\nD 4 (#6390c1)\nL 9 (#03cd40)\nU 8 (#0ffa71)\nL 3 (#151ed2)\nU 12 (#147cd3)\nL 7 (#843392)\nU 6 (#147cd1)\nR 10 (#1299f2)\nU 11 (#0736e1)\nR 11 (#654672)\nU 8 (#2465b3)\nL 4 (#38ffb0)\nD 4 (#0b0b53)\nL 13 (#2a0df2)\nU 4 (#72d2c3)\nL 4 (#2a0df0)\nU 3 (#0f7b73)\nL 12 (#38ffb2)\nU 2 (#2a4ba3)\nL 11 (#66dd22)\nD 3 (#2577b1)\nL 5 (#6a82b2)\nD 3 (#516653)\nL 7 (#093060)\nD 12 (#79ad63)\nL 6 (#093062)\nD 5 (#348f83)\nL 4 (#83e322)\nU 8 (#0da4d3)\nL 4 (#251e42)\nD 6 (#356721)\nL 15 (#060892)\nU 6 (#304a61)\nL 3 (#226e72)\nU 5 (#021bb3)\nR 9 (#7123c2)\nU 4 (#021bb1)\nR 13 (#1a4822)\nU 6 (#304a63)\nL 3 (#157e22)\nU 14 (#56fb41)\nR 7 (#4a75b0)\nU 7 (#4578a1)\nR 6 (#7eeb50)\nU 6 (#3b6d01)\nL 11 (#0ba9e2)\nU 5 (#8a0ba1)\nR 11 (#3ac130)\nU 6 (#3faf21)\nR 6 (#6a9062)\nD 7 (#3460e1)\nR 15 (#6a9060)\nD 3 (#421101)\nL 15 (#3838a0)\nD 7 (#58cf13)\nR 6 (#715670)\nD 7 (#6b3d83)\nR 4 (#715672)\nD 10 (#366933)\nR 8 (#591620)\nU 10 (#2f6453)\nR 4 (#8295e0)\nU 10 (#3cb3f1)\nR 9 (#117f00)\nU 10 (#497781)\nR 5 (#1ce3c2)\nU 6 (#14c9a1)\nL 13 (#1ce3c0)\nU 4 (#3060a1)\nL 8 (#2f4880)\nU 5 (#3b2431)\nL 4 (#2f4882)\nU 12 (#236031)\nL 7 (#251ac0)\nD 5 (#0052b1)\nL 4 (#3d5a30)\nD 4 (#372e11)\nL 12 (#58bf50)\nD 10 (#5db081)\nL 2 (#1582a2)\nD 4 (#60b921)\nL 7 (#39fe02)\nD 8 (#657e23)\nL 3 (#291c72)\nD 7 (#657e21)\nL 2 (#393a72)\nD 7 (#3196a1)\nL 9 (#10c4f2)\nD 8 (#650ba1)\nL 10 (#052452)\nD 4 (#0980a1)\nL 13 (#301942)\nD 10 (#18a341)\nL 4 (#60a8f2)\nD 10 (#73c9b1)\nR 6 (#006630)\nD 6 (#0c04c1)\nR 5 (#35f0f0)\nU 6 (#7db103)\nR 7 (#34c280)\nD 9 (#7db101)\nL 6 (#25a890)\nD 3 (#1bf0a1)\nR 12 (#17e9c2)\nD 8 (#018d01)\nL 12 (#7c9bc2)\nD 6 (#5e24b3)\nL 4 (#0c6ce2)\nU 5 (#4fa793)\nL 4 (#5a26b2)\nU 7 (#14b5d3)\nR 4 (#330d12)\nU 5 (#1b3311)\nL 8 (#20c112)\nD 9 (#1c8ed1)\nL 10 (#0f11d2)\nU 11 (#5f2d41)\nR 4 (#76d4f2)\nU 11 (#2b92f1)\nL 4 (#184052)\nU 8 (#63a4f3)\nL 6 (#24fea2)\nU 3 (#4fc613)\nR 6 (#453b70)\nU 7 (#52e7c3)\nL 3 (#307370)\nU 10 (#1edca3)\nL 7 (#75aee2)\nD 3 (#1d4c43)\nL 3 (#6e3de2)\nU 13 (#10edc3)\nL 8 (#293ec2)\nD 13 (#68f773)\nL 8 (#373872)\nD 6 (#6a7d73)\nL 3 (#5f5212)\nD 5 (#649e01)\nR 6 (#53fed2)\nD 6 (#6bc451)\nR 4 (#04d680)\nU 6 (#25e6f1)\nR 8 (#1e60d0)\nD 3 (#456061)\nR 4 (#1e03f0)\nD 6 (#398451)\nL 11 (#648390)\nD 2 (#15d5f1)\nL 7 (#369602)\nD 10 (#5054a1)\nL 7 (#53bcc2)\nU 15 (#5054a3)\nL 5 (#0fd012)\nD 15 (#2915a1)\nL 4 (#0b9c02)\nD 4 (#0f1fd1)\nR 11 (#79edf2)\nD 5 (#63a4f1)\nR 7 (#011ec2)\nD 5 (#3452d1)\nR 8 (#2ccf12)\nU 5 (#790d01)\nR 8 (#3f7272)\nD 8 (#06e261)\nL 7 (#28e202)\nD 15 (#4c0ee1)\nL 3 (#1161b2)\nU 15 (#433e91)\nL 6 (#74d6b0)\nD 5 (#45c0f1)\nL 14 (#74d6b2)\nD 7 (#339011)\nL 3 (#038fa2)\nU 2 (#2a3983)\nL 10 (#06bc22)\nU 3 (#10e1f3)\nL 10 (#27a8d2)\nU 9 (#19d373)\nL 4 (#748b22)\nD 6 (#65bbe3)\nL 9 (#04b122)\nD 4 (#3d31f1)\nL 13 (#010b22)\nU 10 (#05d651)\nL 4 (#74a772)\nD 14 (#05d653)\nL 8 (#14e752)\nU 3 (#4534a1)\nL 9 (#8a99e0)\nU 11 (#0e0ab1)\nR 5 (#421542)\nU 7 (#2a3981)\nR 8 (#6c1e52)\nU 6 (#1d91e1)\nL 13 (#4a4082)\nU 3 (#3e5c83)\nR 6 (#0ebd62)\nU 10 (#21a223)\nL 5 (#153ec2)\nU 6 (#502eb3)\nR 5 (#1299a2)\nU 13 (#461143)\nL 6 (#691cb2)\nU 3 (#1d6533)\nL 4 (#495722)\nU 2 (#3c61e3)\nL 11 (#0cc8b0)\nD 4 (#1fd653)\nL 7 (#0cc8b2)\nD 5 (#340273)\nR 7 (#2c7ec2)\nD 8 (#699123)\nL 4 (#2c7ec0)\nD 2 (#69da23)\nL 6 (#39c2c2)\nU 15 (#18e693)\nL 5 (#3d0762)\nU 4 (#043ac3)\nL 8 (#49ca22)\nU 7 (#057783)\nL 6 (#57e6a2)\nU 12 (#2d7e63)",
	}
	var res []int
	for _, d := range data {
		input := strings.Split(d, "\n")
		res = append(res, d18p1(input))
		res = append(res, d18p2(input))
	}
	return res
}

func d18p2(input []string) int {
	dir := make([]string, 0, len(input))
	dist := make([]int, 0, len(input))

	for _, line := range input {
		raw := strings.Split(line, " ")
		hex := raw[2][2 : len(raw[2])-2]
		meters, _ := strconv.ParseInt(hex, 16, 64)
		d := raw[2][len(raw[2])-2]
		switch d {
		case '0':
			d = 'R'
		case '1':
			d = 'D'
		case '2':
			d = 'L'
		case '3':
			d = 'U'
		}
		dist = append(dist, int(meters))
		dir = append(dir, string(d))
	}

	return d18main(dir, dist)
}

func d18main(dir []string, dist []int) int {
	corners := make([][2]int, 0, len(dir))
	corners = append(corners, [2]int{0, 0})
	length := 0
	var x, y int
	for i, d := range dir {
		meters := dist[i]
		dx, dy := d18getD(d)
		x += dx * meters
		y += dy * meters
		corners = append(corners, [2]int{x, y})
		length += abs(dx*meters + dy*meters)
	}

	s := 0
	for i := 1; i < len(corners); i++ {
		s += corners[i-1][0]*corners[i][1] - corners[i-1][1]*corners[i][0]
	}
	s += corners[len(corners)-1][0]*corners[0][1] - corners[len(corners)-1][1]*corners[0][0]
	s = abs(s / 2)

	return s + length/2 + 1
}

func d18p1(input []string) int {
	dir := make([]string, 0, len(input))
	dist := make([]int, 0, len(input))
	for _, line := range input {
		raw := strings.Split(line, " ")
		meters, _ := strconv.Atoi(raw[1])
		dir = append(dir, raw[0])
		dist = append(dist, meters)
	}

	return d18main(dir, dist)
}

func d18getD(d string) (int, int) {
	if d == "R" {
		return 1, 0
	}
	if d == "L" {
		return -1, 0
	}
	if d == "U" {
		return 0, -1
	}
	if d == "D" {
		return 0, 1
	}
	panic("no way")
}
