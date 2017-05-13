package parameterbag

import "errors"

//Bag - interface for a parameter bag
type Bag interface {
	//Get a parameter bag value. If value does not exist should return ""
	Get(string) string
	//Set a parameter value
	Set(string, string) error
	//Has a parameter or not?
	Has(string) bool
	//List all parameters that exist in this bag
	ListParameters() []string
	//Freeze the parameter bag, if this is true then Set will return an error
	//One the bag is frozen it cannot be unfrozen but GetMutableCopy() may be called on the ParameterBag
	Freeze()
	//Returns the frozen state of the bag
	IsFrozen() bool
}

//ParameterBag - concrete definition of a parameter bag type
type ParameterBag struct {
	params map[string]string
	frozen bool
}

// NewParameterBag - create a new parameter bag with a blank list of parameters
func NewParameterBag() *ParameterBag {
	bag := make(map[string]string)
	return &ParameterBag{params: bag, frozen: false}

}

// NewParameterBagFromMap - Create a new parameter bag by passing in an existing map[string]string
func NewParameterBagFromMap(params map[string]string) *ParameterBag {
	return &ParameterBag{params: params, frozen: false}
}

func (b *ParameterBag) GetMutableCopy() *ParameterBag {

	mutableCopy := NewParameterBag()
	for k, v := range b.params {
		mutableCopy.params[k] = v
	}
	return mutableCopy
}

//Get returns the specified parameter
func (b *ParameterBag) Get(name string) string {
	return b.params[name]
}

//Set sets the specified parameter
func (b *ParameterBag) Set(name string, value string) error {
	if b.frozen == true {
		return errors.New("Parameter bag is frozen")
	}
	b.params[name] = value
	return nil
}

//Has this bag got the specified parameter?
func (b *ParameterBag) Has(name string) bool {
	return !(b.params[name] == "")
}

//ListParameters gets a slice of all the parameter keys in this bag
func (b *ParameterBag) ListParameters() []string {
	keys := make([]string, len(b.params))
	i := 0
	for k := range b.params {
		keys[i] = k
		i++
	}
	return keys
}

func (b *ParameterBag) Freeze() {
	b.frozen = true
}

func (b *ParameterBag) IsFrozen() bool {
	return b.frozen
}
