package main 

import ( 
		"encoding/json"
		"fmt"
)

type LineItem struct {
		SKU string 
		Price float64 
		Discount float64 
		Quantity int 
}

func NewLineItem() LineItem {
		return LineItem{
			Quantity: 1,
		}
}


func unmarshalLineItem(data []byte) (LineItem, error) {
	li := NewLineItem()
	if err := json.Unmarshal(data, &li); err != nil {
			return LineItem{}, nil
	}
	if li.Quantity < 1 {
			return LineItem{}, fmt.Errorf("bad quantity")
	}
	return li, nil
}


func main() {
		data := []byte(`{"sku":"x3xs","price":1.2}`)

		li, err := unmarshalLineItem(data)
		if err != nil {
				fmt.Println("ERROR:", err)

		} else {
				fmt.Printf("%#v\n", li)
		}

		data = []byte(`{"sku": "x3xs", "price": 1.2, "quantity": 0}`)
		li, err = unmarshalLineItem(data)
		if err != nil {
			fmt.Println("ERROR:", err)
		} else {
				fmt.Printf("%#v\n",li)
		}
}