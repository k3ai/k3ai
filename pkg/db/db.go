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
		"name"	TEXT NOT NULL UNIQUE,
		"ctype"	TEXT NOT NULL,
		"kube"	BLOB,
		"status"	TEXT NOT NULL
	);`

	plugins = `CREATE TABLE "plugins" (
		"name"	TEXT NOT NULL UNIQUE,
		"desc"	TEXT NOT NULL,
		"ptype"	TEXT,
		"tag"	TEXT,
		"version"	TEXT,
		"url"	TEXT NOT NULL,
		"status" 	TEXT NOT NULL
	);`

	currentStatus = `CREATE TABLE "listStatus" (
		"cluster_id"	INTEGER NOT NULL,
		"plugin_id"	INTEGER NOT NULL
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

func UpdatePlugins(plugin [7]string) error {
	db := DbLogin()
	
	name := strings.ToLower(plugin[0])
	desc := strings.ToLower(plugin[1])
	ver := strings.ToLower(plugin[4])
	tag  := strings.ToLower(plugin[3])
	pluginType := strings.ToLower(plugin[2])
	pluginUrl := strings.ToLower(plugin[5])
	pluginStatus := strings.Title(plugin[6])
	
	fillApps := `INSERT OR REPLACE INTO plugins (name,desc,version,tag,ptype,url,status) VALUES (?, ?, ?, ?, ?,?,?);`
	statement, _:= db.Prepare(fillApps)
	_,err := statement.Exec(name,desc,ver,tag,pluginType,pluginUrl,pluginStatus)
	if err != nil {
		fillNew := `INSERT INTO plugins (name,desc,version,tag,ptype,url,status) VALUES ?, ?, ?, ?, ?,?,?);`
		alternative, _:= db.Prepare(fillNew)
		_,err = alternative.Exec(name,desc,ver,tag,pluginType,pluginUrl,pluginStatus)
		if err != nil {
			log.Fatal(err)
		}
		alternative.Close()
	}
	statement.Close()
	defer statement.Close()
	return nil

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

func InsertPlugins(plugin [7]string ) error {
	db := DbLogin()
	
	name := strings.ToLower(plugin[0])
	desc := strings.ToLower(plugin[1])
	ver := strings.ToLower(plugin[4])
	tag  := strings.ToLower(plugin[3])
	pluginType := strings.ToLower(plugin[2])
	pluginUrl := strings.ToLower(plugin[5])
	pluginStatus := strings.Title(plugin[6])
	
	fillApps := `INSERT INTO plugins(name,desc,version,tag,ptype,url,status) VALUES (?, ?, ?, ?, ?,?,?)`
	statement, err:= db.Prepare(fillApps)
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	_,err = statement.Exec(name,desc,ver,tag,pluginType,pluginUrl,pluginStatus)
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
		var name string
		var desc string
		var ptype string
		var tag string
		var version string
		var url string
		var status string
		row.Scan(&name,&desc,&ptype,&tag,&version,&url,&status)
		if ptype == "application"{
			appsResults = append(appsResults,name,desc,ptype,tag,version,status)
		}
		if ptype == "infra" {
			infraResults = append(infraResults,name,desc,ptype,tag,version, status)
		}
		if ptype == "bundles" {
			bundlesResults = append(bundlesResults,name,desc,ptype,tag,version, status)
		}
		if ptype =="comms" {
			commsResults = append(commsResults,name,desc,ptype,tag,version, status)
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
		var name string
		var desc string
		var ptype string
		var tag string
		var version string
		var url string
		var status string
		row.Scan(&name,&desc,&ptype,&tag,&version,&url,&status)
		Results = append(Results,name,desc,ptype,tag,version,status)
		
	}
	return Results
}

func ListClustersByName() (clusterResults[]string) {

	db := DbLogin()
	row, err := db.Query("SELECT * from CLUSTERS;")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var name string
		var ctype string
		var kube string
		var status string
		row.Scan(&name,&ctype,&kube,&status)
		clusterResults = append(clusterResults,name,ctype,status)
	}

		
	return clusterResults
}

func ListClusterByName(target string) (name string,ctype string) {

	db := DbLogin()
	row, err := db.Query("SELECT name,ctype from CLUSTERS WHERE name=?;",target)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	row.Next()
	row.Scan(&name,&ctype)
		
	return name,ctype
}

func CheckClusterName(name string) (resname string, ctype string) {
	db := DbLogin()
	row, _ := db.Query("SELECT name,ctype from CLUSTERS where name=?;",name)

	defer row.Close()
	row.Next()
	err := row.Scan(&name,&ctype)
	if err == nil {
		return "name exist",ctype
	}
	return "",""
}

func InsertCluster(clusterConfig []string ) error {

	db := DbLogin()
	insertStatus := `INSERT INTO clusters VALUES (?,?,?,?);`
	dbState, err := db.Prepare(insertStatus)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	_, err = dbState.Exec(clusterConfig[0],clusterConfig[1],clusterConfig[2],clusterConfig[3])
	return err
}

func DeleteCluster(name string ) error {
	db := DbLogin()
	insertStatus := `DELETE from clusters where name=?;`
	dbState, err := db.Prepare(insertStatus)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	_, err = dbState.Exec(name)
	return err
}

func KubeStr() (kubeStr string) {
	db := DbLogin()
	row, err := db.Query("SELECT kube from CLUSTER;")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var name string
		var ctype string
		var kube string
		var status string
		row.Scan(&name,&ctype,&kube,&status)
		kubeStr = kube
	}
	return kubeStr
}