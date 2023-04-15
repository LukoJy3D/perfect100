#!/bin/bash

# Clone repository if it doesn't exist
if [ ! -d gpt4all ]; then
    git clone https://github.com/nomic-ai/gpt4all.git
else
    echo "Repository already exists"
fi

# Download file unless it already exists
if [ ! -f gpt4all-lora-quantized.bin ]; then
    curl -L -o gpt4all-lora-quantized.bin https://the-eye.eu/public/AI/models/nomic-ai/gpt4all/gpt4all-lora-quantized.bin
else
    echo "File already exists"
fi
