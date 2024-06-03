package migration

import (
	"github.com/spurtcms/member/migration/mysql"
	"github.com/spurtcms/member/migration/postgres"
	"gorm.io/gorm"
)

func AutoMigration(DB *gorm.DB, dbtype any) {

	if dbtype == "postgres" {

		postgres.MigrateTables(DB) //auto migrate table

	} else if dbtype == "mysql" {

		mysql.MigrateTables(DB) //auto migrate table
	}

}
