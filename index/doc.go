package index

//Doc represents a document, that can be indexed.
type Doc struct {
    id string
    data interface{}
}

//Creates a new Doc with the given id and data.
func NewDoc(id string, data interface{}) Doc{
    return Doc{id: id, data: data}
}

//Doc.Id() returns the doc's id.
func (d *Doc) Id() string{
    return d.id
}

//doc.Data() returns the doc's data.
func (d *Doc) Data() interface{} {
    return d.data
}
