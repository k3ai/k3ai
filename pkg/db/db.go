package db

import (
	"context"
	"log"
	"os"
	"strings"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)
var ctx context.Context
const (
	clusters = `CREATE TABLE "clusters" (
		"id"	INTEGER NOT NULL UNIQUE,
		"name"	TEXT NOT NULL UNIQUE,
		"ctype"	TEXT NOT NULL,
		"kube"	BLOB,
		"status"	TEXT NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT)
	);`

	plugins = `CREATE TABLE "plugins" (
		"id"	INTEGER NOT NULL UNIQUE,
		"name"	TEXT NOT NULL UNIQUE,
		"desc"	TEXT NOT NULL,
		"ptype"	TEXT,
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

func List(name string) (url string) {
	db := DbLogin()
	row, err := db.Query("SELECT url from PLUGINS where name=?;",name)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	row.Next()
	row.Scan(&url)
	return url
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
	
	fillApps := `INSERT INTO plugins(name,desc,version,tag,ptype,url) VALUES (?, ?, ?, ?, ?,?)`
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


func ListPlugins() (appsResults []string,infraResults []string,bundlesResults []string,commsResults[]string) {

	db := DbLogin()
	row, err := db.Query("SELECT * FROM plugins ORDER BY ptype;")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var desc string
		var ptype string
		var tag string
		var version string
		var url string
		row.Scan(&id,&name,&desc,&ptype,&tag,&version,&url)
		if ptype == "application"{
			appsResults = append(appsResults,name,desc,ptype,tag,version)
		}
		if ptype == "infra" {
			infraResults = append(infraResults,name,desc,ptype,tag,version)
		}
		if ptype == "bundles" {
			bundlesResults = append(bundlesResults,name,desc,ptype,tag,version)
		}
		if ptype =="comms" {
			commsResults = append(commsResults,name,desc,ptype,tag,version)
		}
		
	}
	
	return appsResults,infraResults,bundlesResults,commsResults
}

func ListPluginsByName(name string) (Results[]string) {

	db := DbLogin()
	row, err := db.Query("SELECT * FROM plugins WHERE name=?;",name)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var desc string
		var ptype string
		var tag string
		var version string
		var url string
		row.Scan(&id,&name,&desc,&ptype,&tag,&version,&url)
		Results = append(Results,name,desc,ptype,tag,version)
		
	}
	return Results
}

func ListClustersByName() (Results[]string) {

	db := DbLogin()
	row, err := db.Query("SELECT * FROM clusters;")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var desc string
		var ptype string
		var tag string
		var version string
		var url string
		row.Scan(&id,&name,&desc,&ptype,&tag,&version,&url)
		Results = append(Results,name,desc,ptype,tag,version)
		
	}
	return Results
}

func InsertCluster(clusterConfig []string ) error {
	db := DbLogin()
	insertStatus := `INSERT INTO clusters VALUES (?,?,?,?,?);`
	dbState, err := db.Prepare(insertStatus)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	var id int
	_, err = dbState.Exec(id,clusterConfig[0],clusterConfig[1],clusterConfig[2],clusterConfig[3])
	return nil
}

func KubeStr() (kubeStr string) {
	db := DbLogin()
	row, err := db.Query("SELECT kube from CLUSTER;")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var ctype string
		var kube string
		var status string
		row.Scan(&id,&name,&ctype,&kube,&status)
		kubeStr = kube
	}
	return kubeStr
}