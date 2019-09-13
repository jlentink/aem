package objects

type Projects struct {
	Project []ProjectRegistered
}

type ProjectRegistered struct {
	Name string
	Path string
}
