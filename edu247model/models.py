from sqlalchemy import Column, Integer, String, JSON, DateTime, Text
from sqlalchemy.sql import func
from db import Base

class AIModel(Base):
    __tablename__ = "ai_models"

    id = Column(Integer, primary_key=True)
    name = Column(String, index=True)
    version = Column(String)
    description = Column(Text)
    created_at = Column(DateTime, server_default=func.now())
    metadata = Column(JSON)
    artifact_path = Column(String)
