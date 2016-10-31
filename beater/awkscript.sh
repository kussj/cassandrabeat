#!/bin/bash
# awk script to parse nodetool cfstats output

####
#
# 6 -> Pending Flushes (int)
# 8 -> SSTable count (int) 
# 9 -> Space used (live) (int)
# 10 -> Space used (total) (int)
# 11 -> Space used by snapshots (total) (int)
# 14 -> Number of keys (estimate) (int)
# 20 -> Local read latency (float)
# 22 -> Local write latency (float)
#
####

nodetool cfstats $1 | awk '
			FNR == 6 { print $3 }
			FNR == 8 { print $3 }
			FNR == 9 { print $4 }
			FNR == 10 { print $4 }
			FNR == 11 { print $6 }
			FNR == 14 { print $5 }
			FNR == 20 { print $4 }
			FNR == 22 { print $4 } '