#!/usr/bin/env bash
set -euo pipefail

VERSION="${NI_FRESH_INSTALL_VERSION:-0.3.0}"
VERSION="${VERSION#v}"
TAG="${NI_FRESH_INSTALL_TAG:-v$VERSION}"
REPO="${NI_FRESH_INSTALL_REPO:-Nam-Cheol/ni}"
RELEASE_BASE_URL="${NI_FRESH_INSTALL_BASE_URL:-https://github.com/$REPO/releases/download/$TAG}"
INSTALLER_URL="${NI_FRESH_INSTALL_INSTALLER_URL:-https://raw.githubusercontent.com/$REPO/main/install.sh}"
CHECK_TMP="$(mktemp -d "${TMPDIR:-/tmp}/ni-fresh-install-check.XXXXXX")"
LAST_STDOUT="$CHECK_TMP/stdout.log"
LAST_STDERR="$CHECK_TMP/stderr.log"

trap 'rm -rf "$CHECK_TMP"' EXIT

need() {
  command -v "$1" >/dev/null 2>&1 || {
    echo "fresh-install-check failed: required command not found: $1" >&2
    exit 1
  }
}

detect_os() {
  case "$(uname -s)" in
    Linux) printf '%s\n' linux ;;
    Darwin) printf '%s\n' darwin ;;
    MINGW*|MSYS*|CYGWIN*) printf '%s\n' windows ;;
    *)
      echo "fresh-install-check failed: unsupported operating system: $(uname -s)" >&2
      exit 1
      ;;
  esac
}

detect_arch() {
  case "$(uname -m)" in
    x86_64|amd64) printf '%s\n' amd64 ;;
    arm64|aarch64) printf '%s\n' arm64 ;;
    *)
      echo "fresh-install-check failed: unsupported architecture: $(uname -m)" >&2
      exit 1
      ;;
  esac
}

sha256_file() {
  if command -v sha256sum >/dev/null 2>&1; then
    sha256sum "$1" | awk '{print $1}'
    return 0
  fi

  if command -v shasum >/dev/null 2>&1; then
    shasum -a 256 "$1" | awk '{print $1}'
    return 0
  fi

  echo "fresh-install-check failed: sha256sum or shasum is required" >&2
  exit 1
}

run_cmd() {
  local label="$1"
  shift
  echo "fresh-install-check: $label" >&2
  : >"$LAST_STDOUT"
  : >"$LAST_STDERR"
  if ! "$@" >"$LAST_STDOUT" 2>"$LAST_STDERR"; then
    echo "fresh-install-check failed: $label" >&2
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
    echo "fresh-install-check failed: expected stdout to contain: $expected" >&2
    echo "--- stdout ---" >&2
    sed -n '1,160p' "$LAST_STDOUT" >&2
    return 1
  fi
}

need curl
need awk
need grep
need sed

OS="$(detect_os)"
ARCH="$(detect_arch)"
case "$OS/$ARCH" in
  linux/amd64|linux/arm64|darwin/amd64|darwin/arm64|windows/amd64) ;;
  windows/arm64)
    echo "fresh-install-check failed: windows arm64 release asset is not configured" >&2
    exit 1
    ;;
  *)
    echo "fresh-install-check failed: unsupported platform: $OS/$ARCH" >&2
    exit 1
    ;;
esac

EXT="tar.gz"
BIN_NAME="namba-intent"
if [[ "$OS" == "windows" ]]; then
  EXT="zip"
  BIN_NAME="namba-intent.exe"
  need unzip
else
  need tar
fi

ASSET="namba-intent_${VERSION}_${OS}_${ARCH}.${EXT}"
CHECKSUMS="namba-intent_${VERSION}_checksums.txt"
ASSET_URL="${RELEASE_BASE_URL%/}/$ASSET"
CHECKSUM_URL="${RELEASE_BASE_URL%/}/$CHECKSUMS"

RELEASE_DIR="$CHECK_TMP/release"
MANUAL_EXTRACT_DIR="$CHECK_TMP/manual-extract"
MANUAL_BIN_DIR="$CHECK_TMP/manual-bin"
CURL_BIN_DIR="$CHECK_TMP/curl-bin"
MANUAL_PROJECT_DIR="$CHECK_TMP/manual-project"
CURL_PROJECT_DIR="$CHECK_TMP/curl-project"
INSTALLER_PATH="$CHECK_TMP/install.sh"

mkdir -p "$RELEASE_DIR" "$MANUAL_EXTRACT_DIR" "$MANUAL_BIN_DIR" "$CURL_BIN_DIR"

