package dao

import (
	"context"
	"time"
)

type DoSomething struct {
	ID                int64     `json:"id" db:"id"`
	Flag              int       `json:"flag" db:"flag"`
	CreateTimeUtc     time.Time `json:"create_time_utc" db:"create_time_utc"`
	LastUpdateTimeUtc time.Time `json:"last_update_time_utc" db:"last_update_time_utc"`
	Who               string    `json:"who" db:"who"`
	Action            string    `json:"action" db:"action"`
	Thing             string    `json:"thing" db:"thing"`
	More              string    `json:"more" db:"more"`
}

func (d *DoSomething) Insert(ctx context.Context) error {
	sqlS := "INSERT INTO do_something(who, `action`, thing, `more`) VALUES(?, ?, ?, ?)"

	res, err := sqlxDB.ExecContext(ctx, sqlS, d.Who, d.Action, d.Thing, d.More)
	if err != nil {
		return err
	}

	d.ID, err = res.LastInsertId()
	return err
}

func (d *DoSomething) Get(ctx context.Context) error {
	sqlS := "SELECT * FROM do_something WHERE who = ? AND `action` = ? AND thing = ?"

	return sqlxDB.GetContext(ctx, d, sqlS, d.Who, d.Action, d.Thing)
}

func (d *DoSomething) Delete(ctx context.Context) error {
	sqlS := "DELETE FROM do_something WHERE who = ? AND `action` = ? AND thing = ?"

	_, err := sqlxDB.ExecContext(ctx, sqlS, d.Who, d.Action, d.Thing)
	if err != nil {
		return err
	}

	return nil
}
