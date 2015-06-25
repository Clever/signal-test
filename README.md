# Overview

`signal-test` is a small container meant to be used to test which signals are being sent to it.  We use it to test that mesos is sending our containers the appropriate signals on shutdown.

# Building

`docker build -t clever/signal-test .`
