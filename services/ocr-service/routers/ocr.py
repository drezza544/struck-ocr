from pathlib import Path
from fastapi import APIRouter, HTTPException, UploadFile, File
from schemas.request import OCRRequest
from schemas.response import OCRResponse
from services.ocr_service import OCRService
from config.settings import settings

router = APIRouter()
ocr_service = OCRService()

@router.post("/v1/ocr", response_model=OCRResponse)
def run_ocr(payload: OCRRequest):
    try:
        result = ocr_service.process_image_url(str(payload.image_url))
        return result
    except Exception as exc:
        raise HTTPException(status_code=500, detail=str(exc))

@router.post("/v1/ocr/upload", response_model=OCRResponse)
async def run_ocr_upload(file: UploadFile = File(...)):
    try:
        suffix = Path(file.filename).suffix.lower()
        if suffix not in [".jpg", ".jpeg", ".png", ".webp"]:
            raise HTTPException(status_code=400, detail="Unsupported file type..")

        temp_dir = Path(settings.temp_dir)
        temp_dir.mkdir(parents=True, exist_ok=True)

        file_path = temp_dir / file.filename
        with open(file_path, "wb") as f:
            f.write(await file.read())

        result = ocr_service.process_local_file(str(file_path))
        return result
    except HTTPException:
        raise
    except Exception as exc:
        raise HTTPException(status_code=500, detail=str(exc))


