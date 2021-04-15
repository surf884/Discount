package main

func main() {
	count := 10
	price := 250
	sum := count * price
	discount := 30
	total := ((sum) / 100) * (100 - discount)
	println(total)
}
