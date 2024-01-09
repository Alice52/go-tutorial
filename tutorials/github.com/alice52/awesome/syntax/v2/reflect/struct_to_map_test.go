package reflect

import (
	"encoding/json"
	"fmt"
)

type TestStruct1 struct {
	TestStruct1Field1 int
}

type TestStruct2 struct {
	TestStruct1       `structtomap:"teststruct1"`
	TestStruct2Field1 int
}

type TestStruct3 struct {
	StringField1  string `structtomap:"stringField1"`
	IntField1     int    `structtomap:"int_field_1"`
	StringField2  string
	ComplexField1 complex128
	TestStruct2
	TestStruct2Field1 TestStruct2
	ArrayField1       [2]int
	SliceField1       []int
	SliceField2       []*TestStruct2
	PointerField1     *int
	PointerField2     *TestStruct2
	MapField1         map[string]*TestStruct2
	MapField2         map[string]map[string]*TestStruct2
	MapField3         map[string]interface{}
}

func main() {
	if v, err := Decode(TestStruct3{
		StringField1:  "string field 1",
		IntField1:     100,
		StringField2:  "string field 2",
		ComplexField1: complex(1, 2),
		TestStruct2: TestStruct2{
			TestStruct2Field1: 1,
		},
		TestStruct2Field1: TestStruct2{
			TestStruct2Field1: 2,
		},
		ArrayField1: [2]int{1, 2},
		SliceField1: []int{1, 2, 3, 4},
		SliceField2: []*TestStruct2{
			{TestStruct2Field1: 1111, TestStruct1: TestStruct1{TestStruct1Field1: 1111}},
			{TestStruct2Field1: 2222},
			nil,
		},
		PointerField1: nil,
		PointerField2: &TestStruct2{
			TestStruct1:       TestStruct1{TestStruct1Field1: 1},
			TestStruct2Field1: 1,
		},
		MapField1: map[string]*TestStruct2{
			"1": {TestStruct2Field1: 1, TestStruct1: TestStruct1{TestStruct1Field1: 2}},
			"2": {TestStruct2Field1: 3},
			"3": nil,
		},
		MapField2: map[string]map[string]*TestStruct2{
			"1": {"11": nil, "12": &TestStruct2{
				TestStruct1:       TestStruct1{TestStruct1Field1: 1},
				TestStruct2Field1: 1,
			}},
		},
		MapField3: map[string]interface{}{
			"1": "1",
			"2": &TestStruct2{
				TestStruct1:       TestStruct1{TestStruct1Field1: 3},
				TestStruct2Field1: 1,
			},
		},
	}); err != nil {
		panic(err.Error())
	} else {
		if buf, err := json.MarshalIndent(v, "", "\t"); err != nil {
			panic(err.Error())
		} else {
			fmt.Println(string(buf))
		}
	}
}
