module 3dact.com/m/v2

go 1.15

replace 3dact.com/blog => ./blog

replace 3dact.com/blog/models => ./blog/models

replace 3dact.com/blog/api => ./blog/api

replace 3dact.com/blog/dao => ./blog/dao

replace 3dact.com/config => ./config

replace 3dact.com/config/dbconnect => ./config/dbconnect

require (
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
	gorm.io/driver/postgres v1.0.5 // indirect
	gorm.io/gorm v1.20.7 // indirect
	3dact.com/blog v0.0.0-00010101000000-000000000000
	3dact.com/blog/api v0.0.0-00010101000000-000000000000
	3dact.com/blog/dao v0.0.0-00010101000000-000000000000
	3dact.com/blog/models v0.0.0-00010101000000-000000000000
	3dact.com/config v0.0.0-00010101000000-000000000000
	xorm.io/xorm v1.0.5 // indirect
)
