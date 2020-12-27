# Algorithms

Implemented so far:

- Brute Force Algorithms (bruteforce.go)
- Divide and Conquer Algorithms (divide.go)
- Backtracking Algorithms (backtrack.go)

## Brute Force Algorithms

Brute Force Algorithms are known for wide applicability and simplicity in solving complex problems.  They do not provide efficient algorithms.  The algorithms are slow and non-performant.

### Use-cases

Searching, string matching, and matrix multiplication.  Single computational tasks can also be solved using brute force algorithms.

## Divide and Conquer Algorithms

Divide and Conquer algorithms break a complex problem into smaller problems to solve, and then solves them.  The approach is to recursively solve the sub-problems and merge the solutions of the sub-problems.

### Examples

Recursion, quick sort, binary search, fast Fourier transform, and merge sort are good examples of divide and conquer algorithms.  
Memory use is efficent with these algorithms.  Performance is sometimes an issue in the case of recursion.  On multiprocessor machines, these algorithms can be executed on different processors after breaking them down into sub-problems.  

## Backtracking Algorithms

A Backtracking algorithm solves a problem by constructing the solution incrementally.  Multiple options are evaluated, and the algorithm chooses to go to the next component of the solution through *recursion*.  Backtracking can be a chronological type or can traverse the paths, depending on the problem you are solving.

### Uses
Backtracking is useful in scenarios such as finding a value in an unordered table.  *It is faster than a brute force algorithm*, which rejects a large number of solutions in an iteration.  Constraint satisfaction problems such as parsing, rules engine, knapsack problems, and combinatorial optimization are solved using backtracking.