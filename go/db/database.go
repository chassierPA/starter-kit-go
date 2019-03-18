package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Database is an *sql.DB object encapsuled
type Database struct {
	DB      *sql.DB
	Host    string `json:"host"`
	Port    string `json:"port"`
	User    string `json:"user"`
	Pass    string `json:"pass"`
	Type    string `json:"type"`
	DBname  string `json:"dbname"`
	SSLmode string `json:"sslmode"`
}

// Open used to connect to your databases, make sure that the information exist (SSLMode "sslmode=disable")
func (db *Database) Open() error {
	var err error

	if &db == nil || db.Port == "" || db.Host == "" {
		return fmt.Errorf("Info of database unknown")
	}
	db.DB, err = sql.Open(db.Type, "user="+db.User+" password="+db.Pass+" host="+db.Host+" port="+db.Port+" dbname="+db.DBname+" sslmode="+db.SSLmode)
	if err != nil {
		return err
	}
	return db.Ping()
}

// Close used to close the connexion with the current DB
func (db *Database) Close() error {
	return db.DB.Close()
}

// QueryRow used to get unique row on databases
func (db *Database) QueryRow(cmdSQL string, arg ...interface{}) *sql.Row {
	return db.DB.QueryRow(cmdSQL, arg...)

}

//Query used to get values in databases
func (db *Database) Query(cmdSQL string) (*sql.Rows, error) {
	return db.DB.Query(cmdSQL)
}

//Exec used to exec req sql on your database
func (db *Database) Exec(cmdSQL string, args ...interface{}) (sql.Result, error) {
	return db.DB.Exec(cmdSQL, args...)
}

//Ping used to ping connection with your databases
func (db *Database) Ping() error {
	return db.DB.Ping()
}

//Prepare used to prepare reqSql in your databases (optional but optimised if your cmdSql will return lot of datas)
func (db *Database) Prepare(cmdSQL string) (*sql.Stmt, error) {
	return db.DB.Prepare(cmdSQL)
}

/*
type Conf struct {
	db Database `json:"db"`
}

func main() {
	var c Conf
	file, err := ioutil.ReadFile("./path/To/JSON")
	if err != nil {
		fmt.Println("Setup fail: ", err.Error())
		os.Exit(1)
	}
	json.Unmarshal(file, &c)

	if err := c.db.Open(); err != nil {
		fmt.Println(err);
	}else {
		fmt.Println("connexion successful")
	}
	if err := c.db.Ping(); err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("Pinged to the databases")
	}
}*/
