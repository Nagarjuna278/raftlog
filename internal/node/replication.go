package node

import (
	"log"
	"time"
)

func (n *Node) runLeader() {
	n.Mu.Lock()
	term := n.currentTerm
	logEntries := n.log
	commitIndex := n.commitIndex
	n.Mu.Unlock()

	applyChan := make(chan LogEntry, 10)
	go func() {
		for entry := range applyChan {
			log.Printf("Node %d applied log entry: %v", n.id, entry)
		}
	}()

	for _, peer := range n.peers {
		if peer == n.id {
			continue
		}
		go func(p int) {
			n.Mu.Lock()
			prevLogIndex := n.nextIndex[p] - 1
			prevLogTerm := 0
			if prevLogIndex >= 0 {
				prevLogTerm = logEntries[prevLogIndex].Term
			}
			entries := logEntries[prevLogIndex+1:]
			args := AppendEntriesArgs{
				Term:         term,
				LeaderID:     n.id,
				PrevLogIndex: prevLogIndex,
				PrevLogTerm:  prevLogTerm,
				Entries:      entries,
				LeaderCommit: commitIndex,
			}
			n.Mu.Unlock()

			reply := n.appendEntries(p, args)
			n.Mu.Lock()
			if reply.Success {
				n.matchIndex[p] = prevLogIndex + len(entries)
				n.nextIndex[p] = n.matchIndex[p] + 1
			} else if reply.Term > n.currentTerm {
				n.currentTerm = reply.Term
				n.state = "Follower"
				n.votedFor = -1
			} else {
				n.nextIndex[p]--
			}
			for N := len(n.log) - 1; N > n.commitIndex; N-- {
				count := 1
				for _, peer := range n.peers {
					if peer != n.id && n.matchIndex[peer] >= N {
						count++
					}
				}
				if count > len(n.peers)/2 && n.log[N].Term == n.currentTerm {
					n.commitIndex = N
					break
				}
			}
			for i := n.lastApplied + 1; i <= n.commitIndex; i++ {
				applyChan <- n.log[i]
				n.lastApplied = i
			}
			n.Mu.Unlock()
		}(peer)
	}
	time.Sleep(50 * time.Millisecond)
}

func (n *Node) appendEntries(peer int, args AppendEntriesArgs) AppendEntriesReply {
	n.Mu.Lock()
	defer n.Mu.Unlock()
	reply := AppendEntriesReply{Term: n.currentTerm}
	if args.Term < n.currentTerm {
		return reply
	}
	if args.Term > n.currentTerm {
		n.currentTerm = args.Term
		n.state = "Follower"
		n.votedFor = -1
	}
	if args.PrevLogIndex >= len(n.log) {
		return reply
	}
	if args.PrevLogIndex >= 0 && n.log[args.PrevLogIndex].Term != args.PrevLogTerm {
		n.log = n.log[:args.PrevLogIndex]
		return reply
	}
	n.log = append(n.log[:args.PrevLogIndex+1], args.Entries...)
	reply.Success = true
	if args.LeaderCommit > n.commitIndex {
		n.commitIndex = min(args.LeaderCommit, len(n.log)-1)
	}
	return reply
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
