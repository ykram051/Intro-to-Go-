package main


import ("fmt"
"example.com/exo5/types")



func main(){

	car := types.Car{
		CarVehicule: types.Vehicule{
			Make:  "Toyota",
			Model: "Corolla",
			Year:  2020,
		},
		NumOfDoors: 4,
	}

	truck := types.Truck{
		TruckVehicule: types.Vehicule{
			Make:  "Ford",
			Model: "F-150",
			Year:  2018,
		},
		PayloadCapacity: 1000,
	}

	fmt.Printf("Car insurance : %f ",car.CalculateInsurance())
	fmt.Printf("Truck insurance : %f ", truck.CalculateInsurance())

	var listOfprintables []types.Printable
	listOfprintables = append(listOfprintables,car,truck)
	types.PrintAll(listOfprintables)
}
