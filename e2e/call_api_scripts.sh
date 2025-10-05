#!/bin/bash

curl -s -X POST "http://localhost:8080/character/new" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "test-character-001",
    "name": "テスト太郎",
    "appearance": "ZWF0dGVzdA=="
  }' 


curl -s -X GET "http://localhost:8080/character/test-character-001"

curl -s -X PATCH "http://localhost:8080/character/test-character-001/sleep"

curl -s -X PATCH "http://localhost:8080/character/test-character-001/wake-up"


curl -s -X POST "http://localhost:8080/character/test-character-001/eat" \
  -H "Content-Type: application/json" \
  -d '{"food": "ラーメン"}'

curl -s -X GET "http://localhost:8080/train-status/1"