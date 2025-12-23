package action

type Action string

const (
	Insert Action = "insert"
	Update Action = "update"
)

func (a Action) String() string {
	return string(a)
}
