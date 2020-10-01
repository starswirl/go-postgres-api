# go-postgres-api

go + gin + GORM + postgresのDocker環境



## 立ち上げ方

1. [Go環境構築](doc/Go環境構築.md)を読む
2. .env.sampleを参考に.env作る

3. 下記コマンド

```bash
cd server
go get
cd ..
docker-compose up
```

8080ポートをローカルに転送
http://localhost:8080/usersにGETリクエスト


## フォーマット

```
cd server
go fmt ./...
```

## コミットテンプレート

```emoji
# ==================== Emojis ====================
# 🎉  :tada: 初めてのコミット（Initial Commit）
# 🔖  :bookmark: バージョンタグ（Version Tag）
# ✨  :sparkles: 新機能（New Feature）
# 🐛  :bug: バグ修正（Bagfix）
# ♻️  :recycle: リファクタリング(Refactoring)
# 📚  :books: ドキュメント（Documentation）
# 🎨  :art: デザインUI/UX(Accessibility)
# 🐎  :horse: パフォーマンス（Performance）
# 🔧  :wrench: ツール（Tooling）
# 🚨  :rotating_light: テスト（Tests）
# 💩  :hankey: 非推奨追加（Deprecation）
# 🗑️  :wastebasket: 削除（Removal）
# 🚧  :construction: WIP(Work In Progress)
```
