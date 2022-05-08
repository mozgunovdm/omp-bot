package mdi

var DataModel = []Bost{
	{Name: "Test1"},
	{Name: "Test2"},
	{Name: "Test3"},
	{Name: "Test4"},
	{Name: "Test5"},
	{Name: "Test6"},
	{Name: "Test7"},
	{Name: "Test8"},
	{Name: "Test9"},
	{Name: "Test11"},
	{Name: "Test12"},
	{Name: "Test13"},
	{Name: "Test14"},
	{Name: "Test15"},
	{Name: "Test16"},
	{Name: "Test17"},
	{Name: "Test18"},
}

type Bost struct {
	Name string
}

func (f Bost) String() string {
	return f.Name
}
