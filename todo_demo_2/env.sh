#!/bin/bash
set -eu


python -m venv env
source env/bin/activate
pip install Flask Flask-WTF psycopg2-binary