package db

import (
	"os"
	"log"
	"strings"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	clusters = `CREATE TABLE "clusters" (
		"id"	INTEGER NOT NULL UNIQUE,
		"name"	TEXT NOT NULL UNIQUE,
		"type"	TEXT NOT NULL,
		"kube"	BLOB,
		"status"	TEXT NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT)
	);`

	plugins = `CREATE TABLE "plugins" (
		"id"	INTEGER NOT NULL UNIQUE,
		"name"	TEXT NOT NULL UNIQUE,
		"desc"	TEXT NOT NULL,
		"type"	TEXT,
		"tag"	TEXT,
		"version"	TEXT,
		"url"	TEXT NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT)
	);`

	currentStatus = `CREATE TABLE "listStatus" (
		"id"	INTEGER NOT NULL UNIQUE,
		"cluster_id"	INTEGER NOT NULL,
		"plugin_id"	INTEGER NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT)
	);`
)

var tables = []string {clusters,plugins,currentStatus}

func DbLogin() (sqlDBConn *sql.DB) {
	homeDir,_ := os.UserHomeDir()
	dbPath := homeDir + "/.k3ai/k3ai.db"
	sqlDBConn, err := sql.Open("sqlite3",dbPath ) // Open the created SQLite File
	if err != nil {
		log.Fatal(err)
	} 
	// defer sqlDBConn.Close() // Defer Closing the database
	// sqlDBConn.Close()
	return sqlDBConn

}


func InitDB(ch chan bool) {
	
	homeDir,_ := os.UserHomeDir()
	if _, err := os.Stat(homeDir + "/.k3ai/k3ai.db"); os.IsNotExist(err) {
		file, err := os.Create(homeDir + "/.k3ai/k3ai.db") // Create SQLite file
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	sqlDBConn := DbLogin()
	err = createTables(ch,sqlDBConn)
	if err != nil {
		log.Fatal(err)
		}
	}
	ch <- true
}

func Insert(sqlDBConn *sql.DB ) error {

	insertStatus := `INSERT INTO listStatus (cluster_id,plugin_id)
	SELECT plugins.id,clusters.id
	FROM plugins
	INNER JOIN clusters
	WHERE plugins.name = ? AND clusters.name= ?;`
	dbState, err := sqlDBConn.Prepare(insertStatus)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	dbState.Exec()
	return nil

}

func List(sqlDBConn *sql.DB ) error {
	listStatusByCluster := `SELECT * FROM listStatus
	WHERE cluster_id IN (
			SELECT clusters.id FROM clusters
			WHERE clusters.name=?
		);`
	dbState, err := sqlDBConn.Prepare(listStatusByCluster)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	dbState.Exec()
	return nil

}

func Update(sqlDBConn *sql.DB ) {


}


func Delete(sqlDBConn *sql.DB ) {

}

func createTables(ch chan bool,sqlDBConn *sql.DB ) error {
	for i:=0;i < len(tables); i++ {
		dbState, err := sqlDBConn.Prepare(tables[i])
		if err != nil {
			log.Fatal(err.Error())
			return err
		}
		dbState.Exec()
	}
	sqlDBConn.Close()
	ch <-true
	return nil
}

func InsertPlugins(plugin [6]string ) error {
	db := DbLogin()
	
	name := strings.ToLower(plugin[0])
	desc := strings.ToLower(plugin[1])
	ver := strings.ToLower(plugin[4])
	tag  := strings.ToLower(plugin[3])
	pluginType := strings.ToLower(plugin[2])
	pluginUrl := strings.ToLower(plugin[5])
	
	fillApps := `INSERT INTO plugins(name,desc,version,tag,type,url) VALUES (?, ?, ?, ?, ?,?)`
	statement, err:= db.Prepare(fillApps)
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	_,err = statement.Exec(name,desc,ver,tag,pluginType,pluginUrl)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil

}


