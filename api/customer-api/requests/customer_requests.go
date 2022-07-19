package requests

type CreateCustomerRequest struct {
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"max=256,min=6"`
	LicenseID   string `json:"license_id"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Active      bool   `json:"active"`
}

type UpdateCustomerRequest struct {
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"max=256,min=6"`
	LicenseID   string `json:"license_id"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Active      bool   `json:"active"`
}

type ChangePasswordRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}
