package models

type User struct {
	BaseModel
	Name        string `json:"name" gorm:"not null"`          // 名前
	RealName    string `json:"real_name"`                     // 本名
	Email       string `json:"email" gorm:"unique;not null"`  // メールアドレス
	Password    string `json:"-" gorm:"not null"`             // パスワード (JSONには含めない)
	Icon        string `json:"icon"`                          // アイコン (URL or path)
	ProfileMemo string `json:"profile_memo" gorm:"type:text"` // プロフィールメモ
}
