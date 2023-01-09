package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := make(map[string]int)

	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, itemExists := bill[item]
	_, unitExists := units[unit]

	if !unitExists {
		return false
	}

	if !itemExists {
		bill[item] = units[unit]

		return true
	} else {
		bill[item] += units[unit]

		return true
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, itemExists := bill[item]
	_, unitExists := units[unit]

	if !itemExists {
		return false
	}

	if !unitExists {
		return false
	}

	tmp := bill[item] - units[unit]

	if tmp < 0 {
		return false
	} else if tmp == 0 {
		delete(bill, item)

		return true
	} else {
		bill[item] -= units[unit]

		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	_, itemExists := bill[item]

	if !itemExists {
		return 0, false
	}

	return bill[item], true
}
