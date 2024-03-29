package log // import "github.com/Lang0808/GolangLibs/log"

type LogEntry struct {
	CmdId int16

	SrcId string

	ExecTime int64

	Params map[string]interface{}

	ExtParams map[string]interface{}
}

func PrintError(err error) string {
	if err == nil {
		return "nil"
	}
	return err.Error()
}
