package main


import "fmt"

type Vehicule struct{
	Make string 
	Model string
	Year int 
}
type Car struct {
	CarVehicule Vehicule
	NumOfDoors int 
}
type Truck struct{
	TruckVehicule Vehicule
	PayloadCapacity int 
}

type Insurable interface {
	CalculateInsurance() float64

}
type Printable interface{
	Details() 
}

func main(){

	car := Car{
		CarVehicule: Vehicule{
			Make:  "Toyota",
			Model: "Corolla",
			Year:  2020,
		},
		NumOfDoors: 4,
	}

	truck := Truck{
		TruckVehicule: Vehicule{
			Make:  "Ford",
			Model: "F-150",
			Year:  2018,
		},
		PayloadCapacity: 1000,
	}

	fmt.Printf("Car insurance : %f ",car.CalculateInsurance())
	fmt.Printf("Truck insurance : %f ", truck.CalculateInsurance())

	var listOfprintables []Printable
	listOfprintables = append(listOfprintables,car,truck)
	PrintAll(listOfprintables)
}

func PrintAll(printables []Printable){
	for i,_ := range printables {
		printables[i].Details()
	}
}


func (t Truck) CalculateInsurance() float64{
	return float64(t.PayloadCapacity*10)
}

func (c Car) CalculateInsurance() float64{
	return float64(c.NumOfDoors*10)
}
func (t Truck) Details() {
	println("Truck Details:")
	println("Make:", t.TruckVehicule.Make)
	println("Model:", t.TruckVehicule.Model)
	println("Year:", t.TruckVehicule.Year)
	println("Payload Capacity:", t.PayloadCapacity)
}

func (c Car) Details() {
	println("Car Details:")
	println("Make:", c.CarVehicule.Make)
	println("Model:", c.CarVehicule.Model)
	println("Year:", c.CarVehicule.Year)
	println("Number of Doors:", c.NumOfDoors)
}
