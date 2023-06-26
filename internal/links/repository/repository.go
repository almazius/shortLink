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

// NewPostgres Create new instance Postgres
func NewPostgres(pool *pgxpool.Pool) links.DatabaseService {
	return &Postgres{
		Pool: pool,
		Log:  log.New(os.Stdout, "Postgres ", log.LstdFlags|log.Lshortfile),
	}
}

// ExistShortLink Checks for a short link in the database
func (p *Postgres) ExistShortLink(ctx context.Context, shortLink string) (bool, *links.MyError) {
	conn, err := p.Pool.Acquire(ctx)
	if err != nil {
		p.Log.Print(err)
		return false, &links.MyError{
			Code: 500,
			Err:  err,
		}
	}
	defer conn.Release()

	res := false
	err = conn.QueryRow(ctx, `select exists(select * from links where shortLink=$1)`, shortLink).Scan(&res)
	if err != nil {
		p.Log.Print(err)
		return false, &links.MyError{
			Code: 500,
			Err:  err,
		}
	}
	return res, nil
}

// GetShortLink Retrieves an abbreviated link from the database.
//
//	If it is not present, returns an empty string and an error
func (p *Postgres) GetShortLink(ctx context.Context, link string) (string, *links.MyError) {
	conn, err := p.Pool.Acquire(ctx)
	if err != nil {
		p.Log.Print(err)
		return "", &links.MyError{
			Code: 500,
			Err:  err,
		}
	}
	defer conn.Release()

	short := ""
	err = conn.QueryRow(ctx, `select shortLink from links where exists(select * from links where link=$1) 
                              and link=$1`, link).Scan(&short)
	if err != nil {
		if err != pgx.ErrNoRows {
			p.Log.Print(err)
		}
		return "", &links.MyError{Code: 500, Err: err}
	}

	return short, nil
}

// AddNote Creates an entry in the database. Returns the record number and error. If an error occurred, returns -1
func (p *Postgres) AddNote(ctx context.Context, link, shortLink string, time time.Time) (int64, *links.MyError) {
	conn, err := p.Pool.Acquire(ctx)
	if err != nil {
		p.Log.Print(err)
		return -1, &links.MyError{
			Code: 500,
			Err:  err,
		}
	}
	defer conn.Release()

	var id int64
	err = conn.QueryRow(ctx, `INSERT INTO links (link, shortLink, creationTime) values ($1, $2, $3) 
                                                  returning id`, link, shortLink, time).Scan(&id)
	if err != nil {
		p.Log.Print(err)

		return -1, &links.MyError{Code: 500, Err: err}
	}
	return id, nil
}

// GetFullLink  Returns the full link for the shortened version.
// In case of an error, returns an empty string and an error
func (p *Postgres) GetFullLink(ctx context.Context, shortLink string) (string, *links.MyError) {
	conn, err := p.Pool.Acquire(ctx)
	if err != nil {
		p.Log.Print(err)
		return "", &links.MyError{
			Code: 500,
			Err:  err,
		}
	}
	defer conn.Release()

	fullLink := ""
	err = conn.QueryRow(ctx, `select link from links where exists(select * from links where shortLink=$1) and shortLink=$1`, shortLink).Scan(&fullLink)
	if err != nil {
		if err != pgx.ErrNoRows {
			p.Log.Print(err)
		}
		return "", &links.MyError{Code: 500, Err: err}
	}

	return fullLink, nil
}
