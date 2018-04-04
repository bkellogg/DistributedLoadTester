# Week 2

Report for week 2 of the independent study.

Meeting: *Wednesday*, April 4 at 3:30pm

## Progress/Accompishments

1. Agent process will now bind and listen on a tcp socket.
2. Master can dial and send a message to an agent and recieve and read responses.
3. Beginnign definition of a new protocol. Maybe BTP (Binary Transfer Protocol)?
3. Master can send a built executable to the agent where the agent will write it to the disk on its host machine and execute it and send output back to the master application.

## BTP (Binary Transfer Protocol)

- Non human readable
- first 8 bytes define the size of the executable that is being sent
- next x bytes (defined by the first 8 bytes) are the bytes of the executable

## Questions/Issues
1. With regards to digitally signing the bytes sent from the master to protect against random people sending executables to the agents, will the signing key need to be built into the agent executable?
2. Bits vs bytes: we can save space by sending the size of the executable as a couple of bits instead of bytes among other things. How do we work with bits in Go?

## Goals For Next Meeting
1. Further define and expand capability of btp
2. Create package/library for btp so lowlevel stuff is done much more easily and in one location, rather than everywhere.
3. Allow execution to continue even is the master process disconnects, OR allow the master process to say that it does not need to see application output from the agents.