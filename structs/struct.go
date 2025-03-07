package structs

type User struct {
	Id          int    `json:"id" db:"id"`
	FirstName   string `json:"firstName" db:"firstName"`
	LastName    string `json:"lastName" db:"lastName"`
	Email       string `json:"email" db:"email"`
	PhoneNumber int    `json:"phoneNumber" db:"phoneNumber"`
	Roles       string `json:"roles" db:"roles"`
	Password    string `json:"password" db:"password"`
}

type Listing struct {
	Id             int     `json:"id"`
	HostId         int     `json:"hostUId"`
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	Location       string  `json:"location"`
	Address        string  `json:"address"`
	MaxPeople      int     `json:"maxPeople"`
	PricePerNight  float64 `json:"pricePerNight"`
	CreatedAt      string  `json:"createdAt"`
	ApprovalStatus string  `json:"approvalStatus"`
}

type Booking struct {
	Id        int    `json:"id"`
	GuestId   int    `json:"guestId"`
	ListingId int    `json:"listingId"`
	Persons   int    `json:"persons"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Status    string `json:"status"`
}

type Approval struct {
	Id             int
	ApprovalTypeId string `json:"approvalTypeId"`
	ApprovalType   string `json:"approvalType"`
	ApproverId     int    `json:"approverId"`
	Status         string `json:"status"`
	CreatedAt      string `json:"createdAt"`
	updatedAt      string `json:"updatedAt"`
}

type Message struct {
	Code    int
	Error   bool
	Message string
	Data    interface{}
}

type JwtData struct {
	Id        int
	FirstName string
	Role      string
	Exp       interface{}
}
