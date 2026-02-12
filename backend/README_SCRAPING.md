# Kakeibo Backend - スクレイピング機能

## 📚 概要

このプロジェクトには、Goで実装されたWebスクレイピング機能が含まれています。
Collyライブラリを使用して、Webページから商品情報などのデータを抽出できます。

## 🚀 セットアップ

### 依存関係のインストール

```bash
go get github.com/gocolly/colly/v2
```

### ビルド

```bash
go build -o main.exe ./cmd/api
```

### Docker で実行

```bash
docker compose up -d --build
```

## 📖 ドキュメント

詳細なスクレイピングガイドは以下を参照してください:

- [スクレイピング完全ガイド](./docs/scraping_guide.md)

## 🔧 使い方

### 1. API経由でスクレイピング

#### ガイドを取得

```bash
GET http://localhost:8080/api/scrape/guide
```

#### スクレイピング実行(カスタムセレクタ)

```bash
POST http://localhost:8080/api/scrape
Content-Type: application/json

{
  "url": "https://example.com/products",
  "item_selector": ".product-item",
  "name_selector": ".product-name",
  "price_selector": ".product-price",
  "link_selector": "a.product-link"
}
```

### 2. コードから直接使用

```go
package main

import (
    "fmt"
    "kakeibo-backend/scraper"
)

func main() {
    // スクレイパーを作成
    ps := scraper.NewProductScraper(nil)
    
    // スクレイピング実行
    products, err := ps.ScrapeExample("https://example.com/products")
    if err != nil {
        panic(err)
    }
    
    // 結果を表示
    for _, product := range products {
        fmt.Printf("%s: %s\n", product.Name, product.Price)
    }
}
```

## 📁 プロジェクト構成

```
backend/
├── scraper/
│   ├── product_scraper.go      # スクレイパー本体
│   ├── product_scraper_test.go # テスト
│   └── examples/
│       └── main.go             # 使用例
├── handlers/
│   └── scraper.go              # APIハンドラー
├── docs/
│   └── scraping_guide.md       # 詳細ガイド
└── cmd/api/
    └── main.go                 # エントリーポイント
```

## ⚠️ 注意事項

### 法的・倫理的注意点

1. **利用規約を確認**
   - スクレイピング対象サイトの利用規約を必ず確認してください
   - スクレイピング禁止のサイトには使用しないでください

2. **robots.txtを尊重**
   - `https://example.com/robots.txt` を確認
   - クロール禁止の範囲は避けてください

3. **サーバーに負荷をかけない**
   - 適切なリクエスト間隔を設定(デフォルト: 2秒)
   - 並行リクエスト数を制限(デフォルト: 1)

4. **個人情報の取り扱い**
   - 個人情報保護法に注意
   - 取得したデータの利用方法に注意

### 技術的制限

- **JavaScript動的コンテンツ**: 基本的なスクレイパーでは取得不可
  - 解決策: chromedpなどのヘッドレスブラウザを使用
- **CAPTCHA保護**: 突破困難
- **ログイン必要**: Cookie/セッション管理が必要

## 🧪 テスト

```bash
# ユニットテスト実行
go test ./scraper/...

# 例を実行
go run ./scraper/examples/main.go
```

## 📝 API エンドポイント

### GET /api/scrape/guide

スクレイピングAPIの使い方ガイドを返します。

**レスポンス例:**

```json
{
  "description": "Web scraping API for extracting product information",
  "endpoint": "POST /api/scrape",
  "example_request": {
    "url": "https://example.com/products",
    "item_selector": ".product-item",
    "name_selector": ".product-name",
    "price_selector": ".product-price"
  }
}
```

### POST /api/scrape

Webページから商品情報をスクレイピングします。

**リクエストボディ:**

```json
{
  "url": "https://example.com/products",
  "item_selector": ".product-item",
  "name_selector": ".product-name",
  "price_selector": ".product-price",
  "link_selector": "a.product-link",
  "allowed_domains": ["example.com"]
}
```

**レスポンス例:**

```json
{
  "success": true,
  "count": 10,
  "products": [
    {
      "name": "商品名",
      "price": "1,000円",
      "url": "https://example.com/product/1",
      "image_url": "https://example.com/image.jpg",
      "scraped_at": "2026-02-11T15:00:00Z"
    }
  ]
}
```

## 🔗 参考リンク

- [Colly公式ドキュメント](http://go-colly.org/)
- [goquery GitHub](https://github.com/PuerkitoBio/goquery)
- [スクレイピング完全ガイド](./docs/scraping_guide.md)

## 📄 ライセンス

このプロジェクトはMITライセンスの下で公開されています。
