package main

/**
 * @param id: store id
 * @param name: store name
 * @param genre: store genre
 * @param averagePrice: average price
 * @param tel: phone number
 * @param lat: latitude
 * @param long: longitude
 * @param isFavourite: if the restaurant in the fav list
 */
type Restaurant struct {
	Id           int
	Name         string
	Address      string
	Genre        string
	AveragePrice float32
	Tel          string
	Lat          float32
	Long         float32
	IsFavourite  bool
}
