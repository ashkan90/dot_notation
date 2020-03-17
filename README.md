# Golang -Struct Based String-Dot-Notation Accessor
Accessing structs with string never been get this much easy!! :D
## Usage 
##### !Note(You need to be careful when you're creating your struct. The fields have to be exported.)
```
type g struct {
	Customer struct {
		Details struct {
			Profile struct {
				Name string
			}
		}
	}
}

func main() {
    ng := &g{
    		Customer: struct{
    			Details struct{
    				Profile struct{
    					Name string
    				}
    			}
    		}{
    			Details: struct{
    				Profile struct{
    					Name string
    				}
    			}{
    				Profile: struct{
    					Name string
    				}{ Name: "Emirhan" }}},
    	}

    fmt.Println(Notate(ng, "Customer.Details.Profile.Name"))

}
```