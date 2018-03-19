package sqlxdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleInStatementHappyPaths(t *testing.T) {
	type inputData struct {
		query string
		args  []interface{}
	}
	type expectedData struct {
		query string
		args  []interface{}
	}

	testCases := []struct {
		description string
		input       inputData
		expected    expectedData
	}{
		{
			description: "no bindvar at all",
			input: inputData{
				query: "SELECT * from my_table",
			},
			expected: expectedData{
				query: "SELECT * from my_table",
			},
		},
		{
			description: "no IN clause at all",
			input: inputData{
				query: "SELECT * FROM my_table WHERE id=$1 and x=$2",
				args:  []interface{}{"id", 2},
			},
			expected: expectedData{
				query: "SELECT * FROM my_table WHERE id=$1 and x=$2",
				args:  []interface{}{"id", 2},
			},
		},
		{
			description: "1 IN",
			input: inputData{
				query: "SELECT * FROM my_table WHERE id IN ($1)",
				args:  []interface{}{[]string{"id1", "id2"}},
			},
			expected: expectedData{
				query: "SELECT * FROM my_table WHERE id IN ($1, $2)",
				args:  []interface{}{"id1", "id2"},
			},
		},
		{
			description: "1 IN with only 1 value",
			input: inputData{
				query: "SELECT * FROM my_table WHERE id IN ($1)",
				args:  []interface{}{[]string{"id1"}},
			},
			expected: expectedData{
				query: "SELECT * FROM my_table WHERE id IN ($1)",
				args:  []interface{}{"id1"},
			},
		},
		{
			description: "1 IN - 1 $x",
			input: inputData{
				query: "SELECT * FROM my_table WHERE id IN ($1) and x=$2",
				args: []interface{}{
					[]string{"id1", "id2"}, "value",
				},
			},
			expected: expectedData{
				query: "SELECT * FROM my_table WHERE id IN ($1, $2) and x=$3",
				args:  []interface{}{"id1", "id2", "value"},
			},
		},
		{
			description: "1 $x - 1 IN",
			input: inputData{
				query: "SELECT * FROM my_table WHERE x=$1 AND id IN ($2)",
				args: []interface{}{
					"value", []string{"id1", "id2"},
				},
			},
			expected: expectedData{
				query: "SELECT * FROM my_table WHERE x=$1 AND id IN ($2, $3)",
				args:  []interface{}{"value", "id1", "id2"},
			},
		},
		{
			description: "1 $x - 1 IN (inversed)",
			input: inputData{
				query: "SELECT * FROM my_table WHERE x=$2 AND id IN ($1)",
				args: []interface{}{
					[]string{"id1", "id2"}, "value",
				},
			},
			expected: expectedData{
				query: "SELECT * FROM my_table WHERE x=$3 AND id IN ($1, $2)",
				args:  []interface{}{"id1", "id2", "value"},
			},
		},
		{
			description: "1 IN - 1 $x - 1 IN",
			input: inputData{
				query: "SELECT * FROM my_table WHERE id IN ($1) AND x=$2 AND y IN ($3)",
				args: []interface{}{
					[]string{"id1", "id2", "id3"}, "value", []int{1, 2},
				},
			},
			expected: expectedData{
				query: "SELECT * FROM my_table WHERE id IN ($1, $2, $3) AND x=$4 AND y IN ($5, $6)",
				args:  []interface{}{"id1", "id2", "id3", "value", 1, 2},
			},
		},
		{
			description: "1 IN - 1 $x - 1 IN (random pos)",
			input: inputData{
				query: "SELECT * FROM my_table WHERE id IN ($2) AND x=$3 AND y IN ($1)",
				args: []interface{}{
					[]int{1, 2}, []string{"id1", "id2", "id3"}, "value",
				},
			},
			expected: expectedData{
				query: "SELECT * FROM my_table WHERE id IN ($3, $4, $5) AND x=$6 AND y IN ($1, $2)",
				args:  []interface{}{1, 2, "id1", "id2", "id3", "value"},
			},
		},
		{
			description: "2 IN - 3 $x - 2 IN - 1 $x",
			input: inputData{
				query: "WHERE a IN ($1) AND b IN ($2) AND c=$3 AND d=$3 AND e=$4 AND f IN ($5) AND g IN ($6) AND h=$7",
				args: []interface{}{
					[]int{1}, []string{"2", "3", "4"}, 5, "6", []int{5}, []string{"6", "7", "8"}, "last",
				},
			},
			expected: expectedData{
				query: "WHERE a IN ($1) AND b IN ($2, $3, $4) AND c=$5 AND d=$5 AND e=$6 AND f IN ($7) AND g IN ($8, $9, $10) AND h=$11",
				args: []interface{}{
					1, "2", "3", "4", 5, "6", 5, "6", "7", "8", "last",
				},
			},
		},
		{
			description: "1 IN that is not a slice",
			input: inputData{
				query: "SELECT * FROM my_table WHERE id IN ($1)",
				args:  []interface{}{"id1"},
			},
			expected: expectedData{
				query: "SELECT * FROM my_table WHERE id IN ($1)",
				args:  []interface{}{"id1"},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			query, args, err := (&Queryable{}).handleInClauses(tc.input.query, tc.input.args)
			require.NoError(t, err, "handleInClauses() should have not returned an error")
			assert.Equal(t, tc.expected.query, query, "unexpected query returned")
			require.Len(t, args, len(tc.expected.args), "unexpected number of args returned")

			for i, val := range args {
				assert.Equal(t, tc.expected.args[i], val, "invalid value for arg %d", i)
			}
		})
	}
}
