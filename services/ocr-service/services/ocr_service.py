# placeholder for OCR service orchestration
from pathlib import Path
import json
import cv2

from infrastructure.storage.local import LocalStorage
from ocr_engine.paddle import PaddleOCREngine
from preprocessing.normalize import to_grayscale
from preprocessing.denoise import apply_denoise
from preprocessing.binarize import apply_threshold
from config.settings import settings

class OCRService:
    def __init__(self) -> None:
        self.storage = LocalStorage()
        self.engine = PaddleOCREngine()

    def _resize(self, image, scale: float = 1.5):
        height, width = image.shape[:2]
        return cv2.resize(
            image,
            (int(width * scale), int(height * scale)),
            interpolation=cv2.INTER_CUBIC,
        )

    def _preprocess(self, input_path: Path) -> Path:
        image = cv2.imread(str(input_path))
        if image is None:
            raise ValueError("Image could not be loaded")
        
        gray = to_grayscale(image)
        denoised = apply_denoise(gray)
        thresholded = apply_threshold(denoised)
        resized = self._resize(thresholded, scale=1.5)

        output_path = Path(settings.processed_dir) / f"processed_{input_path.name}"
        output_path.parent.mkdir(parents=True, exist_ok=True)

        cv2.imwrite(str(output_path), resized)
        return output_path
    
    def _extract_from_dict_payload(self, payload: dict):
        text_blocks = []
        raw_lines = []

        rec_texts = payload.get("rec_texts", [])
        rec_scores = payload.get("rec_scores", [])

        for idx, text in enumerate(rec_texts):
            score = 0.0
            if idx < len(rec_scores):
                try:
                    score = float(rec_scores[idx])
                except Exception:
                    score = 0.0

            text = str(text).strip()
            if not text:
                continue

            text_blocks.append({
                "text": text,
                "confidence": round(score, 4),
            })
            raw_lines.append(text)

        return text_blocks, raw_lines

    def _extract_from_legacy_payload(self, raw_result):
        text_blocks = []
        raw_lines = []

        for group in raw_result:
            if not group:
                continue

            for item in group:
                try:
                    if (
                        isinstance(item, (list, tuple))
                        and len(item) >= 2
                        and isinstance(item[1], (list, tuple))
                        and len(item[1]) >= 2
                    ):
                        text = str(item[1][0]).strip()
                        confidence = float(item[1][1])

                        if not text:
                            continue

                        text_blocks.append({
                            "text": text,
                            "confidence": round(confidence, 4),
                        })
                        raw_lines.append(text)
                except Exception:
                    continue

        return text_blocks, raw_lines

    def _build_response(self, raw_result, original_file_path: str, processed_file_path: str) -> dict:
        text_blocks = []
        raw_lines = []

        if not raw_result:
            return {
                "raw_text": "",
                "text_blocks": [],
                "meta": {
                    "engine": "paddleocr",
                    "preprocessing_applied": True,
                    "original_file_path": original_file_path,
                    "processed_file_path": processed_file_path,
                }
            }

        if isinstance(raw_result, dict):
            payload = raw_result.get("res", raw_result)
            text_blocks, raw_lines = self._extract_from_dict_payload(payload)

        elif isinstance(raw_result, list) and len(raw_result) > 0:
            first = raw_result[0]

            if isinstance(first, dict):
                payload = first.get("res", first)
                text_blocks, raw_lines = self._extract_from_dict_payload(payload)
            else:
                text_blocks, raw_lines = self._extract_from_legacy_payload(raw_result)

        return {
            "raw_text": "\n".join(raw_lines),
            "text_blocks": text_blocks,
            "meta": {
                "engine": "paddleocr",
                "preprocessing_applied": True,
                "original_file_path": original_file_path,
                "processed_file_path": processed_file_path,
            }
        }

    def _save_debug_output(self, original_file_path: str, result: dict):
        output_dir = Path(settings.output_dir)
        output_dir.mkdir(parents=True, exist_ok=True)

        original_name = Path(original_file_path).stem
        output_file = output_dir / f"{original_name}.json"

        with open(output_file, "w", encoding="utf-8") as f:
            json.dump(result, f, ensure_ascii=False, indent=2)

    def process_image_url(self, image_url: str) -> dict:
        downloaded_path = self.storage.download_image(image_url)
        processed_path = self._preprocess(downloaded_path)
        raw_result = self.engine.extract(str(processed_path))

        result = self._build_response(
            raw_result=raw_result,
            original_file_path=str(downloaded_path),
            processed_file_path=str(processed_path),
        )
        self._save_debug_output(str(downloaded_path), result)
        return result

    def process_local_file(self, file_path: str) -> dict:
        original_path = Path(file_path)
        processed_path = self._preprocess(original_path)
        raw_result = self.engine.extract(str(processed_path))

        result = self._build_response(
            raw_result=raw_result,
            original_file_path=str(original_path),
            processed_file_path=str(processed_path),
        )
        self._save_debug_output(str(original_path), result)
        return result