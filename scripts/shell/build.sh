#! /bin/bash

declare -A apps=(["api"]="api" ["drive"]="cli" ["worker"]="worker")

for app in "${!apps[@]}"; do
  echo "building $app..."
  go build -o $app ./cmd/${apps[$app]}

done 
