package shared

import (
	"os"
	"time"
	"strings"
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
	name		string
	desc		string
	ver			string
	tag			string
	pluginType 	string
	url			string
}
var appsRows appTable

func AppsDisplaySQL() (result []string){
	homeDir,_ := os.UserHomeDir()
	dbPath := homeDir + "/" + homeK3ai + "/" + "k3ai.db"
	db, _ := sql.Open("sqlite3",dbPath ) // Open the created SQLite File
	row, _ := db.Query("SELECT * FROM plugins WHERE type='Application' ORDER BY name ;")
	defer row.Close()
	var id int
	var name string
	var desc string
	var ver string
	var tag string
	var pluginType string
	var resource string
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&id, &name, &desc,&ver,&tag,&pluginType,&resource)
		appsRows.name = name
		appsRows.desc = desc
		appsRows.ver = ver
		appsRows.tag = tag
		result = append(result,name,desc,ver,tag)
	}
	return result
}

func InfraDisplaySQL() (result []string){
	homeDir,_ := os.UserHomeDir()
	dbPath := homeDir + "/" + homeK3ai + "/" + "k3ai.db"
	db, _ := sql.Open("sqlite3",dbPath ) // Open the created SQLite File
	row, _ := db.Query("SELECT * FROM plugins WHERE type='Infra' ORDER BY name ;")
	defer row.Close()
	var id int
	var name string
	var desc string
	var ver string
	var tag string
	var pluginType string
	var resource string
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&id, &name, &desc,&ver,&tag,&pluginType,&resource)
		appsRows.name = name
		appsRows.desc = desc
		appsRows.ver = ver
		appsRows.tag = tag
		result = append(result,name,desc,ver,tag)
	}
	
	// return result
	return result
}

func BundleDisplaySQL() (result []string){
	homeDir,_ := os.UserHomeDir()
	dbPath := homeDir + "/" + homeK3ai + "/" + "k3ai.db"
	db, _ := sql.Open("sqlite3",dbPath ) // Open the created SQLite File
	row, _ := db.Query("SELECT * FROM plugins WHERE type='Bundle' ORDER BY name ;")
	defer row.Close()
	var id int
	var name string
	var desc string
	var ver string
	var tag string
	var pluginType string
	var resource string
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&id, &name, &desc,&ver,&tag,&pluginType,&resource)
		appsRows.name = name
		appsRows.desc = desc
		appsRows.ver = ver
		appsRows.tag = tag
		result = append(result,name,desc,ver,tag)
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

	//Create relationship table between clusters and plugins to store config status
		appsTable := `CREATE TABLE cluster_plugins (
			"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
			"clusId" INT ,
			"plugId" INT
			);`
			dbState, err := db.Prepare(appsTable)
			if err != nil {
				log.Fatal(err.Error())
				return db,err
			}
			dbState.Exec()

	// Create CLUSTERS table to register clusters
	appsTable = `CREATE TABLE clusters (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		"name" TEXT ,
		"type" TEXT
		);`
		dbState, err = db.Prepare(appsTable)
		if err != nil {
			log.Fatal(err.Error())
			return db,err
		}
		dbState.Exec()
	//Create Application plugins list
	appsTable = `CREATE TABLE plugins (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		"name" TEXT UNIQUE,
		"desc" TEXT,
		"ver" TEXT,
		"tag" TEXT,
		"type" TEXT,
		"url" TEXT
		);`
		dbState, err = db.Prepare(appsTable)
		if err != nil {
			log.Fatal(err.Error())
			return db,err
		}
		dbState.Exec()

		return db,err
}


func FillPluginTables(data *data.K3ai, url string) error {
	homeDir,_ := os.UserHomeDir()
	dbPath := homeDir + "/" + homeK3ai + "/" + "k3ai.db"
	db, _ := sql.Open("sqlite3",dbPath ) // Open the created SQLite File
	defer db.Close()
	name := strings.ToLower(data.Metadata.Name)
	desc := strings.ToLower(data.Metadata.Desc)
	ver := strings.ToLower(data.Metadata.Version)
	tag  := strings.ToLower(data.Metadata.Tag)
	pluginType := data.Kind
	urlDB := url
	
	fillApps := `INSERT INTO plugins(name,desc,ver,tag,type,url) VALUES (?, ?, ?, ?, ?,?)`
	statement, err:= db.Prepare(fillApps)
	if err != nil {
		log.Error(err)
	}
	defer statement.Close()
	_,err = statement.Exec(name,desc,ver,tag,pluginType,urlDB)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}


func UpdatePluginTables(data *data.K3ai, url string) error {
	homeDir,_ := os.UserHomeDir()
	dbPath := homeDir + "/" + homeK3ai + "/" + "k3ai.db"
	db, err := sql.Open("sqlite3",dbPath ) // Open the created SQLite File
	if err != nil {
		log.Info("Did you run k3ai init first? We need to creat the db first..")
	}
	defer db.Close()
	name := strings.ToLower(data.Metadata.Name)
	desc := strings.ToLower(data.Metadata.Desc)
	ver := strings.ToLower(data.Metadata.Version)
	tag  := strings.ToLower(data.Metadata.Tag)
	pluginType := data.Kind
	urlDB := url
	
	fillApps := `INSERT OR REPLACE INTO plugins(name,desc,ver,tag,type,url) VALUES (?, ?, ?, ?, ?,?);`
	statement, err:= db.Prepare(fillApps)
	if err != nil {
		log.Error(err)
	}
	defer statement.Close()
	_,err = statement.Exec(name,desc,ver,tag,pluginType,urlDB)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func SelectPlugin(pluginName string) (pluginType string, pluginUrl string) {
	homeDir,_ := os.UserHomeDir()
	dbPath := homeDir + "/" + homeK3ai + "/" + "k3ai.db"
	db, _ := sql.Open("sqlite3",dbPath ) // Open the created SQLite File
	query := "SELECT type,url FROM plugins WHERE name='" + pluginName +"';"
	row := db.QueryRow(query,pluginName)
	err := row.Scan(&pluginType,&pluginUrl)
	if err != nil {
		log.Error(err)
	}

	return pluginType, pluginUrl
}

func RegisterCluster(pluginName string) (pluginType string, pluginUrl string) {
	homeDir,_ := os.UserHomeDir()
	dbPath := homeDir + "/" + homeK3ai + "/" + "k3ai.db"
	db, _ := sql.Open("sqlite3",dbPath ) // Open the created SQLite File
	query := "SELECT type,url FROM plugins WHERE name='" + pluginName +"';"
	row := db.QueryRow(query,pluginName)
	err := row.Scan(&pluginType,&pluginUrl)
	if err != nil {
		log.Error(err)
	}

	return pluginType, pluginUrl
}
