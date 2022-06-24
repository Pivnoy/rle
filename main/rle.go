package main

import (
	"strconv"
	"strings"
	"unicode"
)

type RleInterface interface {
	encodeString(inputValue chan string, outputValue chan string)
	decodeString(inputValue chan string, outputValue chan string)
	RunLengthEncode(value []string) []string
	RunLengthDecode(value []string) []string
}

type RleDecoder struct {
}

func NewRleDecoder() *RleDecoder {
	return &RleDecoder{}
}

func (rl *RleDecoder) encodeString(inputValue chan string, outputValue chan string) {
	input := <-inputValue
	var result strings.Builder
	for len(input) > 0 {
		firstLetter := input[0]
		inputLength := len(input)
		input = strings.TrimLeft(input, string(firstLetter))
		if counter := inputLength - len(input); counter > 1 {
			result.WriteString(strconv.Itoa(counter))
		}
		result.WriteString(string(firstLetter))
	}
	outputValue <- result.String()
}

func (rl *RleDecoder) decodeString(inputValue chan string, outputValue chan string) {
	input := <-inputValue
	var result strings.Builder
	for len(input) > 0 {
		letterIndex := strings.IndexFunc(input, func(r rune) bool { return !unicode.IsDigit(r) })
		multiply := 1
		if letterIndex != 0 {
			multiply, _ = strconv.Atoi(input[:letterIndex])
		}
		result.WriteString(strings.Repeat(string(input[letterIndex]), multiply))
		input = input[letterIndex+1:]
	}
	outputValue <- result.String()
}

func (rl *RleDecoder) RunLengthEncode(value []string) []string {
	var result []string
	inputCh := make(chan string)
	outputCh := make(chan string)
	for _, str := range value {
		go rl.encodeString(inputCh, outputCh)
		inputCh <- str
		resultStr := <-outputCh
		result = append(result, resultStr)
	}
	return result
}

func (rl *RleDecoder) RunLengthDecode(value []string) []string {
	var result []string
	inputCh := make(chan string)
	outputCh := make(chan string)
	for _, str := range value {
		go rl.decodeString(inputCh, outputCh)
		inputCh <- str
		resultStr := <-outputCh
		result = append(result, resultStr)
	}
	return result
}
