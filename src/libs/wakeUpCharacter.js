/**
 * キャラクターを就寝させる
 * @returns {Promise<{success: boolean}>}
 */
export const sleepCharacter = async () => {
  // localStorageからcharacterIdを取得
  const characterId = localStorage.getItem('characterId');

  if (!characterId) {
    throw new Error('キャラクターIDが見つかりません');
  }

  // PATCH /character/{id}/sleep を呼び出し
  const response = await fetch(`/api/character/${characterId}/sleep`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }

  return {
    success: true,
  };
};

/**
 * キャラクターを起床させる
 * @returns {Promise<{jobId: string}>} ジョブID
 */
export const wakeUpCharacter = async () => {
  // localStorageからcharacterIdを取得
  const characterId = localStorage.getItem('characterId');

  if (!characterId) {
    throw new Error('キャラクターIDが見つかりません');
  }

  // PATCH /character/{id}/wake-up を呼び出し
  const response = await fetch(`/api/character/${characterId}/wake-up`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }

  const data = await response.json();
  return {
    jobId: data.jobId,
  };
};

/**
 * キャラクター情報を取得（就寝時間などを含む）
 * @returns {Promise<Object>} キャラクター情報
 */
export const getCharacter = async () => {
  // localStorageからcharacterIdを取得
  const characterId = localStorage.getItem('characterId');

  if (!characterId) {
    throw new Error('キャラクターIDが見つかりません');
  }

  // GET /character/{id} を呼び出し
  const response = await fetch(`/api/character/${characterId}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }

  const data = await response.json();
  return data;
};
