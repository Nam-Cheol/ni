# Windows PowerShell Alias Fix

## Observed Issue

PowerShell에는 built-in alias가 있다:

```powershell
ni -> New-Item
```

`install.ps1`가 `%LOCALAPPDATA%\ni\bin`을 User PATH에 추가한 뒤에도
PowerShell은 PATH의 `ni.exe`보다 built-in alias `ni`를 먼저 resolve할 수 있었다.

Observed results:

- `ni --help`가 `--help`라는 file을 만들었다.
- `ni version`이 `version`이라는 file을 만들었다.

이는 PowerShell이 `ni.exe`가 아니라 `New-Item`을 실행했다는 뜻이다.

## Root Cause

PowerShell command resolution은 PATH command보다 alias를 먼저 본다. Install
directory를 User PATH에 추가하는 것은 필요하지만, built-in alias가 남아 있으면
short command name `ni`에는 충분하지 않다.

## Installer Fix

`install.ps1`는 `$PROFILE`이 필요할 때만 생성하고, existing profile content를
보존하며, 다음 ni-managed block을 한 번만 추가한다:

```powershell
# >>> ni installer >>>
Remove-Item Alias:ni -Force -ErrorAction SilentlyContinue
# <<< ni installer <<<
```

이 block은 새 session에서 PowerShell alias를 제거해서 사용자가 `ni`를 입력하면
PATH의 `ni.exe`가 resolve되게 한다.

Profile update가 실패하면 installer는 command-name resolution이 완료됐다고
조용히 claim하지 않고, exact manual block과 diagnostic commands를 출력한다.

## Uninstall Behavior

`install.ps1 -Uninstall`은 다음을 제거한다:

- `%LOCALAPPDATA%\ni\bin\ni.exe`
- installer-managed User PATH entry
- ni-managed PowerShell profile block만

Unrelated PATH entries와 unrelated PowerShell profile content는 보존한다.

## Manual Diagnostics

Install 후 새 PowerShell session을 열고 다음을 실행한다:

```powershell
Get-Command ni -All
ni --help
ni version
```

Profile block이 load된 뒤 `Get-Command ni`는 `ni.exe`로 resolve되어야 한다. 여전히
`New-Item`으로 resolve되면 `$PROFILE`을 확인하고 ni-managed block이 정확히 한 번
있는지 확인한다.

Uninstall 후 새 PowerShell session을 열고 다음을 실행한다:

```powershell
Get-Command ni -All
```

Ni-managed binary path가 더 이상 없어야 한다.

## Claim Boundary

이 fix는 installer에 구현됐고 repository static check로 커버된다. Windows real-host
verification은 real Windows PowerShell install, 새 session의
`Get-Command ni -All`, `ni --help`, `ni version`, uninstall transcript가 생길
때까지 pending이다.

이 fix는 kernel behavior를 바꾸지 않는다. `ni run`은 여전히 bounded downstream
handoff prompt를 compile하며 downstream work를 실행하지 않는다.
