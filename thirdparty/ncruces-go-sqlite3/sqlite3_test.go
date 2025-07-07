package ncruces_sqlite3

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func TestInMemory(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	_, err = db.ExecContext(ctx, `
CREATE TABLE users (id INTEGER PRIMARY KEY, email TEXT);
INSERT INTO users(email) VALUES ('foo@example.com'), ('bar@example.com'), ('baz@example.com');
`)
	if err != nil {
		t.Fatal(err)
	}

	rows, err := db.QueryContext(ctx, `SELECT email FROM users;`)
	if err != nil {
		t.Fatal(err)
	}

	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			t.Fatal(err)
		}
		t.Logf("email: %s", email)
	}

	rerr := rows.Close()
	if rerr != nil {
		t.Fatal(rerr)
	}
}
