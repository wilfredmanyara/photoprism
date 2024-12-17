#!/usr/bin/env bash

# This installs latest Go on Linux.
# bash <(curl -s https://raw.githubusercontent.com/photoprism/photoprism/develop/scripts/dist/install-go.sh)

PATH="/usr/local/sbin:/usr/sbin:/sbin:/usr/local/bin:/usr/bin:/bin:/scripts:$PATH"

DESTDIR=$(realpath "${1:-/usr/local}")

# Query version.
if [[ -z $GOLANG_VERSION ]]; then
  GOLANG_VERSION=$(curl -fsSL https://go.dev/VERSION?m=text | head -n 1)
fi

echo "Installing ${GOLANG_VERSION} in \"${DESTDIR}\"..."

set -e

# Determine the system architecture.
if [[ $PHOTOPRISM_ARCH ]]; then
  SYSTEM_ARCH=$PHOTOPRISM_ARCH
else
  SYSTEM_ARCH=$(uname -m)
fi

DESTARCH=${BUILD_ARCH:-$SYSTEM_ARCH}

sudo mkdir -p "$DESTDIR"

set -eux;

case $DESTARCH in
  amd64 | AMD64 | x86_64 | x86-64)
    URL="https://go.dev/dl/${GOLANG_VERSION}.linux-amd64.tar.gz"
    ;;

  arm64 | ARM64 | aarch64)
    URL="https://go.dev/dl/${GOLANG_VERSION}.linux-arm64.tar.gz"
    ;;

  arm | ARM | aarch | armv7l | armhf)
    URL="https://go.dev/dl/${GOLANG_VERSION}.linux-armv6l.tar.gz"
    ;;

  *)
    echo "Unsupported Machine Architecture: \"$DESTARCH\"" 1>&2
    exit 1
    ;;
esac

# Replace current installation in "/usr/local/go".
echo "Installing Go for ${DESTARCH^^} from \"$URL\". Please wait."
sudo rm -rf /usr/local/go
wget --inet4-only -c "$URL" -O - | sudo tar -xz -C /usr/local

# Add symlink to go binary.
echo "Adding symbolic links for go and gofmt."
sudo ln -sf /usr/local/go/bin/go /usr/local/bin/go
sudo ln -sf /usr/local/go/bin/gofmt /usr/local/bin/gofmt

# Test go command by showing installed Go version. Telemetry in Go >= 1.23 should be set to "off" in
# ~/.config/go/telemetry, see https://go.dev/doc/telemetry. You can otherwise run "go telemetry off".
go version

echo "Done."