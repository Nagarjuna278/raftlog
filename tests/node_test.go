package tests

import (
	"raftlog/internal/node"
	"testing"
	"time"
)

func TestLeaderElection(t *testing.T) {
	peers := []int{0, 1, 2}
	nodes := make([]*node.Node, 3)
	for i := 0; i < 3; i++ {
		nodes[i] = node.NewNode(i, peers)
		nodes[i].Start()
	}
	time.Sleep(1 * time.Second)
	leaders := 0
	for _, n := range nodes {
		if n.State() == "Leader" {
			leaders++
		}
	}
	if leaders != 1 {
		t.Fatalf("Expected 1 leader, got %d", leaders)
	}
}
