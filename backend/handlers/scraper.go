package handlers

import (
	"kakeibo-backend/models"
	"kakeibo-backend/scraper"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ScraperHandlers struct {
	DB *gorm.DB
}

// ScrapeRequest はスクレイピングリクエストの構造
type ScrapeRequest struct {
	URL            string   `json:"url" validate:"required,url"`
	ItemSelector   string   `json:"item_selector"`
	NameSelector   string   `json:"name_selector"`
	PriceSelector  string   `json:"price_selector"`
	LinkSelector   string   `json:"link_selector"`
	AllowedDomains []string `json:"allowed_domains"`
	Save           bool     `json:"save"`         // DBに保存するかどうか
	WebsiteName    string   `json:"website_name"` // サイト名（保存時用）
}

// ScrapeResponse はスクレイピング結果のレスポンス
type ScrapeResponse struct {
	Success  bool              `json:"success"`
	Count    int               `json:"count"`
	Products []scraper.Product `json:"products"`
	Message  string            `json:"message,omitempty"`
}

// ScrapeProducts はWebページから商品情報をスクレイピング
// POST /api/scrape
func (h *ScraperHandlers) ScrapeProducts(c echo.Context) error {
	var req ScrapeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request",
		})
	}

	// URLは必須
	if req.URL == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "url is required",
		})
	}

	// スクレイパー設定
	config := scraper.DefaultConfig()
	if len(req.AllowedDomains) > 0 {
		config.AllowedDomains = req.AllowedDomains
	}
	config.Delay = 2 * time.Second
	config.Parallelism = 1

	ps := scraper.NewProductScraper(config)

	var products []scraper.Product
	var err error

	// カスタムセレクタが指定されている場合
	if req.ItemSelector != "" {
		products, err = ps.ScrapeWithCustomSelector(
			req.URL,
			req.ItemSelector,
			req.NameSelector,
			req.PriceSelector,
			req.LinkSelector,
		)
	} else {
		// デフォルトのセレクタを使用
		products, err = ps.ScrapeExample(req.URL)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ScrapeResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	// DBへの保存処理
	if req.Save {
		for _, p := range products {
			item := models.ScrapedItem{
				URL:         p.URL,
				Name:        p.Name,
				Price:       p.Price,
				ImageURL:    p.ImageURL,
				Description: p.Description,
				WebsiteName: req.WebsiteName,
			}
			if err := h.DB.Create(&item).Error; err != nil {
				// エラーログを出力するが処理は続行
				// 実際の運用ではエラートラッキングシステムに通知するなど
				c.Logger().Errorf("Failed to save product: %v", err)
			}
		}
	}

	return c.JSON(http.StatusOK, ScrapeResponse{
		Success:  true,
		Count:    len(products),
		Products: products,
	})
}

// GetScrapingGuide はスクレイピングの使い方ガイドを返す
// GET /api/scrape/guide
func (h *ScraperHandlers) GetScrapingGuide(c echo.Context) error {
	guide := map[string]interface{}{
		"description": "Web scraping API for extracting product information",
		"endpoint":    "POST /api/scrape",
		"example_request": map[string]interface{}{
			"url":            "https://example.com/products",
			"item_selector":  ".product-item",
			"name_selector":  ".product-name",
			"price_selector": ".product-price",
			"link_selector":  "a.product-link",
			"save":           true,
			"website_name":   "Example Shop",
		},
		"notes": []string{
			"スクレイピング対象サイトの利用規約を必ず確認してください",
			"robots.txtを尊重してください",
			"サーバーに過度な負荷をかけないよう注意してください",
			"個人情報の取り扱いに注意してください",
		},
		"selectors_help": map[string]string{
			"item_selector":  "商品アイテム全体を囲む要素のCSSセレクタ",
			"name_selector":  "商品名のCSSセレクタ",
			"price_selector": "価格のCSSセレクタ",
			"link_selector":  "商品リンクのCSSセレクタ",
		},
	}

	return c.JSON(http.StatusOK, guide)
}
