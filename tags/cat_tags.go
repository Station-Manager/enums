package tags

type CatStateTag string

func (ca CatStateTag) String() string {
	return string(ca)
}

const (
	Identity CatStateTag = "IDENTITY"
	VfoAFreq CatStateTag = "VFOAFREQ"
	VfoBFreq CatStateTag = "VFOBFREQ"
	Split    CatStateTag = "SPLIT"
	Select   CatStateTag = "SELECT"
	MainMode CatStateTag = "MAINMODE"
	SubMode  CatStateTag = "SUBMODE"
	TxPwr    CatStateTag = "TXPWR"
)

var AllCatStateTags = []struct {
	Value  CatStateTag
	TSName string
}{
	{Identity, Identity.String()},
	{VfoAFreq, VfoAFreq.String()},
	{VfoBFreq, VfoBFreq.String()},
	{Split, Split.String()},
	{Select, Select.String()},
	{MainMode, MainMode.String()},
	{SubMode, SubMode.String()},
	{TxPwr, TxPwr.String()},
}
