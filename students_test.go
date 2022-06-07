package coverage

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

const (
	data        = "1 2 3 4 5\n6 7 8 9 10\n11 12 13 14 15\n16 17 18 19 20\n21 22 23 24 25"
	badData     = "1 2 3 4 5\n6 7 8 9 10\n11 12 13 14\n16 17 18 19 20\n21 22 23 24 25"
	badDataChar = "1 2 3 4 5\n6 7 8 9 10\n11 12 13 14 A\n16 17 18 19 20\n21 22 23 24 25"
	errBadData  = "bad data"
)

func TestCols(t *testing.T) {
	tests := []struct {
		name string
		data string
		want [][]int
	}{
		{
			"rows",
			data,
			[][]int{
				{1, 6, 11, 16, 21},
				{2, 7, 12, 17, 22},
				{3, 8, 13, 18, 23},
				{4, 9, 14, 19, 24},
				{5, 10, 15, 20, 25},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := New(tt.data)
			cols := m.Cols()
			assert.Nil(t, err, errBadData)
			assert.Equal(t, cols, tt.want, fmt.Errorf("cols() = %v, want %v", cols, tt.want))
		})
	}
}

func TestRows(t *testing.T) {
	tests := []struct {
		name string
		data string
		want [][]int
	}{
		{
			"cols",
			data,
			[][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := New(tt.data)
			rows := m.Rows()
			assert.Nil(t, err, errBadData)
			assert.Equal(t, tt.want, rows, fmt.Errorf("rows() = %v, want %v", rows, tt.want))
		})
	}
}

func TestSet(t *testing.T) {
	type args struct {
		row   int
		col   int
		value int
	}
	tests := []struct {
		name       string
		data       string
		args       args
		wantMatrix *Matrix
		want       bool
	}{
		{
			"ok",
			data,
			args{3, 3, 44},
			&Matrix{
				5,
				5,
				[]int{
					1, 2, 3, 4, 5,
					6, 7, 8, 9, 10,
					11, 12, 13, 14, 15,
					16, 17, 18, 44, 20,
					21, 22, 23, 24, 25,
				},
			},
			true,
		},
		{
			"not ok row",
			data,
			args{5, 3, 44},
			&Matrix{
				5,
				5,
				[]int{
					1, 2, 3, 4, 5,
					6, 7, 8, 9, 10,
					11, 12, 13, 14, 15,
					16, 17, 18, 19, 20,
					21, 22, 23, 24, 25,
				},
			},
			false,
		},
		{
			"not ok col",
			data,
			args{3, 5, 44},
			&Matrix{
				5,
				5,
				[]int{
					1, 2, 3, 4, 5,
					6, 7, 8, 9, 10,
					11, 12, 13, 14, 15,
					16, 17, 18, 19, 20,
					21, 22, 23, 24, 25,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := New(tt.data)
			assert.Nil(t, err, errBadData)
			got := m.Set(tt.args.row, tt.args.col, tt.args.value)
			assert.Equal(t, tt.want, got, fmt.Errorf("set() = %v, want %v", got, tt.want))

			assert.Equal(t, m, tt.wantMatrix)
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		want    *Matrix
		wantErr error
	}{
		{
			"5x5",
			data,
			&Matrix{
				5,
				5,
				[]int{
					1, 2, 3, 4, 5,
					6, 7, 8, 9, 10,
					11, 12, 13, 14, 15,
					16, 17, 18, 19, 20,
					21, 22, 23, 24, 25,
				},
			},
			nil,
		},
		{
			"bad data",
			badData,
			nil,
			errors.New("Rows need to be the same length"),
		},
		{
			"bad data char",
			badDataChar,
			nil,
			errors.New("invalid syntax"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.data)
			if e, ok := err.(*strconv.NumError); ok {
				err = e.Err
			}
			assert.Equal(t, err, tt.wantErr)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		name string
		p    People
		want int
	}{
		{
			"len = 3",
			People{
				{"Oleg", "Romanov", time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)},
				{"Sasha", "Ivanov", time.Date(2003, 2, 1, 12, 30, 0, 0, time.UTC)},
				{"Dima", "Kozlov", time.Date(1998, 2, 1, 12, 30, 0, 0, time.UTC)},
			},
			3,
		},
		{
			"len = 5",
			People{
				{"Oleg", "Romanov", time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)},
				{"Sasha", "Ivanov", time.Date(2003, 2, 1, 12, 30, 0, 0, time.UTC)},
				{"Dima", "Kozlov", time.Date(1998, 2, 1, 12, 30, 0, 0, time.UTC)},
				{"Sveta", "Petrova", time.Date(1988, 2, 1, 12, 30, 0, 0, time.UTC)},
				{"Lena", "Sidorova", time.Date(1998, 2, 1, 12, 30, 0, 0, time.UTC)},
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLess(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    People
		args args
		want bool
	}{
		{
			"true - lastName",
			People{
				{"Oleg", "Ivanov", time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)},
				{"Oleg", "Romanov", time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)},
			},
			args{0, 1},
			true,
		},
		{
			"false - lastName",
			People{
				{"Oleg", "Romanov", time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)},
				{"Oleg", "Ivanov", time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)},
			},
			args{0, 1},
			false,
		},
		{
			"true - firstName",
			People{
				{"Oleg", "Ivanov", time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)},
				{"Roma", "Ivanov", time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)},
			},
			args{0, 1},
			true,
		},
		{
			"false - firstName",
			People{
				{"Roma", "Ivanov", time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)},
				{"Oleg", "Ivanov", time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)},
			},
			args{0, 1},
			false,
		},
		{
			"true - birthDay",
			People{
				{"Oleg", "Romanov", time.Date(2003, 2, 1, 12, 30, 0, 0, time.UTC)},
				{"Oleg", "Romanov", time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)},
			},
			args{0, 1},
			true,
		},
		{
			"false - birthDay",
			People{
				{"Oleg", "Romanov", time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)},
				{"Oleg", "Romanov", time.Date(2003, 2, 1, 12, 30, 0, 0, time.UTC)},
			},
			args{0, 1},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    People
		args args
	}{
		{
			"swap ok",
			People{
				{"Oleg", "Romanov", time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)},
				{"Sasha", "Ivanov", time.Date(2003, 2, 1, 12, 30, 0, 0, time.UTC)},
			},
			args{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Swap(tt.args.i, tt.args.j)
		})
	}
}
