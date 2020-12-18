package fuzzysearcher

import (
	"bufio"
	"github.com/armon/go-radix"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// TODO suggestion/autocomplete endpoint,
// TODO 	returns array of strings from radixTree current node to be used as autocomplete values
// TODO tests

type CompleteWorks struct {
	keys []string
	wordIndex map[string][]int
	lineIndex map[int]string
	radixTree* radix.Tree
}

type FuzzySearcher struct {
	completeWorks* CompleteWorks
}

// searches for the closest key to the provided string
// returns a list of lines including the search result and a line number for later reference
func (fs FuzzySearcher) Search(toSearch string) [][]string {
	var matchingLines [][]string

	_, linesInterface, _ := fs.completeWorks.radixTree.LongestPrefix(toSearch)
	switch t := linesInterface.(type) {
		case []int:
			for _, value := range t {
				pair := []string{}
				pair = append(pair,
					fs.completeWorks.lineIndex[value-1] + "\n" +
					fs.completeWorks.lineIndex[value] + "\n" +
					fs.completeWorks.lineIndex[value+1])
				pair = append(pair, strconv.Itoa(value))
				matchingLines = append(matchingLines, pair)
			}
	}

	return matchingLines
}

// populates word, line, and radix tree indexes based on the provided text file filename
func (fs* FuzzySearcher) Load(filename string) error {

	fs.completeWorks = &CompleteWorks{
		wordIndex: map[string][]int{},
		lineIndex: map[int]string{},
		radixTree: radix.New()}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// scan over file and create indexes for words and lines
	lineCount := 1
	for scanner.Scan() {
		line := scanner.Text()
		fs.completeWorks.lineIndex[lineCount] = line

		// Make a Regex to say we only want letters and numbers
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			log.Fatal(err)
		}
		processedString := reg.ReplaceAllString(strings.ToLower(line), " ")
		words := strings.Fields(processedString)

		// create word to line index
		for _, word := range words {
			fs.completeWorks.wordIndex[word] = append(fs.completeWorks.wordIndex[word], lineCount)
		}

		lineCount++
	}
	// insert word to line index into radix tree
	for k, v := range fs.completeWorks.wordIndex {
		fs.completeWorks.radixTree.Insert(k, v)
	}

	return nil
}

// returns a larger context of lines relative to the line number provided
func (fs* FuzzySearcher) GetLineContext(lineNumber int) string {
	var values []string
	for i := 0; i < 51; i++ {
		values = append(values, fs.completeWorks.lineIndex[lineNumber-25+i])
	}
	context := strings.Join(values, "\n")

	return context
}