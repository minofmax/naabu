package runner

import (
	"github.com/minofmax/naabu/v2/pkg/port"
	"github.com/minofmax/naabu/v2/pkg/protocol"
	"testing"

	"github.com/minofmax/naabu/v2/pkg/result"
	"github.com/minofmax/naabu/v2/pkg/scan"
	"github.com/stretchr/testify/assert"
)

func TestHandleNmap(t *testing.T) {
	// just attempt to start nmap
	var r Runner
	r.options = &Options{NmapCLI: "nmap -sV -oX - -Pn"}
	// nmap with empty cli shouldn't trigger any error
	res := result.NewResult()
	r.scanner = &scan.Scanner{}
	r.scanner.ScanResults = res
	var err error
	err, _ = r.handleNmap()
	assert.Nil(t, err)
	// nmap syntax error (this test might fail if nmap is not installed on the box)
	//assert.Nil(t, r.handleNmap())
	r.scanner.ScanResults.AddPort("127.0.0.1", &port.Port{Port: 8080, Protocol: protocol.TCP})
	//r.scanner.ScanResults.SetPorts("127.0.0.1", []*port.Port{{Port: 8080, Protocol: protocol.TCP}})
	err, _ = r.handleNmap()
	assert.Nil(t, err)
}
