package objects

// Projects Registered projects
type Projects struct {
	Project []ProjectRegistered
}

// ProjectRegistered registered project
type ProjectRegistered struct {
	Name string
	Path string
}
