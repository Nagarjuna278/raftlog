package node

import (
	"sync"
)

// Node represents a Raft server
type Node struct {
	Mu          sync.Mutex
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
	n.Mu.Lock()
	defer n.Mu.Unlock()
	return n.state
}

func (n *Node) Append(command string) {
	n.Mu.Lock()
	if n.state != "Leader" {
		n.Mu.Unlock()
		return
	}
	entry := LogEntry{
		Index:   len(n.log),
		Term:    n.currentTerm,
		Command: command,
	}
	n.log = append(n.log, entry)
	n.Mu.Unlock()
}

func (n *Node) run() {
	for {
		n.Mu.Lock()
		state := n.state
		n.Mu.Unlock()

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
