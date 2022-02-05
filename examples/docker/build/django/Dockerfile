FROM python:3.8.12-alpine3.15 as compile
ENV PATH="/app/venv/bin:$PATH"
WORKDIR /app
COPY . .
RUN python -m venv /app/venv && \
    python -m pip install --upgrade pip && \
    python -m pip install -r requirements.txt

FROM python:3.8.12-alpine3.15
ENV PATH="/app/venv/bin:$PATH"
WORKDIR /app
COPY --from=compile /app /app
EXPOSE 8000
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
