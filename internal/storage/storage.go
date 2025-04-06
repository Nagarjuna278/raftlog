package storage

import (
	"log"
	"raftlog/internal/node"
)

// Persist simulates saving state to disk
func Persist(n *node.Node) {
	n.Lock()
	defer n.Unlock()
	log.Printf("Node %d persisted state: term=%d, log=%v", n.id, n.currentTerm, n.log)
}
