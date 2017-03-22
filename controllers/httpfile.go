package controllers


//上传文件集合
type HttpFileCollection []HttpPostedFile

type HttpPostedFile struct {
	ContentLength int64
	ContentType string
	FileName string
	InputStream []byte
}

func (file *HttpPostedFile) SaveAs(p string) error {
	return nil
}