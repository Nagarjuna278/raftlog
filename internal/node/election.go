package node

import (
	"log"
	"math/rand"
	"time"
)

func (n *Node) runFollower() {
	timeout := time.Duration(0+rand.Intn(150)) * time.Millisecond
	timer := time.NewTimer(timeout)
	select {
	case <-timer.C:
		n.mu.Lock()
		n.state = "Candidate"
		n.currentTerm++
		n.votedFor = n.id
		n.mu.Unlock()
	case <-time.After(50 * time.Millisecond): // Simulate heartbeat
	}
}

func (n *Node) runCandidate() {
	n.mu.Lock()
	term := n.currentTerm
	lastLogIndex := len(n.log) - 1
	lastLogTerm := 0
	if lastLogIndex >= 0 {
		lastLogTerm = n.log[lastLogIndex].Term
	}
	n.mu.Unlock()

	votes := 1
	voteChan := make(chan bool, len(n.peers))
	for _, peer := range n.peers {
		if peer == n.id {
			continue
		}
		go func(p int) {
			args := RequestVoteArgs{
				Term:         term,
				CandidateID:  n.id,
				LastLogIndex: lastLogIndex,
				LastLogTerm:  lastLogTerm,
			}
			reply := n.requestVote(p, args)
			voteChan <- reply.VoteGranted
		}(peer)
	}

	for i := 0; i < len(n.peers)-1; i++ {
		if <-voteChan {
			votes++
		}
	}

	n.mu.Lock()
	if votes > len(n.peers)/2 && n.state == "Candidate" && n.currentTerm == term {
		n.state = "Leader"
		for _, peer := range n.peers {
			n.nextIndex[peer] = len(n.log)
			n.matchIndex[peer] = 0
		}
	}
	n.mu.Unlock()
}

func (n *Node) requestVote(peer int, args RequestVoteArgs) RequestVoteReply {
	n.mu.Lock()
	defer n.mu.Unlock()
	reply := RequestVoteReply{Term: n.currentTerm}
	if args.Term < n.currentTerm {
		log.Printf("Node %d received RequestVote from %d in term %d, denied vote (current term is higher), response: {Term: %d, VoteGranted: %t}", n.id, args.CandidateID, args.Term, reply.Term, reply.VoteGranted)
		return reply
	}
	if args.Term > n.currentTerm {
		n.currentTerm = args.Term
		n.state = "Follower"
		n.votedFor = -1
	}
	if n.votedFor == -1 || n.votedFor == args.CandidateID {
		lastIndex := len(n.log) - 1
		lastTerm := 0
		if lastIndex >= 0 {
			lastTerm = n.log[lastIndex].Term
		}
		if args.LastLogTerm > lastTerm || (args.LastLogTerm == lastTerm && args.LastLogIndex >= lastIndex) {
			n.votedFor = args.CandidateID
			reply.VoteGranted = true
			log.Printf("Node %d granted vote to %d in term %d, response: {Term: %d, VoteGranted: %t}", n.id, args.CandidateID, args.Term, reply.Term, reply.VoteGranted)
			return reply
		}
	}
	log.Printf("Node %d denied vote to %d in term %d, response: {Term: %d, VoteGranted: %t}", n.id, args.CandidateID, args.Term, reply.Term, reply.VoteGranted)
	return reply
}
