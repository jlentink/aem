package pom

// Open pom file
func Open(path string) (*Pom, error) {
	pom := Pom{}
	err := pom.Open(path)

	return &pom, err
}
