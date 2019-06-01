package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPadSpace(t *testing.T) {
	assert.Equal(t, []string{"あいうえお", "か　　　　"}, padSpace([]string{"あいうえお", "か"}))
	assert.Equal(t, []string{"あ　　　　", "かきくけこ"}, padSpace([]string{"あ", "かきくけこ"}))
	assert.Equal(t, []string{"a　　　　", "かきくけこ"}, padSpace([]string{"a", "かきくけこ"}))

	s := []string{"a", "かきくけこ"}
	padSpace(s)
	assert.Equal(t, []string{"a", "かきくけこ"}, s, "破壊的変更がされていないことの検証")
}

func TestReverse(t *testing.T) {
	assert.Equal(t, []string{"b", "a"}, reverse([]string{"a", "b"}))
	assert.Equal(t, []string{"a"}, reverse([]string{"a"}))
	assert.Equal(t, []string{}, reverse([]string{}))
}

func TestToVertical(t *testing.T) {
	assert.Equal(t, []string{"a", "b"}, toVertical([]string{"ab"}))
	assert.Equal(t, []string{"ca", "db"}, toVertical([]string{"ab", "cd"}))
	assert.Equal(t, []string{"うあ", "えい"}, toVertical([]string{"あい", "うえ"}))
	assert.Equal(t, []string{"｜あ", "｜い"}, toVertical([]string{"あい", "ーー"}), "伸ばし棒は縦棒に変換")
	assert.Equal(t, []string{"一あ", "｜い"}, toVertical([]string{"あい", "一ー"}), "漢字の一は変わらない")
	assert.Equal(t, []string{}, toVertical([]string{}))
}
