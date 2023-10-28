package handler

import "fmt"

func errParamRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreatedOpening
type CreatedOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreatedOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Salary <= 0 && r.Link == "" {
		return fmt.Errorf("malformed request body")
	}
	if r.Role == "" {
		return errParamRequired("role", "string")
	}
	if r.Company == "" {
		return errParamRequired("company", "string")
	}
	if r.Location == "" {
		return errParamRequired("location", "string")
	}
	if r.Remote == nil {
		return errParamRequired("remote", "bool")
	}
	if r.Link == "" {
		return errParamRequired("link", "string")
	}
	if r.Salary <= 0 {
		return errParamRequired("salary", "int64")
	}

	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	//If any field  is provided, validation is truthy
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Salary > 0 || r.Link != "" {
		return nil
	}
	//If none of the fields were provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")
}
