#!/bin/bash

# Pull the specified model
echo "Pulling model: ${MODEL_NAME}"
ollama pull ${MODEL_NAME}

# Start the Ollama server
echo "Starting Ollama server..."
exec ollama run ${MODEL_NAME}
