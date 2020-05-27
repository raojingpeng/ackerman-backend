package Models

type Query interface {
	QueryById()
	QueryByFirst()
	QueryByFind()
}
