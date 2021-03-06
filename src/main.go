package main

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/cloud-barista/cb-tumblebug/src/core/common"
	"github.com/cloud-barista/cb-tumblebug/src/core/mcis"

	grpcserver "github.com/cloud-barista/cb-tumblebug/src/api/grpc/server"
	restapiserver "github.com/cloud-barista/cb-tumblebug/src/api/rest/server"
)

// Main Body

// @title CB-Tumblebug REST API
// @version 0.2.0
// @description CB-Tumblebug REST API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://cloud-barista.github.io
// @contact.email contact-to-cloud-barista@googlegroups.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /tumblebug

// @securityDefinitions.basic BasicAuth

func main() {
	common.SPIDER_REST_URL = os.Getenv("SPIDER_REST_URL")
	common.DRAGONFLY_REST_URL = os.Getenv("DRAGONFLY_REST_URL")
	common.DB_URL = os.Getenv("DB_URL")
	common.DB_DATABASE = os.Getenv("DB_DATABASE")
	common.DB_USER = os.Getenv("DB_USER")
	common.DB_PASSWORD = os.Getenv("DB_PASSWORD")

	// load config
	//masterConfigInfos = confighandler.GetMasterConfigInfos()

	//Ticker for MCIS status validation
	validationDuration := 60000 //ms
	ticker := time.NewTicker(time.Millisecond * time.Duration(validationDuration))
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
			mcis.ValidateStatus()
		}
	}()
	defer ticker.Stop()

	var err error
	/*
		common.MYDB, err = sql.Open("mysql", //"root:pwd@tcp(127.0.0.1:3306)/testdb")
			common.DB_USER+":"+
				common.DB_PASSWORD+"@tcp("+
				common.DB_URL+")/"+
				common.DB_DATABASE)
	*/
	common.MYDB, err = sql.Open("sqlite3", "file:../meta_db/dat/cbtumblebug.s3db")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database access info set successfully")
	}

	_, err = common.MYDB.Exec("USE " + common.DB_DATABASE)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB selected successfully..")
	}

	stmt, err := common.MYDB.Prepare("CREATE Table IF NOT EXISTS spec(" +
		"id varchar(50) NOT NULL," +
		"connectionName varchar(50) NOT NULL," +
		"cspSpecName varchar(50) NOT NULL," +
		"name varchar(50)," +
		"os_type varchar(50)," +
		"num_vCPU varchar(50)," +
		"num_core varchar(50)," +
		"mem_GiB varchar(50)," +
		"mem_MiB varchar(50)," +
		"storage_GiB varchar(50)," +
		"description varchar(50)," +
		"cost_per_hour varchar(50)," +
		"num_storage varchar(50)," +
		"max_num_storage varchar(50)," +
		"max_total_storage_TiB varchar(50)," +
		"net_bw_Gbps varchar(50)," +
		"ebs_bw_Mbps varchar(50)," +
		"gpu_model varchar(50)," +
		"num_gpu varchar(50)," +
		"gpumem_GiB varchar(50)," +
		"gpu_p2p varchar(50)," +
		"PRIMARY KEY (id));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Table spec created successfully..")
	}

	stmt, err = common.MYDB.Prepare("CREATE Table IF NOT EXISTS image(" +
		"id varchar(50) NOT NULL," +
		"name varchar(50)," +
		"connectionName varchar(50) NOT NULL," +
		"cspImageId varchar(400) NOT NULL," +
		"cspImageName varchar(400) NOT NULL," +
		"creationDate varchar(50) NOT NULL," +
		"description varchar(400) NOT NULL," +
		"guestOS varchar(50) NOT NULL," +
		"status varchar(50) NOT NULL," +
		"PRIMARY KEY (id));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Table image created successfully..")
	}

	//defer db.Close()

	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		restapiserver.ApiServer()
		wg.Done()
	}()

	go func() {
		grpcserver.RunServer()
		//fmt.Println("gRPC server started on " + grpcserver.Port)
		wg.Done()
	}()

	wg.Wait()
}
