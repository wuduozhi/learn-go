package models

import (
	"fmt"
	"crypto/rand"
	"crypto/sha1"
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-web/pkg/logging"
)

var Db *sql.DB

func init(){
	var err error
	Db,err = sql.Open("mysql", "root:WudUozhI@tcp(127.0.0.1:3306)/goweb?parseTime=true")
	if err != nil {
		logging.Warn(err)
	}
	return
}

func createUUID() (uuid string) {
	u := new([16]byte)
	_,err := rand.Read(u[:])
	if err != nil {
		logging.Warn("Cannot generate UUID")
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
