// Code generated by https://github.com/j178/leetgo

package main

import (
    "testing"

    . "github.com/j178/leetgo/testutils/go"
)

var testcases = `
input:
[-1,0,3,5,9,12]
9
output:
4

input:
[-1,0,3,5,9,12]
2
output:
-1
`

func Test_search(t *testing.T) {
    targetCaseNum := 0
    // targetCaseNum := -1
    if err := RunTestsWithString(t, search, testcases, targetCaseNum); err != nil {
        t.Fatal(err)
    }
}