#!/usr/bin/bash

pkgDir=$(basename $1)
prefix="natstool"
tarname=$prefix.$(date "+%Y%m%d").tar.gz

echo "pkging $pkgDir to $tarname ..."
tar -zcf  $tarname \
    --exclude nats/nats \
    --exclude tools \
    --exclude ms-client/ms-client \
    $pkgDir
