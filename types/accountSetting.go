package types

type AccountSetting struct {
	AccountUrl       string              `json:"account_url"`
	Email            string              `json:"email"`
	AlumPrivacy      string              `json:"album_privacy"`
	BlockedUsers     []map[string]string `json:"blocked_users"`
	ActiveEmails     []string            `json:"active_emails"`
	PublicImages     bool                `json:"public_images"`
	MessagingEnabled bool                `json:"messaging_enabled"`
	ShowMature       bool                `json:"show_mature"`
	FirstParty       bool                `json:"first_party"`
}

type RawAccountSetting struct {
	AccountSetting `json:"data"`
}
