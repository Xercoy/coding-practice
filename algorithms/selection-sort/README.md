# Selection Sort

https://iq.opengenus.org/time-complexity-of-selection-sort/

https://www.geeksforgeeks.org/selection-sort/

Selection sort is one of the most basic sorting algorithms out there. It works
as you might think if you were to physically sort a list of some criteria where
there is a 'minimum' and a 'maximum' within a collection that can be compared.

Start at the first element, compare that current element with the rest of the
elements in the list. If the current element is greater than the one being
compared, switch the values. By switching, you make it so that the lower value
is moved to the left and the higher value is moved to the right. Proceed to do this comparison betwen the current element (starting at 0) and the rest of the nodes.

When this is done completely for the first node (element 0), it is guaranteed that
the node in the first element is the minimum. This is because the algorithm is
essentially 'finding the lowest value' within a list. The algorithm SELECTED the
lowest value.

This notion of 'finding the lowest value' becomes a sorting algorithm because
in subsequent passes/iterations, you mark the next element in the list as the
current element, THEN you find the minimum value between the current element
and the ones that are remaining. When you find the minimum for every element in
the list, you can be sure that the entire list has been sorted from min to max.

The other way you can think about it is like this:

(considering a list with a length of n, where n is at least 3):
```
- pass 0: find the lowest value in the list from elements 0-n and put it in element 0

- pass 1: find the second lowest value in the list from elements 1-n and put it
in element 1

- pass 2: find the third lowest value in the list from elements 2-n and put it in element 2
```

See how that works? Every single time you move through the list to find the lowest value, second lowest value, or third lowest value, the number of elements that you have to compare,
AKA the number of elements that are unsorted, shrinks.

Because this is algorithm contains a nested loop and doesnt end early, n^2 is the best,
worst, and average time complexity. Since the collection/list is static, the space complexity
is O(1).

This algo iterates upon itself compeletely for every element in the list, more proof that it is n * n = n^2... I'll do a better job of explaining that.
