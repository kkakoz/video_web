package safe

func Get[T any](f func() T) T {
	defer func() {
		recover()
	}()
	return f()
}

func GetDef[T any](f func() T, defValue T) (t T) {
	defer func() {
		recover()
		t = defValue
	}()
	return f()
}

type Repo[T any] struct {
}

func (Repo[T]) Get() T {
	return *new(T)
}

type User struct {
	Name string `json:"name"`
}

type Class struct {
	ID string `json:"name"`
}

type UserRepo struct {
	Repo[User]
}

type ClassRepo struct {
	Repo[Class]
}

func TestV() {

}
