package node

import (
	"sync"
)

// Node represents a Raft server
type Node struct {
	mu          sync.Mutex
	id          int
	state       string
	currentTerm int
	votedFor    int
	log         []LogEntry
	commitIndex int
	lastApplied int
	nextIndex   map[int]int
	matchIndex  map[int]int
	peers       []int
}

func NewNode(id int, peers []int) *Node {
	return &Node{
		id:          id,
		state:       "Follower",
		currentTerm: 0,
		votedFor:    -1,
		log:         make([]LogEntry, 0),
		commitIndex: 0,
		lastApplied: 0,
		nextIndex:   make(map[int]int),
		matchIndex:  make(map[int]int),
		peers:       peers,
	}
}

func (n *Node) Start() {
	go n.run()
}

func (n *Node) State() string {
	n.mu.Lock()
	defer n.mu.Unlock()
	return n.state
}

func (n *Node) Append(command string) {
	n.mu.Lock()
	if n.state != "Leader" {
		n.mu.Unlock()
		return
	}
	entry := LogEntry{
		Index:   len(n.log),
		Term:    n.currentTerm,
		Command: command,
	}
	n.log = append(n.log, entry)
	n.mu.Unlock()
}

func (n *Node) run() {
	for {
		n.mu.Lock()
		state := n.state
		n.mu.Unlock()

		switch state {
		case "Follower":
			n.runFollower()
		case "Candidate":
			n.runCandidate()
		case "Leader":
			n.runLeader()
		}
	}
}

// Lock locks the node mutex lock
func (n *Node) Lock() {
	n.mu.Lock()
}

// Unlock unlocks the mutex lock
func (n *Node) Unlock() {
	n.Unlock()
}
