package main

import (
	"context"
	"database/sql"
	"testing"
)

func TestApp(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	app := App{db: db}
	ctx := context.Background()
	_, err = app.db.ExecContext(ctx, schema)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := app.db.ExecContext(ctx, `INSERT OR IGNORE INTO visits VALUES (1, 0);`); err != nil {
		t.Fatal(err)
	}

	t.Run("initial count is zero", func(t *testing.T) {
		var got int
		want := 0
		err := app.db.QueryRowContext(ctx, "SELECT count FROM visits;").Scan(&got)
		if err != nil {
			t.Fatal(err)
		}

		if got != want {
			t.Fatalf("want %v, got %v", want, got)
		}
	})
}
