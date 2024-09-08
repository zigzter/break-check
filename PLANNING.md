# Planning

## Purpose
Let developers easily get a summary of all breaking changes that occur in a package from the current installed version up to the latest version.
The user can either specify which packages they want this summary for, or which packages they want ignored.

## User Flow
- Run app in a project folder
    - Flags are passed to control results, pass list of packages to either include or exclude
    - If no flags are passed, check all packages. If both are passed, either throw error, or prefer one over the other
- Prompt for stdout or file for results (default to stdout)
- App pipes result to above choice

## Tech Options
- Go + Bubbletea
    - Good excuse to further improve my Go knowledge and comfort
- TypeScript + Ink
    - Easiest choice for dev work, but user would need Node
- Rust + Ratatui
    - Would require learning Rust

## Requirements
- App is added to PATH
- App is run from the folder the user wants to check for breaking changes
- Can be run with the following flags:
    - `--include`: packages to include, comma separated
    - `--exclude`: packages to exclude, comma separated
    - `--format=x`: stdout or file, if present skip the choice prompt in app
- Check GitHub API auth status on launch, prompt for auth if required
- If GH auth present, present choice of stdout or file, default to stdout
- Loop through list of packages to check, requesting update info from GH API
- Descriptions of updates must be parsed to grab all breaking changes
- Breaking changes get collected and separated by their respective version numbers, then printed to chosen output

## Planning To-Do
- Figure out GitHub API auth choices
- Possibly support [NPM API](https://registry.npmjs.org) fallback if GitHub doesn't show changes
- Look into populating list of installed packages in app itself, and allow include/exclude within the app
