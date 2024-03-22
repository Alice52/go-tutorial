package main

import "gorm.io/gen"

// https://gorm.io/gen/sql_annotation.html
type UserInterface interface {
	// Where("name=@name and age=@age")
	SimpleFindByNameAndAge(name string, age int) (gen.T, error)

	// select * from users where id=@id
	FindUserToMap(id int) (gen.M, error)

	// sql(insert into @@table (name,age) values (@name,@age) )
	InsertValue(age int, name string) error

	// select name from @@table where id=@id
	FindNameById(id int) string

	// select * from @@table
	//  {{where}}
	//      id>0
	//      {{if cond}}id=@id {{end}}
	//      {{if key!="" && value != ""}} or @@key=@value{{end}}
	//  {{end}}
	FindByIDOrCustom(cond bool, id int, key, value string) ([]gen.T, error)

	// update @@table
	//  {{set}}
	//      update_time=now()
	//      {{if name != ""}}
	//          name=@name
	//      {{end}}
	//  {{end}}
	//  {{where}}
	//      id=@id
	//  {{end}}
	UpdateName(name string, id int) (gen.RowsAffected, error)
}
