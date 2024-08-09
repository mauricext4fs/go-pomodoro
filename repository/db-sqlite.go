package repository

import (
	"database/sql"
	"errors"
	"time"
)

type SQLiteRepository struct {
	Conn *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		Conn: db,
	}
}

func (repo *SQLiteRepository) Migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS activities (
		id integer primary key autoincrement,
		activity_type int not null,
		start_timestamp integer not null,
		end_timestamp integer DEFAULT 0
	);
	
	CREATE TABLE IF NOT EXISTS activity_type (
		id integer primary key autoincrement,
		title text not null,
		type text not null
	);
	
	INSERT OR IGNORE INTO activity_type (id, title, type) VALUES (100, 'Standard Pomodoro (25 Minutes)', 'Pomodoro');
	INSERT OR IGNORE INTO activity_type (id, title, type) VALUES (200, '5 Minutes Break', 'Break');
	INSERT OR IGNORE INTO activity_type (id, title, type) VALUES (500, '20 Minutes Break', 'Break');
	`

	_, err := repo.Conn.Exec(query)
	return err
}

func (repo *SQLiteRepository) StartActivity(activities Activities) (*Activities, error) {
	stmt := "INSERT INTO Activities (activity_type, start_timestamp) values (?, ?)"

	res, err := repo.Conn.Exec(stmt, activities.ActivityType, activities.StartTimestamp.Unix())
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	activities.ID = id

	return &activities, nil
}

func (repo *SQLiteRepository) AllActivities() ([]Activities, error) {
	query := "SELECT id, activity_type, start_timestamp, end_timestamp FROM activities ORDER BY start_timestamp DESC"
	rows, err := repo.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []Activities
	for rows.Next() {
		var a Activities
		var startUnixTime int64
		var endUnixTime int64
		err := rows.Scan(
			&a.ID,
			&a.ActivityType,
			&startUnixTime,
			&endUnixTime,
		)
		if err != nil {
			return nil, err
		}
		a.StartTimestamp = time.Unix(startUnixTime, 0)
		a.EndTimestamp = time.Unix(endUnixTime, 0)
		all = append(all, a)
	}

	return all, nil
}

func (repo *SQLiteRepository) CountCompletedPomodoro() (int64, error) {
	query := "SELECT COUNT(*) as ct FROM Activities WHERE end_timestamp != 0 AND activity_type = 100"
	row := repo.Conn.QueryRow(query)

	var a int64
	var c Count
	err := row.Scan(
		&c.Count,
	)

	if err != nil {
		return 0, err
	}
	a = c.Count

	return a, nil
}

func (repo *SQLiteRepository) AllActivityType() ([]ActivityType, error) {
	query := "SELECT id, title, type FROM activity_type ORDER BY id ASC"
	rows, err := repo.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []ActivityType
	for rows.Next() {
		var a ActivityType
		err := rows.Scan(
			&a.ID,
			&a.Title,
			&a.Type,
		)
		if err != nil {
			return nil, err
		}
		all = append(all, a)
	}

	return all, nil
}

func (repo *SQLiteRepository) GetActivityByID(id int) (*Activities, error) {
	row := repo.Conn.QueryRow("SELECT id, activity_type, start_timestamp, end_timestamp FROM activities WHERE id = ?", id)

	var a Activities
	var startUnixTime int64
	var endUnixTime int64
	err := row.Scan(
		&a.ID,
		&a.ActivityType,
		&startUnixTime,
		&endUnixTime,
	)

	if err != nil {
		return nil, err
	}

	a.StartTimestamp = time.Unix(startUnixTime, 0)
	a.EndTimestamp = time.Unix(endUnixTime, 0)

	return &a, nil
}

func (repo *SQLiteRepository) UpdateActivity(id int64, updated Activities) error {
	if id == 0 {
		return errors.New("Invalid Updated ID")
	}

	stmt := "UPDATE activities SET end_timestamp = ? WHERE id = ?"
	res, err := repo.Conn.Exec(stmt, updated.EndTimestamp.Unix(), id)

	if err != nil {
		return err
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return errUpdateFailed
	}

	return nil

}

func (repo *SQLiteRepository) DeleteActivity(id int64) error {
	res, err := repo.Conn.Exec("DELETE FROM activities WHERE id = ?", id)
	if err != nil {
		return err

	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errDeleteFailed
	}

	return nil

}
