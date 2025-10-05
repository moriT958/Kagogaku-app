/**
 * キャラクター画像変換ジョブの進捗を確認
 * @param {string} jobId - ジョブID
 * @returns {Promise<Object>} ジョブのステータス情報（完了 or 未完了）
 */
export const getTrainStatus = async (jobId) => {
  // バリデーション
  if (!jobId || !jobId.trim()) {
    throw new Error('ジョブIDを指定してください');
  }

  // GET /train-status/{jobId} を呼び出し
  const response = await fetch(`/api/train-status/${jobId}`, {
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
