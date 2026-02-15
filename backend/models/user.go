package models

// User ユーザーモデル
type User struct {
	BaseModel
	Name        string `json:"name" gorm:"not null"`
	Email       string `json:"email" gorm:"unique;not null"`
	Password    string `json:"-" gorm:"not null"` // パスワードは絶対に出さない
	Icon        string `json:"icon"`
	ProfileMemo string `json:"profile_memo" gorm:"type:text"`

	// User has many Expenses
	// User側にIDを持たせるのではなく、リレーションとして定義する
	Expenses      []Expense      `json:"expenses"`
	Subscriptions []Subscription `json:"subscriptions"`
}