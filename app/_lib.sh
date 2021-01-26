#!/bin/bash
####################################
#####
###
##
#
#
# Startup Script for the Application
####################################


function RootCheck() {
  if [ "$EUID" -ne 0 ]; then
    echo "Permission Denied."
    exit 1
  fi
}
