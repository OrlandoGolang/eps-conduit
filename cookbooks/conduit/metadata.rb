name             'conduit'
maintainer       'Orlando Golan Meetup'
long_description IO.read(File.join(File.dirname(__FILE__), 'README.md'))

recipe "conduit::golang", "Installs golang 1.5.1"
recipe "conduit::hello_world", "Installs and configures supervisor for hello world application"
