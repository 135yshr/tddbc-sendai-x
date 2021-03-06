package ichigo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIchigo(t *testing.T) {
	t.Run("品種にさちおとめとサイズにLをわたしたときに`さちおとめ: L`が取得できること", func(t *testing.T) {
		sut := Ichigo{Name: "さちおとめ", Size: "L"}
		assert.Equal(t, "さちおとめ: L", sut.String())
	})
}
