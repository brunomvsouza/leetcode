// Code generated by https://github.com/j178/leetgo

package main

import (
    "testing"

    . "github.com/j178/leetgo/testutils/go"
)

var testcases = `
input:
[1,2,4]
[1,3,4]
output:
[1,1,2,3,4,4]

input:
[]
[]
output:
[]

input:
[]
[0]
output:
[0]
`

func Test_mergeTwoLists(t *testing.T) {
    targetCaseNum := 0
    // targetCaseNum := -1
    if err := RunTestsWithString(t, mergeTwoLists, testcases, targetCaseNum); err != nil {
        t.Fatal(err)
    }
}
