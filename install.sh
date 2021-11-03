#!/bin/bash

echo "Downloading k3ai..."
sleep 1

wget https://github.com/k3ai/k3ai/releases/download/1.0/k3ai

chmod +x k3ai

sudo mv k3ai /usr/local/bin

echo "Done. Check our Docs at https://k3ai.in to start."
