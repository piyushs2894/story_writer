package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"story_writer/src/common/config"
)

type MasterSlave struct {
	Master *DB
	Slave  *DB
}

var DBConnMap map[string]*MasterSlave

//TODO: Set prometheus

//DB configuration
type DB struct {
	DBConnection  *sqlx.DB
	DBString      string
	RetryInterval int
	MaxConn       int
	doneChannel   chan bool
}

var (
	dbTicker *time.Ticker
)

func Init(cfgs *config.Config, driver string) {
	DBConnMap = make(map[string]*MasterSlave)

	for k, cfg := range cfgs.Database {
		masterDsn := cfg.Master
		slaveDsn := cfg.Slave
		var maxConnSlave, maxConnMaster int
		if maxConnMaster = cfg.MaxMasterConn; maxConnMaster == 0 {
			maxConnMaster = 80
		}
		if maxConnSlave = cfg.MaxSlaveConn; maxConnSlave == 0 {
			maxConnSlave = 100
		}

		Master := &DB{
			DBString:      masterDsn,
			RetryInterval: 10,
			MaxConn:       maxConnMaster,
			doneChannel:   make(chan bool),
		}
		Master.ConnectAndMonitor(driver)

		Slave := &DB{
			DBString:      slaveDsn,
			RetryInterval: 10,
			MaxConn:       maxConnSlave,
			doneChannel:   make(chan bool),
		}
		Slave.ConnectAndMonitor(driver)

		DBConnMap[k] = &MasterSlave{
			Master: Master,
			Slave:  Slave,
		}

	}

	dbTicker = time.NewTicker(time.Second * 2)
}

// GetOpenConnections return open connections to db
func (d *DB) GetOpenConnections() int64 {
	return d.GetOpenConnections()
}

// Connect to database
func (d *DB) Connect(driver string) error {
	var db *sqlx.DB
	var err error

	db, err = sqlx.Open(driver, d.DBString)

	if err != nil {
		log.Println("[Error]: DB open connection error", err.Error())
	} else {
		d.DBConnection = db
		err = db.Ping()
		if err != nil {
			log.Println("[Error]: DB connection error", err.Error())
		}
		return err
	}

	db.SetMaxOpenConns(d.MaxConn)

	return err
}

// ConnectAndMonitor to database
func (d *DB) ConnectAndMonitor(driver string) {
	err := d.Connect(driver)

	if err != nil {
		log.Printf("Not connected to database %s, trying \n", d.DBString)
	} else {
		log.Printf("Success connecting to database %s \n", d.DBString)
	}

	ticker := time.NewTicker(time.Duration(d.RetryInterval) * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				if d.DBConnection == nil {
					d.Connect(driver)
				} else {
					err := d.DBConnection.Ping()
					if err != nil {
						log.Println("[Error]: DB reconnect error", err.Error())
					}
				}
			case <-d.doneChannel:
				return
			}
		}
	}()
}

// DoneConnectAndMonitor to exit connect and monitor
func (d *DB) DoneConnectAndMonitor() {
	d.doneChannel <- true
}

//Prepare query for database queries
func (d *DB) Prepare(query string) *sql.Stmt {
	statement, err := d.DBConnection.Prepare(query)

	if err != nil {
		log.Printf("Failed to prepare query: %s. Error: %s\n", query, err.Error())
	}

	return statement
}

//Preparex query for database queries
func (d *DB) Preparex(query string) *sqlx.Stmt {
	if d == nil {
		log.Fatalf("Failed to prepare query, database object is nil. Query: %s\n", query)
		return nil
	}

	statement, err := d.DBConnection.Preparex(query)

	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s\n", query, err.Error())
	}

	return statement
}
