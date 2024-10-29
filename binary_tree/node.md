# Binary Tree

- A tree is a data structure that store data in hirerchy

- We refer to this element as nodes and the line it is connecte is edges

- Each node contain value (object)

- Top node of tree is called `Root`

- Bottom node dosen't have children is know as leaf node

- Tree usage
    - Represent hierarchical data
    - Database
    - Autocompletion
    - Compiler (Syntax tree)
    - Compression algo (JPEG, MP3)

- Binary Search Tree (BST) `left < node < right`

- Time complexity Logarithmic time complaxity
    - Lookup `O(log n)`
    - Insert `O(log n)`
    - Delete `O(log n)`

- Tree traversal
    - Breadth first (level order)
    - Depth first 
        - Pre-order Root, Left, Right
        - In-order Left, Root, Right
        - Post-order Left, Right, Root

- Properties of tree node
    - depth - counting the number of edges from root to target node
    - height - counting the number of edges from leaf (longest node) to the target node
    - height of root node is also called the height of the tree


- Formula for calculating the tree height using post traversal

```
1 + max(height(L), height(R))
```
