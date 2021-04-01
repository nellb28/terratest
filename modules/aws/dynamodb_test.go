package aws

import (
	"fmt"
	"testing"
)

func TestGetContStatusE(t *testing.T) {
	t.Parallel()
	region := "us-east-1"
	tableName := "datastore-rrri"
	contStatus, err := GetContinuousBackupsStatusE(t, region, tableName)
	fmt.Println(contStatus, err)
}

func TestGetContStatus(t *testing.T) {
	t.Parallel()
	region := "us-east-1"
	tableName := "datastore-rrri"
	contStatus := GetContinuousBackupsStatus(t, region, tableName)
	fmt.Println(contStatus)
}
