# RaftLog: A Raft Consensus Algorithm Simulation

This project is a Go-based simulation of the Raft consensus algorithm. It demonstrates the core concepts of Raft, including leader election and log replication, within a simplified environment.

## Overview

This simulation creates a cluster of Raft nodes that communicate with each other to elect a leader and maintain a consistent log. The `cmd/main.go` file sets up the simulation, creates the nodes, starts them, and then simulates a client appending a log entry to the leader.

## Components

*   **`internal/node`:** Contains the core Raft node implementation.
    *   `node.go`: Defines the `Node` struct and its methods for handling different Raft states (Follower, Candidate, Leader).
    *   `types.go`: Defines the data structures used for communication between nodes (e.g., `LogEntry`, `RequestVoteArgs`, `AppendEntriesArgs`).
*   **`cmd/main.go`:** The main entry point of the simulation.  It sets up the cluster, starts the nodes, and simulates a client appending a log entry.

## Usage

1.  **Prerequisites:**

    *   Go (version 1.16 or later)

2.  **Clone the repository:**

    ```bash
    git clone <repository_url>
    cd raftlog
    ```

3.  **Run the simulation:**

    ```bash
    go run cmd/main.go
    ```

4.  **Observe the output:**

    The simulation will print log messages to the console, showing the state of each node, the leader election process, and the log replication.

## Simulation Details

*   A 5-node cluster is created.
*   Each node starts as a Follower.
*   The simulation waits for a few seconds to allow leader election to occur.
*   A single log entry ("set x 10") is appended to the leader's log.
*   The simulation runs indefinitely, allowing you to observe the cluster's behavior.

## Future Enhancements

*   Implement the full Raft algorithm, including handling of network partitions and node failures.
*   Add a user interface for visualizing the cluster and its log.
*   Implement a persistent log storage mechanism.
*   Add more sophisticated client interactions and error handling.

## License

This project is licensed under the MIT License - see the `LICENSE` file for details.

## Contact

Maintained by [Nagarjuna278](https://github.com/Nagarjuna278) - iamnagarjunaguntaka@gmail.com.  Feel free to reach out with questions or suggestions!