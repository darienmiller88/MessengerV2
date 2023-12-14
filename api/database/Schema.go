package database

import(

)

var schema string = `
CREATE TABLE users (
    
);

CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer
)`

func GetSchema() string{
	return schema
}