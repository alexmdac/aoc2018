package util

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadStrings(t *testing.T) {
	for _, c := range []struct {
		desc string
		in   string
		want []string
	}{
		{
			desc: "no newline at end",
			in: `foo
bar

baz`,
			want: []string{"foo", "bar", "", "baz"},
		},
		{
			desc: "newline at end",
			in: `foo
bar
`,
			want: []string{"foo", "bar"},
		},
	} {
		t.Run(c.desc, func(t *testing.T) {
			r := strings.NewReader(c.in)
			strs, err := ReadLines(r)
			require.NoError(t, err)
			assert.Equal(t, c.want, strs)
		})
	}
}

func TestReadInts(t *testing.T) {
	for _, c := range []struct {
		desc    string
		in      string
		want    []int
		wantErr bool
	}{
		{
			desc: "newline at end",
			in: `1
+2
-3
`,
			want: []int{1, 2, -3},
		},
		{
			desc: "no newline at end",
			in: `1
+2
-3`,
			want: []int{1, 2, -3},
		},
		{
			desc: "invalid input",
			in: `foo
1
2`,
			wantErr: true,
		},
	} {
		t.Run(c.desc, func(t *testing.T) {
			r := strings.NewReader(c.in)
			ints, err := ReadInts(r)
			if c.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, c.want, ints)
		})
	}
}
