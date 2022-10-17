package dao

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"hogsback.com/whatsapp/pkg/config"
	"hogsback.com/whatsapp/pkg/dto"

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
	conf := config.GetConfig()
	dsn := fmt.Sprintf("postgres://%s:@%s:%s/%s", conf.User, conf.Host, conf.Port, conf.Dbname)

	psqlconn := pgdriver.NewConnector(pgdriver.WithDSN(dsn), pgdriver.WithPassword(conf.Password))
	psqldb := sql.OpenDB(psqlconn)

	db = bun.NewDB(psqldb, pgdialect.New())

	err := db.Ping()
	checkError(err)

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	createTable(context.TODO())
}

func createTable(ctx context.Context) {
	// _, _ = db.NewDropTable().Model((*Message)(nil)).Exec(ctx)
	res, err := db.NewCreateTable().Model((*Message)(nil)).IfNotExists().Exec(ctx)
	checkError(err)
	println(res)
}

func AddToTable(params dto.Parameter, ctx context.Context) {
	msg := &Message{
		Parameters: params,
	}
	res, err := db.NewInsert().Model(msg).Exec(ctx)

	checkError(err)
	println(res)
}

func BunDisconnect() {
	db.Close()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
