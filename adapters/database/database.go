package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database[T any] struct {
	connection *mongo.Database
}

func New[T any](db *mongo.Database) *Database[T] {
	return &Database[T]{
		connection: db,
	}
}

func (db *Database[T]) InsertOne(collName string, payload interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return db.connection.Collection(collName).InsertOne(ctx, payload)
}

func (db *Database[T]) ListAll(collName string) ([]*T, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cursor, err := db.connection.Collection(collName).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var output []*T

	for cursor.Next(ctx) {
		var doc *T
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}

		output = append(output, doc)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if output == nil {
		return []*T{}, nil
	}

	return output, nil
}

func (db *Database[T]) GetOne(collName string, filter primitive.M) (*T, error) {
	collection := db.connection.Collection(collName)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Consulte os documentos com base no filtro
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Itere sobre os documentos e imprima-os
	var doc *T
	for cursor.Next(ctx) {
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return doc, nil
}

// func (db *Database[T]) QueryCount(query string, args ...interface{}) (int, error) {
// 	row := db.pool.QueryRow(query, args...)

// 	var count int
// 	err := row.Scan(&count)
// 	if errors.Is(err, sql.ErrNoRows) {
// 		err = nil
// 	}

// 	return count, err
// }
