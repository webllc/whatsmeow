// Copyright (c) 2021 Tulir Asokan
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package whatsmeow implements a client for interacting with the WhatsApp web multidevice API.
package whatsmeow

import (
	"bytes"
	"encoding/xml"
	"fmt"
	waBinary "go.mau.fi/whatsmeow/binary"
	"go.mau.fi/whatsmeow/types"
)


type OrderDetailType struct {
	Order struct {
		CreationTs string `xml:"creation_ts,attr"`
		ID         string `xml:"id,attr"`
		Product    []struct {
			ID         string `xml:"id"`
			RetailerID string `xml:"retailer_id"`
			ImageUrl   string `xml:"image>url"`
			Price    string `xml:"price"`
			Currency string `xml:"currency"`
			Name     string `xml:"name"`
			Quantity string `xml:"quantity"`
		} `xml:"product"`
		Catalog struct {
			ID   string `xml:"id"`
		} `xml:"catalog"`
		Price struct {
			Subtotal    string `xml:"subtotal"`
			Currency    string `xml:"currency"`
			Total       string `xml:"total"`
			PriceStatus string `xml:"price_status"`
		} `xml:"price"`
	} `xml:"order"`
}

func (cli *Client) GetCatalog(jid types.JID, limit int) (*waBinary.Node, error) {
	catalogNode, err := cli.sendIQ(infoQuery{
		Namespace: "w:biz:catalog",
		Type:      "get",
		To:        types.ServerJID,
		Content: []waBinary.Node{
			{
				Tag: "product_catalog",
				Attrs: waBinary.Attrs{
					"jid":               jid,
					"allow_shop_source": true,
				},
				Content: []waBinary.Node{
					{
						Tag:     "limit",
						Attrs:   nil,
						Content: []byte("10"),
					},
					{
						Tag:     "width",
						Attrs:   nil,
						Content: []byte("100"),
					}, {
						Tag:     "height",
						Attrs:   nil,
						Content: []byte("100"),
					},
				},
			},
		},
	})
	return catalogNode, err
}



func (cli *Client) AddProduct(product types.Product) (*waBinary.Node, error) {
	var content [] waBinary.Node
	content = append(content,
		waBinary.Node{
			Tag: "name",
			Attrs: nil,
			Content: []byte(product.Name),
		})
	content = append(content,
		waBinary.Node{
			Tag: "description",
			Attrs: nil,
			Content: []byte(product.Description),
		})
	content = append(content,
		waBinary.Node{
			Tag: "price",
			Attrs: nil,
			Content: []byte(product.Price),
		})
	content = append(content,
		waBinary.Node{
			Tag: "retailer_id",
			Attrs: nil,
			Content: []byte(product.RetailerId),
		})

	content = append(content,
		waBinary.Node{
			Tag: "currency",
			Attrs: nil,
			Content: []byte(product.Currency),
		})

	content = append(content,
		waBinary.Node{
			Tag: "media",
			Attrs: nil,
			Content: waBinary.Node{
				Tag: "image",
				Attrs: nil,
				Content: []byte(product.Url),
			},
		})

	query := infoQuery{
		Namespace: "w:biz:catalog",
		Type:      "set",
		To:        types.ServerJID,
		Content: []waBinary.Node{
			{
				Tag: "product_catalog_add",
				Attrs: waBinary.Attrs{
					"v": "1",
				},
				Content: []waBinary.Node{
					{
						Tag:     "product",
						Attrs:   waBinary.Attrs{
							"is_hidden": "false",
						},
						Content: content,
					},
				},
			},
		},
	}


	fmt.Printf("Query: %+v", query)
	fmt.Println("")

	productNode, err := cli.sendIQ(query)
	return productNode, err
}




func (cli *Client) GetOrderDetails(orderId, tokenBase64 string) (*OrderDetailType, error) {

	detailsNode, err := cli.sendIQ(infoQuery{
		Namespace: "fb:thrift_iq",
		Type:      "get",
		To:        types.ServerJID,
		SmaxId:    "5",
		Content: []waBinary.Node{
			{
				Tag: "order",
				Attrs: waBinary.Attrs{
					"op": "get",
					"id": orderId,
				},
				Content: []waBinary.Node{
					{
						Tag:   "image_dimensions",
						Attrs: nil,
						Content: []waBinary.Node{
							{
								Tag:     "width",
								Attrs:   nil,
								Content: []byte("100"),
							}, {
								Tag:     "height",
								Attrs:   nil,
								Content: []byte("100"),
							},
						},
					},
					{
						Tag:     "token",
						Attrs:   nil,
						Content: []byte(tokenBase64),
					},
				},
			},
		},
	})
	
	OrderDetail := &OrderDetailType{}
	d := xml.NewDecoder(bytes.NewReader([]byte(detailsNode.XMLString())))
	d.Strict = false
	err = d.Decode(&OrderDetail)
	if err != nil {
		cli.Log.Infof("Order Details response Error: %v",err)
	}

	return OrderDetail, err
}
