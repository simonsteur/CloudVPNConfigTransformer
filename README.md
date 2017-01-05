# CloudVPNConfigTransformer
CVCT is a tool that takes the generated vpn configuration from your cloud provider and allows you to transform it to your environment.

#Why
The idea comes from a time when I was working for a company that was starting to utilize different aws regions at a rapid pace. Being the person to both build the infrastructure in aws and also the network administrator for our on-premises environment; a tool like this became very useful for me.

#Use 
Download the version for your platform, add it to your path. After that you can run "cvct" to test out if everything is working.


#Example
```bash
cvct screenos --file <path to generated config> --if1 <1st tunnel interface> --if2 <2nd tunnel interface> --zone <zone to bind the tunnel interfaces to> --eif <external interface for vpn> --cidr <network range for routing> --nc
```
#Help
Each command has a help function to use just use the `-h` flag
```bash
cvct -h

NAME:
   cvct - transform generated vpn configuration from cloud providers to represent your environment

USAGE:
   cvct [global options] command [command options] [arguments...]

VERSION:
   0.1

COMMANDS:
     screenos  use when dealing with ScreenOS configuration
     junos     use when dealing with junos configuration
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

```