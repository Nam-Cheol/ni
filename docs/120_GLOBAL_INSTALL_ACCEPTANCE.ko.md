# Global Install Acceptance

`ni` global install은 사용자가 새 terminal 또는 PowerShell session을 열고 어느
directory에서든 install path를 입력하지 않고 `ni --help`와 `ni version`을 실행할
수 있다는 뜻이다.

이 문서는 install acceptance만 정의한다. Kernel boundary는 바뀌지 않는다:
`ni run`은 bounded handoff prompt를 compile하며 downstream work를 실행하지 않는다.

## Current Status

| Area | Status | Notes |
| --- | --- | --- |
| macOS/Linux curl installer global PATH handling | Available in installer code | `install.sh`는 user-writable directory에 설치하고, opt-in으로 reversible zsh/bash PATH block을 추가할 수 있다. |
| macOS local global command verification | Verified locally | Repository check는 temporary bin directory에 설치한 뒤 fresh shell process에서 command name으로 `ni --help`와 `ni version`을 검증한다. |
| Windows PowerShell installer | Available in installer code | `install.ps1`는 기본적으로 `%LOCALAPPDATA%\ni\bin`에 설치하고 User PATH만 업데이트하며, built-in `ni` alias를 위해 ni-managed PowerShell profile block을 추가한다. |
| Windows execution verification | Not verified on this macOS host | Static safety check는 있지만 real Windows install, new PowerShell, uninstall verification은 Windows host transcript가 필요하다. |
| Homebrew | Planned / v0.5 candidate | 이 문서는 Homebrew Available을 claim하지 않는다. |

## macOS Install Success

macOS install은 다음이 모두 true일 때만 성공이다:

- `ni` binary가 user-writable location에 설치된다. 기본값은
  `$HOME/.local/bin/ni`이다.
- Install directory가 이미 `PATH`에 있거나, 사용자가 safe installer-managed
  PATH update에 명시적으로 opt in한다.
- 새 shell에서 command name으로 `ni --help`를 실행할 수 있다.
- 새 shell에서 command name으로 `ni version`을 실행할 수 있다.
- Uninstall이 installed binary를 제거한다.
- Uninstall이 installer가 추가한 PATH configuration만 제거한다.

`install.sh --update-path`는 zsh 또는 bash profile에 다음처럼 명확히 표시된 block을
추가할 수 있다:

```sh
# >>> ni installer >>>
export PATH="$HOME/.local/bin:$PATH"
# <<< ni installer <<<
```

`BINDIR`를 지정하면 실제 path는 달라질 수 있다. Installer는 marker 없이 shell
file을 silently edit하면 안 되며, uninstall은 marked block만 제거해야 한다.

## Windows Install Success

Windows install은 다음이 모두 true일 때만 성공이다:

- `ni.exe`가 user-writable location, 되도록
  `%LOCALAPPDATA%\ni\bin\ni.exe`에 설치된다.
- Install directory가 기본적으로 System PATH가 아니라 User PATH에 추가된다.
- Installer가 PowerShell built-in `ni -> New-Item` alias를 처리하기 위해
  `$PROFILE`에 다음 ni-managed block을 한 번만 추가한다:

```powershell
# >>> ni installer >>>
Remove-Item Alias:ni -Force -ErrorAction SilentlyContinue
# <<< ni installer <<<
```

- Existing PowerShell profile content가 보존된다.
- 새 PowerShell session의 `Get-Command ni -All`이 profile block load 이후
  `ni.exe`를 보여준다.
- 새 PowerShell session에서 command name으로 `ni --help`를 실행할 수 있다.
- 새 PowerShell session에서 command name으로 `ni version`을 실행할 수 있다.
- Uninstall이 `ni.exe`를 제거한다.
- Uninstall이 User PATH에서 `ni` bin directory entry만 제거한다.
- Uninstall이 ni-managed PowerShell profile block만 제거한다.
- Installer가 unrelated PATH entries를 보존하고 PATH를 truncate하지 않는다.

Windows installer는 User PATH를 다음 방식으로 읽고 써야 한다:

```powershell
[Environment]::GetEnvironmentVariable("Path", "User")
[Environment]::SetEnvironmentVariable("Path", $newPath, "User")
```

Blind `setx PATH "%PATH%;..."`를 사용하면 안 되고, default install에 admin을
요구하면 안 되며, default로 Machine PATH를 수정하면 안 된다.

## Verification Standard

Install verification은 absolute path execution만이 아니라 command-name resolution을
증명해야 한다.

가능한 required checks:

- Temporary `HOME`, `BINDIR`, 또는 `%LOCALAPPDATA%`에 install한다.
- Installed binary가 존재하는지 확인한다.
- PATH가 이미 존재하거나 managed block으로 추가되었거나 manual follow-up으로
  명확히 문서화되었는지 확인한다.
- Windows에서는 새 PowerShell session에서 `Get-Command ni -All`을 확인한다.
- Expected PATH context를 가진 fresh shell 또는 PowerShell process를 launch한다.
- Command name으로 `ni --help`를 실행한다.
- Command name으로 `ni version`을 실행한다.
- Uninstall한다.
- Installer-managed binary와 PATH entry가 제거되었는지 확인한다.

Non-Windows host에서는 real Windows host를 사용하지 않는 한 Windows installer check는
static only다. Windows PowerShell install, new-session verification, uninstall
transcript가 있기 전에는 Windows execution verified라고 claim하지 않는다.

## No-Overclaim Guard

이 acceptance 문서는 다음을 claim하지 않는다:

- Homebrew Available.
- macOS에서 Windows execution verified.
- no-terminal deterministic validation.
- `ni run`이 downstream work를 실행한다.
- benchmark evidence가 implementation quality를 증명한다.
- fixture relock이 project-root relock이다.

Skills are UX; CLI is authority.
