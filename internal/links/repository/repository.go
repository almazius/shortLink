package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"links/internal/links"
	"log"
	"os"
	"time"
)

type Postgres struct {
	Pool *pgxpool.Pool
	Log  *log.Logger
}

func NewPostgres(pool *pgxpool.Pool) links.DatabaseService {
	return &Postgres{
		Pool: pool,
		Log:  log.New(os.Stdout, "Postgres ", log.LstdFlags|log.Lshortfile),
	}
}

func (p *Postgres) ExistShortLink(ctx context.Context, shortLink string) (bool, *links.MyError) {
	res := false
	err := p.Pool.QueryRow(ctx, `select exists(select * from links where shortLink=$1)`, shortLink).Scan(&res)
	if err != nil {
		p.Log.Print(err)
		return false, &links.MyError{
			Code: 500,
			Err:  err,
		}
	}
	return res, nil
}

func (p *Postgres) GetShortLink(ctx context.Context, link string) (string, *links.MyError) {
	short := ""
	err := p.Pool.QueryRow(ctx, `select shortLink from links where exists(select * from links where link=$1) and link=$1`, link).Scan(&short)
	if err != nil {
		if err != pgx.ErrNoRows {
			p.Log.Print(err)
		}
		return "", &links.MyError{Code: 500, Err: err}
	}

	return short, nil
}

func (p *Postgres) AddNote(ctx context.Context, link, shortLink string, time time.Time) (int64, *links.MyError) {
	var id int64
	err := p.Pool.QueryRow(ctx, `INSERT INTO links (link, shortLink, creationTime) values ($1, $2, $3) returning id`, link, shortLink, time).Scan(&id)
	if err != nil {
		p.Log.Print(err)

		return -1, &links.MyError{Code: 500, Err: err}
	}
	return id, nil
}

func (p *Postgres) GetFullLink(ctx context.Context, shortLink string) (string, *links.MyError) {
	fullLink := ""
	err := p.Pool.QueryRow(ctx, `select link from links where exists(select * from links where shortLink=$1) and shortLink=$1`, shortLink).Scan(&fullLink)
	if err != nil {
		if err != pgx.ErrNoRows {
			p.Log.Print(err)
		}
		return "", &links.MyError{Code: 500, Err: err}
	}

	return fullLink, nil
}
