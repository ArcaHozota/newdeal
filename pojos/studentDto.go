package pojos

// StudentDTO は Java の StudentDto に相当する Go 構造体。
type StudentDTO struct {
	ID           string `json:"id"`           // ID
	LoginAccount string `json:"loginAccount"` // アカウント
	Username     string `json:"username"`     // ユーザ名称
	Password     string `json:"password"`     // パスワード
	Email        string `json:"email"`        // メール
	DateOfBirth  string `json:"dateOfBirth"`  // 生年月日
}
