# starter-kit-go
Simple starter pack for gopher !

In this package there are three useful little packages which are **output log, universal json loader**, and **database**.

****

# Output package

Used to normalize your log of your program

**Example**
```go
func main() {
	InitLog()

	DisplayStart([]string{"Pierre-Alexandre CHASSIER"}, "0.0.1")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
	Debug.Println("Debuging Information")
}
```
*The output of the example is:*
![output package loader](./logger.png)

***

# DB package

Database is an *sql.DB object encapsuled

```go
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
```
*Here all basic function encapsuled:*
```go
// Open used to connect to your databases, make sure that the information exist (SSLMode "sslmode=disable")
func (db *Database) Open() error
```
```go
// Close used to close the connexion with the current DB
func (db *Database) Close() error
```

```go
// QueryRow used to get unique row on databases
func (db *Database) QueryRow(cmdSQL string, arg ...interface{}) *sql.Row
```

```go
//Query used to get values in databases
func (db *Database) Query(cmdSQL string) (*sql.Rows, error)
```

```go
//Exec used to exec req sql on your database
func (db *Database) Exec(cmdSQL string, args ...interface{}) (sql.Result, error)
```

```go
//Ping used to ping connection with your databases
func (db *Database) Ping() error
```

```go
//Prepare used to prepare reqSql in your databases (optional but optimised if your cmdSql will return lot of datas)
func (db *Database) Prepare(cmdSQL string) (*sql.Stmt, error)
```

**Example**

```go
//Example of structure which contain Database object
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
	}
	if err := c.db.Ping(); err != nil {
		fmt.Println(err)
	}
```
In your JSON file you must have :

```json
{
   "db": 
   {
       "host": "",
       "port" : "",
       "user": "",
       "pass": "",
       "type" : "",
       "dbname" : "",
       "sslmode" : ""
   }
}

```

***

# JSON Loader package

This package is able to load any object for JSON file.

**Example**

```go
//Conf is an example of structure
type Conf struct {
	db Database `json:"db"`
}

func main() {
    var db Database
	err := loader.LoadJSON(pathConfig, &db)
	if err != nil {
		output.LogPrintln(output.LTError, "setup fail %s", err.Error())
		os.Exit(1)
    }
    //Your object db Database is now loaded
    //...
}
```