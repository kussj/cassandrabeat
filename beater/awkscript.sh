#!/bin/bash
# awk script to parse nodetool cfstats output

nodetool cfstats $1 | awk '
			FNR == 20 { print $4 }
			FNR == 22 { print $4 } '

