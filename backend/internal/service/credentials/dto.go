package credential

type CredentialDTO struct {
	ID          uint    `json:"id"`
	UserID      uint    `json:"user_id"`
	ServiceName string  `json:"service_name"`
	Username    string  `json:"username"`
	Password    string  `json:"password,omitempty"` // decrypted
	Label       *string `json:"label"`
	Notes       *string `json:"notes"`
}
