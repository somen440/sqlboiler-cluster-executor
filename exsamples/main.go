package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	cluster "github.com/somen440/sqlboiler-cluster-executor"
	"github.com/somen440/sqlboiler-cluster-executor/exsamples/models"
	"github.com/volatiletech/sqlboiler/boil"
)

func init() {
	cluster := cluster.New("mysql", []string{
		"writer-user:writer-password@tcp(0.0.0.0:3312)/sample?parseTime=true",
		"reader-user:reader-password@tcp(0.0.0.0:3313)/sample?parseTime=true",
	})
	boil.SetDB(cluster)
}

func main() {
	ctx := context.Background()

	sample, err := models.Samples().OneG(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tbl: " + sample.Name)

	current := boil.GetDB()
	defer func() {
		boil.SetDB(current)
	}()
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	boil.SetDB(tx)

	entity := &models.Sample{
		Name: "tekitou",
	}
	if err := entity.InsertG(ctx, boil.Infer()); err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
