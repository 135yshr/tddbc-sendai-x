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
			args: Ichigo{Variety: "とちおとめ", Size: "L"},
			want: "とちおとめ: L",
		},
		"品種にあまおうとサイズにLをわたしたときに`あまおう: L`が取得できる": {
			args: Ichigo{Variety: "あまおう", Size: "L"},
			want: "あまおう: L",
		},
		"品種にもういっことサイズにLをわたしたときに`もういっこ: L`が取得できる": {
			args: Ichigo{Variety: "もういっこ", Size: "L"},
			want: "もういっこ: L",
		},
		"品種にとちおとめとサイズにSをわたしたときに`とちおとめ: S`が取得できる": {
			args: Ichigo{Variety: "とちおとめ", Size: "S"},
			want: "とちおとめ: S",
		},
		"品種にとちおとめとサイズにMをわたしたときに`とちおとめ: M`が取得できる": {
			args: Ichigo{Variety: "とちおとめ", Size: "M"},
			want: "とちおとめ: M",
		},
		"品種にとちおとめとサイズにLLをわたしたときに`とちおとめ: LL`が取得できる": {
			args: Ichigo{Variety: "とちおとめ", Size: "LL"},
			want: "とちおとめ: LL",
		},
	}
	for tcName, tt := range tests {
		t.Run(tcName, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.args.String())
		})
	}
}
