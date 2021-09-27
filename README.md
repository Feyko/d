# D
###### D stands for Dev(elopment). It was brought down to a minimal name to ease its use at it aims at being called often from the terminal
D is a cli tool aimed at boosting productivity. The features it currently provides are administration tasks I often find myself doing manually  
I have created this for myself and will continue adding features I find are missing. I believe it may be of use to others once it has grown a bit  
Its ultimate goal is to allow for complete  configuration of a fresh distro, so you can feel at home on any fresh installation after running just one command, provided you have filled extensive configuration pre-emptively  
Contributions are more than welcome <3 Please open an issue first so that we can discuss the features you want to add

Disclaimer: D is built for UNIX systems and has only been tested on Pop!_OS 21.04. It should work without a problem on Debian and Ubuntu systems. No guarantees about Mac. And don't even think about Windows!

# Features
## FMAs
D can add and remove FileManager Actions. You can configure their names, execution command and on what they act
## Aliases
D can add, remove and check for the existence of aliases. This is all done via the .bashrc file for the current user
## PATH
D can manage the PATH. You can add, remove and list parts of the PATH. This is also all done via the .bashrc file for the current user
## .bashrc
D can add line to the .bashrc file for the calling user. It is planned to allow for removal and checking of lines too (though I don't see the use myself)

# Roadmap
This roadmap provides no assurance that any of it will be implemented one day. It only describes what I would like this tool to take care of in the future. It also serves as a repository for my implementation ideas, which are subject to change in the future. Feedback welcomed!

## Configuration
This tool should provide a way to save and load configurations. For example, running the tool on a config file could make sure that all your aliases are set, certain programs are installed, some FMAs are present and all your programs are up to date  
This may be achieved by YAML files describing actions to take and optional checks used to gate these actions. These actions could simply be commands to call, naturally allowing the user to make use of D's systems to set everything up etc.  
Allowing these configuration files to be read from a remote source would allow for the kind of "one command setup" this tool aims to ultimately provide

## Syncing
Syncing files and folders across machines via git, S3 or other storage methods could be instrumental for the goal of this tool  
The current idea is pretty simple, simply add certain files and/or folders to a D configuration file, choose events/intervals for syncing (pushing/pulling) the files and a source to sync to/from

## SSH
Managing SSH has to be the most friction I've had when developing on Linux. I haven't explored other SSH management avenues and may drop this from projected feature set of D, but this seems to fit the bill well for this tool

## Running commands during certain events
This is another instrumental part needed for D. Easily running commands on common events is important for easy but complete management of a system