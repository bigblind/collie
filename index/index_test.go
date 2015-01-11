package index

import (
    "testing"
    )

func NewTestDoc() Doc{
    type X struct{}
    x := X{}
    return NewDoc("doc_id", x)
}

func TestAddDocWithoutFieldsDoesNotRaiseError(t *testing.T) {
    index := NewIndex("index")
    doc := NewTestDoc()
    error := index.Put(&doc)
    if (error != nil){
        t.Fatalf("Indexing a valid doc raised an error: %v", error)
    }
}

func TestAddAndGetDoc(t *testing.T) {
    index := NewIndex("index")
    doc := NewTestDoc()
    index.Put(&doc)
    doc2 := index.Get("doc_id")
    if (&doc != doc2) {
        t.Fatal("document that was Put is unequal to the doc that was received from Get.")
    }
}
