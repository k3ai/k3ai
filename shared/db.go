package shared

import (
	"os"
	"time"
	"database/sql"

	"github.com/alefesta/k3ai/log"
	_ "github.com/mattn/go-sqlite3"
	data "github.com/alefesta/k3ai/config"
)

const (
	repoOwner = "k3ai"
	repoRoot = "plugins"
	repoApps = "apps"
	repoComm = "common"
	repoInfra = "infra"
	repoBundle = "bundles"
	homeK3ai = ".k3ai"
	k3aiDb = "k3ai.db"
	k3aiFile = "k3ai.yaml"
)

type appTable struct {
	name	string
	desc	string
	ver		string
	tag		string
}
var appsRows appTable

func AppsDisplaySQL() (result []string){
	homeDir,_ := os.UserHomeDir()
	dbPath := homeDir + "/" + homeK3ai + "/" + "k3ai.db"
	db, _ := sql.Open("sqlite3",dbPath ) // Open the created SQLite File
	row, _ := db.Query("SELECT * FROM apps ORDER BY name")
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var desc string
		var ver string
		var tag string
		var pluginType string
		var resource string
		row.Scan(&id, &name, &desc,&ver,&tag,&pluginType,&resource)
		appsRows.name = name
		appsRows.desc = desc
		appsRows.ver = ver
		appsRows.tag = tag
		result := []string{name,desc,ver,tag}
		return result
	}
	return result
}

func InfraDisplaySQL() (result []string){
	homeDir,_ := os.UserHomeDir()
	dbPath := homeDir + "/" + homeK3ai + "/" + "k3ai.db"
	db, _ := sql.Open("sqlite3",dbPath ) // Open the created SQLite File
	row, _ := db.Query("SELECT * FROM infra ORDER BY name")
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var desc string
		var ver string
		var tag string
		var pluginType string
		var resource string
		row.Scan(&id, &name, &desc,&ver,&tag,&pluginType,&resource)
		appsRows.name = name
		appsRows.desc = desc
		appsRows.ver = ver
		appsRows.tag = tag
		result := []string{name,desc,ver,tag}
		return result
	}
	return result
}

func BundleDisplaySQL() (result []string){
	homeDir,_ := os.UserHomeDir()
	dbPath := homeDir + "/" + homeK3ai + "/" + "k3ai.db"
	db, _ := sql.Open("sqlite3",dbPath ) // Open the created SQLite File
	row, _ := db.Query("SELECT * FROM bundles ORDER BY name")
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var desc string
		var ver string
		var tag string
		var pluginType string
		var resource string
		row.Scan(&id, &name, &desc,&ver,&tag,&pluginType,&resource)
		appsRows.name = name
		appsRows.desc = desc
		appsRows.ver = ver
		appsRows.tag = tag
		result := []string{name,desc,ver,tag}
		return result
	}
	return result
}

//DCreate is used to create a sqllite db un the previously created directory
func DbCreate() (db *sql.DB,err error){
	homeDir,_ := os.UserHomeDir()
	if _, err := os.Stat(homeDir + "/" + homeK3ai + "/" + "k3ai.db"); os.IsNotExist(err) {
		file, err := os.Create(homeDir + "/" + homeK3ai + "/" + "k3ai.db") // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		dbPath := homeDir + "/" + homeK3ai + "/" + "k3ai.db"
		sqliteDatabase, err := sql.Open("sqlite3",dbPath ) // Open the created SQLite File
		if err != nil {
			return sqliteDatabase,err
		} else {
			defer sqliteDatabase.Close() // Defer Closing the database
			sqliteDatabase,err = createTables(sqliteDatabase)
			if err != nil {
				log.Error(err)
			}
			sqliteDatabase.Close()
			return sqliteDatabase,err
		}
		
	} else {
		log.Error("A previous version of k3ai DB exist in the folder")
		time.Sleep(1 * time.Second)
		log.Error("Please run \"k3ai init update\" to overwrite or \"k3ai init delete\" to remove it.")
		os.Exit(1)
		return nil,err
	}
}

//createTables add the minimal set of tables to the db
func createTables(db *sql.DB) (dbConn *sql.DB,err error) {
	//Create Application plugins list
	appsTable := `CREATE TABLE apps (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"desc" TEXT,
		"ver" TEXT,
		"tag" TEXT,
		"type" TEXT,
		"resources" TEXT
		);`
		dbState, err := db.Prepare(appsTable)
		if err != nil {
			log.Fatal(err.Error())
			return db,err
		}
		dbState.Exec()
	//Create Infra plugins list
	infraTable := `CREATE TABLE infra (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"desc" TEXT,
		"ver" TEXT,
		"tag" TEXT,
		"type" TEXT,
		"resources" TEXT
		);`
		dbState = nil
		dbState, err = db.Prepare(infraTable)
		if err != nil {
			log.Fatal(err.Error())
			return db,err
		} else {
			dbState.Exec()
		}
		
	//Create Common plugins list
	commonTable := `CREATE TABLE commons (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"desc" TEXT,
		"ver" TEXT,
		"tag" TEXT
		);`
		dbState = nil
		dbState, err = db.Prepare(commonTable)
		if err != nil {
			log.Fatal(err.Error())
			return db,err
		} else {
			dbState.Exec()
		}
		
	//Create Bundles plugins list
	bundlesTable := `CREATE TABLE bundles (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"desc" TEXT,
		"ver" TEXT,
		"tag" TEXT
		);`
		dbState = nil
		dbState, err = db.Prepare(bundlesTable)
		if err != nil {
			log.Fatal(err.Error())
			return db,err
		} else {
			dbState.Exec()
		}
		return db,err
}


func FillTables(data *data.K3ai, url string) error {
	homeDir,_ := os.UserHomeDir()
	dbPath := homeDir + "/" + homeK3ai + "/" + "k3ai.db"
	db, _ := sql.Open("sqlite3",dbPath ) // Open the created SQLite File
	defer db.Close()
	name := data.Metadata.Name
	desc := data.Metadata.Desc
	ver := data.Metadata.Version
	tag  := data.Metadata.Tag
	pluginType := data.Metadata.Type
	resources := url
	
	fillApps := `INSERT INTO apps(name,desc,ver,tag,type,resources) VALUES (?, ?, ?, ?, ?, ?)`
	statement, err:= db.Prepare(fillApps)
	if err != nil {
		log.Error(err)
	}
	defer statement.Close()
	_,err = statement.Exec(name,desc,ver,tag,pluginType,resources)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
