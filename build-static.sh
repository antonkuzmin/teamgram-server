#!/bin/sh
set -e

PWD=`pwd`
TEAMGRAMAPP=${PWD}"/app"
INSTALL=${PWD}"/teamgramd"
LDFLAGS="-s -w -linkmode external -extldflags '-static'"

echo "build idgen ..."
cd ${TEAMGRAMAPP}/service/idgen/cmd/idgen
go build -ldflags "${LDFLAGS}" -o ${INSTALL}/bin/idgen

echo "build status ..."
cd ${TEAMGRAMAPP}/service/status/cmd/status
go build -ldflags "${LDFLAGS}" -o ${INSTALL}/bin/status

echo "build dfs ..."
cd ${TEAMGRAMAPP}/service/dfs/cmd/dfs
go build -ldflags "${LDFLAGS}" -o ${INSTALL}/bin/dfs

echo "build media ..."
cd ${TEAMGRAMAPP}/service/media/cmd/media
go build -ldflags "${LDFLAGS}" -o ${INSTALL}/bin/media

echo "build authsession ..."
cd ${TEAMGRAMAPP}/service/authsession/cmd/authsession
go build -ldflags "${LDFLAGS}" -o ${INSTALL}/bin/authsession

echo "build biz ..."
cd ${TEAMGRAMAPP}/service/biz/biz/cmd/biz
go build -ldflags "${LDFLAGS}" -o ${INSTALL}/bin/biz

echo "build msg ..."
cd ${TEAMGRAMAPP}/messenger/msg/cmd/msg
go build -ldflags "${LDFLAGS}" -o ${INSTALL}/bin/msg

echo "build sync ..."
cd ${TEAMGRAMAPP}/messenger/sync/cmd/sync
go build -ldflags "${LDFLAGS}" -o ${INSTALL}/bin/sync

echo "build bff ..."
cd ${TEAMGRAMAPP}/bff/bff/cmd/bff
go build -ldflags "${LDFLAGS}" -o ${INSTALL}/bin/bff

echo "build session ..."
cd ${TEAMGRAMAPP}/interface/session/cmd/session
go build -ldflags "${LDFLAGS}" -o ${INSTALL}/bin/session

echo "build gnetway ..."
cd ${TEAMGRAMAPP}/interface/gnetway/cmd/gnetway
go build -ldflags "${LDFLAGS}" -o ${INSTALL}/bin/gnetway

echo "build otp-tg-sender ..."
cd ${TEAMGRAMAPP}/service/otp-tg-sender
go build -ldflags "${LDFLAGS}" -o ${INSTALL}/bin/otp-tg-sender .

echo ""
echo "=== verifying binaries ==="
BINARIES="idgen status dfs media authsession biz msg sync bff session gnetway"
FAILED=""
for b in $BINARIES; do
    if [ -x "${INSTALL}/bin/$b" ]; then
        echo "  OK  $b"
    else
        echo "  FAIL  $b  (not found or not executable)"
        FAILED="$FAILED $b"
    fi
done

if [ -n "$FAILED" ]; then
    echo ""
    echo "ERROR: missing binaries:${FAILED}"
    exit 1
fi
echo "=== all binaries OK ==="
