package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/deliveryhero/pipeline/v2"
)

func main() {

	transform := pipeline.NewProcessor(func(_ context.Context, s string) ([]string, error) {
		return strings.Split(s, ","), nil
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

	input := "1,2,3,4,5"
	//input := "1"

	for out := range pipeline.Process(context.Background(), apply, pipeline.Emit(input)) {
		for j := range out {
			fmt.Printf("process: %s\n", out[j])
		}
	}
}
