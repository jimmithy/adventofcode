#!/bin/bash

first=""
second=""
third=""
lastWindow=""
total=""
count=0
while IFS= read -r line; do
  if [[ -z "$first" ]]; then
    first="$line"
  elif [[ -z "$second" ]]; then
    second="$line"
  elif [[ -z "$third" ]]; then
    third="$line"
    ((lastWindow = first + second + third))
  else
    first="$second"
    second="$third"
    third="$line"
    ((total = first + second + third))

    if [[ $lastWindow -lt $total ]]; then
      ((count = count + 1))
    fi

    lastWindow="$total"
  fi
done <"$1"

echo "$count"