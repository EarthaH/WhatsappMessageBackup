package dao

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"hogsback.com/whatsapp/pkg/config"
	"hogsback.com/whatsapp/pkg/dto"
	"hogsback.com/whatsapp/pkg/logger"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type Message struct {
	bun.BaseModel `bun:"table:messages,alias:m"`

	ID         int64 `bun:",pk,autoincrement"`
	Parameters dto.Parameter
}

var db *bun.DB

func BunConnect() {
	createDBConnection()
	pingDB()
	logger.Info("Successfully Connected to PostgreSQL DB.")

	addQueryHookToDB()

	createTable(context.TODO())
	logger.Info("DB Tables ready for use.")
}

func AddToTable(params dto.Parameter, ctx context.Context) sql.Result {
	msg := &Message{
		Parameters: params,
	}
	res, err := db.NewInsert().Model(msg).Exec(ctx)

	logger.HandleError(err, false)
	return res
}

func BunDisconnect() {
	db.Close()
	logger.Info("PostgreSQL DB Connection Closed.")
}

func createDBConnection() {
	conf := config.GetConfig()
	dsn := fmt.Sprintf("postgres://%s:@%s:%s/%s", conf.User, conf.Host, conf.Port, conf.Dbname)

	psqlconn := pgdriver.NewConnector(pgdriver.WithDSN(dsn), pgdriver.WithPassword(conf.Password))
	psqldb := sql.OpenDB(psqlconn)

	db = bun.NewDB(psqldb, pgdialect.New())
}

func pingDB() {
	err := db.Ping()
	logger.HandleError(err, true)
}

func addQueryHookToDB() {
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
}

func createTable(ctx context.Context) {
	_, err := db.NewCreateTable().Model((*Message)(nil)).IfNotExists().Exec(ctx)
	logger.HandleError(err, true)
}
