package SQLClient

import (
	utils "blood/Utils"
	"database/sql"

	_ "github.com/microsoft/go-mssqldb"

	"fmt"
	"log"
)

var (
	port     = 1433
	user     = "khairi"
	password = "@Azure2023"
	server   = "dtf-db-server.database.windows.net"
	database = "cdt-db"
)

type SQLClient struct {
	db *sql.DB
}

const (
	TBN_PLAYERS string = "PLAYERS"

	TB_PLAYERS string = "PLAYERS (PLAYER_ID, EMAIL, PLAYER, LEVEL, UPDATED_AT)  "
)

func OpenConn() *SQLClient {

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	var err error

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	err = db.Ping()
	var counter uint
	fmt.Println(err)
	for err != nil {
		fmt.Printf("Coonection Timed-out to '%s'. Attempt N°%d\n", server, counter)
		counter++
		err = db.Ping()
	}

	log.Printf("Connected to DB. <%s@%s:%d>\n", user, database, port)
	return &SQLClient{
		db,
	}
}

func (sc *SQLClient) Query(query string) {
	rows, err := sc.db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	rows.Close()

}

func (sc *SQLClient) GetRows(query string) *sql.Rows {
	rows, err := sc.db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	return rows
}

func (sc *SQLClient) PushRows(rows [][]any, table string) {
	query := fmt.Sprintf("INSERT INTO %s VALUES ", table)
	for _, row := range rows {
		query += utils.FormatValues(row) + ","

	}
	query = query[:len(query)-1]
	sc.Push(query)
}

func (sc *SQLClient) Push(query string) {
	_, err := sc.db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}
