# go-random
Random Go modules - An attempt to write random set of utilities to learn go mechanics

### release 1.0
In release 1.0 we have prilimanry utility which creates an REST api service and related command line.

It goes through the way we should organize code in the form of 
- Models - they are type structures on which actions are going to be taken.
- Interfaces - to capature the actions
- Use model in creation of HTTP API service and various server side error catching in `service/echo/echo.go`  and tie that back to its entry point at `cmd/api/main.go`.
- Organizing CLI interface for the above API service  in the form of its subcommands `cmd/cli/command` and its entry point at `cmd/cli/main.go`.
- Take a note of `cmd/cli/command/exec.go` to see the spine of execution which calls subcommand specific functions.

References 
- [Intro to HTTP handlers](https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go) 
- [Building-Modern-CLI-Applications-in-Go](https://github.com/PacktPublishing/Building-Modern-CLI-Applications-in-Go/tree/main/Chapter03/audiofile)
