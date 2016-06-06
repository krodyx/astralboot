#!/bin/sh
echo Building static spawn
CGO_ENABLED=0 go build -ldflags '-extldflags "-static"'
echo Building RKT ACI
acbuild begin
acbuild set-name astralboot/rocket/spawn
acbuild copy spawn /bin/spawn
acbuild set-exec /bin/spawn
acbuild label add version 0.0.1
acbuild label add arch amd64
acbuild label add os linux
acbuild annotation add authors "Simon Kirkby tigger@interhingy.com"
acbuild write spawn-0.0.1-linux-amd64.aci
acbuild end
echo Finished
