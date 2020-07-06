package utils

import(
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
)



// Choose Driver
func GetConnect() (*gorm.DB, error){

	viper.SetConfigType("json")
    viper.AddConfigPath(".")
    viper.SetConfigName("app.config")

    err := viper.ReadInConfig()
    if err != nil {
        log.Fatal(err)
	}
	
	dns := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?parseTime=true", 
		viper.GetString("database.user"), 
		viper.GetString("database.pass"),
		viper.GetString("database.host"),
		viper.GetString("database.port"), 
		viper.GetString("database.dbname"),
	)

	db, err := gorm.Open("mysql", dns)
	if err != nil{
		return nil, err
	}

	return db, nil
}