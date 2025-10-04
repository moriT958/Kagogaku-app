# バックエンド API の仕様

## エンドポイント

- `POST /character/new`: キャラクターを作成

リクエスト
```json
{
    "id": "uuid",
    "name": "character name",
    "appearance": "base64 image"
}
```

- `GET /character/{id}`: キャラクタ情報を取得 
  - id はフロントでローカルストレージに保存されている

- `PATCH /character/{id}/sleep`: 寝る
- `PATCH /character/{id}/wake-up`: 起きる 

- `GET /train-status/{jobId}`: キャラ画像変換ジョブの進捗確認 (完了 or 未完了)

- `POST /character/{id}/eat`: ご飯を食べる

リクエスト
```json
{
    "food": "ハンバーグ"
}
```