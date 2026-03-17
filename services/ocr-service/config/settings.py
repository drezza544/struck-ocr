from pydantic_settings import BaseSettings, SettingsConfigDict

class Settings(BaseSettings):
    env: str = "development"
    log_level: str = "info"

    go_api_port: int = 8080
    go_api_base_url: str = "http://localhost:8080"

    ocr_service_port: int = 8000
    ocr_service_base_url: str = "http://ocr-service:8000"
    ocr_engine: str = "paddleocr"
    ocr_lang: str = "en"
    ocr_use_gpu: bool = False

    postgres_host: str = "postgres"
    postgres_port: int = 5432
    postgres_db: str = "struck_ocr"
    postgres_user: str = "reza"
    postgres_password: str = "123456"

    storage_base_path: str = "./storage"
    temp_dir: str = "./storage/temp"
    processed_dir: str = "./storage/processed"
    output_dir: str = "./storage/output"

    request_timeout_seconds: int = 30

    model_config = SettingsConfigDict(
        env_file=".env",
        env_file_encoding="utf-8",
        extra="ignore",
    )

settings = Settings()
