---
title: Testing and profiling go apps.
layout: post
published: true
category: programming
tags: [golang]
comments: true
---

## Profiling & Optimising in Go

A very simple application, we will use for demo. 

``` go

package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func Map(sentence string) []string {
	r := regexp.MustCompile("[^\\s]+")
	return r.FindAllString(sentence, -1)
}

func reduce(reverseWords []string) string {
	return strings.Join(reverseWords, " ")
}

func process(words []string) []string {
	nosOfWords := len(words)
	buffChannel := make(chan string, nosOfWords)
	var task sync.WaitGroup
	task.Add(nosOfWords)

	for _, word := range words {
		go func(word string) {
			defer task.Done()
			buffChannel <- reverse(word)
			counter++
		}(word)
	}
	task.Wait()
	close(buffChannel)

	rwords := make([]string, 0)
	for rword := range buffChannel {
		rwords = append(rwords, rword)
	}
	return rwords
}
```

### Run It!
    go run main.go
    or
    go build

## Test
Create a file and save it as `main_test.go`
Execute test case as 

	go test

Its time to add first test case, before we move ahead let me tell there are three types of test that we can write in golang

* Unit Test
* Benchmark Test
* Example Test

## Unit Test

Lets write first test case 

``` go
func TestMap(t *testing.T) {
	
	//Arrange
	phrase := "The quick brown fox jumps over the lazy dog"
	
	//Act
	slicedWord := Map(phrase)
	
	//Assert
	if len(slicedWord) != 9 {
		t.Log("Test failed, nos of words returned are incorrect")
		t.Fail()
	}
}
```

A short representaion dipicting effect of log and Fail/Failnow combinations

|   Log()    | Fail()/FailNow() | Net Effect  | Desc                                    |
| ---------- |:---------------- |------------ | ----------------------------------------|
| t.Log      |  t.Fail()        |  t.Error()  | signals failure                         |
| t.Log      |  t.FailNow()     |  t.Fatal()  | abort further execution of test cases   |



```go

func TestMapForBlankSentence(t *testing.T) {
	words := Map("")
	if len(words) != 0 {
		t.Error("TestMapForBlankSentence failed")
	}
}

func TestReverse(t *testing.T) {
	if reverse("Hello") != "olleH" {
		t.Fatal("Reverse is incorrect")
	}
}

```	
