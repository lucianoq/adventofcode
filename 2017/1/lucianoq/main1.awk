#!/bin/awk -f
BEGIN {
	FS=""
	sum=0
}
{
	for (i=1;i<=NF;i++) {
		#print $i
		print $(i+1)
		if ( $i == $(i+1) ) sum += $i
	}
}
END {
	print $sum
}

