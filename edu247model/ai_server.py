from fastapi import FastAPI, Depends
from pydantic import BaseModel
from sqlalchemy import insert, select
from db import get_db
from models import AIEmbedding
from embeddings import embed_text

app = FastAPI()

# ---------------------------
# 1. Chat endpoint
# ---------------------------
class ChatRequest(BaseModel):
    user_id: str
    message: str

@app.post("/ai/chat")
async def ai_chat(req: ChatRequest, db=Depends(get_db)):
    reply = f"AI: {req.message}"

    stmt = insert(AIEmbedding).values(
        user_id=req.user_id,
        text=req.message,
        embedding=embed_text(req.message)
    )
    await db.execute(stmt)
    await db.commit()

    return {"reply": reply}

# ---------------------------
# 2. Search endpoint (ADD HERE)
# ---------------------------
class SearchRequest(BaseModel):
    user_id: str
    query: str
    limit: int = 5

@app.post("/ai/search")
async def ai_search(req: SearchRequest, db=Depends(get_db)):
    qvec = embed_text(req.query)

    stmt = (
        select(AIEmbedding)
        .where(AIEmbedding.user_id == req.user_id)
        .order_by(AIEmbedding.embedding.cosine_distance(qvec))
        .limit(req.limit)
    )

    rows = (await db.execute(stmt)).scalars().all()

    return [{"text": r.text} for r in rows]
