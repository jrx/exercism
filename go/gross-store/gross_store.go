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
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {

	unit_value, unit_exists := units[unit]

	if !unit_exists {
		return false
	}

	bill[item] += unit_value
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {

	item_value, item_exits := bill[item]
	unit_value, unit_exists := units[unit]

	if !item_exits || !unit_exists || (unit_value > item_value) {
		return false
	} else if unit_value == item_value {
		delete(bill, item)
	} else {
		bill[item] -= unit_value
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	item_value, item_exits := bill[item]
	return item_value, item_exits
}
