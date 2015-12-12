name             'conduit'
maintainer       'Orlando Golan Meetup'
long_description IO.read(File.join(File.dirname(__FILE__), 'README.md'))

recipe "conduit::golang", "Installs golang 1.5.1"
recipe "conduit::supervisor", "Installs and configures supervisor for hello world application"
recipe "conduit::hello_world", "Simple hello world application to test eps-conduit"
recipe "conduit::compile_daemon", "Installs Compile Daemon for easy development"
recipe "conduit::build", "Runs initial build of project"
recipe "conduit::supervisor", "Install and runs supervisor"
