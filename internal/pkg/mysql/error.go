package mysql

import "errors"

const (
	mysqlDuplicateEntryErrCode = 1062
)

func GetMysqlSpecificError(errNumber int, err error) error {
	switch errNumber {
	case mysqlDuplicateEntryErrCode:
		return errors.New("duplicate entry")
	default:
		return err
	}
}
