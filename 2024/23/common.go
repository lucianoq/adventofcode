package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

type Set map[string]struct{}

func (s *Set) ToKey(others ...string) string {
	list := make([]string, 0, len(*s)+len(others))
	for k := range *s {
		list = append(list, k)
	}
	for _, other := range others {
		list = append(list, other)
	}
	sort.Strings(list)
	return strings.Join(list, ",")
}

func (s *Set) FromKey(key string) {
	m := Set{}
	for _, s := range strings.Split(key, ",") {
		m[s] = struct{}{}
	}
	*s = m
}

type Graph map[string]Set

func (g Graph) AddEdge(from, to string) {
	if _, ok := g[from]; !ok {
		g[from] = Set{}
	}
	g[from][to] = struct{}{}
}

func (g Graph) Connected(from, to string) bool {
	if f, ok := g[from]; ok {
		_, ok := f[to]
		return ok
	}
	return false
}

func (g Graph) Keys() Set {
	keys := Set{}
	for n := range g {
		keys[n] = struct{}{}
	}
	return keys
}

func parseInput() Graph {
	graph := Graph{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ff := strings.Split(scanner.Text(), "-")
		graph.AddEdge(ff[0], ff[1])
		graph.AddEdge(ff[1], ff[0])
	}
	return graph
}

func (g Graph) Expand(clusters Set) Set {
	res := Set{}
	for clusterStr := range clusters {
		cluster := Set{}
		cluster.FromKey(clusterStr)
		for node := range g[clusterStr[:2]] {
			if _, ok := cluster[node]; ok {
				continue
			}
			if !g.connectedToAll(node, cluster) {
				continue
			}
			res[cluster.ToKey(node)] = struct{}{}
		}
	}
	return res
}

func (g Graph) connectedToAll(node string, cluster Set) bool {
	for old := range cluster {
		if !g.Connected(old, node) {
			return false
		}
	}
	return true
}
