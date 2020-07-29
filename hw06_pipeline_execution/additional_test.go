package hw06_pipeline_execution //nolint:golint,stylecheck

import (
	"errors"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
	"time"
)

var stageError = errors.New("Error in stage")

func TestAdditional(t *testing.T) {
	// Stage generator
	g := func(name string, f func(v interface{}) interface{}) Stage {
		return func(in In) Out {
			out := make(Bi)
			go func() {
				defer close(out)
				for v := range in {
					time.Sleep(sleepPerStage)
					out <- f(v)
				}
			}()
			return out
		}
	}

	type Results struct {
		v   interface{}
		err error
	}

	stages := []Stage{
		g("Dummy", func(v interface{}) interface{} {
			return Results{v: v, err: nil}
		}),
		g("Divider (10/x)", func(v interface{}) interface{} {
			if v.(Results).v.(int) == 0 || v.(Results).err != nil {
				return Results{v: nil, err: stageError}
			}
			return Results{v: 10 / v.(Results).v.(int), err: nil}
		}),
		g("Multiplier (* 2)", func(v interface{}) interface{} {
			if v.(Results).err != nil {
				return Results{v: nil, err: stageError}
			}
			return Results{v: v.(Results).v.(int) * 2, err: nil}
		}),
		g("Adder (+ 100)", func(v interface{}) interface{} {
			if v.(Results).err != nil {
				return Results{v: nil, err: stageError}
			}
			return Results{v: v.(Results).v.(int) + 100, err: nil}
		}),
		g("Stringifier", func(v interface{}) interface{} {
			if v.(Results).err != nil {
				return Results{v: nil, err: stageError}
			}
			return Results{v: strconv.Itoa(v.(Results).v.(int)), err: nil}
		}),
	}

	t.Run("error case", func(t *testing.T) {
		in := make(Bi)
		data := []int{1, 0, 2}

		go func() {
			for _, v := range data {
				in <- v
			}
			close(in)
		}()

		result := make([]Results, 0, 10)
		start := time.Now()
		for s := range ExecutePipeline(in, nil, stages...) {
			result = append(result, s.(Results))
		}
		elapsed := time.Since(start)

		require.Equal(t, result, []Results{{"120", nil}, {nil, stageError}, {"110", nil}})
		require.Less(t,
			int64(elapsed),
			// ~0.8s for processing 5 values in 4 stages (100ms every) concurrently
			int64(sleepPerStage)*int64(len(stages)+len(data)-1)+int64(fault))
	})
}
