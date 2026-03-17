from pydantic import BaseModel, HttpUrl

class OCRRequest(BaseModel):
    image_url: HttpUrl
