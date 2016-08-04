package structs

import "fmt"

type Account struct {
	UserId      int
	accessToken string
}

func (m Account) GetToken() string {
	return m.accessToken
}

func (m *Account) SetToken(token string) {
	m.accessToken = token
}

type Request struct {
	resource string
}

type AuthRequest struct {
	Request
	username, password string
}

type T struct {
	I  int
	S  string
	xs []int
}

func main() {

	fmt.Println("\n==== Anonymous Struct ====")
	anonymousStruct()

	fmt.Println("\n==== Exported vs Unexported (Public vs Private) ====")
	member := Account{10, "abcd10"}
	publicVsPrivate(member)

	fmt.Println("\n==== Methods ====")
	methods(&member)

	fmt.Println("\n==== Composition ====")
	composition()

	fmt.Println("\n==== Copy struct ====")
	copyStruct()
}

func anonymousStruct() {
	data := struct {
		Number int
		Text   string
	}{
		42,
		"Moonmeander",
	}
	fmt.Println("", data)
}

func publicVsPrivate(member Account) {
	_, err := fmt.Println(" Accessing Public Variable (UerId)... ", member.UserId)
	if err != nil {
		panic(" Can't Access Public Variable")
	}

	_, err = fmt.Println(" Accessing Private Variable (accessToken)... ", member.accessToken)
	if err != nil {
		panic(" Can't Access Private Variable")
	}
}

func methods(m *Account) {
	token := "new token"
	fmt.Println(" Previous Token : ", m.GetToken())
	fmt.Println(" Setting Token", token, "...")
	m.SetToken(token)
	fmt.Println(" Current Token : ", m.GetToken())
}

func composition() {
	var ar AuthRequest
	var r Request
	r.resource = "stackoverflow.com/radbrawler"
	ar.Request = r
	ar.username = "radbrawler"
	ar.password = "password"
	fmt.Println(" Composite Stucture is: ", ar)
}

func copyStruct() {
	old := T{I: 1, S: "one", xs: []int{1, 2, 3}}
	new := old

	fmt.Println(" Value of old:", old)
	fmt.Println(" Value of new:", new)
	fmt.Println("\n Updating new.xs[1] = 500 ...")
	new.xs[1] = 500
	fmt.Println(" Value of old:", old)
	fmt.Println(" Value of new:", new)

	fmt.Println("\n Explicitly Initializing new's slice field {xs field (slice of int)} ...")
	new.xs = make([]int, len(old.xs))
	fmt.Println(" Updating 'new' value from 'old' using copy() ...")
	copy(new.xs, old.xs)
	fmt.Println(" Updating new.xs[1] = 500 ... \n ")
	new.xs[1] = 100

	fmt.Println(" Value of old:", old)
	fmt.Println(" Value of new:", new)

}
