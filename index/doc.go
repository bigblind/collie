package index

type Doc struct {
    id string
    version uint64
    data interface{}
    }

    func NewDoc(id string, data interface{}) Doc{
        return Doc{id: id, data: data, version: 0}
    }

    func (d *Doc) Id() string{
        return d.id
    }

    func (d *Doc) Version() uint64{
        return d.version
    }

    func (d *Doc) updateVersion(){
        d.version += 1
    }
    
