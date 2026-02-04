from pydantic import BaseModel
from typing import List, Optional, Dict, Any
from enum import Enum


class ItemType(str, Enum):
    QUESTION = "question"
    ANSWER = "answer"
    NOTE = "note"
    POST = "post"
    VILLAGE = "village"


class ActionType(str, Enum):
    VIEW = "view"
    LIKE = "like"
    COMMENT = "comment"
    SHARE = "share"
    COLLECT = "collect"


class RecommendationItem(BaseModel):
    id: int
    type: ItemType
    title: Optional[str] = None
    content: str
    score: float
    reason: Optional[str] = None


class RecommendationRequest(BaseModel):
    user_id: int
    count: int = 10
    filters: Optional[Dict[str, Any]] = None


class RecommendationResponse(BaseModel):
    items: List[RecommendationItem]
    total: int


class UserBehaviorRequest(BaseModel):
    user_id: int
    item_id: int
    item_type: ItemType
    action: ActionType
    metadata: Optional[Dict[str, Any]] = None


class SimilarityRequest(BaseModel):
    content1: str
    content2: str
