#!/bin/bash
ps aux | grep chrom | grep -v grep | awk '{print $2}' | xargs kill -9
