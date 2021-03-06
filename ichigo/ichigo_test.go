package ichigo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIchigo(t *testing.T) {
	t.Run("品種にとちおとめとサイズにLをわたしたときに`とちおとめ: L`が取得できること", func(t *testing.T) {
		sut := Ichigo{Name: "とちおとめ", Size: "L"}
		assert.Equal(t, "とちおとめ: L", sut.String())
	})

	t.Run("品種にあまおうとサイズにLをわたしたときに`あまおう: L`が取得できること", func(t *testing.T) {
		sut := Ichigo{Name: "あまおう", Size: "L"}
		assert.Equal(t, "あまおう: L", sut.String())
	})

	t.Run("品種にもういっことサイズにLをわたしたときに`もういっこ: L`が取得できること", func(t *testing.T) {
		sut := Ichigo{Name: "もういっこ", Size: "L"}
		assert.Equal(t, "もういっこ: L", sut.String())
	})
}
