package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"
)

func main() {
	dsnr := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dbr, err := gorm.Open("mysql", dsnr)

	if err != nil {
		panic(err.Error())
	}
	err = querySomething(dbr)
	if err != nil {
		fmt.Printf("original err:%T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
		os.Exit(1)
	}
}

func querySomething(dbr *gorm.DB) error {
	var isRankingTemp int64
	err := dbr.Table("room_pk_pool").Select("rank_status").Where("union_id=222").Row().Scan(&isRankingTemp)
	if err != nil {
		if err != sql.ErrNoRows {
			//个人认为如果该数据对业务来说很关键，并且这个错误会导致业务逻辑上出现问题，需要上报
			return errors.Wrap(err, "err no rows")
		}
		return errors.Wrap(err, err.Error())
	}

	return nil
}
