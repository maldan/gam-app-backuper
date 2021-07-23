package backuper

type TestApi struct {
	Table string
}

func (f TestApi) GetIndex() {
	return "Test"
}
