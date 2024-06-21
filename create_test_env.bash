#!/bin/bash

# Create root directory
mkdir -p project

# Create files in the root directory
touch project/script.py
touch project/module.py
touch project/README.md
touch project/requirements.txt
touch project/.env

# Create .py_env directory and files
mkdir -p project/.py_env
touch project/.py_env/config.py
touch project/.py_env/settings.py

# Create src directory and files
mkdir -p project/src
touch project/src/main.py
touch project/src/helper.py
touch project/src/utils.py

# Create src/data directory and files
mkdir -p project/src/data
touch project/src/data/data1.csv
touch project/src/data/data2.csv
touch project/src/data/processed_data.py.csv

# Create tests directory and files
mkdir -p project/tests
touch project/tests/test_main.py
touch project/tests/test_helper.py

# Create tests/data directory and files
mkdir -p project/tests/data
touch project/tests/data/test_data1.csv
touch project/tests/data/test_data2.csv
touch project/tests/data/test_script.py.csv

# Create docs directory and files
mkdir -p project/docs
touch project/docs/index.md
touch project/docs/tutorial.md

# Create docs/images directory and files
mkdir -p project/docs/images
touch project/docs/images/diagram.py.png
touch project/docs/images/flowchart.png

# Create assets directory and files
mkdir -p project/assets
touch project/assets/logo.py.svg
touch project/assets/style.css
touch project/assets/script.js

# Create logs directory and files
mkdir -p project/logs
touch project/logs/app.log
touch project/logs/error.log

# Create build/output directory and files
mkdir -p project/build/output
touch project/build/output/app.py.exe
touch project/build/output/library.so

# Create build/temp directory and files
mkdir -p project/build/temp
touch project/build/temp/temp_file.tmp

echo "File tree created successfully!"
