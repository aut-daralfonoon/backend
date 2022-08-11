package datastore

import "gorm.io/gorm"

type config struct {
}

func (c config) toConnStr() string {
	return ""
}

func NewDBConn() {

}

func GetorCreateDBConn() *gorm.DB {
	return nil
}
