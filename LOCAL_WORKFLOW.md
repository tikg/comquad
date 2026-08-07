# Local Development Workflow

This document describes my local workflow for maintaining my fork while staying in sync with the upstream repository.

---

# Repository Layout

- `origin` → My fork
- `upstream` → Original repository (`Inoriol/comquad`)

Verify:

```bash
git remote -v
```

Expected:

```text
origin    https://github.com/tikg/comquad.git
upstream  https://github.com/Inoriol/comquad.git
```

---

# Sync With Upstream

Before starting any work:

```bash
git fetch upstream
git switch main
git rebase upstream/main
git push origin main
```

This keeps my fork up to date with the latest changes from upstream.

---

# Create a Feature Branch

Never develop directly on `main`.

```bash
git switch -c feature/<feature-name>
```

Example:

```bash
git switch -c feature/install-dependency-check
```

---

# Development

Make changes.

Check status:

```bash
git status
```

Review changes:

```bash
git diff
```

---

# Commit

Stage only the files related to the change.

```bash
git add <files>
```

Commit:

```bash
git commit -m "Describe the change"
```

---

# Push

```bash
git push -u origin feature/<feature-name>
```

Subsequent pushes:

```bash
git push
```

---

# Open Pull Request

Create a Pull Request from:

```
tikg/comquad
```

to

```
Inoriol/comquad
```

---

# Updating a Feature Branch

If upstream changes while I'm working:

```bash
git fetch upstream
git rebase upstream/main
git push --force-with-lease
```

---

# Useful Commands

Current branch:

```bash
git branch --show-current
```

Remotes:

```bash
git remote -v
```

Commit graph:

```bash
git log --graph --oneline --decorate --all
```

Status:

```bash
git status
```

---

# Project Notes

## Planned Contribution

- Add dependency detection.
- Offer automatic installation of missing dependencies.
- Initial support:
  - macOS
  - Linux
- Windows support deferred.

## Dependency Detection

Target dependencies:

- Podman
- podlet
- systemd (Linux)
- Go (build only)

Desired UX:

```text
Missing dependency: podlet

podlet is required to transpile compose.yaml into Quadlets.

Install podlet now? [Y/n]
```
