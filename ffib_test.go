/*
 * Тесты
 */

package fib_test

import (
	"./gRPC/grpcserver"
	"./gRPC/grpcserver/fibonacci"
	"./rest_http/restserver"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

//структура теста
type testC struct {
	name      string
	start     int32
	end       int32
	exp       map[int64]int64
	wantError bool
}

var testCases = []testC{
	{
		name:  "x=0, y=1",
		start: 0,
		end:   1,
		exp:   map[int64]int64{0: 0, 1: 1},
	},
	{
		name:  "x=0, y=4",
		start: 0,
		end:   4,
		exp:   map[int64]int64{0: 0, 1: 1, 2: 1, 3: 2, 4: 3},
	},
	{
		name:      "x=0, y=0",
		start:     0,
		end:       0,
		wantError: true,
	},

	{
		name:      "x=-1, y=0",
		start:     -1,
		end:       0,
		wantError: true,
	},
}

//Тестирование функции, возвращающей последовательность чисел фибоначчи между номерами
func TestFib(t *testing.T) {
	for _, testCase := range testCases {
		fib, err := fibonacci.Fibonacci(int64(testCase.start), int64(testCase.end))
		if err != nil && !testCase.wantError {
			t.Fatalf("Unexpected error: %s", err)
		} else {
			assert.Equal(t, testCase.exp, fib)
		}
	}
}

//Тестирование gRPC
func TestGRPS(t *testing.T) {
	s := GRPCapi.GRPCServer{}

	for _, testCase := range testCases {
		req := &GRPCapi.FibRequest{X: testCase.start, Y: testCase.end}
		fib, err := s.Get(context.Background(), req)
		if err != nil {
			if !testCase.wantError {
				t.Fatalf("Unexpected error: %s", err)
			}
		} else {
			assert.Equal(t, testCase.exp, fib.Result)
		}
	}
}

//Тестирование Rest
func TestREST(t *testing.T) {
	for _, testCase := range testCases {
		req, err := http.NewRequest(http.MethodPost, "/calc/?range="+strconv.Itoa(int(testCase.start))+","+strconv.Itoa(int(testCase.end)), nil)
		if err != nil && !testCase.wantError {
			t.Fatalf("Unexpected error: %s", err)
		}

		w := httptest.NewRecorder()
		RESTapi.CalcHandler(w, req)
		res := w.Result()
		defer res.Body.Close()

		file, err := os.Open("data.json")
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		var fibdata map[int64]int64
		json.NewDecoder(file).Decode(&fibdata)
		assert.Equal(t, testCase.exp, fibdata)
		file.Close()

		os.Remove("data.json")
		os.Create("data.json")
	}
}
