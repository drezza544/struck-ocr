from fastapi import FastAPI
from config.settings import settings
from routers.health import router as health_router
from routers.ocr import router as ocr_router

app = FastAPI(
    title="struck-ocr-ocr-service",
    version="0.1.0",
)

app.include_router(health_router)
app.include_router(ocr_router)
