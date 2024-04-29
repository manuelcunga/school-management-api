package util

type ErrorCode string

const (
	SchoolAlredayExists ErrorCode = "ups, esta escola já existe"
	NotFoundSchool      ErrorCode = "ups, esta escola já existe"
	InternalServerError ErrorCode = "ups, houve um erro interno tente novamente mais tarde"
)
