# Polymorphism - OOP in Go

Polymorphism in Go is achieved with the help of [interfaces](https://golangbot.com/interfaces-part-1/). As we have already discussed, [interfaces](https://golangbot.com/interfaces-part-1/) are implicitly implemented in Go. A type implements an interface if it provides definitions for all the [methods](https://golangbot.com/methods/) declared in the interface. Let’s see how polymorphism is achieved in Go with the help of interfaces.

## Polymorphism using an interface[](https://golangbot.com/polymorphism/#polymorphism-using-an-interface)

Any type which provides definition for all the methods of an interface is said to implicitly implement that interface. This will be more clear as we discuss an example of polymorphism shortly.

**A variable of type interface can hold any value which implements the interface. This property of interfaces is used to achieve polymorphism in Go.**

Let’s understand polymorphism in Go with the help of a program that calculates the net income of an organization. For simplicity, let’s assume that this imaginary organization has income from two kinds of projects viz. _fixed billing_ and _time and material_. The net income of the organization is calculated by the sum of the incomes from these projects. To keep this tutorial simple, we will assume that the currency is dollars and we will not deal with cents. It will be represented using `int`. (I recommend reading [this](https://forum.golangbridge.org/t/what-is-the-proper-golang-equivalent-to-decimal-when-dealing-with-money/413) to learn how to represent cents.

Let’s first define an interface `Income`.

```go
type Income interface {
	calculate() int
	source() string
}
```

The `Income` interface defined above contains two methods `calculate()` which calculates and returns the income from the source and `source()` which returns the name of the source.

Next, let’s define a struct for `FixedBilling` project type.

```go
type FixedBilling struct {
	projectName string
	biddedAmount int
}
```

The `FixedBilling` project has two fields `projectName` which represents the name of the project and `biddedAmount` which is the amount that the organization has bid for the project.

The `TimeAndMaterial` struct will represent projects of Time and Material type.

```go
type TimeAndMaterial struct {
	projectName string
	noOfHours  int
	hourlyRate int
}
```

The `TimeAndMaterial` struct has three fields names `projectName`, `noOfHours` and `hourlyRate`.

The next step would be to define methods on these struct types which calculate and return the actual income and source of income.

```go
func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}
```

In the case of `FixedBilling` projects, the income is just the amount bid for the project. Hence we return this from the `calculate()` method of `FixedBilling` type.

In the case of `TimeAndMaterial` projects, the income is the product of the `noOfHours` and `hourlyRate`. This value is returned from the `calculate()` method with receiver type `TimeAndMaterial`.

We return the name of the project as the source of income from the `source()` method.

Since both `FixedBilling` and `TimeAndMaterial` structs provide definitions for the `calculate()` and `source()` methods of the `Income` interface, both structs implement the `Income` interface.

Let’s declare the `calculateNetIncome` function which will calculate and print the total income.

```go
func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organization = $%d", netincome)
}
```

The `calculateNetIncome` [function](https://golangbot.com/functions/) above accepts a [slice](https://golangbot.com/arrays-and-slices/#slices) of `Income` interfaces as argument. It calculates the total income by iterating over the slice and calling `calculate()` method on each of its items. It also displays the income source by calling `source()` method. Depending on the concrete type of the `Income` interface, different `calculate()` and `source()` methods will be called. Thus we have achieved polymorphism in the `calculateNetIncome` function.

In the future, if a new kind of income source is added by the organization, this function will still calculate the total income correctly without a single line of code change :).

The only part remaining in the program is the main function.

```go
func main() {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	incomeStreams := []Income{project1, project2, project3}
	calculateNetIncome(incomeStreams)
}
```

In the `main` function above we have created three projects, two of type `FixedBilling` and one of type `TimeAndMaterial`. Next, we create a slice of type `Income` with these 3 projects. Since each of these projects has implemented the `Income` interface, it is possible to add all three projects to a slice of type `Income`. Finally, we call `calculateNetIncome` function and pass this slice as a parameter. It will display the various income sources and the income from them.

Here is the full program for your reference.

```go
package main

import (
	"fmt"
)

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours  int
	hourlyRate int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organization = $%d", netincome)
}

func main() {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	incomeStreams := []Income{project1, project2, project3}
	calculateNetIncome(incomeStreams)
}
```

This program will output

```
Income From Project 1 = $5000
Income From Project 2 = $10000
Income From Project 3 = $4000
Net income of organization = $19000
```

## Adding a new income stream to the above program[](https://golangbot.com/polymorphism/#adding-a-new-income-stream-to-the-above-program)

Let’s say the organization has found a new income stream through advertisements. Let’s see how simple it is to add this new income stream and calculate the total income without making any changes to the `calculateNetIncome` function. This becomes possible because of polymorphism.

Lets first define the `Advertisement` type and the `calculate()` and `source()` methods on the `Advertisement` type.

```go
type Advertisement struct {
	adName     string
	CPC        int
	noOfClicks int
}

func (a Advertisement) calculate() int {
	return a.CPC * a.noOfClicks
}

func (a Advertisement) source() string {
	return a.adName
}
```

The `Advertisement` type has three fields `adName`, `CPC` (cost per click) and `noOfClicks` (number of clicks). The total income from ads is the product of `CPC` and `noOfClicks`.

Let’s modify the `main` function a little to include this new income stream.

```go
func main() {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	bannerAd := Advertisement{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
	popupAd := Advertisement{adName: "Popup Ad", CPC: 5, noOfClicks: 750}
	incomeStreams := []Income{project1, project2, project3, bannerAd, popupAd}
	calculateNetIncome(incomeStreams)
}
```

We have created two ads namely `bannerAd` and `popupAd`. The `incomeStreams` slice includes the two ads we just created.

Here is the full program after adding Advertisement.

```go
package main

import (
	"fmt"
)

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

type Advertisement struct {
	adName     string
	CPC        int
	noOfClicks int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func (a Advertisement) calculate() int {
	return a.CPC * a.noOfClicks
}

func (a Advertisement) source() string {
	return a.adName
}
func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organization = $%d", netincome)
}

func main() {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	bannerAd := Advertisement{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
	popupAd := Advertisement{adName: "Popup Ad", CPC: 5, noOfClicks: 750}
	incomeStreams := []Income{project1, project2, project3, bannerAd, popupAd}
	calculateNetIncome(incomeStreams)
}
```

The above program will output,

```
Income From Project 1 = $5000
Income From Project 2 = $10000
Income From Project 3 = $4000
Income From Banner Ad = $1000
Income From Popup Ad = $3750
Net income of organization = $23750
```

You would have noticed that we did not make any changes to the `calculateNetIncome` function though we added a new income stream. It just worked because of polymorphism. Since the new `Advertisement` type also implemented the `Income` interface, we were able to add it to the `incomeStreams` slice. The `calculateNetIncome` function also worked without any changes as it was able to call the `calculate()` and `source()` methods of the `Advertisement` type.
