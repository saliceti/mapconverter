package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// GenericMap is a generic map with keys as strings
type GenericMap map[string]interface{}

// ConverterConfig is a struct to store the map converter configuration
type ConverterConfig struct {
	PullFrom string
	DumpTo   string
	PushTo   string
}

func main() {
	config := getConfig()

	inputString := readFromOutput(config.PullFrom)

	inputMap := loadMapFromString(inputString)

	outputString := dumpMapToString(config.DumpTo, inputMap)

	pushToOutput(config.PushTo, outputString)
}

func getConfig() ConverterConfig {
	config := ConverterConfig{}
	flag.StringVar(&config.PullFrom, "l", "stdin", "Pull from input (ex: stdin)")
	flag.StringVar(&config.DumpTo, "d", "yaml", "Map dump format (ex: yaml, json)")
	flag.StringVar(&config.PushTo, "s", "stdout", "Push to output (ex: stdout)")
	flag.Parse()
	return config
}

// Pull
func readFromOutput(input string) string {
	var inputString string
	switch input {
	case "stdin":
		inputString = readFromStdin()
	default:
		log.Fatalf("Unknown input: %s", input)
	}
	return inputString
}

func readFromStdin() string {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}

// Load
func loadMapFromString(inputString string) GenericMap {
	inputMap, err := loadFromJSON(inputString)
	if err != nil {
		inputMap, err = loadFromYAML(inputString)
	}
	if err != nil {
		log.Fatal("Cannot decode input")
	}
	return inputMap
}

func loadFromYAML(inString string) (GenericMap, error) {
	out := make(GenericMap)
	err := yaml.Unmarshal([]byte(inString), &out)

	return out, err
}

func loadFromJSON(inString string) (GenericMap, error) {
	out := make(GenericMap)
	err := json.Unmarshal([]byte(inString), &out)

	return out, err
}

// Dump
func dumpMapToString(format string, inputMap GenericMap) string {
	var outputString string

	switch format {
	case "yaml":
		outputString = dumpToYAML(inputMap)
	case "json":
		outputString = dumpToJSON(inputMap)
	default:
		log.Fatalf("Unknown dump format: %s", format)
	}

	return outputString
}

func dumpToJSON(inputMap GenericMap) string {
	j, err := json.Marshal(inputMap)
	if err != nil {
		log.Fatal(err)
	}
	return string(j)
}

func dumpToYAML(inputMap GenericMap) string {
	j, err := yaml.Marshal(inputMap)
	if err != nil {
		log.Fatal(err)
	}
	return string(j)
}

// Push
func pushToOutput(output string, outputString string) {
	switch output {
	case "stdout":
		writeToStdout(outputString)
	default:
		log.Fatalf("Unknown output: %s", output)
	}
}

func writeToStdout(inputString string) {
	fmt.Println(inputString)
}
