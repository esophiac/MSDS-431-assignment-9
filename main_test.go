package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	trimmed9 "github.com/esophiac/trimmed"
)

// test the sliceFloat function
func TestSliceFloat(t *testing.T) {
	rand.Seed(1)

	in1, in2 := 5, 5
	expected := []float64{3.023301439898098, 4.702545440225062, 3.322800266092452, 2.188570935934901, 2.1231874853563286}
	out := sliceFloat(in1, in2)

	if reflect.DeepEqual(expected, out) != true {
		t.Errorf("Expected %v, got %v", expected, out)
	}
}

// test the sliceInt function
func TestSliceInt(t *testing.T) {
	rand.Seed(1)

	in1, in2 := 5, 5
	expected := []float64{1, 2, 2, 4, 1}
	out := sliceInt(in1, in2)

	if reflect.DeepEqual(expected, out) != true {
		t.Errorf("Expected %v, got %v", expected, out)
	}
}

// testing the sampledSlice function in main.go
func TestSampledSlice(t *testing.T) {

	rand.Seed(1)

	in := []float64{5, 4, 3, 2, 1}
	expected := []float64{4, 3, 3, 1, 4}
	out := sampledSlice(in)

	if reflect.DeepEqual(expected, out) != true {
		t.Errorf("Expected %v, got %v", expected, out)
	}
}

// benchmarking the workflow for int values
func BenchmarkIntVals(b *testing.B) {

	testInt := sliceInt(100, 100)
	trimSlice := []float64{0.05}

	intChannel := make(chan []float64)
	go boots(intChannel, testInt, 100)
	testMean := bootsMean(intChannel, trimSlice)
	finalMean := trimmed9.Mean(testMean)
	fmt.Println("Bootstrap Sampled Mean, Ints: ", finalMean)

}

// benchmarking the workflow for float values
func BenchmarkFloatVals(b *testing.B) {

	testFloat := sliceFloat(100, 100)
	trimSlice := []float64{0.05}

	floatChannel := make(chan []float64)
	go boots(floatChannel, testFloat, 100)
	floatMean := bootsMean(floatChannel, trimSlice)
	finalFloat := trimmed9.Mean(floatMean)

	fmt.Println("Bootstrap Sampled Mean, Floats: ", finalFloat)

}
