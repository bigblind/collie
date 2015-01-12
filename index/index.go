package index

//Index represents a search index, a searchable store of documents.
type Index struct {
    Name string
    docs map[string]*Doc
}

//NewIndex creates a new Index with the given name.
func NewIndex(name string) Index {
    return Index{Name: name, docs: make(map[string]*Doc)}
}

//Index.Put puts a new document in the index, returns an error if the document isn't valid, nil otherwise.
func (i *Index) Put(d *Doc) error {
    i.docs[d.Id()] = d
    return nil
}

//Index.Get gets the document with the given id from the index, returns nil if no such doc exists.
func (i *Index) Get(id string) *Doc {
    return i.docs[id]
}

//Deletes the document with the given id from the index. If no such doc is found, this function is a noop.
func (i *Index) Delete(id string) error {
	delete(i.docs, id)
	return nil
}

//Index.Search searches the index using a query string.
func(i *Index) Searchs(query string) []Doc {
	r := []Doc{}
	for _, doc := range i.docs {
		r = append(r, *doc)
	}
	return r
}


