package models

import (
	"lastlegends-proto-library/product-proto/pb"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Product struct {
	Id               bson.ObjectID     `bson:"_id" json:"id"`
	Name             string            `bson:"name" json:"name"`
	ShortDescription string            `bson:"shortdescription" json:"shortdescription"`
	Description      string            `bson:"description" json:"description"`
	Price            float64           `bson:"price" json:"price"`
	Sizes            []string          `bson:"sizes" json:"sizes"`
	Colors           []string          `bson:"colors" json:"colors"`
	Images           map[string]string `bson:"images" json:"images"`
	Tag              string            `bson:"tag" json:"tag"`
	Band             string            `bson:"band" json:"band"`
}

func NewProduct() *Product {
	return new(Product)
}

func (p *Product) ToProto() *pb.Product {
	return &pb.Product{
		Id:               p.Id.Hex(),
		Name:             p.Name,
		ShortDescription: p.ShortDescription,
		Description:      p.Description,
		Price:            p.Price,
		Sizes:            p.Sizes,
		Colors:           p.Colors,
		Images:           p.Images,
		Tag:              p.Tag,
		Band:             p.Band,
	}
}

func (p *Product) FromProto(product *pb.Product) {
	p.Id, _ = bson.ObjectIDFromHex(product.Id)
	p.Name = product.Name
	p.ShortDescription = product.ShortDescription
	p.Description = product.Description
	p.Price = product.Price
	p.Sizes = product.Sizes
	p.Colors = product.Colors
	p.Images = product.Images
	p.Tag = product.Tag
	p.Band = product.Band
}

func (p *Product) Collection() string {
	return "mt_products"
}
