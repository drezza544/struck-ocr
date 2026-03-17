# placeholder for local storage
from pathlib import Path
import requests
from urllib.parse import urlparse
from config.settings import settings

class LocalStorage:
    def __init__(self) -> None:
        self.base_path = Path(settings.storage_base_path)
        self.temp_dir = Path(settings.temp_dir)
        self.processed_dir = Path(settings.processed_dir)
        self.output_dir = Path(settings.output_dir)

        self.base_path.mkdir(parents=True, exist_ok=True)
        self.temp_dir.mkdir(parents=True, exist_ok=True)
        self.processed_dir.mkdir(parents=True, exist_ok=True)
        self.output_dir.mkdir(parents=True, exist_ok=True)

    def download_image(self, image_url: str) -> Path:
        parsed = urlparse(image_url)
        filename = Path(parsed.path).name or "receipt.jpg"
        output_path = self.temp_dir / filename

        response = requests.get(image_url, timeout=settings.request_timeout_seconds)
        response.raise_for_status()

        output_path.write_bytes(response.content)
        return output_path