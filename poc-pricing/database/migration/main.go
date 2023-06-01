// Package main, work as CLI for db migration.
// How to run this CLI:
// - go inside the directory, `cd database/migration`
// `go run main.go -dsn={value} -driver={value}`
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"math"
	"os"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	flagReverse := flag.Bool("reverse", false, "To reverse migration")
	flagStep := flag.Int("step", 0, "How many step(s) want when reversal migrate happens")

	// TODO: add more description of usage this flag, if only want use the other driver(s)
	flagDSN := flag.String("dsn",
		"root:@tcp(localhost:3308)/poc-pricing-db?multiStatements=true",
		"DSN for database. By default, the value DSN will be filled for mysql driver\n"+
			"Format :\n"+
			"- mysql/mariadb: user:password@tcp(host:port)/db_name",
	)
	flagDriver := flag.String(
		"driver",
		"mysql",
		"Driver for database.\n"+
			"List of drivers:\n"+
			"- mysql (mariadb use this too)\n",
	)
	flag.Parse()

	m, err := newMigrate(*flagDSN, *flagDriver)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if *flagReverse {
		err = down(m, *flagStep)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		os.Exit(0)
	}

	err = up(m)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

func newMigrate(dsn, driver string) (*migrate.Migrate, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("error when open [%s] connection: %v", driver, err)
	}

	// TODO: add switch for other driver(s)
	dbd, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, fmt.Errorf("error when create instance [%s]: %v", driver, err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://./ddl", driver, dbd)
	if err != nil {
		return nil, fmt.Errorf("error when migrate instance [%s]: %v", driver, err)
	}

	return m, nil
}

func up(m *migrate.Migrate) error {
	err := m.Up()
	if err != nil {
		return fmt.Errorf("error when migrate up: %v", err)
	}

	return nil
}

func down(m *migrate.Migrate, steps int) error {
	err := m.Steps(int(math.Abs(float64(steps))) * -1)
	if err != nil {
		return fmt.Errorf("error when migrate down: %v", err)
	}

	return nil
}
