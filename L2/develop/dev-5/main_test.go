package main

import (
	"bufio"
	"dev-5/config"
	"os"
	"testing"
)

func TestTask(t *testing.T) {
	testTable := []struct {
		input    config.Config
		expected []string
	}{
		{
			input: config.Config{
				A:         1,
				B:         0,
				C:         0,
				RegString: "рамки",
				Count:     false,
				I:         false,
				V:         false,
				F:         false,
				N:         false,
			},
			expected: []string{"A: [С другой стороны постоянное информационно-пропагандистское обеспечение нашей деятельности " +
				"обеспечивает широкому кругу (специалистов) участие в формировании позиций, занимаемых участниками в " +
				"отношении поставленных задач. Идейные соображениявысшего]"},
		},
		{
			input: config.Config{
				A:         0,
				B:         1,
				C:         0,
				RegString: "",
				Filename:  "",
				Count:     false,
				I:         false,
				V:         false,
				F:         false,
				N:         false,
			},
			expected: []string{"B: [С другой стороны постоянное информационно-пропагандистское обеспечение нашей " +
				"деятельности обеспечивает широкому кругу]"},
		},
		{
			input: config.Config{
				A:         0,
				B:         0,
				C:         1,
				RegString: "р",
				Filename:  "",
				Count:     false,
				I:         false,
				V:         false,
				F:         false,
				N:         false,
			},
			expected: []string{"C: [С другой стороны постоянное информационно-пропагандистское обеспечение нашей деятельности " +
				"обеспечивает широкому кругу (специалистов) участие в формировании позиций, занимаемых участниками в " +
				"отношении поставленных задач. Идейные соображениявысшего]"},
		},
		{
			input: config.Config{
				A:         0,
				B:         0,
				C:         0,
				RegString: "р",
				Filename:  "",
				Count:     true,
				I:         false,
				V:         true,
				F:         false,
				N:         false,
			},
			expected: []string{"Count -c -v (invert):-33"},
		},
		{
			input: config.Config{
				A:         0,
				B:         0,
				C:         0,
				RegString: "р",
				Filename:  "",
				Count:     true,
				I:         true,
				V:         false,
				F:         true,
				N:         true,
			},
			expected: []string{"Count -c:0", "Match nunbers -n:[]"},
		},
		{
			input: config.Config{
				A:         0,
				B:         0,
				C:         0,
				RegString: "",
				Filename:  "",
				Count:     true,
				I:         true,
				V:         false,
				F:         false,
				N:         true,
			},
			expected: []string{"Count -c:42", "Match nunbers -n:[1 2 3 4 5 6 7 8 9]"},
		},
	}

	f, err := os.Open("file.txt")
	if err != nil {
		t.Errorf("error with open file")
	}
	text := make([]string, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		t.Errorf("error with scanner")
	}
	for _, tt := range testTable {
		got := task(text, tt.input)
		if Equal(got, tt.expected) {
			t.Errorf("Expected %s, got %s", tt.expected, got)
		}
	}
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
