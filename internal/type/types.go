package typesFile

type GeneralError struct {
	StatusCode int    `json:"statusCode"` //This json:""  set karta h response me field kese dikhegi
	Error      string `json:"error"`
}
type Student struct {
	Id    int
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Age   int    `json:"age" validate:"required"`
}
