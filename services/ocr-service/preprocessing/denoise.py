# placeholder for denoise preprocessing
import cv2
import numpy as np

def apply_denoise(image: np.ndarray) -> np.ndarray:
    return cv2.fastNlMeansDenoising(image, None, 15, 7, 21)