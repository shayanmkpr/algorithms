0. **Binary Search**  
Binary search is used when data is sorted, or when the answer space has a monotonic condition like “too small, too small, valid, valid.” Implementation usually keeps `left` and `right` boundaries, checks `mid`, then moves one boundary to discard half the search space. It is used for sorted arrays, lower/upper bounds, search suggestions, and “minimum value that satisfies X” problems. Time complexity is `O(log n)`.

1. **Merge Sort**  
Merge sort is a real production-relevant sorting algorithm, especially when stable sorting or predictable performance matters. It works by recursively splitting the array into halves until the pieces are size 1, then merging sorted halves back together using two pointers. It always runs in `O(n log n)`, but usually needs `O(n)` extra memory. It is also useful for linked lists and external sorting where data may not fit in memory.

2. **Quick Sort**  
Quick sort is a fast in-memory sorting algorithm. It chooses a pivot, partitions the array so smaller items go before the pivot and larger items after it, then recursively sorts both sides. Its average time complexity is `O(n log n)`, but bad pivots can make it `O(n²)`, so real implementations use randomized pivots or hybrid strategies. It is important because partitioning shows up in many other problems too.

3. **Two Pointers**  
Two pointers is a technique where two indexes move through an array/string to avoid nested loops. A common implementation starts one pointer at the beginning and one at the end, then moves one side based on the current condition. It is heavily used with sorted arrays, pair-sum problems, removing duplicates, reversing arrays, palindrome checks, and partitioning. Most solutions are `O(n)`.

4. **Sliding Window**  
Sliding window is used when the problem asks about a contiguous subarray or substring. Implementation usually has `left` and `right` pointers: expand `right` to include new elements, and move `left` forward when the window becomes invalid or too large. You maintain state such as count, sum, frequency map, or max value as the window moves. It is used for longest substring, shortest subarray, rate limiting, and streaming-window problems. Usually `O(n)`.

5. **Prefix Sum**  
Prefix sum precomputes cumulative totals so range queries become fast. Implementation creates an array where `prefix[i]` stores the sum before or up to index `i`; then range sum `l..r` is calculated by subtracting two prefix values. It is used in analytics, range queries, subarray-sum problems, and 2D matrix sums. Building takes `O(n)`, each query is `O(1)`.

6. **BFS - Breadth-First Search**  
BFS explores a graph or tree level by level. Implementation uses a queue: push the start node, repeatedly pop the front, visit its neighbors, and mark them as seen. BFS is the standard choice for shortest path in unweighted graphs because the first time you reach a node is through the minimum number of edges. Time complexity is `O(V + E)`.

7. **DFS - Depth-First Search**  
DFS explores one path as deeply as possible before backtracking. It can be implemented recursively with a call stack or iteratively with an explicit stack. DFS is used for tree traversal, graph reachability, connected components, cycle detection, parsing structures, and many recursive search problems. Time complexity is `O(V + E)` for graphs.

8. **Topological Sort**  
Topological sort orders nodes so dependencies come before dependents. A common implementation is Kahn’s algorithm: compute each node’s indegree, push nodes with indegree `0` into a queue, remove them one by one, and reduce the indegree of their neighbors. It only works on directed acyclic graphs. It is used in build systems, dependency resolution, task scheduling, migrations, and course-prerequisite problems. Time complexity is `O(V + E)`.

9. **Heap / Priority Queue**  
A heap keeps the smallest or largest item available without fully sorting all data. Implementation usually uses an array where each node’s children are at predictable indexes, and operations “bubble up” or “heapify down” to restore order. It is used in schedulers, top-K queries, merging sorted lists, event loops, and graph algorithms like Dijkstra. Push and pop are `O(log n)`, peek is `O(1)`.

10. **Dijkstra’s Algorithm**  
Dijkstra finds shortest paths from one source when edge weights are non-negative. Implementation keeps a distance map and a priority queue of the next cheapest node; when a shorter path to a neighbor is found, you update its distance and push it into the queue. It is used in routing, maps, network latency, dependency cost, and game pathfinding. With a heap, time complexity is usually `O((V + E) log V)`.

11. **Union-Find / Disjoint Set**  
Union-Find manages groups of connected items. Implementation stores a `parent` array where each item points to its group representative, plus optimizations like path compression and union by rank/size. It supports `find(x)` to get the group leader and `union(a, b)` to merge groups. It is used for connectivity, clustering, Kruskal’s MST, and detecting cycles in undirected graphs. Operations are almost `O(1)` amortized.

12. **Backtracking**  
Backtracking is controlled brute force. Implementation uses recursion: choose an option, add it to the current state, recurse, then undo the choice before trying the next option. It becomes practical when you can prune invalid paths early. It is used for permutations, combinations, Sudoku, N-Queens, search with constraints, and generating valid configurations. Time complexity is usually exponential.

13. **Greedy Algorithms**  
Greedy algorithms make the best local choice at each step, but they are only valid when that choice can be proven to lead to the global optimum. Implementation is often simple: sort or prioritize items, then scan and accept/reject choices based on a rule. They are used in interval scheduling, activity selection, Huffman coding, some cache/queue policies, and minimum spanning tree algorithms. Complexity is often `O(n)` or `O(n log n)`.

13. **Dynamic Programming**  
Dynamic programming solves problems with overlapping subproblems by storing previous results. Implementation can be top-down with recursion plus memoization, or bottom-up with a `dp` array/table. The key is defining the state, the transition, and the base case. It is used for knapsack, edit distance, longest common subsequence, coin change, sequence optimization, and many planning problems. Complexity depends on states and transitions, often `O(n)`, `O(n²)`, or `O(n * capacity)`.

14. **Bellman-Ford**  
Bellman-Ford finds shortest paths even when negative edge weights exist. Implementation initializes distances, then relaxes every edge `V - 1` times; one extra pass can detect a negative cycle. It is slower than Dijkstra but handles cases Dijkstra cannot. It is used in financial arbitrage detection, routing with penalties, and graphs where negative weights are meaningful. Time complexity is `O(VE)`.

15. **Floyd-Warshall**  
Floyd-Warshall computes shortest paths between every pair of nodes. Implementation uses a 2D distance matrix and tries each node as an intermediate point between every pair `i` and `j`. It is simple and powerful for small or dense graphs, but too expensive for large graphs. Time complexity is `O(V³)`, space is `O(V²)`.

16. **Minimum Spanning Tree - Kruskal / Prim**  
A minimum spanning tree connects all nodes with minimum total edge cost and no cycles. Kruskal’s algorithm sorts all edges and adds the cheapest edge that does not create a cycle, usually using Union-Find. Prim’s algorithm starts from one node and repeatedly adds the cheapest edge that expands the tree, usually using a priority queue. These are used in network design, clustering, infrastructure planning, and optimization over connected weighted graphs. Kruskal is usually `O(E log E)`, Prim is usually `O(E log V)`.
