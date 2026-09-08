from sqlalchemy import Column, Integer, Text, String
from sqlalchemy.orm import declarative_base
from pgvector.sqlalchemy import Vector
from db import Base

class AIEmbedding(Base):
    __tablename__ = "ai_embeddings"

    id = Column(Integer, primary_key=True)
    user_id = Column(String, index=True)
    text = Column(Text)
    embedding = Column(Vector(384))   # match your model dimension
