# uptimo (Uplink uptime Monitor)

Simple http base uptime / downtime monitor

I needed a tool to simply monitor the availability of my internet uplink.
It will report the end time of a downtime and it's duration. To measure
the uptime http requests are send to a remote host.

        $ ./uptimo -duration 1s -remote-host 10.0.1.254
        2018-03-09T21:06:10+01:00       1m54.0893118
