package index

type Doc struct {
    id string
    data interface{}
}

func NewDoc(id string, data interface{}) Doc{
    return Doc{id: id, data: data}
}

func (d *Doc) Id() string{
    return d.id
}

func (d *Doc) Data() interface{} {
    return d.data
}
