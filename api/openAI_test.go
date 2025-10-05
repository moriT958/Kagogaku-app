package api

import (
	"log"
	"path"
	"testing"
)

// 手動で実行したいのでスキップ
func TestUpdateCharacterImage(t *testing.T) {
	t.Skip()
	t.Run("画像変換の動作確認", func(t *testing.T) {
		outImgPath, err := updateCharacterImage(path.Join("..", "images", "body-lotion.png"), "Create a lovely gift basket with these four items in it")
		if err != nil {
			t.Errorf("Failed to edit image: %v", err)
		}
		log.Println(outImgPath)
	})
}
