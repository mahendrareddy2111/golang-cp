package knapsackproblem

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKnapsackProblemCase1(t *testing.T) {
	expected := []interface{}{10, []int{1, 3}}
	output := KnapsackProblem([][]int{{1, 2}, {4, 3}, {5, 6}, {6, 7}}, 10)
	require.Equal(t, expected, output)
}

func TestKnapsackProblemCase2(t *testing.T) {
	expected := []interface{}{1500, []int{ 3,12,14}}
	output := KnapsackProblem([][]int{{465, 100}, {400, 85},{255, 55},{350, 45},{650,130},{1000,190},{455,100},{100,25},{1200,190},{320,65},{750,100},{50,45},{550,65},{100,50},{600,70},{240,40}}, 200)
	require.Equal(t, expected, output)
}

func TestKnapsackProblemCase3(t *testing.T) {
	expected := []interface{}{101, []int{0,2,3}}
	output := KnapsackProblem([][]int{{2,1},{70,70},{30,30},{69,69},{100,100}},100)
	require.Equal(t, expected, output)
}
