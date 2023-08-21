package entities

type GoogleNames struct {
	Name string `json:"displayName"`
}

type GooglePhotos struct {
	Image string `json:"url"`
}

type GoogleEmailAddress struct {
	Email string `json:"value"`
}

type GoogleUserInformation struct {
	Name  []GoogleNames        `json:"names"`
	Photo []GooglePhotos       `json:"photos"`
	Email []GoogleEmailAddress `json:"emailAddresses"`
}

func (obj *GoogleUserInformation) ConvertToUser() *User {
	return &User{
		Name:  obj.Name[0].Name,
		Email: obj.Email[0].Email,
		Image: obj.Photo[0].Image,
	}
}
