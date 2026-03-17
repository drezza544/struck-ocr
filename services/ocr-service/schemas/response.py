from pydantic import BaseModel
from typing import List

class OCRTextBlock(BaseModel):
    text: str
    confidence: float

class OCRMeta(BaseModel):
    engine: str
    preprocessing_applied: bool
    original_file_path: str
    processed_file_path: str

class OCRResponse(BaseModel):
    raw_text: str
    text_blocks: List[OCRTextBlock]
    meta: OCRMeta
