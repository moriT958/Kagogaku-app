/**
 * キャラクターにご飯を食べさせる
 * @param {string} food - 食べ物の名前（例: "ハンバーグ"）
 * @returns {Promise<{success: boolean}>}
 */
export const eatCharacter = async (food) => {
  // バリデーション
  if (!food || !food.trim()) {
    throw new Error('食べ物を指定してください');
  }

  // localStorageからcharacterIdを取得
  const characterId = localStorage.getItem('characterId');

  if (!characterId) {
    throw new Error('キャラクターIDが見つかりません');
  }

  // POST /character/{id}/eat を呼び出し
  const response = await fetch(`/api/character/${characterId}/eat`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      food: food.trim(),
    }),
  });

  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }

  return {
    success: true,
  };
};
