package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitQuantity, exists := units[unit]
	if !exists {
		return false
	}

	bill[item] += unitQuantity
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitQuantity, unitExists := units[unit]
	if !unitExists {
		return false
	}

	_, itemExists := bill[item]
	if !itemExists {
		return false
	}

	bill[item] -= unitQuantity
	if bill[item] < 0 {
		bill[item] += unitQuantity
		return false
	}

	if bill[item] == 0 {
		delete(bill, item)
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemQuantity, itemExists := bill[item]
	return itemQuantity, itemExists
}
