package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var libraryDB *gorm.DB

func Init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "library_rw:library_rw4HomeLib@tcp(127.0.0.1:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		os.Exit(-1)
	}
	libraryDB = db
}

/*
CREATE DATABASE IF NOT EXISTS library DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;

create user 'library_rw'@'%' identified by 'library_rw4HomeLib';
create user 'library_r'@'%' identified by 'library_r4HomeLib';

//grant select,insert,update,delete,create on library.* to library_rw;
grant all privileges on library.* to library_rw;
grant select on library.* to library_r;

flush  privileges;


>> mysql -ulibrary_rw -plibrary_rw4HomeLib
*/

/*
ALTER USER 'library_rw'@'%' IDENTIFIED WITH mysql_native_password BY 'library_rw4HomeLib';
ALTER USER 'library_r'@'%' IDENTIFIED WITH mysql_native_password BY 'library_r4HomeLib';


*/
