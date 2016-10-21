#!/bin/sh
echo Building static spawn
CGO_ENABLED=0 go build -ldflags '-extldflags "-static"'
echo Building RKT ACI
tag=`git describe --tags 2>/dev/null`
acbuild begin
acbuild set-name astralboot/rocket/spawn
acbuild copy spawn /bin/spawn
acbuild set-exec /bin/spawn
acbuild label add version $tag
acbuild label add arch amd64
acbuild label add os linux
acbuild annotation add authors "Simon Kirkby tigger@interhingy.com"
acbuild write --overwrite spawn-$tag-linux-amd64.aci
acbuild end
gpg --sign --detach-sig -a spawn-$tag-linux-amd64.aci
echo Finished
