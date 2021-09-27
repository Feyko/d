# D
D is a cli tool aimed at boosting productivity. The features it currently provides are administration tasks I often find myself doing manually.  
I have created this for myself and will continue adding features that I find myself missing. I believe it may be of use to others once it has grown a bit  
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