package log // import "github.com/Lang0808/GolangLibs/log"

type LogEntry struct {
	cmdId int16

	srcId int32

	execTime int64

	params map[string]interface{}

	extParams map[string]interface{}
	
}
