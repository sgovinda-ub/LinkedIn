/*
Design a data structure that will be initialized with a string array, and then it should answer queries of the shortest distance between
two different strings from the array.

Implement the WordDistance class:

WordDistance(String[] wordsDict) initializes the object with the strings array wordsDict.
int shortest(String word1, String word2) returns the shortest distance between word1 and word2 in the array wordsDict.

Example 1:

Input
["WordDistance", "shortest", "shortest"]
[[["practice", "makes", "perfect", "coding", "makes"]], ["coding", "practice"], ["makes", "coding"]]
Output
[null, 3, 1]

Explanation
WordDistance wordDistance = new WordDistance(["practice", "makes", "perfect", "coding", "makes"]);
wordDistance.shortest("coding", "practice"); // return 3
wordDistance.shortest("makes", "coding");    // return 1

Constraints:

    1 <= wordsDict.length <= 3 * 104
    1 <= wordsDict[i].length <= 10
    wordsDict[i] consists of lowercase English letters.
    word1 and word2 are in wordsDict.
    word1 != word2
    At most 5000 calls will be made to shortest.
*/

type WordDistance struct {
    wordIndexer map[string][]int
}


func Constructor(wordsDict []string) WordDistance {
    indexer := make(map[string][]int)
    for idx, word := range wordsDict {
        if val, ok := indexer[word]; ok {
            val = append(val, idx)
            indexer[word] = val
        } else {
            indexer[word] = []int{idx}
        }
    }
    fmt.Println(indexer)
    return WordDistance{indexer}
}

func min(x int, y int) int {
    if x < y {
        return x
    }
    return y
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func (this *WordDistance) Shortest(word1 string, word2 string) int {
    i := 0
    j := 0
    result := math.MaxInt32
    list1 := this.wordIndexer[word1]
    list2 := this.wordIndexer[word2]
    for i < len(list1) && j < len(list2) {
        result = min(result, abs(list1[i]-list2[j]))
        if list1[i] < list2[j] {
            i++
        } else {
            j++
        }
    }
    return result
}


/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */
