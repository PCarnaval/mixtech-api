package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

//Create request

type CreateOpeningRequest struct {
	Role    string `json:"role"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Salary  int64  `json:"salary"`
	Remote  *bool  `json:"remote"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Name == "" && r.Company == "" && r.Email == "" && r.Salary == 0 && r.Remote == nil {
		return fmt.Errorf("request body is nil or malformed")
	}

	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "boolean")
	}
	return nil
}
