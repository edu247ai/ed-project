import time
from fastapi import Request

async def request_logger(request: Request, call_next):
    start = time.time()
    response = await call_next(request)
    duration = round(time.time() - start, 4)

    print(f"[API] {request.method} {request.url.path} - {duration}s")

    return response
