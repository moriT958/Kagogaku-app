# バックエンド API の仕様

## エンドポイント


- `GET /sleep`: 寝る
- `GET /wake-up`: 起きる 

- `POST /eat`: ご飯を食べる

```json
{
    "food": "ハンバーグ"
}
```

- `POST /train`: キャラの体調更新 

```json
{
    "appearance": "base64 形式の画像"
}
```