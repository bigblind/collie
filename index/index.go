package index

type Index struct {
    Name string
    docs map[string]*Doc
}

func NewIndex(name string) Index {
    return Index{Name: name, docs: make(map[string]*Doc)}
}


func (i *Index) Put(d *Doc) error {
    i.docs[d.Id()] = d
    return nil
}

func (i *Index) Get(id string) *Doc {
    return i.docs[id]
}


