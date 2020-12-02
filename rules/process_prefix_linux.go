package rules

import (
	"fmt"
	"strings"

	C "github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/log"
)

type ProcessPrefix struct {
	Process
}

func (p *ProcessPrefix) RuleType() C.RuleType {
	return C.ProcessPrefix
}

func (p *ProcessPrefix) Match(metadata *C.Metadata) bool {
	key := fmt.Sprintf("%s:%s:%s", metadata.NetWork.String(), metadata.SrcIP.String(), metadata.SrcPort)
	cached, hit := processCache.Get(key)
	if !hit {
		processName, err := resolveProcessName(metadata)
		if err != nil {
			log.Debugln("[%s] Resolve process of %s failure: %s", C.Process.String(), key, err.Error())
		}

		processCache.Set(key, processName)

		cached = processName
	}
	if len(cached.(string)) < len(p.process) {
		return false
	}

	return strings.EqualFold(cached.(string)[:len(p.process)], p.process)
}

func NewProcessPrefix(process string, adapter string) (*ProcessPrefix, error) {
	return &ProcessPrefix{
		Process: Process{
			adapter: adapter,
			process: process,
		},
	}, nil
}
