Given two linked lists where the value of each node is a single digit
integer representing a number:

```
# 321
List 1: 1 -> 2 -> 3

# 998
List 2: 8 -> 9 -> 9
```

Create a function that will add the nodes of each cell in each position
together and return a third linked list in the same format:

```
#  321
#+ 998
#-----
# 1319

List 3: 9, 1, 3, 1
```

## Notes

The key points here is that there needs to be a helper func that will be
able to carry over any overflow. The recursive function will call the next
nodes and add the carryover value to the next iteration.

The other key is being able to handle the base cases:
- what happens if the first list is empty? including carryover.
- what happens if the second list is empty? including carryover.
- what happens if both lists are empty and there's carryover?

The main case is that both nodes are not empty and that they need to be added
together + carryover is handled.

The chain that starts the recursion is assigning the current node's Next() to
be a function call where the arguments for each list is the next node + carryover.

```
func AddList(list1, list2 *Node)
    // call helper func that includes carryover

func AddListHelper(list1, list2 *Node, int carryover)
    // case: list1 empty, carryover
    // case: list2 empty, carryover
    // case: both empty, carryover
    // case: both non nil, carryover
```
