package svgparser

// FindID finds the first child with the specified ID.
func (e *Element) FindID(id string) *Element {
	for _, child := range e.Children {
		for _, attr := range child.Attributes {
			if attr.Name.Local == "id" && attr.Value == id {
				return child
			}
		}
		if element := child.FindID(id); element != nil {
			return element
		}
	}
	return nil
}

// FindAll finds all children with the given name.
func (e *Element) FindAll(name string) []*Element {
	var elements []*Element
	for _, child := range e.Children {
		if child.Name.Local == name {
			elements = append(elements, child)
		}
		elements = append(elements, child.FindAll(name)...)
	}
	return elements
}

// FindAllBySpaceAndLocalName finds all children with the specific space and local name.
func (e *Element) FindAllBySpaceAndLocalName(space, localName string) []*Element {
	var elements []*Element
	for _, child := range e.Children {
		if child.Name.Space == space && child.Name.Local == localName {
			elements = append(elements, child)
		}
		elements = append(elements, child.FindAllBySpaceAndLocalName(space, localName)...)
	}
	return elements
}
