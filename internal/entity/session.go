package entity

import "time"

type Session struct {
	ID               string
	UserId           string
	DeviceName       string
	DeviceId         string
	IPAddress        string `gorm:"ip_address"`
	Platform         string
	AccessToken      string
	RefreshToken     *string
	FCMToken         *string `gorm:"column:fcm_token"`
	GoogleOAuthToken *string `gorm:"column:google_oauth_token"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	User             User
}
