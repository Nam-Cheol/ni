#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
TEST_TMP="$(mktemp -d "${TMPDIR:-/tmp}/namba-intent-install-sh-test.XXXXXX")"
LAST_STDOUT="$TEST_TMP/stdout.log"
LAST_STDERR="$TEST_TMP/stderr.log"

trap 'rm -rf "$TEST_TMP"' EXIT

cd "$ROOT"

run_cmd() {
  local label="$1"
  shift
  echo "test-install-sh: $label" >&2
  : >"$LAST_STDOUT"
  : >"$LAST_STDERR"
  if ! "$@" >"$LAST_STDOUT" 2>"$LAST_STDERR"; then
    echo "test-install-sh failed: $label" >&2
    echo "--- stdout ---" >&2
    sed -n '1,160p' "$LAST_STDOUT" >&2
    echo "--- stderr ---" >&2
    sed -n '1,160p' "$LAST_STDERR" >&2
    return 1
  fi
}

run_must_fail() {
  local label="$1"
  shift
  echo "test-install-sh: $label" >&2
  : >"$LAST_STDOUT"
  : >"$LAST_STDERR"
  if "$@" >"$LAST_STDOUT" 2>"$LAST_STDERR"; then
    echo "test-install-sh failed: command unexpectedly passed: $label" >&2
    echo "--- stdout ---" >&2
    sed -n '1,160p' "$LAST_STDOUT" >&2
    echo "--- stderr ---" >&2
    sed -n '1,160p' "$LAST_STDERR" >&2
    return 1
  fi
}

require_stdout() {
  local expected="$1"
  if ! grep -Fq "$expected" "$LAST_STDOUT"; then
    echo "test-install-sh failed: expected stdout to contain: $expected" >&2
    echo "--- stdout ---" >&2
    sed -n '1,160p' "$LAST_STDOUT" >&2
    return 1
  fi
}

require_stderr() {
  local expected="$1"
  if ! grep -Fq "$expected" "$LAST_STDERR"; then
    echo "test-install-sh failed: expected stderr to contain: $expected" >&2
    echo "--- stderr ---" >&2
    sed -n '1,160p' "$LAST_STDERR" >&2
    return 1
  fi
}

detect_os() {
  case "$(uname -s)" in
    Linux) printf '%s\n' linux ;;
    Darwin) printf '%s\n' darwin ;;
    *) printf '%s\n' unsupported ;;
  esac
}

detect_arch() {
  case "$(uname -m)" in
    x86_64|amd64) printf '%s\n' amd64 ;;
    arm64|aarch64) printf '%s\n' arm64 ;;
    *) printf '%s\n' unsupported ;;
  esac
}

sha256_file() {
  if command -v sha256sum >/dev/null 2>&1; then
    sha256sum "$1" | awk '{print $1}'
  else
    shasum -a 256 "$1" | awk '{print $1}'
  fi
}

OS="$(detect_os)"
ARCH="$(detect_arch)"

if [[ "$OS" == "unsupported" || "$ARCH" == "unsupported" ]]; then
  echo "test-install-sh: skipping unsupported host $(uname -s)/$(uname -m)" >&2
  exit 0
fi

VERSION="0.0.0-test"
ASSET="namba-intent_${VERSION}_${OS}_${ARCH}.tar.gz"
CHECKSUMS="namba-intent_${VERSION}_checksums.txt"
RELEASE_DIR="$TEST_TMP/release"
PAYLOAD_DIR="$TEST_TMP/payload"
INSTALL_DIR="$TEST_TMP/bin"

mkdir -p "$RELEASE_DIR" "$PAYLOAD_DIR"

cat >"$PAYLOAD_DIR/namba-intent" <<'SH'
#!/usr/bin/env sh
case "${1:-}" in
  --help)
    echo "Namba Intent is a Project Intent Compiler for AI Agents."
    ;;
  version)
    echo "0.0.0-test"
    ;;
  *)
    echo "test namba-intent: expected --help or version" >&2
    exit 2
    ;;
esac
SH
chmod 0755 "$PAYLOAD_DIR/namba-intent"

tar -czf "$RELEASE_DIR/$ASSET" -C "$PAYLOAD_DIR" namba-intent
SHA="$(sha256_file "$RELEASE_DIR/$ASSET")"
printf '%s  %s\n' "$SHA" "$ASSET" >"$RELEASE_DIR/$CHECKSUMS"

BASE_URL="file://$RELEASE_DIR"

run_cmd "dry-run selects the local test asset" env \
  NI_INSTALL_BASE_URL="$BASE_URL" \
  BINDIR="$INSTALL_DIR" \
  sh ./install.sh --dry-run --version "$VERSION"
require_stdout "mode:       dry-run"
require_stdout "$ASSET"
require_stdout "$INSTALL_DIR/namba-intent"

run_cmd "dry-run strips v prefix for asset names" env \
  NI_INSTALL_BASE_URL="$BASE_URL" \
  BINDIR="$INSTALL_DIR" \
  sh ./install.sh --dry-run --version "v$VERSION"
require_stdout "$ASSET"

run_cmd "linux amd64 override dry-run" env \
  NI_INSTALL_BASE_URL="$BASE_URL" \
  NI_INSTALL_OS=linux \
  NI_INSTALL_ARCH=amd64 \
  BINDIR="$INSTALL_DIR" \
  sh ./install.sh --dry-run --version "$VERSION"
