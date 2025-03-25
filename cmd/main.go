package main

import (
	"log"
	"math/rand"
	"raftlog/internal/node"
	"time"
)

// main function is the entry point of the Raft cluster simulation.
func main() {
	// Seed the random number generator for randomness in elections and other operations.
	rand.Seed(time.Now().UnixNano())

	// Define the IDs of the nodes in the Raft cluster (a 5-node cluster in this case).
	peers := []int{0, 1, 2, 3, 4}

	// Create a slice to hold the Node instances.
	nodes := make([]*node.Node, 5)

	// Initialize and start each node in the cluster.
	for i := 0; i < 5; i++ {
		nodes[i] = node.NewNode(i, peers) // Create a new node with ID 'i' and the list of peer IDs.
		nodes[i].Start()                  // Start the node's main loop in a goroutine.
	}

	// Log a message indicating that the Raft cluster has been started.
	log.Println("Raft cluster started")

	// Simulate a client appending a log entry to the leader.
	time.Sleep(3 * time.Second) // Wait for leader election to occur.
	log.Println("Raft cluster waited for 3 sec")

	// Variable to store the current leader node.
	var currentLeader *node.Node

	// Loop to find the leader node.
	for {
		time.Sleep(5 * time.Second) // Wait for a short duration to allow leader election to converge.

		// Iterate through the nodes to find the leader.
		for _, n := range nodes {
			if n.State() == "Leader" { // Check if the node's state is "Leader".
				currentLeader = n    // Assign the leader node to the 'currentLeader' variable.
				n.Append("set x 10") // Append a command to the leader's log.
				break                // Exit the inner loop once the leader is found.
			}
			log.Printf("node %v ", n.State()) // Log the current state of the node
		}
		log.Printf("Current leader: %v\n", currentLeader) // Log the current leader

		// If a leader has been found, exit the outer loop.
		if currentLeader != nil {
			break
		}
	}

	// Keep the main function running indefinitely to allow the Raft cluster to operate.
	select {}
}
