# TO-DO

## Improve user experience

### More binary executable tools

~~Currently only have the ability to run server, everything else is done externally~~

Server-facing UX SHOULD be simple and minimalistic.
Ideally, usage should be:

- config stays in `.env` file:
  - supply a `.env.example` file as documentation
  - if no `.env` exists, fall back defaults in the programs
  - else, use `.env`
- tools accessible via flags (e.g. `--version`, `--port`)

**Actual TODO:**

- [x] Rename binary to something like `skelbiu-api`
- [ ] add tools
  - [ ] run - start serving traffic
  - [ ] *MAYYYBE* add `man` documentation
- [ ] add flags
  - [ ] standard UNIX flags such as `--version`, `--help`, etc.
  - [x] override ENV variables (just use standard Bash `VAR=VALUE`)

### Better release system

- [x] Generalize compilation in Makefile
- [ ] Remove Pascal code in installer script
- [ ] Add rolling release with compiled installer and Linux binaries

### More standard API interface

>[!IMPORTANT]
> Only release this as part of v2.0

- [ ] Make API use skelbiu url names, not Go struct fields

## Proper backend deployment

- [ ] Add graceful shutdowns for server (test with launching a compiled executable or something)
- [ ] Add handling for Network interruptions by starting server again automatically

## Industry standards / security

### API keys / auth

Research standard API security and implement auth of some kind

Look into:

- [ ] generating free API keys
- [ ] rate limiting on said free keys
- [ ] idk
