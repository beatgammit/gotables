gotables
========

`gotables` is a very simple replacement for iptables for simple port forwarding. It listens on a given host and port and pipes all traffic to a remote host and port.

installation
============

go get github.com/beatgammit/gotables

configuration
=============

The configuration file format is a very simple space-delimited format of `<src> <dest>`. For example, the following will forward all traffic received by the host (the computer `gotables` is running on) on port 10000 to 192.168.1.100 on port 20000:

    # forward all traffic from port 10000 to 192.168.1.100 on port 20000
    0.0.0.0:10000 192.168.1.100:20000

Comments consist of a leading `#`, as above.

license
=======

`gotables` is licensed under the BSD 2-clause license. Please check LICENSE.BSD for specifics.
