package main

import (
	"fmt"
	"math/rand"

	trimmed9 "github.com/esophiac/trimmed"
	"github.com/seehuhn/mt19937"
)

// define an interface to hold all of the types of number you can use in the functions
type AllNum interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
}

// create a slice of random floats that is lngth long
func sliceFloat(lngth int, max int) []float64 {

	var result []float64

	for i := 0; i < lngth; i++ {
		result = append(result, rand.Float64()*float64(max)) // Generates random integers between 0 and max
	}

	return result
}

// create a slice of random ints that is lngth long
func sliceInt(lngth int, max int) []float64 {

	var result []float64

	for i := 0; i < lngth; i++ {
		result = append(result, float64(rand.Intn(max))) // Generates random integers between 0 and max
	}

	return result
}

// given a slice of numbers, create a new slice that is a sample with replacement
// the same size as the initial slice
func sampledSlice[T AllNum](inital []T) (final []T) {

	sampleRange := len(inital)

	for i := 0; i < sampleRange; i++ {
		final = append(final, inital[rand.Intn(sampleRange)])
	}

	return final
}

// Resample the data with replacement n times and send with a channel
func boots[T AllNum](ch chan []T, fSlice []T, n int) {

	for i := 0; i < n; i++ {

		// create a new slice with replacement
		sample := sampledSlice(fSlice)

		// add it to the result
		ch <- sample
	}
	close(ch)
}

// Compute median n times to generate a distribution of estimated statistics
// Receivedd channel
func bootsMean(ch chan []float64, trimVals []float64) (meanSlice []float64) {

	for data := range ch {

		newData := trimmed9.TrimMean(data, trimVals)
		meanSlice = append(meanSlice, newData)
	}

	return meanSlice
}

func main() {

	rand.New(mt19937.New())

	// running the program to create samples of at least 100 integers with symmetric trimming results (0.05)
	// with bootstrapping

	testInt := sliceInt(100, 100)
	trimSlice := []float64{0.05}

	newInt := trimmed9.TrimMean(testInt, trimSlice)
	fmt.Println("base Trimmed Mean, Ints: ", newInt)

	intChannel := make(chan []float64)
	go boots(intChannel, testInt, 100)
	testMean := bootsMean(intChannel, trimSlice)
	finalMean := trimmed9.Mean(testMean)

	fmt.Println("Bootstrap Sampled Mean, Ints: ", finalMean)

	// running the program to create samples of at least 100 floats with symmetric trimming results (0.05)
	// with bootstrapping. Uses the same trimSlice as the Int test.
	testFloat := sliceFloat(100, 100)

	newFloat := trimmed9.TrimMean(testFloat, trimSlice)
	fmt.Println("base Trimmed Mean, Floats: ", newFloat)

	floatChannel := make(chan []float64)
	go boots(floatChannel, testFloat, 100)
	floatMean := bootsMean(floatChannel, trimSlice)
	finalFloat := trimmed9.Mean(floatMean)

	fmt.Println("Bootstrap Sampled Mean, Floats: ", finalFloat)

}
