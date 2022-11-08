package postgresql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Postgresql struct {
	DB       *gorm.DB
	Entities []interface{}
}

func (p *Postgresql) Connect() {
	connection, err := gorm.Open(postgres.Open(os.Getenv("DSN_STRING")), &gorm.Config{})
	if err != nil {
		log.Fatalln("wrong database url")
	}

	sqlDb, _ := connection.DB()

	err = sqlDb.Ping()
	if err != nil {
		fmt.Println("can not connect to the database")
	} else {
		p.DB = connection
		fmt.Println("connected to database")
	}
}

func (p *Postgresql) Migrate() {
	_ = p.DB.AutoMigrate(p.Entities...)
}

func (p *Postgresql) AddToMigrate(entities ...interface{}) {
	p.Entities = append(p.Entities, entities...)
}
