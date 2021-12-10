package hash

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func Test_ConsistentHash(t *testing.T) {
	var (
		node1 = "a"
		node2 = "b"
		node3 = "c"
	)

	nodeWeight := make(map[string]int)
	nodeWeight[node1] = 2
	nodeWeight[node2] = 2
	nodeWeight[node3] = 2

	hash := NewHashRing(DefaultVirtualSpots)

	hash.AddNodes(nodeWeight)

	m := make(map[string]int)

	var nodeMatch string

	for i := 0; i < 1000; i++ {
		s := rand.Int()
		nodeMatch = hash.GetNode(strconv.Itoa(s))
		count := m[nodeMatch]
		m[nodeMatch] = count + 1
	}
	fmt.Println(m)
	fmt.Println("node1 shut down")
	hash.RemoveNode(node1)

	m = make(map[string]int)
	for i := 0; i < 1000; i++ {
		nodeMatch = hash.GetNode(strconv.Itoa(i))
		count := m[nodeMatch]
		m[nodeMatch] = count + 1
	}

	fmt.Println(m)
}
