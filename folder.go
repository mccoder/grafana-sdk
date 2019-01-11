package sdk

var (
	folderID uint
)

type Folder struct {
	ID        uint   `json:"id"`
	UID       string `json:"uid"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	HasAcl    bool   `json:"hasAcl"`
	CanSave   bool   `json:"canSave"`
	CanEdit   bool   `json:"canEdit"`
	CanAdmin  bool   `json:"canAdmin"`
	CreatedBy string `json:"createdBy"`
	Created   string `json:"created"`
	UpdatedBy string `json:"updatedBy"`
	Updated   string `json:"updated"`
	Version   int    `json:"version"`
}

func CreateFolder() *Folder {
	folderID += 1
	return &Folder{
		ID: folderID,
	}
}

func (f *Folder) SetTitle(title string) *Folder {
	f.Title = title
	return f
}

func (f *Folder) SetUID(uid string) *Folder {
	f.UID = uid
	return f
}
