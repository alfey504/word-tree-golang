package main

import (
	"time"

	"practice.com/word_storage_tree/random_string"
	"practice.com/word_storage_tree/word_tree"
)

// Globals -----------------------------------
const testCount = 170000

var testString = random_string.GenerateRandomStrings(testCount, 10)
var randomIndexes = func() []int {
	values := make([]int, testCount)
	for pos := range values {
		values[pos] = random_string.RandomNumber(0, testCount-1)
	}
	return values
}()

var wordTree = func() word_tree.WordTree {
	wordTree := word_tree.CreateWordTree()
	for _, word := range testString {
		wordTree.AddWord(word)
	}
	return wordTree
}()

// Main Function ----------------------------
func main() {
	Test()
}

// Benchmark Function -----------------------
func BenchmarkFindOnTree() {
	for _, index := range randomIndexes {
		word := testString[index]
		wordTree.DoesExist(word)
	}
}

func BenchmarkLinerSearch() {
	for _, index := range randomIndexes {
		word := testString[index]
		for _, wordTest := range testString {
			if word == wordTest {
				break
			}
		}
	}
}

// Test -------------------------------------
func Test() {
	timeBefore := time.Now()
	BenchmarkFindOnTree()
	timeAfter := time.Now()
	Log("Tree", timeBefore, timeAfter)

	timeBefore = time.Now()
	BenchmarkLinerSearch()
	timeAfter = time.Now()
	Log("Linear Search", timeBefore, timeAfter)
}

func Log(title string, timeBefore time.Time, timeAfter time.Time) {
	println(title)
	runTime := (timeAfter.Sub(timeBefore).Microseconds())
	println("________________________________________")
	println("Run time: ", runTime)
	print("\n\n")
}
