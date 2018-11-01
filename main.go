package main

import (
	"fmt"
	"io/ioutil"

	"github.com/jbrukh/bayesian"
)

const (
	// Scam represents all scammy ICOs
	Scam bayesian.Class = "Scam"
	// NotScam represents all non-scammy ICOs
	NotScam bayesian.Class = "Not Scam"

	scamDirectoryLocation    = "abstracts/scam"
	notScamDirectoryLocation = "abstracts/notscam"
	testFileLocation         = "abstracts/test.txt"
)

var scamFiles []string
var notScamFiles []string

func main() {
	classifier := bayesian.NewClassifier(Scam, NotScam)

	enumerateClasses()
	learn(classifier)
	predict(testFileLocation, classifier)
}

func predict(location string, classifier *bayesian.Classifier) {
	probabilities, _, _ := classifier.ProbScores([]string{readFile(location)})

	fmt.Println(probabilities)
}

func learn(classifier *bayesian.Classifier) {
	classifier.Learn(scamFiles, Scam)
	classifier.Learn(notScamFiles, NotScam)
}

func enumerateClasses() {
	scamDirectory := enumerateDirectory(scamDirectoryLocation)

	for _, file := range scamDirectory {
		fileContent := readFile(file)
		scamFiles = append(scamFiles, fileContent)
	}

	notScamDirectory := enumerateDirectory(notScamDirectoryLocation)

	for _, file := range notScamDirectory {
		fileContent := readFile(file)
		notScamFiles = append(notScamFiles, fileContent)
	}
}

func enumerateDirectory(location string) []string {
	var files []string
	listing, err := ioutil.ReadDir(location)
	if err != nil {
		panic(err)
	}

	for _, l := range listing {
		files = append(files, location+"/"+l.Name())
	}

	return files
}

func readFile(location string) string {
	file, err := ioutil.ReadFile(location)
	if err != nil {
		panic(err)
	}

	return string(file)
}
