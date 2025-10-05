export const convertImageToBase64 = (file) => {
  return new Promise((resolve, reject) => {
    if (!file) {
      reject(new Error('ファイルが選択されていません'));
      return;
    }

    const reader = new FileReader();
    reader.onload = (e) => {
      resolve(e.target.result);
    };
    reader.onerror = (error) => {
      reject(error);
    };
    reader.readAsDataURL(file);
  });
};

/**
 * キャラクターを新規作成します
 * @param {string} name - キャラクター名
 * @param {string} base64Image - Base64形式の画像データ
 * @returns {Promise<{id: string, success: boolean}>} キャラクターIDと成功フラグ
 */
export const createCharacter = async (name, base64Image) => {
  // 入力チェック
  if (!name || !name.trim()) {
    throw new Error('キャラクター名を入力してください');
  }
  if (!base64Image) {
    throw new Error('画像を選択してください');
  }

  // UUID生成
  const characterId = crypto.randomUUID();

  // APIリクエスト
  const response = await fetch('/api/character/new', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      id: characterId,
      name: name.trim(),
      appearance: base64Image,
    }),
  });

  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }

  // localStorageに保存
  localStorage.setItem('characterId', characterId);

  return {
    id: characterId,
    success: true,
  };
};
