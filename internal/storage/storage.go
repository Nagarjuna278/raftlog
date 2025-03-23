package storage

import (
	"log"
	"raftlog/internal/node"
)

// Persist simulates saving state to disk
func Persist(n *node.Node) {
	n.mu.Lock()
	defer n.mu.Unlock()
	log.Printf("Node %d persisted state: term=%d, log=%v", n.id, n.currentTerm, n.log)
}
