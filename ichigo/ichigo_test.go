package ichigo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIchigo_String(t *testing.T) {
	tests := map[string]struct {
		args Ichigo
		want string
	}{
		"品種に`とちおとめ`とサイズに`L`をわたしたときに`とちおとめ: L`が取得できる": {
			args: Ichigo{Variety: "とちおとめ", Size: "L"},
			want: "とちおとめ: L",
		},
		"品種に`あまおう`とサイズに`L`をわたしたときに`あまおう: L`が取得できる": {
			args: Ichigo{Variety: "あまおう", Size: "L"},
			want: "あまおう: L",
		},
		"品種に`もういっこ`とサイズに`L`をわたしたときに`もういっこ: L`が取得できる": {
			args: Ichigo{Variety: "もういっこ", Size: "L"},
			want: "もういっこ: L",
		},
		"品種に`とちおとめ`とサイズに`S`をわたしたときに`とちおとめ: S`が取得できる": {
			args: Ichigo{Variety: "とちおとめ", Size: "S"},
			want: "とちおとめ: S",
		},
		"品種に`とちおとめ`とサイズに`M`をわたしたときに`とちおとめ: M`が取得できる": {
			args: Ichigo{Variety: "とちおとめ", Size: "M"},
			want: "とちおとめ: M",
		},
		"品種に`とちおとめ`とサイズに`LL`をわたしたときに`とちおとめ: LL`が取得できる": {
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

func TestIchigo_ConvertSizeFormWeight(t *testing.T) {
	t.Run("品種に`とちおとめ`と重さgに`8g`を渡したときに`とちおとめ: S`が取得できる", func(t *testing.T) {
		sut := NewWithVarietyAndWeight("とちおとめ", 8)
		assert.Equal(t, "とちおとめ: S", sut.String())
	})
	t.Run("品種に`とちおとめ`と重さgに`10g`を渡したときに`とちおとめ: M`が取得できる", func(t *testing.T) {
		sut := NewWithVarietyAndWeight("とちおとめ", 8)
		assert.Equal(t, "とちおとめ: M", sut.String())
	})
}
