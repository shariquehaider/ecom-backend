package models

import (
	"context"

	"github.com/shariquehaider/ecom-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type Address struct {
	AddressType string `json:"addressType,omitempty"`
	AddressLine string `json:"addressLine,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	Country     string `json:"country,omitempty"`
	PinCode     string `json:"pinCode,omitempty"`
	Phone       string `json:"phone,omitempty"`
}

func UpdateAddress(id string, addressDetails *Address) (int64, error) {
	objId, err := utils.VerifyObjectId(id)
	if err != nil {
		return 0, err
	}
	var update bson.M
	if addressDetails.AddressType == "billingAddress" {
		update = bson.M{"$set": bson.M{
			"billingAddress.addressLine": addressDetails.AddressLine,
			"billingAddress.city":        addressDetails.City,
			"billingAddress.state":       addressDetails.State,
			"billingAddress.country":     addressDetails.Country,
			"billingAddress.pinCode":     addressDetails.PinCode,
			"billingAddress.phone":       addressDetails.Phone,
		}}
	} else {
		update = bson.M{"$set": bson.M{
			"shippingAddress.addressLine": addressDetails.AddressLine,
			"shippingAddress.city":        addressDetails.City,
			"shippingAddress.state":       addressDetails.State,
			"shippingAddress.country":     addressDetails.Country,
			"shippingAddress.pinCode":     addressDetails.PinCode,
			"shippingAddress.phone":       addressDetails.Phone,
		}}
	}

	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objId}, update)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}
