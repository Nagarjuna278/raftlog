package main

import (
	"log"
	"math/rand"
	"raftlog/internal/node"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	peers := []int{0, 1, 2, 3, 4} // 5-node cluster
	nodes := make([]*node.Node, 5)
	for i := 0; i < 5; i++ {
		nodes[i] = node.NewNode(i, peers)
		nodes[i].Start()
	}
	log.Println("Raft cluster started")

	// Simulate client appending a log
	time.Sleep(3 * time.Second) // Wait for leader election
	log.Println("Raft cluster waited for 1 sec")

	var currentLeader *node.Node
	for {
		time.Sleep(5 * time.Second)
		for _, n := range nodes {
			if n.State() == "Leader" {
				currentLeader = n
				n.Append("set x 10")
				break
			}
			log.Printf("node %v ", n.State())
		}
		log.Printf("Current leader: %v\n", currentLeader)

		if currentLeader != nil {
			break
		}
	}

	select {} // Keep running
}
