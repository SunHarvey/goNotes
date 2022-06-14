package main

func LitersToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func MillilitersToGallons(m Milliliters) Gallons {
	return Gallons(m * 0.000264)
}

func GallonsToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func GallonsToMilliliters(g Gallons) Milliliters {
	return Milliliters(g * 3785.41)
}
