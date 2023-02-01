// Code generated by https://github.com/j178/leetgo

package main

import (
    "testing"

    . "github.com/j178/leetgo/testutils/go"
)

var testcases = `
input:
[4,2,7,1,3,6,9]
output:
[4,7,2,9,6,3,1]

input:
[2,1,3]
output:
[2,3,1]

input:
[]
output:
[]
`

func Test_invertTree(t *testing.T) {
    targetCaseNum := 0
    // targetCaseNum := -1
    if err := RunTestsWithString(t, invertTree, testcases, targetCaseNum); err != nil {
        t.Fatal(err)
    }
}