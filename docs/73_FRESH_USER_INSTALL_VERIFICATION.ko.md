# Fresh User Install 검증

날짜: 2026-05-28

범위: 새 사용자가 Go를 모르거나 설치하지 않아도 게시된 `v0.3.0` release에서
`ni`를 설치하고 실행할 수 있는지 검증한다.

이 검증은 release binary path와 curl installer path만 다룬다. Homebrew, Scoop,
package-manager distribution, model skills, Codex, Claude, downstream agents,
runtime execution behavior는 테스트하지 않는다.

## 검증 Script

실행:

```bash
bash scripts/fresh-install-check.sh
```

스크립트는 temporary directories만 사용한다. 현재 플랫폼의 published `v0.3.0`
archive와 `ni_0.3.0_checksums.txt`를 download하고, archive checksum을 검증한
뒤, binary를 temporary install directory에 extract하고 다음 명령을 실행한다.

```bash
ni --help
ni version
ni init --dir <temporary-project> --profile prototype
ni status --dir <temporary-project>
```

그다음 public curl installer를 temporary directory로 download하고, 같은
`v0.3.0` binary를 temporary `BINDIR`에 install한 뒤 같은 help/version/init/status
checks를 반복한다.

## Expected Output

출력에는 다음이 포함되어야 한다.

```text
fresh-install-check: manual release binary path passed
fresh-install-check: curl installer path passed
fresh-install-check: ni --help, ni version, ni init, and ni status passed without Go
```

`ni status`는 새 temporary project에 대해 `BLOCKED`를 출력하는 것이 정상이다.
`ni init`은 open intent gaps가 있는 draft planning workspace를 만들기 때문이다.
이 blocked result는 fresh-user proof로 성공이다. Installed CLI가 Go 없이도
project를 initialize하고 deterministic readiness gate를 실행할 수 있음을
보여준다.

## 검증한 Paths

| Path | Status | Proof |
| --- | --- | --- |
| Manual release binary | Verified by script | Release archive와 checksum file을 download하고 SHA-256을 검증한 뒤 temporary directory에서 binary를 실행한다. |
| Curl installer | Verified by script | `install.sh`를 download하고 temporary `BINDIR`를 사용하며 checksum output을 확인한 뒤 installed binary를 실행한다. |
| Homebrew | Not tested | Homebrew는 Planned 상태이며 이 check로 Available 표시를 하면 안 된다. |

## 결과

Non-Go user는 verified `v0.3.0` release binary path 또는 verified curl installer
path로 `ni`를 install한 뒤 `ni --help`, `ni version`, `ni init`, `ni status`를
실행할 수 있다.

이 검증은 new release를 publish하지 않고, tag를 push하지 않으며, downstream
work를 execute하거나 agent를 call하거나 runtime execution behavior를 추가하지
않는다.