echo "fresh-install-check: version $TAG"
echo "fresh-install-check: platform $OS/$ARCH"
echo "fresh-install-check: temp root $CHECK_TMP"

run_cmd "download release archive" curl -fsSL "$ASSET_URL" -o "$RELEASE_DIR/$ASSET"
run_cmd "download release checksums" curl -fsSL "$CHECKSUM_URL" -o "$RELEASE_DIR/$CHECKSUMS"

EXPECTED_SHA="$(awk -v asset="$ASSET" '
  {
    name = $2
    sub(/^\*/, "", name)
    sub(/^\.\//, "", name)
    if (name == asset) {
      print $1
      exit
    }
  }
' "$RELEASE_DIR/$CHECKSUMS")"

if [[ "$EXPECTED_SHA" == "" ]]; then
  echo "fresh-install-check failed: checksum file does not contain $ASSET" >&2
  exit 1
fi

ACTUAL_SHA="$(sha256_file "$RELEASE_DIR/$ASSET")"
if [[ "$EXPECTED_SHA" != "$ACTUAL_SHA" ]]; then
  echo "fresh-install-check failed: checksum mismatch for $ASSET" >&2
  exit 1
fi
echo "fresh-install-check: verified checksum for $ASSET"

if [[ "$EXT" == "zip" ]]; then
  run_cmd "extract release archive" unzip -q "$RELEASE_DIR/$ASSET" -d "$MANUAL_EXTRACT_DIR"
else
  run_cmd "extract release archive" tar -xzf "$RELEASE_DIR/$ASSET" -C "$MANUAL_EXTRACT_DIR"
fi

FOUND_BIN="$(find "$MANUAL_EXTRACT_DIR" -type f -name "$BIN_NAME" | sed -n '1p')"
if [[ "$FOUND_BIN" == "" ]]; then
  echo "fresh-install-check failed: release archive did not contain $BIN_NAME" >&2
  exit 1
fi

cp "$FOUND_BIN" "$MANUAL_BIN_DIR/$BIN_NAME"
chmod 0755 "$MANUAL_BIN_DIR/$BIN_NAME"

run_cmd "manual release binary namba-intent --help" "$MANUAL_BIN_DIR/$BIN_NAME" --help
require_stdout "Namba Intent is a Project Intent Compiler for AI Agents."

run_cmd "manual release binary namba-intent version" "$MANUAL_BIN_DIR/$BIN_NAME" version
require_stdout "$VERSION"

run_cmd "manual release binary namba-intent init" "$MANUAL_BIN_DIR/$BIN_NAME" init --dir "$MANUAL_PROJECT_DIR" --profile prototype
require_stdout "initialized Namba Intent planning workspace"

run_cmd "manual release binary namba-intent status" "$MANUAL_BIN_DIR/$BIN_NAME" status --dir "$MANUAL_PROJECT_DIR"
require_stdout "BLOCKED"

run_cmd "download curl installer" curl -fsSL "$INSTALLER_URL" -o "$INSTALLER_PATH"

run_cmd "curl installer dry-run" env BINDIR="$CURL_BIN_DIR" sh "$INSTALLER_PATH" --dry-run --version "$VERSION"
require_stdout "$ASSET"
require_stdout "$CURL_BIN_DIR/$BIN_NAME"

run_cmd "curl installer temporary install" env BINDIR="$CURL_BIN_DIR" sh "$INSTALLER_PATH" --version "$VERSION"
require_stdout "Verified checksum for $ASSET"
require_stdout "Installed ni to $CURL_BIN_DIR/$BIN_NAME"

run_cmd "curl installed namba-intent --help" "$CURL_BIN_DIR/$BIN_NAME" --help
require_stdout "Namba Intent is a Project Intent Compiler for AI Agents."

run_cmd "curl installed namba-intent version" "$CURL_BIN_DIR/$BIN_NAME" version
require_stdout "$VERSION"

run_cmd "curl installed namba-intent init" "$CURL_BIN_DIR/$BIN_NAME" init --dir "$CURL_PROJECT_DIR" --profile prototype
require_stdout "initialized Namba Intent planning workspace"

run_cmd "curl installed namba-intent status" "$CURL_BIN_DIR/$BIN_NAME" status --dir "$CURL_PROJECT_DIR"
require_stdout "BLOCKED"

echo "fresh-install-check: manual release binary path passed"
echo "fresh-install-check: curl installer path passed"
echo "fresh-install-check: namba-intent --help, namba-intent version, namba-intent init, and namba-intent status passed without Go"
