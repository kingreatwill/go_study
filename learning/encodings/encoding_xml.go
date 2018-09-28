// GO语言中的xml对象的编解码(encoding/xml包);
package encodings

import "fmt"
import "encoding/xml"
import "github.com/kingreatwill/go_study/learning/common"

// 结构体转换为XML;
// func Marshal(v interface{}) ([]byte, error);
func struct2xml() {
	common.Println("encoding_json.go")

	type Address struct {
		City, State string
	}

	type Person struct {
		XMLName   xml.Name `xml:"person"`
		Id        int      `xml:"id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Age       int      `xml:"age"`
		Height    float32  `xml:"height,omitempty"`
		Married   bool
		Address
		Comment string `xml:",comment"`
	}

	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}
	//output, err := xml.Marshal(v)
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println(string(output))
	// <person id="13"><name><first>John</first><last>Doe</last></name><age>42</age><Married>false</Married><City>Hanga Roa</City><State>Easter Island</State><!-- Need more details. --></person>
	// <person id="13">
	//   <name>
	//     <first>John</first>
	//     <last>Doe</last>
	//   </name>
	//   <age>42</age>
	//   <Married>false</Married>
	//   <City>Hanga Roa</City>
	//   <State>Easter Island</State>
	//   <!-- Need more details. -->
	// </person>
}

// XML转换为结构体;
// func Unmarshal(data []byte, v interface{}) error;
func xml2struct() {
	type Email struct {
		Where string `xml:"where,attr"`
		Addr  string
	}

	type Address struct {
		City, State string
	}

	type Result struct {
		XMLName xml.Name `xml:"Person"`
		Name    string   `xml:"FullName"`
		Phone   string
		Email   []Email
		Groups  []string `xml:"Group>Value"`
		Address
	}

	v := Result{Name: "none", Phone: "none"}
	data := `
		<Person>
			<FullName>Grace R. Emlin</FullName>
			<Company>Example Inc.</Company>
			<Email where="home">
				<Addr>gre@example.com</Addr>
			</Email>
			<Email where='work'>
				<Addr>gre@work.com</Addr>
			</Email>
			<Group>
				<Value>Friends</Value>
				<Value>Squash</Value>
			</Group>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
		</Person>
	`
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("Name: %q\n", v.Name)
	fmt.Printf("Phone: %q\n", v.Phone)
	fmt.Printf("Email: %v\n", v.Email)
	fmt.Printf("Groups: %v\n", v.Groups)
	fmt.Printf("Address: %v\n", v.Address)
}
