package main

import (
	"context"
	"fmt"

	"github.com/deliveryhero/pipeline/v2"
)

func main() {

	transform := pipeline.NewProcessor(func(_ context.Context, s []string) ([]string, error) {
		return s, nil
		//return strings.Split(s, ","), nil
	}, nil)

	double := pipeline.NewProcessor(func(_ context.Context, s string) (string, error) {
		return s + s, nil
	}, nil)

	addLeadingZero := pipeline.NewProcessor(func(_ context.Context, s string) (string, error) {
		return "0" + s, nil
	}, nil)

	apply := pipeline.Apply(
		transform,
		pipeline.Sequence(
			double,
			addLeadingZero,
			double,
		),
	)

	//input := "1,2,3,4,5"
	// where input must be a slice
	//var input = []string{"1", "2", "3", "4", "5"}
	var input = []string{"1"}
	//input := "1"

	for out := range pipeline.Process(context.Background(), apply, pipeline.Emit(input)) {
		for j := range out {
			fmt.Printf("process: %s\n", out[j])
		}
	}
}
