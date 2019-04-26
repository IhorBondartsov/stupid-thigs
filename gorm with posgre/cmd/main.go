package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func main() {
	log.Println("RUN APP")
	defer fmt.Println("STOP APP")
	db, err := gorm.Open("postgres", "host=postgres port=5432 user=postgres dbname=TEST_SM password=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = db.DB().Ping()
	if err != nil {
		log.Fatal(err)
	}

	MigrateInformationTable(db)
	err =TransactionCreateInformation(db)
	if err != nil {
		log.Println(err)
	}
	Create(db)
	items := Read(db)
	if items != nil{
		for k, _ := range items{
			Delete(db, items[k].ID)
		}
	}

	defer db.Close()

}

//---------------------------------------------------------------------
//                             CRUD
//---------------------------------------------------------------------

func Create(db *gorm.DB) {
	log.Println("Create Row")
	i := Information{
		FirstName: "John",
		LastName:  "Smith",
	}
	resdb := db.Create(&i)
	err := resdb.Error
	if err != nil {
		log.Println(err)
		return
	}
	affected := db.RowsAffected
	log.Println("Affected rows:", affected, "LastID", i.ID)
}

func Read(db *gorm.DB) []*Information{
	items := []*Information{}
	resdb := db.Find(&items)
	err := resdb.Error
	if err != nil {
		log.Println(err)
		return nil
	}

	for _, v := range items{
		log.Println(*v)
	}
	return items
}

func Delete(db *gorm.DB, id uint) {
	log.Println("Delete id row")
	resdb := db.Delete(&Information{ID:id})
	err := resdb.Error
	if err != nil {
		log.Println(err)
		return
	}
}

//---------------------------------------------------------------------
//                             Transaction
//---------------------------------------------------------------------
func TransactionCreateInformation(db *gorm.DB) error {
log.Println("Start transaction")
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Create(&Information{LastName: "Giraffe"}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(&Information{FirstName: "Lion"}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

//---------------------------------------------------------------------
//                             TABLE
//---------------------------------------------------------------------
func MigrateInformationTable(db *gorm.DB) {
	log.Println("Create Table")
	//WARNING: AutoMigrate will ONLY create tables, missing columns and missing indexes,
	// and WON’T change existing column’s type or delete unused columns to protect your data.
	resdb := db.AutoMigrate(NewTableInformation(AdminRole))
	err := resdb.Error
	if err != nil {
		log.Println(err)
		return
	}
	affected := db.RowsAffected
	log.Println("Affected rows:", affected)
}

func CreateTable(db *gorm.DB){
	// Create table for model `User`
	//db.CreateTable(&User{})

	// will append "ENGINE=InnoDB" to the SQL statement when creating table `users`
	//db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
}

func DropTable(db *gorm.DB) {
	//// Drop model `User`'s table
	//db.DropTable(&User{})
	//
	//// Drop table `users`
	//db.DropTable("users")
	//
	//// Drop model's `User`'s table and table `products`
	//db.DropTableIfExists(&User{}, "products")
}

//---------------------------------------------------------------------
//                             Column
//---------------------------------------------------------------------
func ModifyColumn(db *gorm.DB) {
	// change column description's data type to `text` for model `User`
	//db.Model(&User{}).ModifyColumn("description", "text")
}

func DropColumn(db *gorm.DB){
	// Drop column description from model `User`
	//db.Model(&User{}).DropColumn("description")
}

//---------------------------------------------------------------------
//                             Index
//---------------------------------------------------------------------

func AddIndexes(db *gorm.DB) {
	// Add index for columns `name` with given name `idx_user_name`
	//db.Model(&User{}).AddIndex("idx_user_name", "name")

	// Add index for columns `name`, `age` with given name `idx_user_name_age`
	//db.Model(&User{}).AddIndex("idx_user_name_age", "name", "age")

	// Add unique index
	//db.Model(&User{}).AddUniqueIndex("idx_user_name", "name")

	// Add unique index for multiple columns
	//db.Model(&User{}).AddUniqueIndex("idx_user_name_age", "name", "age")
}

func RemoveIndex(db *gorm.DB){
	// Remove index
	//db.Model(&User{}).RemoveIndex("idx_user_name")

}


//---------------------------------------------------------------------
//                             Connection Pool
//---------------------------------------------------------------------
func ConnectionPool(db *gorm.DB){
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	//db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	//db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	//db.DB().SetConnMaxLifetime(time.Hour)

}
