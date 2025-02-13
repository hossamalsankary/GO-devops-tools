FROM ollama/ollama:latest

# Set your model name as an environment variable.
ENV MODEL_NAME=llama3

# Preload the model during the build.
# This RUN step starts the Ollama service temporarily, waits for it to be up,
# pulls the model (which downloads the model files into the container’s filesystem),
# then kills the temporary server.
RUN set -ex; \
    echo "Starting Ollama server temporarily..."; \
    ollama serve & \
    pid=$!; \
    echo "Waiting for the server to be ready..."; \
    # You might adjust this sleep if needed or use a loop checking a health endpoint.
    sleep 10; \
    echo "Pulling model ${MODEL_NAME}..."; \
    ollama pull ${MODEL_NAME}; \
    echo "Stopping temporary server..."; \
    kill $pid || true; \
    wait $pid || true

# Expose the port for Ollama.
EXPOSE 11434

# When the container is run, simply start the Ollama service.
