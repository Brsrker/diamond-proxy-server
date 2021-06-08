package data

import (
	"context"

	"brsrker.com/diamond/proxyserver/internal/logger"
	"brsrker.com/diamond/proxyserver/pkg/clientapp"
)

type ClientAppRepository struct {
	Data *Data
}

const TAG = "clients_repository"

func (ur *ClientAppRepository) GetByClientCodeOrigin(ctx context.Context, clientCode string, appCode string) (clientapp.ClientApp, error) {

	q := `
	select
		clients."code"              as "clientCode",
		apps.name                   as "appName",
		apps.url                    as "url",
		apps."path"            		as "path",
		apps."localPath"            as "localPath",
		apps."localPort"            as "localPort",
		contracts."maxClientApp"    as "maxClientApp",
		contracts."maxBytesPkg"     as "maxBytesPkg",
		contracts."maxBytesMonth"   as "maxBytesMonth",
		contracts."maxSessions"     as "maxSessions",
		contracts."sessionCooldown" as "sessionCooldown",
		"appStats"."startAt"        as "startAt",
		"appStats"."endAt" 			as "endAt",
		"appStats"."bytesConsumed" 	as "bytesConsumed"
	from apps
			 left join "appStats" on "appStats"."appId" = apps.id
			 left join clients on clients.id = apps."clientId"
			 left join "contracts" on contracts.id = clients."contractId"
	where
			clients.code = $1 and
			apps.code = $2 and
			apps.enable = true and
			clients.enable = true and
			contracts.enable = true 
	limit 1;
`

	row := ur.Data.DB.QueryRowContext(ctx, q, clientCode, appCode)

	var c clientapp.ClientApp
	err := row.Scan(&c.ClientCode, &c.AppName, &c.Url, &c.Path, &c.LocalPath,
		&c.LocalPort, &c.MaxClientApp, &c.MaxBytesPkg, &c.MaxBytesMonth,
		&c.MaxSessions, &c.SessionCooldown, &c.StartAt, &c.EndAt, &c.BytesConsumed)
	if err != nil {
		logger.Error(TAG, err)
		return clientapp.ClientApp{}, err
	}

	return c, nil
}
