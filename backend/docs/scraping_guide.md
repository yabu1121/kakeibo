# Goスクレイピング完全ガイド

## 📖 目次
1. [スクレイピングの基礎](#スクレイピングの基礎)
2. [できること・できないこと](#できること・できないこと)
3. [主要ライブラリ比較](#主要ライブラリ比較)
4. [実装パターン](#実装パターン)
5. [ベストプラクティス](#ベストプラクティス)

---

## スクレイピングの基礎

### スクレイピングとは?
Webページから自動的にデータを抽出する技術です。

### 基本的な流れ
```
1. HTTPリクエストを送信
   ↓
2. HTMLレスポンスを取得
   ↓
3. HTMLを解析(パース)
   ↓
4. 必要なデータを抽出
   ↓
5. データを保存/処理
```

---

## できること・できないこと

### ✅ できること

#### 1. 静的HTMLからのデータ抽出
```go
// 例: ニュースサイトから記事タイトルを取得
titles := doc.Find("h2.article-title").Each(func(i int, s *goquery.Selection) {
    fmt.Println(s.Text())
})
```

**具体例:**
- ニュース記事のタイトル、本文、公開日
- ECサイトの商品名、価格、在庫状況
- 不動産サイトの物件情報
- 天気予報データ
- 株価情報

#### 2. 複数ページの巡回
```go
// ページネーション対応
for page := 1; page <= 10; page++ {
    url := fmt.Sprintf("https://example.com/items?page=%d", page)
    // スクレイピング処理
}
```

#### 3. データの構造化
```go
type Product struct {
    Name  string  `json:"name"`
    Price float64 `json:"price"`
    URL   string  `json:"url"`
}
```

#### 4. 定期実行・監視
- Cronジョブで定期的に実行
- 価格変動の追跡
- 新着情報の通知

### ❌ できないこと/難しいこと

#### 1. JavaScript動的コンテンツ
**問題:**
```html
<!-- ページロード時は空 -->
<div id="products"></div>

<!-- JavaScriptで後から追加される -->
<script>
  fetch('/api/products').then(data => {
    document.getElementById('products').innerHTML = data;
  });
</script>
```

**解決策:**
- ヘッドレスブラウザ(chromedp)を使用
- APIを直接叩く(開発者ツールでネットワークタブを確認)

#### 2. CAPTCHA・bot対策
- reCAPTCHA
- Cloudflare保護
- IP制限

**対策:**
- 基本的に突破は困難
- 公式APIの利用を検討

#### 3. ログイン認証
**問題:**
- Cookie/セッション管理が必要
- CSRF トークン対応

**解決策:**
```go
// Cookieを保持するHTTPクライアント
jar, _ := cookiejar.New(nil)
client := &http.Client{Jar: jar}
```

#### 4. 頻繁な構造変更
- サイトのHTML構造が変わるとスクレイパーも修正が必要
- 定期的なメンテナンスが必要

---

## 主要ライブラリ比較

### 1. Colly (推奨)
```go
import "github.com/gocolly/colly/v2"
```

**特徴:**
- ✅ 初心者に優しい
- ✅ 並行処理対応
- ✅ リトライ、キャッシュ機能内蔵
- ✅ robots.txt自動チェック

**使用例:**
```go
c := colly.NewCollector()

c.OnHTML("h1", func(e *colly.HTMLElement) {
    fmt.Println(e.Text)
})

c.Visit("https://example.com")
```

### 2. goquery
```go
import "github.com/PuerkitoBio/goquery"
```

**特徴:**
- ✅ jQueryライクなセレクタ
- ✅ 軽量
- ❌ HTTP処理は別途必要

**使用例:**
```go
doc, _ := goquery.NewDocument("https://example.com")
doc.Find("a").Each(func(i int, s *goquery.Selection) {
    href, _ := s.Attr("href")
    fmt.Println(href)
})
```

### 3. chromedp (動的コンテンツ用)
```go
import "github.com/chromedp/chromedp"
```

**特徴:**
- ✅ JavaScript実行可能
- ✅ スクリーンショット取得
- ❌ 重い(Chromeを起動)
- ❌ 複雑

**使用例:**
```go
ctx, cancel := chromedp.NewContext(context.Background())
defer cancel()

var html string
chromedp.Run(ctx,
    chromedp.Navigate("https://example.com"),
    chromedp.WaitVisible("#content"),
    chromedp.OuterHTML("body", &html),
)
```

---

## 実装パターン

### パターン1: シンプルなスクレイピング
```go
package main

import (
    "fmt"
    "log"
    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector()

    // HTMLセレクタでデータ抽出
    c.OnHTML("h2.title", func(e *colly.HTMLElement) {
        fmt.Println("Title:", e.Text)
    })

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL)
    })

    c.Visit("https://example.com")
}
```

### パターン2: データを構造化して保存
```go
type Article struct {
    Title   string
    URL     string
    Date    string
}

var articles []Article

c.OnHTML("article", func(e *colly.HTMLElement) {
    article := Article{
        Title: e.ChildText("h2"),
        URL:   e.ChildAttr("a", "href"),
        Date:  e.ChildText("time"),
    }
    articles = append(articles, article)
})
```

### パターン3: ページネーション対応
```go
c.OnHTML("a.next-page", func(e *colly.HTMLElement) {
    nextPage := e.Attr("href")
    c.Visit(e.Request.AbsoluteURL(nextPage))
})
```

### パターン4: レート制限
```go
c := colly.NewCollector(
    colly.Async(true),
)

// 並行リクエスト数を制限
c.Limit(&colly.LimitRule{
    DomainGlob:  "*",
    Parallelism: 2,
    Delay:       1 * time.Second, // 1秒待機
})
```

---

## ベストプラクティス

### 1. User-Agentを設定
```go
c := colly.NewCollector(
    colly.UserAgent("Mozilla/5.0 (compatible; MyBot/1.0)"),
)
```

### 2. エラーハンドリング
```go
c.OnError(func(r *colly.Response, err error) {
    log.Printf("Request URL: %s failed with response: %v\nError: %v", 
        r.Request.URL, r, err)
})
```

### 3. リトライ設定
```go
c.OnError(func(r *colly.Response, err error) {
    if r.StatusCode == 429 { // Too Many Requests
        time.Sleep(5 * time.Second)
        r.Request.Retry()
    }
})
```

### 4. robots.txtを尊重
```go
c := colly.NewCollector(
    colly.AllowedDomains("example.com"),
    colly.IgnoreRobotsTxt(false), // robots.txtをチェック
)
```

### 5. キャッシュ利用
```go
c := colly.NewCollector(
    colly.CacheDir("./cache"),
)
```

---

## セキュリティ・法的注意事項

### ⚠️ 必ず確認すること

1. **利用規約を読む**
   - スクレイピング禁止の記載がないか確認

2. **robots.txtを確認**
   ```
   https://example.com/robots.txt
   ```

3. **APIの有無を確認**
   - 公式APIがあればそちらを使用

4. **個人情報の取り扱い**
   - 個人情報保護法に注意

5. **サーバー負荷**
   - 適切な間隔でリクエスト
   - 並行数を制限

### 推奨設定
```go
c.Limit(&colly.LimitRule{
    DomainGlob:  "*example.com",
    Parallelism: 1,              // 同時リクエスト数
    Delay:       2 * time.Second, // 2秒間隔
    RandomDelay: 1 * time.Second, // ランダム遅延
})
```

---

## トラブルシューティング

### 問題1: データが取得できない
**原因:**
- JavaScriptで動的生成されている
- セレクタが間違っている

**解決:**
1. ブラウザの開発者ツールで要素を確認
2. chromedpを使用
3. APIを直接叩く

### 問題2: 403 Forbidden
**原因:**
- User-Agentがない
- bot判定されている

**解決:**
```go
c.OnRequest(func(r *colly.Request) {
    r.Headers.Set("User-Agent", "Mozilla/5.0...")
    r.Headers.Set("Referer", "https://example.com")
})
```

### 問題3: 文字化け
**原因:**
- 文字エンコーディングの問題

**解決:**
```go
import "golang.org/x/text/encoding/japanese"

decoder := japanese.ShiftJIS.NewDecoder()
utf8Text, _ := decoder.String(shiftJISText)
```

---

## 参考リンク

- [Colly公式ドキュメント](http://go-colly.org/)
- [goquery GitHub](https://github.com/PuerkitoBio/goquery)
- [chromedp GitHub](https://github.com/chromedp/chromedp)
