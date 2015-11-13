package settings

type Langs struct {
	Name        string
	Type        string
	Precommand  string
	Command     string
	Postcommand string
}

type Specs struct {
	Type   string
	Method string
	Data   interface{}
	Input  string
	Output string
	Result interface{}
}

type Codes struct {
	Type     string
	Contents string
}

type Funcs struct {
	Name string
	Code []Codes
}

type Settings struct {
	Languages []Langs
	Tests     []Specs
	Functions []Funcs
}
