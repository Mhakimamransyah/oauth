package entities

type GithubUserInformation struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"avatar_url"`
}

func (obj *GithubUserInformation) ConvertToUser() *User {
	return &User{
		Name:  obj.Name,
		Email: obj.Email,
		Image: obj.Image,
	}
}
