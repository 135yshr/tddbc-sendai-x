package ichigo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIchigo(t *testing.T) {
	tests := map[string]struct {
		args Ichigo
		want string
	}{
		"品種にとちおとめとサイズにLをわたしたときに`とちおとめ: L`が取得できる": {
			args: Ichigo{Name: "とちおとめ", Size: "L"},
			want: "とちおとめ: L",
		},
		"品種にあまおうとサイズにLをわたしたときに`あまおう: L`が取得できる": {
			args: Ichigo{Name: "あまおう", Size: "L"},
			want: "あまおう: L",
		},
		"品種にもういっことサイズにLをわたしたときに`もういっこ: L`が取得できる": {
			args: Ichigo{Name: "もういっこ", Size: "L"},
			want: "もういっこ: L",
		},
		"品種にとちおとめとサイズにSをわたしたときに`とちおとめ: S`が取得できる": {
			args: Ichigo{Name: "とちおとめ", Size: "S"},
			want: "とちおとめ: S",
		},
	}
	for tcName, tt := range tests {
		t.Run(tcName, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.args.String())
		})
	}
}
