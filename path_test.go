package jsonast_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/berquerant/jsonast"
)

func TestPath(t *testing.T) {
	for _, tc := range []struct {
		path jsonast.Path
		want string
	}{
		{
			want: ".",
			path: jsonast.NewPath(),
		},
		{
			want: ".x",
			path: jsonast.NewPath(
				jsonast.SimplePath("x"),
			),
		},
		{
			want: ".x[1]",
			path: jsonast.NewPath(
				jsonast.SimplePath("x"),
				jsonast.NewIndexPath(1),
			),
		},
		{
			want: ".x.y.z",
			path: jsonast.NewPath(
				jsonast.SimplePath("x"),
				jsonast.SimplePath("y"),
				jsonast.SimplePath("z"),
			),
		},
		{
			want: `.x.y["z"]`,
			path: jsonast.NewPath(
				jsonast.SimplePath("x"),
				jsonast.SimplePath("y"),
				jsonast.NewIndexPath("z"),
			),
		},
	} {
		t.Run(tc.want, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.path.AsPath())
		})
	}
}
