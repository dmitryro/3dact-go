module 3dact.com/m/v2

go 1.15

replace 3dact.com/actions => ./actions

replace 3dact.com/actions/models => ./actions/models

replace 3dact.com/actions/dao => ./actions/dao

replace 3dact.com/actions/api => ./actions/api

replace 3dact.com/blog => ./blog

replace 3dact.com/blog/models => ./blog/models

replace 3dact.com/blog/api => ./blog/api

replace 3dact.com/blog/dao => ./blog/dao

replace 3dact.com/channels => ./channels

replace 3dact.com/channels/models => ./channels/models

replace 3dact.com/channels/dao => ./channels/dao

replace 3dact.com/channels/api => ./channels/api

replace 3dact.com/chat => ./chat

replace 3dact.com/chat/models => ./chat/models

replace 3dact.com/chat/dao => ./chat/dao

replace 3dact.com/chat/api => ./chat/api

replace 3dact.com/config => ./config

replace 3dact.com/config/dbconnect => ./config/dbconnect

replace 3dact.com/crawler => ./crawler

replace 3dact.com/crawler/models => ./crawler/models

replace 3dact.com/crawler/dao => ./crawler/dao

replace 3dact.com/crawler/api => ./crawler/api

replace 3dact.com/events => ./events

replace 3dact.com/events/models => ./events/models

replace 3dact.com/events/dao => ./events/dao

replace 3dact.com/events/api => ./events/api

replace 3dact.com/log => ./log

replace 3dact.com/log/models => ./log/models

replace 3dact.com/log/dao => ./log/dao

replace 3dact.com/log/api => ./log/api

replace 3dact.com/messaging => ./messaging

replace 3dact.com/messaging/models => ./messaging/models

replace 3dact.com/messaging/dao => ./messaging/dao

replace 3dact.com/messaging/api => ./messaging/api

replace 3dact.com/metrics => ./metrics

replace 3dact.com/metrics/models => ./metrics/models

replace 3dact.com/metrics/dao => ./metrics/dao

replace 3dact.com/metrics/api => ./metrics/api

replace 3dact.com/settings => ./settings

replace 3dact.com/settings/models => ./settings/models

replace 3dact.com/settings/dao => ./settings/dao

replace 3dact.com/settings/api => ./settings/api

replace 3dact.com/user => ./user

replace 3dact.com/user/models => ./user/models

replace 3dact.com/user/dao => ./user/dao

replace 3dact.com/user/api => ./user/api

replace 3dact.com/workflows => ./workflows

replace 3dact.com/workflows/models => ./workflows/models

replace 3dact.com/workflows/dao => ./workflows/dao

replace 3dact.com/workflows/api => ./workflows/api

require (
	3dact.com/blog v0.0.0-00010101000000-000000000000
	3dact.com/blog/api v0.0.0-00010101000000-000000000000
	3dact.com/blog/dao v0.0.0-00010101000000-000000000000
	3dact.com/blog/models v0.0.0-00010101000000-000000000000
	3dact.com/config v0.0.0-00010101000000-000000000000
	3dact.com/config/dbconnect v0.0.0-00010101000000-000000000000 // indirect
	3dact.com/user v0.0.0-00010101000000-000000000000
	3dact.com/user/api v0.0.0-00010101000000-000000000000
	3dact.com/user/dao v0.0.0-00010101000000-000000000000
	3dact.com/user/models v0.0.0-00010101000000-000000000000
    3dact.com/workflows v0.0.0-00010101000000-000000000000
    3dact.com/workflows/api v0.0.0-00010101000000-000000000000
    3dact.com/workflows/dao v0.0.0-00010101000000-000000000000
    3dact.com/workflows/models v0.0.0-00010101000000-000000000000
    3dact.com/actions v0.0.0-00010101000000-000000000000
    3dact.com/actions/api v0.0.0-00010101000000-000000000000
    3dact.com/actions/dao v0.0.0-00010101000000-000000000000
    3dact.com/actions/models v0.0.0-00010101000000-000000000000
    3dact.com/events v0.0.0-00010101000000-000000000000
    3dact.com/events/api v0.0.0-00010101000000-000000000000
    3dact.com/events/dao v0.0.0-00010101000000-000000000000
    3dact.com/events/models v0.0.0-00010101000000-000000000000
    3dact.com/channels v0.0.0-00010101000000-000000000000
    3dact.com/channels/api v0.0.0-00010101000000-000000000000
    3dact.com/channels/dao v0.0.0-00010101000000-000000000000
    3dact.com/channels/models v0.0.0-00010101000000-000000000000
    3dact.com/crawler v0.0.0-00010101000000-000000000000
    3dact.com/crawler/api v0.0.0-00010101000000-000000000000
    3dact.com/crawler/dao v0.0.0-00010101000000-000000000000
    3dact.com/crawler/models v0.0.0-00010101000000-000000000000
    3dact.com/messaging v0.0.0-00010101000000-000000000000
    3dact.com/messaging/api v0.0.0-00010101000000-000000000000
    3dact.com/messaging/dao v0.0.0-00010101000000-000000000000
    3dact.com/messaging/models v0.0.0-00010101000000-000000000000
    3dact.com/metrics v0.0.0-00010101000000-000000000000
    3dact.com/metrics/api v0.0.0-00010101000000-000000000000
    3dact.com/metrics/dao v0.0.0-00010101000000-000000000000
    3dact.com/metrics/models v0.0.0-00010101000000-000000000000
    3dact.com/settings v0.0.0-00010101000000-000000000000
    3dact.com/settings/api v0.0.0-00010101000000-000000000000
    3dact.com/settings/dao v0.0.0-00010101000000-000000000000
    3dact.com/settings/models v0.0.0-00010101000000-000000000000
    3dact.com/chat v0.0.0-00010101000000-000000000000
    3dact.com/chat/api v0.0.0-00010101000000-000000000000
    3dact.com/chat/dao v0.0.0-00010101000000-000000000000
    3dact.com/chat/models v0.0.0-00010101000000-000000000000
    3dact.com/log v0.0.0-00010101000000-000000000000
    3dact.com/log/api v0.0.0-00010101000000-000000000000
    3dact.com/log/dao v0.0.0-00010101000000-000000000000
    3dact.com/log/models v0.0.0-00010101000000-000000000000
	github.com/callicoder/go-docker-compose v0.0.0-20190725022912-cfca182529bc
	github.com/go-chi/chi v1.5.0 // indirect
	github.com/go-chi/render v1.0.1 // indirect
	github.com/go-pg/pg/v10 v10.6.2
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-swagger/go-swagger v0.25.0 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/jackc/pgx v3.6.2+incompatible // indirect
	github.com/labstack/echo v3.3.10+incompatible // indirect
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
    github.com/gocolly/colly/v2 latest
	gorm.io/driver/postgres v1.0.5 // indirect
	gorm.io/gorm v1.20.7 // indirect
	xorm.io/xorm v1.0.5 // indirect
)
