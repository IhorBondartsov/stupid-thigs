package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	// Unmarshal
	jsonUser := `{"created_at":1556787245,"id":1,"name":"Ken"}`
	pm := MyUser{}
	err := json.Unmarshal([]byte(jsonUser), &pm)
	if err != nil {
		fmt.Println("Unmarshal: ", err)
	}
	fmt.Println("Unmarshal: ", pm)

	// Marshaling
	m := MyUser{1, "Ken", time.Now()}
	b, err := json.Marshal(&m)
	if err != nil {
		fmt.Println("Marshal: ", err)
	}
	fmt.Println("Marshal: ", string(b))

	// Unmarshal using like method
	pm = MyUser{}
	err = pm.UnmarshalJSON([]byte(jsonUser))
	if err != nil {
		fmt.Println("UnmarshalJSON: ", err)
	}
	fmt.Println("UnmarshalJSON: ", pm)

	// MarshalJSON
	b, err = m.MarshalJSON()
	if err != nil {
		fmt.Println("MarshalJSON: ", err)
	}
	fmt.Println("MarshalJSON: ", string(b))

	// Make big problem

	// UnmarshalJSON
	// in this case we have
	//runtime: goroutine stack exceeds 1000000000-byte limit
	//fatal error: stack overflow
	//-------------------------------------------------------------
	wu := MyWrongUser{}
	err = wu.UnmarshalJSON([]byte(jsonUser))
	if err != nil {
		fmt.Println("UnmarshalJSON Wrong User: ", err)
	}
	fmt.Println("UnmarshalJSON Wrong User: ", wu)
	//-------------------------------------------------------------

	// MarshalJSON
	// in this case we have
	//runtime: goroutine stack exceeds 1000000000-byte limit
	//fatal error: stack overflow
	//-------------------------------------------------------------
	wu2 := MyWrongUser{
		CreatedAt: time.Now(),
		Name:      "User",
		ID:        1}
	b, err = wu2.MarshalJSON()
	if err != nil {
		fmt.Println("MarshalJSON Wrong User: ", err)
	}
	fmt.Println("MarshalJSON Wrong User: ", string(b))
}

type MyUser struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *MyUser) MarshalJSON() ([]byte, error) {
	type Alias MyUser
	return json.Marshal(&struct {
		CreatedAt int64 `json:"created_at"`
		*Alias
	}{
		CreatedAt: u.CreatedAt.Unix(),
		Alias:     (*Alias)(u),
	})
}

func (u *MyUser) UnmarshalJSON(data []byte) error {
	type Alias MyUser
	aux := &struct {
		CreatedAt int64 `json:"created_at"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	u.CreatedAt = time.Unix(aux.CreatedAt, 0)
	return nil
}

type MyWrongUser struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *MyWrongUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(u)
}

func (u *MyWrongUser) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, u); err != nil {
		return err
	}
	return nil
}
