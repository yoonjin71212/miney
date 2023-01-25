#!/bin/bash
DIRECTORY="$1"
cd "$DIRECTORY"
systemctl restart nginx
