package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Order represents the structure of an order in MongoDB
type Order struct {
    ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    CustomerID  string             `json:"customer_id" bson:"customer_id"`
    ProductID   string             `json:"product_id" bson:"product_id"`
    Quantity    int                `json:"quantity" bson:"quantity"`
    TotalPrice  float64            `json:"total_price" bson:"total_price"`
    OrderDate   string             `json:"order_date" bson:"order_date"`
}