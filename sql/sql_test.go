package sql_test

import (
	dbsql "database/sql"
	"testing"

	"github.com/gtramontina/go-extlib/sql"
	"github.com/gtramontina/go-extlib/testing/assert"
	_ "modernc.org/sqlite"
)

func TestSQL(t *testing.T) {
	db, err := dbsql.Open("sqlite", ":memory:")
	assert.NoError(t, err)
	assert.NoError(t, db.Ping())

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS cars (
			id		INTEGER 	PRIMARY KEY AUTOINCREMENT,
			model	TEXT 		NOT NULL,
			make	TEXT 		NOT NULL,
			year	INTEGER 	NOT NULL
		);
		INSERT INTO cars (model, make, year) VALUES ('Civic', 'Honda', 2019);
		INSERT INTO cars (model, make, year) VALUES ('Accord', 'Honda', 2018);
		INSERT INTO cars (model, make, year) VALUES ('Camry', 'Toyota', 2017);
		INSERT INTO cars (model, make, year) VALUES ('Corolla', 'Toyota', 2016);
		INSERT INTO cars (model, make, year) VALUES ('Fusion', 'Ford', 2015);
		INSERT INTO cars (model, make, year) VALUES ('Focus', 'Ford', 2014);
	`)
	assert.NoError(t, err)

	type Car struct {
		ID    int
		Model string
		Make  string
		Year  int
	}

	t.Run("QueryRow does not find anything", func(t *testing.T) {
		_, found := sql.QueryRow[Car](db, "SELECT * FROM cars WHERE id = -1")
		assert.False(t, found)
	})

	t.Run("QueryRow finds and returns an item", func(t *testing.T) {
		{
			car, found := sql.QueryRow[Car](db, "SELECT * FROM cars WHERE id = 1")
			assert.True(t, found)
			assert.DeepEqual(t, car, Car{ID: 1, Model: "Civic", Make: "Honda", Year: 2019})
		}
		{
			car, found := sql.QueryRow[Car](db, "SELECT * FROM cars WHERE id = 4")
			assert.True(t, found)
			assert.DeepEqual(t, car, Car{ID: 4, Model: "Corolla", Make: "Toyota", Year: 2016})
		}
	})

	t.Run("QueryRow finds and returns a pointer to an item", func(t *testing.T) {
		{
			car, found := sql.QueryRow[*Car](db, "SELECT * FROM cars WHERE id = 2")
			assert.True(t, found)
			assert.DeepEqual(t, car, &Car{ID: 2, Model: "Accord", Make: "Honda", Year: 2018})
		}
		{
			car, found := sql.QueryRow[*Car](db, "SELECT * FROM cars WHERE id = 5")
			assert.True(t, found)
			assert.DeepEqual(t, car, &Car{ID: 5, Model: "Fusion", Make: "Ford", Year: 2015})
		}
	})

	t.Run("Query returns an empty stream of items", func(t *testing.T) {
		stream := sql.Query[Car](db, "SELECT * FROM cars WHERE id = -1")

		_, found := stream.Next()
		assert.False(t, found)
	})

	t.Run("Query returns a stream of items", func(t *testing.T) {
		stream := sql.Query[Car](db, "SELECT * FROM cars")

		var cars []Car
		for next, found := stream.Next(); found; next, found = stream.Next() {
			cars = append(cars, next)
		}

		assert.DeepEqual(t, cars, []Car{
			{ID: 1, Model: "Civic", Make: "Honda", Year: 2019},
			{ID: 2, Model: "Accord", Make: "Honda", Year: 2018},
			{ID: 3, Model: "Camry", Make: "Toyota", Year: 2017},
			{ID: 4, Model: "Corolla", Make: "Toyota", Year: 2016},
			{ID: 5, Model: "Fusion", Make: "Ford", Year: 2015},
			{ID: 6, Model: "Focus", Make: "Ford", Year: 2014},
		})
	})

	t.Run("Query returns a stream of pointers to items", func(t *testing.T) {
		stream := sql.Query[*Car](db, "SELECT * FROM cars WHERE make = 'Honda'")

		var cars []*Car
		for next, found := stream.Next(); found; next, found = stream.Next() {
			cars = append(cars, next)
		}

		assert.DeepEqual(t, cars, []*Car{
			{ID: 1, Model: "Civic", Make: "Honda", Year: 2019},
			{ID: 2, Model: "Accord", Make: "Honda", Year: 2018},
		})
	})
}
