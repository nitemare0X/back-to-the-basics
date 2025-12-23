# Challenge: Combine Them!

Now that you have both filter and map, let's do something practical:

## Problem: Process Numbers in a Pipeline

Goal: Given a list of numbers, get the squares of only the even numbers.

Example:

Input: [1, 2, 3, 4, 5, 6]
Filter evens: [2, 4, 6]
Square them: [4, 16, 36]
Output: [4, 16, 36]

Your Task:

Write code that:

1. Uses your filter function (the original one that checks conditions!)
2. Uses your map function (rename from filter to map first!)
3. Chains them together to solve the problem

Hint: You'll call one function, then pass its result to the other:

```javascript
const result = map(filter(numbers, isEven), square);
```

Test with: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

Expected: [4, 16, 36, 64, 100]

This teaches you function composition - combining simple functions to solve complex problems. It's how professional code is written! 🚀
