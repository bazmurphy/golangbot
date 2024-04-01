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

// Let’s say the organization has found a new income stream through advertisements.
// Let’s see how simple it is to add this new income stream and calculate the total income without making any changes to the calculateNetIncome function.
// This becomes possible because of polymorphism.

// we add an Advertisement
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

// we define the calculate method on Advertisement
func (a Advertisement) calculate() int {
	return a.CPC * a.noOfClicks
}

// we define the source method on Advertisement
func (a Advertisement) source() string {
	return a.adName
}

// You would have noticed that we did not make any changes to the calculateNetIncome function though we added a new income stream. It just worked because of polymorphism. Since the new Advertisement type also implemented the Income interface, we were able to add it to the incomeStreams slice. The calculateNetIncome function also worked without any changes as it was able to call the calculate() and source() methods of the Advertisement type.
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
	// We have created two ads namely bannerAd and popupAd.
	bannerAd := Advertisement{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
	popupAd := Advertisement{adName: "Popup Ad", CPC: 5, noOfClicks: 750}
	// The incomeStreams slice includes the two ads we just created.
	incomeStreams := []Income{project1, project2, project3, bannerAd, popupAd}
	calculateNetIncome(incomeStreams)
}

// Income From Project 1 = $5000
// Income From Project 2 = $10000
// Income From Project 3 = $4000
// Income From Banner Ad = $1000
// Income From Popup Ad = $3750
// Net income of organization = $23750
