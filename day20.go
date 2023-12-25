package main

import (
	"strings"
)

func D20() []int {
	data := []string{
		"broadcaster -> a, b, c\n%a -> b\n%b -> c\n%c -> inv\n&inv -> a",
		"broadcaster -> a\n%a -> inv, con\n&inv -> b\n%b -> con\n&con -> output",
		"%fx -> kh, hl\n%nv -> fx, hl\n&vm -> zg\n%rr -> rl, ql\n&bc -> rn, pm, rp, xg, fv, tn\n%xg -> xm\n%kl -> hl\n&hq -> xt, sp, gl, jd, jl, ss, mc\n%xn -> mv, bc\n%dp -> hh, ql\n%mv -> bc, pp\n%rg -> hq, zq\n&lm -> zg\n%mc -> xt\n%ct -> rd\n%ss -> mz\n%rt -> kk, hl\n%mz -> hq, sp\n%zq -> hq, qj\n%rn -> pm\n%kk -> ld\n%hh -> ql, vb\n%kc -> kl, hl\n%pm -> tt\n%fh -> bc, jn\n&jd -> zg\nbroadcaster -> rt, jr, rp, jl\n%cp -> ql, ln\n&fv -> zg\n&ql -> ln, jr, xs, mg, vm\n%xm -> bc, xn\n%xt -> ss\n%mg -> lt\n%ln -> rr\n%qj -> hq\n%ld -> nv\n%pp -> bc\n%gl -> mc\n%rd -> hl, zz\n%jl -> js, hq\n%jr -> xs, ql\n%jn -> bc, tn\n%sp -> gh\n%vb -> ql\n%gh -> rg, hq\n%vh -> ql, dp\n%js -> gl, hq\n%kh -> ck\n&hl -> kh, rt, ct, kk, lm, ld\n%rp -> fh, bc\n%tt -> xg, bc\n%xs -> mg\n%lt -> ql, cp\n%zz -> hl, kc\n%tn -> rn\n%ck -> hl, ct\n%rl -> ql, vh\n&zg -> rx",
	}

	var res []int
	for _, d := range data {
		input := strings.Split(d, "\n")
		res = append(res, d20p1(input))
		res = append(res, d20p2(input))
	}

	return res
}

type d20op struct {
	t          byte
	state      bool
	recipients []string
}

type d20n struct {
	from   string
	name   string
	signal bool
}

func d20p2(input []string) int {
	m, inputSignals, inputs, broadcast := d20parse(input)
	const EXIT_NODE = "rx"
	target := make(map[string]int)
	for node := range inputs[EXIT_NODE] {
		for n := range inputs[node] {
			target[n] = -1
		}
	}
	cntFound := 0

	for i := 1; i < 10000; i++ {
		nodes := make([]d20n, 0)
		for _, b := range broadcast {
			nodes = append(nodes, d20n{name: b, signal: false, from: "broadcast"})
		}
		for len(nodes) > 0 {
			n := nodes[0]

			nodes = nodes[1:]
			for node, val := range target {
				if val == -1 && inputSignals[node] {
					cntFound++
					target[node] = i
				}
			}
			if cntFound == len(target) {
				cur := -1
				for _, num := range target {
					if cur != -1 {
						cur = lcm(cur, num)
					} else {
						cur = num
					}
				}
				return cur
			}

			hasOutput, signalToProvide := d20processSignal(m[n.name], inputs, n, inputSignals)
			if hasOutput {
				for _, r := range m[n.name].recipients {
					nodes = append(nodes, d20n{name: r, signal: signalToProvide, from: n.name})
				}
			}
		}
	}
	return -1
}

func d20p1(input []string) int {
	m, inputSignals, inputs, broadcast := d20parse(input)

	signalsProvided := make(map[bool]int)
	for i := 0; i < 1000; i++ {
		nodes := make([]d20n, 0)
		signalsProvided[false]++
		for _, b := range broadcast {
			signalsProvided[false]++
			nodes = append(nodes, d20n{name: b, signal: false, from: "broadcast"})
		}
		for len(nodes) > 0 {
			n := nodes[0]

			nodes = nodes[1:]
			node := m[n.name]

			hasOutput, signalToProvide := d20processSignal(node, inputs, n, inputSignals)
			if hasOutput {
				for _, r := range node.recipients {
					signalsProvided[signalToProvide]++
					nodes = append(nodes, d20n{name: r, signal: signalToProvide, from: n.name})
				}
			}
		}
	}
	return signalsProvided[false] * signalsProvided[true]
}

func d20processSignal(node *d20op, inputs map[string]map[string]bool, n d20n, inputSignals map[string]bool) (bool, bool) {
	if node == nil {
		return false, false
	}
	var signalToProvide bool
	if node.t == '&' {
		q := false
		for inp := range inputs[n.name] {
			if !inputSignals[inp] {
				q = true
				break
			}
		}
		signalToProvide = q
	}
	if node.t == '%' {
		if n.signal {
			return false, false
		}
		node.state = !node.state
		signalToProvide = node.state
	}

	inputSignals[n.name] = signalToProvide
	return true, signalToProvide
}

func d20parse(input []string) (map[string]*d20op, map[string]bool, map[string]map[string]bool, []string) {
	m := make(map[string]*d20op)
	inputSignals := make(map[string]bool)
	inputs := make(map[string]map[string]bool)
	var broadcast []string

	for _, s := range input {
		line := strings.Split(s, " -> ")
		op := line[0]
		recipients := strings.Split(line[1], ", ")
		if op == "broadcaster" {
			broadcast = recipients
			continue
		}
		t := op[0]
		name := op[1:]
		m[name] = &d20op{t: t, recipients: recipients}
		inputSignals[name] = false
		for _, r := range recipients {
			if inputs[r] == nil {
				inputs[r] = make(map[string]bool)
			}
			inputs[r][name] = true
		}
	}
	return m, inputSignals, inputs, broadcast
}