require_stdout "platform:   linux/amd64"
require_stdout "namba-intent_${VERSION}_linux_amd64.tar.gz"

run_cmd "linux arm64 override dry-run" env \
  NI_INSTALL_BASE_URL="$BASE_URL" \
  NI_INSTALL_OS=linux \
  NI_INSTALL_ARCH=arm64 \
  BINDIR="$INSTALL_DIR" \
  sh ./install.sh --dry-run --version "$VERSION"
require_stdout "platform:   linux/arm64"
require_stdout "namba-intent_${VERSION}_linux_arm64.tar.gz"

run_cmd "darwin amd64 override dry-run" env \
  NI_INSTALL_BASE_URL="$BASE_URL" \
  NI_INSTALL_OS=darwin \
  NI_INSTALL_ARCH=amd64 \
  BINDIR="$INSTALL_DIR" \
  sh ./install.sh --dry-run --version "$VERSION"
require_stdout "platform:   darwin/amd64"
require_stdout "namba-intent_${VERSION}_darwin_amd64.tar.gz"

run_cmd "darwin arm64 override dry-run" env \
  NI_INSTALL_BASE_URL="$BASE_URL" \
  NI_INSTALL_OS=darwin \
  NI_INSTALL_ARCH=arm64 \
  BINDIR="$INSTALL_DIR" \
  sh ./install.sh --dry-run --version "$VERSION"
require_stdout "platform:   darwin/arm64"
require_stdout "namba-intent_${VERSION}_darwin_arm64.tar.gz"

run_cmd "windows amd64 override dry-run" env \
  NI_INSTALL_BASE_URL="$BASE_URL" \
  NI_INSTALL_OS=windows \
  NI_INSTALL_ARCH=amd64 \
  BINDIR="$INSTALL_DIR" \
  sh ./install.sh --dry-run --version "$VERSION"
require_stdout "platform:   windows/amd64"
require_stdout "namba-intent_${VERSION}_windows_amd64.zip"
require_stdout "$INSTALL_DIR/namba-intent.exe"

if [[ -e "$INSTALL_DIR/namba-intent" ]]; then
  echo "test-install-sh failed: dry-run created $INSTALL_DIR/namba-intent" >&2
  exit 1
fi

run_cmd "install from a local release asset" env \
  NI_INSTALL_BASE_URL="$BASE_URL" \
  BINDIR="$INSTALL_DIR" \
  sh ./install.sh --version "$VERSION"
require_stdout "Verified checksum for $ASSET"
require_stdout "Installed namba-intent to $INSTALL_DIR/namba-intent"

run_cmd "installed namba-intent --help" "$INSTALL_DIR/namba-intent" --help
require_stdout "Namba Intent is a Project Intent Compiler for AI Agents."

run_cmd "installed namba-intent version" "$INSTALL_DIR/namba-intent" version
require_stdout "$VERSION"

run_cmd "fresh shell resolves installed namba-intent by command name" env \
  PATH="$INSTALL_DIR:$PATH" \
  sh -c 'command -v namba-intent && namba-intent --help && namba-intent version'
require_stdout "Namba Intent is a Project Intent Compiler for AI Agents."
require_stdout "$VERSION"

PROFILE="$TEST_TMP/home/.zshrc"
UPDATE_INSTALL_DIR="$TEST_TMP/update-bin"
run_cmd "install with managed PATH block" env \
  NI_INSTALL_BASE_URL="$BASE_URL" \
  NI_INSTALL_SHELL_PROFILE="$PROFILE" \
  HOME="$TEST_TMP/home" \
  SHELL="/bin/zsh" \
  BINDIR="$UPDATE_INSTALL_DIR" \
  sh ./install.sh --update-path --version "$VERSION"
require_stdout "Added Namba Intent PATH block to $PROFILE"

if ! grep -Fq "# >>> namba-intent installer >>>" "$PROFILE"; then
  echo "test-install-sh failed: managed PATH block was not written" >&2
  exit 1
fi

run_cmd "uninstall removes binary and managed PATH block" env \
  NI_INSTALL_SHELL_PROFILE="$PROFILE" \
  HOME="$TEST_TMP/home" \
  SHELL="/bin/zsh" \
  BINDIR="$UPDATE_INSTALL_DIR" \
  sh ./install.sh --uninstall
require_stdout "Removed $UPDATE_INSTALL_DIR/namba-intent"
require_stdout "Removed Namba Intent PATH block from $PROFILE"

if [[ -e "$UPDATE_INSTALL_DIR/namba-intent" ]]; then
  echo "test-install-sh failed: uninstall left installed binary" >&2
  exit 1
fi

if grep -Fq "# >>> namba-intent installer >>>" "$PROFILE"; then
  echo "test-install-sh failed: uninstall left managed PATH block" >&2
  exit 1
fi

printf 'bad  %s\n' "$ASSET" >"$RELEASE_DIR/$CHECKSUMS"
run_must_fail "checksum mismatch fails" env \
  NI_INSTALL_BASE_URL="$BASE_URL" \
  BINDIR="$TEST_TMP/bad-bin" \
  sh ./install.sh --version "$VERSION"
require_stderr "checksum mismatch"

echo "test-install-sh: install.sh dry-run, checksum, BINDIR, and help/version checks passed"
