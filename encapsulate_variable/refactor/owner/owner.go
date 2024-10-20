package owner

type Owner struct {
	firstName string
	lastName  string
}

var defaultOwner = Owner{firstName: "Martin", lastName: "Fowler"}

func GetDefaultOwner() Owner {
	return defaultOwner
}

func SetDefaultOwner(owner Owner) Owner {
	defaultOwner = owner
	return defaultOwner
}

func New(firstName, lastName string) Owner {
	return Owner{
		firstName: firstName,
		lastName:  lastName,
	}
}
