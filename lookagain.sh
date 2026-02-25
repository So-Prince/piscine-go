#!/bin/bash
find . -type f -name "*.sh" | sed 's/\.sh$//; s|.*/||' | sort -r
