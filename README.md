# Break Check

Very early on in development, check back later!

## Piping Input
Use `xargs` to pipe a list of package names to check, for example: `jq -r '.dependencies | keys[]' package.json | xargs bc run`
