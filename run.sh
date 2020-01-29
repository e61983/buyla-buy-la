#!/bin/bash

echo "======================="
echo "url: ${TEST_URL}:${PORT}"
echo "======================="

ssh -R ${TEST_URL}:${PORT}:localhost:8080 serveo.net -l buyla
