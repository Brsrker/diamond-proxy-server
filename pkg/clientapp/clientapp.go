package clientapp

import (
	"context"
	"database/sql"
)

type ClientApp struct {
	ClientCode      string         `json:"clientCode,omitempty"`
	AppName         string         `json:"appName,omitempty"`
	Url             string         `json:"url,omitempty"`
	Path            string         `json:"path,omitempty"`
	LocalPath       string         `json:"localPath,omitempty"`
	LocalPort       int            `json:"localPort,omitempty"`
	MaxClientApp    int            `json:"maxClientApp,omitempty"`
	MaxBytesPkg     int            `json:"maxBytesPkg,omitempty"`
	MaxBytesMonth   int            `json:"maxBytesMonth,omitempty"`
	MaxSessions     int            `json:"maxSessions,omitempty"`
	SessionCooldown string         `json:"sessionCooldown,omitempty"`
	StartAt         sql.NullTime   `json:"startAt,omitempty"`
	EndAt           sql.NullTime   `json:"endAt,omitempty"`
	BytesConsumed   sql.NullString `json:"bytesConsumed,omitempty"`
}

type Repository interface {
	GetByClientCodeOrigin(ctx context.Context, clientCode string, appCode string) (ClientApp, error)
}
