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
- Go + Bubbletea + Cobra
    - Good excuse to further improve my Go knowledge and comfort
- TypeScript + Ink
    - Easiest choice for dev work, but user would need Node
- Rust + Ratatui
    - Would require learning Rust

## Requirements & Considerations
- Start with support for Node packages, potentially add more languages in the future
    - If more added, app should auto-detect what type of project it is
- Support at least Linux & MacOS
- App is added to PATH
- App is run from the folder the user wants to check for breaking changes
- Can be run with the following flags:
    - `--include`: packages to include, comma separated
    - `--exclude`: packages to exclude, comma separated
    - `--format=x`: stdout or file, if present skip the choice prompt in app
- If `package.json` (or other dependecy file) is missing, exit and print error
- Check GitHub API auth status on launch, prompt for auth if required
- If GH auth present, present choice of stdout or file, default to stdout
- Loop through list of packages to check, requesting update info from GH API
    - Cache these results per package
        - The program is meant to run once and exit, so we can't cache in memory
        - Probable solution is to store these in a temp file, and clean the file if the program is run with a temp file older than 24h
        - Allow a command to allow the user to clear this cache manually
    - Run requests in parallel
    - Limit requests to API limit
- Descriptions of updates must be parsed to grab all breaking changes
- Handle packages that don't follow semver
- Breaking changes get collected and separated by their respective version numbers, then printed to chosen output
- Print out any packages that were missing changelogs
- Print out error messages to a debug file to allow for easier issue resolving

## Planning To-Do
- Figure out GitHub API auth choices
- Possibly support [NPM API](https://registry.npmjs.org) fallback if GitHub doesn't show changes
- Look into populating list of installed packages in app itself, and allow include/exclude within the app
- Figure out best way to present data
- Support JSON output? Would allow user to pipe output to something like `jq`
