PYV ?= python3.11
PY ?= .venv/bin/python
PIP ?= .venv/bin/pip

deps:
	$(PIP) install -r requirements-dev.txt
	$(PIP) install --upgrade pip
	xargs poetry add < requirements.txt

format:
	$(PY) -m black -l 79 --fast src tests
	$(PY) -m isort src tests

lint:
	$(PY) -m pylint src tests
	$(PY) -m mypy src tests

venv:
	$(PYV) -m venv .venv
