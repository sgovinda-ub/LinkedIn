/*
You are given a nested list of integers nestedList. Each element is either an integer or a list whose elements may also be integers or other lists.
The depth of an integer is the number of lists that it is inside of. For example, the nested list [1,[2,2],[[3],2],1] has each integer's value set
to its depth. Let maxDepth be the maximum depth of any integer.

The weight of an integer is maxDepth - (the depth of the integer) + 1.

Return the sum of each integer in nestedList multiplied by its weight.

Example 1:

Input: nestedList = [[1,1],2,[1,1]]
Output: 8
Explanation: Four 1's with a weight of 1, one 2 with a weight of 2.
1*1 + 1*1 + 2*2 + 1*1 + 1*1 = 8

Example 2:

Input: nestedList = [1,[4,[6]]]
Output: 17
Explanation: One 1 at depth 3, one 4 at depth 2, and one 6 at depth 1.
1*3 + 4*2 + 6*1 = 17

Constraints:

    1 <= nestedList.length <= 50
    The values of the integers in the nested list is in the range [-100, 100].
    The maximum depth of any integer is less than or equal to 50.
*/
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func max(x int, y int) int {
    if x > y {
        return x
    }    
    return y
}

func calculateMaxDepth(nestedList []*NestedInteger, depth int, maxDepth *int) int {
    depth = depth + 1
    *maxDepth = max(*maxDepth, depth)
    for _, item := range nestedList { 
        if len(item.GetList()) != 0 {
            calculateMaxDepth(item.GetList(), depth, maxDepth)
        }
    }
    return *maxDepth
}

func calculateTotalWeight(nestedList []*NestedInteger, depth int, maxDepth int) int {
    sum := 0
    depth = depth + 1
    for _, item := range nestedList { 
        if len(item.GetList()) != 0 {
            sum = sum + calculateTotalWeight(item.GetList(), depth, maxDepth)
        } else {
            value := item.GetInteger()
            sum = sum + value * (maxDepth - depth + 1)
        }
    }
    return sum 
}

func depthSumInverse(nestedList []*NestedInteger) int {
    maxDepth := 0
    calculateMaxDepth(nestedList, 0, &maxDepth)
    totalWeight := calculateTotalWeight(nestedList, 0, maxDepth)
    return totalWeight
}
