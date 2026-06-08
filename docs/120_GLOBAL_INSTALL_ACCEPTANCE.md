# Global Install Acceptance

`ni` global install means a user can open a new terminal or PowerShell session
and run `ni --help` and `ni version` from any directory without typing the
install path.

This document defines install acceptance only. It does not change the kernel
boundary: `ni run` compiles a bounded handoff prompt and does not execute
downstream work.

## Current Status

| Area | Status | Notes |
| --- | --- | --- |
| macOS/Linux curl installer global PATH handling | Available in installer code | `install.sh` can install to a user-writable directory and optionally add a reversible zsh/bash PATH block. |
| macOS local global command verification | Verified locally | Repository checks install into a temporary bin directory and verify `ni --help` and `ni version` by command name through a fresh shell process. |
| Windows PowerShell installer | Available in installer code | `install.ps1` installs to `%LOCALAPPDATA%\ni\bin` by default, updates User PATH only, and adds a ni-managed PowerShell profile block for the built-in `ni` alias. |
| Windows execution verification | Not verified on this macOS host | Static safety checks exist; real Windows install, new PowerShell, and uninstall verification still require a Windows host transcript. |
| Homebrew | Planned / v0.5 candidate | No Homebrew Available claim is made by this document. |

## macOS Install Success

A macOS install is successful only when all of these are true:

- The `ni` binary is installed to a user-writable location, defaulting to
  `$HOME/.local/bin/ni`.
- The install directory is already on `PATH`, or the user explicitly opts into a
  safe installer-managed PATH update.
- A new shell can run `ni --help` by command name.
- A new shell can run `ni version` by command name.
- Uninstall removes the installed binary.
- Uninstall removes only the PATH configuration added by the installer.

`install.sh --update-path` may add this clearly marked shell profile block for
zsh or bash:

```sh
# >>> ni installer >>>
export PATH="$HOME/.local/bin:$PATH"
# <<< ni installer <<<
```

The actual path may differ when `BINDIR` is set. The installer must not silently
edit shell files without the marker, and uninstall must remove only the marked
block.

## Windows Install Success

A Windows install is successful only when all of these are true:

- `ni.exe` is installed to a user-writable location, preferably
  `%LOCALAPPDATA%\ni\bin\ni.exe`.
- The install directory is added to User PATH, not System PATH, by default.
- The installer handles the PowerShell built-in `ni -> New-Item` alias by
  adding this ni-managed block to `$PROFILE` only once:

```powershell
# >>> ni installer >>>
Remove-Item Alias:ni -Force -ErrorAction SilentlyContinue
# <<< ni installer <<<
```

- Existing PowerShell profile content is preserved.
- `Get-Command ni -All` in a new PowerShell session shows `ni.exe` after the
  profile block loads.
- A new PowerShell session can run `ni --help` by command name.
- A new PowerShell session can run `ni version` by command name.
- Uninstall removes `ni.exe`.
- Uninstall removes only the `ni` bin directory entry from User PATH.
- Uninstall removes only the ni-managed PowerShell profile block.
- The installer preserves unrelated PATH entries and does not truncate PATH.

The Windows installer must read and write User PATH with:

```powershell
[Environment]::GetEnvironmentVariable("Path", "User")
[Environment]::SetEnvironmentVariable("Path", $newPath, "User")
```

It must not use blind `setx PATH "%PATH%;..."`, must not require admin for the
default install, and must not modify Machine PATH by default.

## Verification Standard

Install verification should prove command-name resolution, not only absolute
path execution.

Required checks where possible:

- Install to a temporary `HOME`, `BINDIR`, or `%LOCALAPPDATA%`.
- Confirm the installed binary exists.
- Confirm PATH is already present, added through a managed block, or explicitly
  documented as a manual follow-up.
- On Windows, inspect `Get-Command ni -All` in a new PowerShell session.
- Launch a fresh shell or PowerShell process with the expected PATH context.
- Run `ni --help` by command name.
- Run `ni version` by command name.
- Uninstall.
- Confirm the installer-managed binary and PATH entry are removed.

On non-Windows hosts, Windows installer checks are static only unless a real
Windows host is used. Do not claim Windows execution verified until a Windows
PowerShell install, new-session verification, and uninstall transcript exists.

## No-Overclaim Guard

This acceptance document does not claim:

- Homebrew Available.
- Windows execution verified on macOS.
- no-terminal deterministic validation.
- `ni run` executes downstream work.
- benchmark evidence proves implementation quality.
- fixture relock is project-root relock.

Skills are UX; CLI is authority.
