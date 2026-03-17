# placeholder for paddle adapter
from paddleocr import PaddleOCR

class PaddleOCREngine:
    def __init__(self) -> None:
        self.client = PaddleOCR(
            use_angle_cls=True,
            lang="en",
        )
    
    def extract(self, image_path: str):
        result = self.client.ocr(image_path)

        print("PADDLE RAW RESULT TYPE: ", type(result))
        print("PADDLE RAW RESULT: ", result)

        return result