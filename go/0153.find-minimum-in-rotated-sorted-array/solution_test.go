// Code generated by https://github.com/j178/leetgo.

package main

import (
    "testing"

    . "github.com/j178/leetgo/testutils/go"
)

var testcases = `
input:
[3,4,5,1,2]
output:
1

input:
[4,5,6,7,0,1,2]
output:
0

input:
[11,13,15,17]
output:
11
`

func Test_findMin(t *testing.T) {
    targetCaseNum := 0
    // targetCaseNum := -1
    if err := RunTestsWithString(t, findMin, testcases, targetCaseNum); err != nil {
        t.Fatal(err)
    }
}
