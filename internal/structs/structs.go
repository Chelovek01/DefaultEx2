package structs

type Person struct {
	User_id int32  `json:"user_id,omitempty" db:"user_id"`
	Name    string `json:"name,omitempty" db:"name"`
	Adress  string `json:"adress,omitempty" db:"adress"`
	Balance int32  `json:"balance,omitempty" db:"balance"`
}

type Reserve struct {
	User_id      int32  `json:"user_id,omitempty" db:"user_id"`
	Name         string `json:"name,omitempty" db:"name"`
	Order_id     int32  `json:"order_id,omitempty" db:"order_id"`
	Service_id   int32  `json:"service_id,omitempty" db:"service_id"`
	Cost         int32  `json:"cost,omitempty" db:"cost"`
	Сonfirmation string `json:"сonfirmation,omitempty" db:"сonfirmation"`
}

type Profit struct {
	User_id    int32  `json:"user_id,omitempty" db:"user_id"`
	Name       string `json:"name,omitempty" db:"name"`
	Order_id   int32  `json:"order_id,omitempty" db:"order_id"`
	Service_id int32  `json:"service_id,omitempty" db:"service_id"`
	Sum        int32  `json:"sum,omitempty" db:"sum"`
}
