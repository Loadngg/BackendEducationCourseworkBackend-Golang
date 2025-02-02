// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type lectureMaterialsTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *lectureMaterialsTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("lecture_materials").
func (v *lectureMaterialsTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *lectureMaterialsTableType) Columns() []string {
	return []string{
		"id",
		"title",
		"file_link",
		"lecture_id",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *lectureMaterialsTableType) NewStruct() reform.Struct {
	return new(LectureMaterials)
}

// NewRecord makes a new record for that table.
func (v *lectureMaterialsTableType) NewRecord() reform.Record {
	return new(LectureMaterials)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *lectureMaterialsTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// LectureMaterialsTable represents lecture_materials view or table in SQL database.
var LectureMaterialsTable = &lectureMaterialsTableType{
	s: parse.StructInfo{
		Type:    "LectureMaterials",
		SQLName: "lecture_materials",
		Fields: []parse.FieldInfo{
			{Name: "ID", Type: "int32", Column: "id"},
			{Name: "Title", Type: "string", Column: "title"},
			{Name: "FileLink", Type: "string", Column: "file_link"},
			{Name: "LectureID", Type: "*int32", Column: "lecture_id"},
		},
		PKFieldIndex: 0,
	},
	z: new(LectureMaterials).Values(),
}

// String returns a string representation of this struct or record.
func (s LectureMaterials) String() string {
	res := make([]string, 4)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Title: " + reform.Inspect(s.Title, true)
	res[2] = "FileLink: " + reform.Inspect(s.FileLink, true)
	res[3] = "LectureID: " + reform.Inspect(s.LectureID, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *LectureMaterials) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Title,
		s.FileLink,
		s.LectureID,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *LectureMaterials) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Title,
		&s.FileLink,
		&s.LectureID,
	}
}

// View returns View object for that struct.
func (s *LectureMaterials) View() reform.View {
	return LectureMaterialsTable
}

// Table returns Table object for that record.
func (s *LectureMaterials) Table() reform.Table {
	return LectureMaterialsTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *LectureMaterials) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *LectureMaterials) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *LectureMaterials) HasPK() bool {
	return s.ID != LectureMaterialsTable.z[LectureMaterialsTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.ID = pk.
func (s *LectureMaterials) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = LectureMaterialsTable
	_ reform.Struct = (*LectureMaterials)(nil)
	_ reform.Table  = LectureMaterialsTable
	_ reform.Record = (*LectureMaterials)(nil)
	_ fmt.Stringer  = (*LectureMaterials)(nil)
)

func init() {
	parse.AssertUpToDate(&LectureMaterialsTable.s, new(LectureMaterials))
}
