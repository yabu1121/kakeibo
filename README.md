# 家計簿アプリ (Kakeibo)

Next.js + Go (Echo + GORM) で構築された家計簿管理アプリケーション

## 🚀 機能

- ✅ 消費記録の管理
- ✅ サブスクリプション管理
- ✅ カテゴリー分類
- ✅ ユーザー管理
- ✅ RESTful API
- 🔄 gRPC対応 (セットアップ中)

## 📁 プロジェクト構成

```
kakeibo/
├── backend/              # Go バックエンド
│   ├── cmd/api/         # エントリーポイント
│   ├── models/          # データモデル
│   ├── handlers/        # APIハンドラー
│   ├── proto/           # gRPC定義
│   ├── Dockerfile
│   └── docker-compose.yml
└── frontend/            # Next.js フロントエンド
    ├── app/            # App Router
    ├── lib/            # ユーティリティ
    └── .env.local
```

## 🛠️ 技術スタック

### バックエンド

- **言語**: Go 1.24
- **フレームワーク**: Echo v4
- **ORM**: GORM
- **データベース**: PostgreSQL 15
- **API**: REST + gRPC
- **コンテナ**: Docker

### フロントエンド

- **フレームワーク**: Next.js 15 (App Router)
- **言語**: TypeScript
- **スタイリング**: Tailwind CSS
- **アイコン**: Lucide React

## 🚀 クイックスタート

### 1. バックエンドの起動

```bash
cd backend
docker-compose up --build
```

バックエンドは `http://localhost:8080` で起動します。

#### エンドポイント

- `GET /health` - ヘルスチェック
- `GET /api/users` - ユーザー一覧
- `POST /api/users` - ユーザー作成
- `GET /api/categories` - カテゴリー一覧
- `POST /api/categories` - カテゴリー作成
- `GET /api/expenses` - 消費一覧
- `POST /api/expenses` - 消費記録
- `GET /api/subscriptions` - サブスク一覧
- `POST /api/subscriptions` - サブスク作成

### 2. フロントエンドの起動

```bash
cd frontend
npm install
npm run dev
```

フロントエンドは `http://localhost:3000` で起動します。

## 📊 データベーススキーマ

### Users (ユーザー)

- ID
- Name (名前)
- RealName (本名)
- Email (メールアドレス)
- Password (パスワード)
- Icon (アイコン)
- ProfileMemo (プロフィールメモ)

### Categories (カテゴリー)

- ID
- Name (カテゴリー名)

### Expenses (消費)

- ID
- UserID (ユーザーID)
- Amount (金額)
- Date (日付)
- CategoryID (カテゴリーID)

### Subscriptions (サブスク)

- ID
- UserID (ユーザーID)
- ProductName (商品名)
- CategoryID (カテゴリーID)
- Frequency (頻度: monthly/yearly)

## 🔧 開発

### バックエンド開発

```bash
cd backend

# 依存関係の追加
go get <package>

# モジュールの整理
go mod tidy

# ローカルで実行 (Dockerなし)
go run cmd/api/main.go
```

### フロントエンド開発

```bash
cd frontend

# 開発サーバー起動
npm run dev

# ビルド
npm run build

# 本番サーバー起動
npm start
```

## 🔐 環境変数

### バックエンド (docker-compose.yml)

```yaml
DB_HOST=db
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=kakeibo
DB_PORT=5432
PORT=8080
```

### フロントエンド (.env.local)

```
NEXT_PUBLIC_API_URL=http://localhost:8080/api
```

## 📝 gRPCのセットアップ

gRPCを使用する場合は、`GRPC_SETUP.md` を参照してください。

現在はRESTful APIが完全に動作しています。

## 🎨 UIの特徴

- **モダンなデザイン**: グラデーション、シャドウ、アニメーション
- **レスポンシブ**: モバイル・タブレット・デスクトップ対応
- **直感的なUX**: タブ切り替え、フォーム表示/非表示
- **視覚的なフィードバック**: ホバー効果、ローディング表示

## 📚 学習ポイント

### 1. **GORMのリレーション**

```go
type Expense struct {
    CategoryID uint
    Category   Category `gorm:"foreignKey:CategoryID"`
}
```

### 2. **Echoのミドルウェア**

```go
e.Use(middleware.Logger())
e.Use(middleware.Recover())
e.Use(middleware.CORS())
```

### 3. **Next.jsのApp Router**

- `app/page.tsx` - ルートページ
- `app/layout.tsx` - レイアウト
- `'use client'` - クライアントコンポーネント

### 4. **TypeScriptの型安全性**

```typescript
interface Expense {
  id: number;
  amount: number;
  date: string;
}
```

## 🐛 トラブルシューティング

### バックエンドが起動しない

```bash
# Dockerコンテナを削除して再起動
docker-compose down -v
docker-compose up --build
```

### フロントエンドがAPIに接続できない

- `.env.local` のURLを確認
- CORSが有効か確認 (バックエンドで設定済み)
- バックエンドが起動しているか確認

### データベースに接続できない

- PostgreSQLコンテナが起動しているか確認
- 環境変数が正しいか確認

## 📦 今後の拡張案

- [ ] ユーザー認証 (JWT)
- [ ] グラフ・チャート表示
- [ ] 予算設定機能
- [ ] エクスポート機能 (CSV, PDF)
- [ ] 通知機能
- [ ] ダークモード
- [ ] 多言語対応
- [ ] gRPCの完全実装

## 📄 ライセンス

MIT

## 👨‍💻 開発者

あなたの名前
