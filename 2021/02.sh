#!/bin/bash

forward=0
depth=0
aim=0
while IFS= read -r line; do
  if [[ $line == "forward"* ]]; then
    count=${line#forward }
    ((forward = forward + count))
    ((depth = depth + aim * count))
  elif [[ $line == "up"* ]]; then
    count=${line#up }
    ((aim = aim - count))
  elif [[ $line == "down"* ]]; then
    count=${line#down }
    ((aim = aim + count))
  fi

done <"$1"

echo "$forward * $depth"

