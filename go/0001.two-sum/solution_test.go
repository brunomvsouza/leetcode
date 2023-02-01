// Code generated by https://github.com/j178/leetgo

package main

import (
	"testing"

	. "github.com/j178/leetgo/testutils/go"
)

var testcases = `
input:
[2,7,11,15]
9
output:
[0,1]

input:
[3,2,4]
6
output:
[1,2]

input:
[3,3]
6
output:
[0,1]
`

func Test_twoSum(t *testing.T) {
	targetCaseNum := 0
	// targetCaseNum := -1
	if err := RunTestsWithString(t, twoSum, testcases, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
