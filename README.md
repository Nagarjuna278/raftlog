# RaftLog: A Raft Consensus Algorithm Simulation

This project is a Go-based simulation of the Raft consensus algorithm. It demonstrates the core concepts of Raft, 
including leader election and log replication, within a simplified environment.

## Overview

This system creates a cluster of nodes that communicate to maintain a consistent log. One node acts as the leader, 
appending log entries and replicating them to followers. If the leader fails (e.g., due to a network change or crash), 
the remaining nodes elect a new leader to continue replication. The `cmd/server/main.go` file sets up the cluster, 
starts the nodes, and runs the replication process.

## Components
            
*   **`internal/node`:** Contains the core node implementation.
    *   `node.go`: Defines the `Node` struct and its methods for handling leader and follower roles, including the main loop for replication and failure detection.
    *   `election.go`: Implements basic leader election logic, triggered when the current leader is lost.
    *   `rpc.go`: Placeholder for RPC handlers (to be implemented for real network communication).
*   **`internal/log`:** Manages the log structure and replication.
    *   `log.go`: Defines the `Log` struct and methods for appending and reading log entries.
    *   `sync.go`: Handles log replication from the leader to followers.
*   **`internal/network`:** Placeholder for network communication logic.
    *   `network.go`: To be implemented for real node-to-node communication (e.g., heartbeats, log replication).
*   **`internal/config`:** Configuration management.
    *   `config.go`: Defines the `Config` struct for node IDs and peer addresses.
*   **`cmd/server/main.go`:** The main entry point. Sets up the cluster and starts the nodes.

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
*   Implement real network communication in network.go and rpc.go (e.g., using gRPC for heartbeats and log replication).
*   Add a user interface for visualizing the cluster and its log.
*   Implement a persistent log storage mechanism.
*   Add more sophisticated client interactions and error handling.

## License

This project is licensed under the MIT License - see the `LICENSE` file for details.

## Contact

Maintained by [Nagarjuna278](https://github.com/Nagarjuna278) - iamnagarjunaguntaka@gmail.com.  Feel free to reach out with questions or suggestions!