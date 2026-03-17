# placeholder for binarize preprocessing
import cv2
import numpy as np

def apply_threshold(image: np.ndarray) -> np.ndarray:
    return cv2.adaptiveThreshold(
        image,
        255,
        cv2.ADAPTIVE_THRESH_GAUSSIAN_C,
        cv2.THRESH_BINARY,
        31,
        11
    )
