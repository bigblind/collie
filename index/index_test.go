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

func TestGetNonexistantDoc(t *testing.T){
	index := NewIndex("index")
	result := index.Get("DoesNotExist")
	if(result != nil){
		t.Fatalf("Expected index.Get of a doc that doesn't exist to return nil, got %v", result)
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

func TestDeletingADoc(t *testing.T){
	index := NewIndex("index")
	doc := NewTestDoc()
	index.Put(&doc)
	err := index.Delete(doc.Id())
	if (err != nil){
		t.Fatalf("Deleting an existing doc returned an error: %v", err)
	}

	result := index.Get(doc.Id())
	if (result != nil){
		t.Fatalf("expected the result of getting a deleted doc to be nil, got %v", result)
	}
}

func TestEmptySearchEmptyIndex(t *testing.T){
	index := NewIndex("index")
	result := index.Searchs("")
	if (len(result) != 0){
		t.Fatalf("An empty search of an empty index should return results of length 0.")
	}
}

func TestEmptySearchOnIndexWithDoc(t *testing.T){
	index := NewIndex("index")
	type x struct{}
	doc1 := NewDoc("doc1", x{})
	doc2 := NewDoc("doc2", x{})
	index.Put(&doc1)
	index.Put(&doc2)
	result := index.Searchs("")
	if(len(result) != 2){
		t.Fatalf("Expected search to return 2 results, got %v", result)
	}
}
