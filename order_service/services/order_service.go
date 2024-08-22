package services

import (
	"context"
	"errors"
	"order_service/config"
	"order_service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB connection and collection variables
var (
	client     *mongo.Client
	collection *mongo.Collection
)

func init() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		panic(err)
	}

	collection = client.Database("orderdb").Collection("orders")
}

// GetAllOrders retrieves all orders from the database
func GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var order models.Order
		if err = cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

// GetOrderByID retrieves an order by its ID from the database
func GetOrderByID(id string) (*models.Order, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid order ID")
	}

	var order models.Order
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &order, nil
}

// CreateOrder creates a new order in the database
func CreateOrder(order *models.Order) error {
	order.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.TODO(), order)
	return err
}

// UpdateOrder updates an existing order in the database
func UpdateOrder(id string, updatedOrder *models.Order) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid order ID")
	}

	update := bson.M{
		"$set": updatedOrder,
	}

	result, err := collection.UpdateByID(context.TODO(), objectID, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("order not found")
	}

	return nil
}

// DeleteOrder deletes an order by its ID from the database
func DeleteOrder(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid order ID")
	}

	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("order not found")
	}

	return nil
}
