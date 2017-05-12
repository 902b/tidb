// Copyright 2017 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package json

import (
	"github.com/pingcap/tidb/mysql"
	"github.com/pingcap/tidb/terror"
)

// JSON is for MySQL JSON type.
type JSON interface {
	// ParseFromString parses a json from string.
	ParseFromString(s string) error
	// DumpToString dumps itself to string.
	DumpToString() string
	// Serialize means serialize itself into bytes.
	Serialize() []byte
	// Deserialize means deserialize a json from bytes.
	Deserialize(bytes []byte)
}

// CreateJSON will create a json with bson as serde format and nil as data.
func CreateJSON(j interface{}) JSON {
	return &jsonImpl{
		json: j,
	}
}

var (
	// ErrInvalidJSONText means invalid JSON text.
	ErrInvalidJSONText = terror.ClassJSON.New(mysql.ErrInvalidJSONText, mysql.MySQLErrName[mysql.ErrInvalidJSONText])
	// ErrInvalidJSONPath means invalid JSON path.
	ErrInvalidJSONPath = terror.ClassJSON.New(mysql.ErrInvalidJSONPath, mysql.MySQLErrName[mysql.ErrInvalidJSONPath])
	// ErrInvalidJSONData means invalid JSON data.
	ErrInvalidJSONData = terror.ClassJSON.New(mysql.ErrInvalidJSONData, mysql.MySQLErrName[mysql.ErrInvalidJSONData])
)

func init() {
	terror.ErrClassToMySQLCodes[terror.ClassJSON] = map[terror.ErrCode]uint16{
		mysql.ErrInvalidJSONText: mysql.ErrInvalidJSONText,
		mysql.ErrInvalidJSONPath: mysql.ErrInvalidJSONPath,
		mysql.ErrInvalidJSONData: mysql.ErrInvalidJSONData,
	}
}
