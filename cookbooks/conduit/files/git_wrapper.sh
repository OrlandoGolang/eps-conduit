#!/bin/bash
exec ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i "/home/vagrant/.ssh/eps-conduit-hello" "$@"
