# Windows PowerShell Alias Fix

## Observed Issue

PowerShell defines a built-in alias:

```powershell
ni -> New-Item
```

After `install.ps1` added `%LOCALAPPDATA%\ni\bin` to User PATH, PowerShell could
still resolve `ni` to `New-Item` before resolving `ni.exe` from PATH.

Observed results:

- `ni --help` created a file named `--help`.
- `ni version` created a file named `version`.

That means PowerShell invoked `New-Item`, not `ni.exe`.

## Root Cause

PowerShell command resolution checks aliases before PATH commands. Adding the
install directory to User PATH is necessary, but it is not sufficient for the
short command name `ni` while the built-in alias exists.

## Installer Fix

`install.ps1` now creates `$PROFILE` only if needed, preserves existing profile
content, and adds this ni-managed block only once:

```powershell
# >>> ni installer >>>
Remove-Item Alias:ni -Force -ErrorAction SilentlyContinue
# <<< ni installer <<<
```

The block removes the PowerShell alias in new sessions so PATH can resolve
`ni.exe` when the user types `ni`.

If the profile update fails, the installer prints the exact manual block and
the diagnostic commands instead of silently claiming that command-name
resolution is complete.

## Uninstall Behavior

`install.ps1 -Uninstall` removes:

- `%LOCALAPPDATA%\ni\bin\ni.exe`
- the installer-managed User PATH entry
- only the ni-managed PowerShell profile block

It preserves unrelated PATH entries and unrelated PowerShell profile content.

## Manual Diagnostics

After install, open a new PowerShell session and run:

```powershell
Get-Command ni -All
ni --help
ni version
```

`Get-Command ni` should resolve to `ni.exe` after the profile block loads. If it
still resolves to `New-Item`, inspect `$PROFILE` and confirm the ni-managed block
is present exactly once.

After uninstall, open a new PowerShell session and run:

```powershell
Get-Command ni -All
```

The ni-managed binary path should no longer be present.

## Claim Boundary

This fix is implemented in the installer and covered by static repository
checks. Windows real-host verification remains pending until a real Windows
PowerShell install, new-session `Get-Command ni -All`, `ni --help`, `ni version`,
and uninstall transcript exists.

This fix does not change kernel behavior. `ni run` still compiles a bounded
downstream handoff prompt and does not execute downstream work.
