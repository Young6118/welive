import numpy as np
from typing import List, Optional, Dict, Any
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import cosine_similarity
import redis
import json

from app.config import settings
from app.models.schemas import RecommendationItem, ItemType


class RecommenderService:
    def __init__(self):
        self.vectorizer = TfidfVectorizer(max_features=5000)
        self.redis_client = redis.Redis(
            host=settings.REDIS_HOST,
            port=settings.REDIS_PORT,
            password=settings.REDIS_PASSWORD,
            db=settings.REDIS_DB,
            decode_responses=True
        )
    
    async def recommend_questions(
        self,
        user_id: int,
        count: int = 10,
        filters: Optional[Dict[str, Any]] = None
    ) -> List[RecommendationItem]:
        """
        基于协同过滤和内容相似度推荐问题
        """
        # TODO: 实现真实推荐逻辑
        # 目前返回模拟数据
        return [
            RecommendationItem(
                id=i,
                type=ItemType.QUESTION,
                title=f"推荐问题 {i}",
                content=f"这是推荐问题 {i} 的内容",
                score=0.9 - i * 0.05,
                reason="基于您的兴趣"
            )
            for i in range(1, min(count + 1, 11))
        ]
    
    async def recommend_notes(
        self,
        user_id: int,
        count: int = 10,
        filters: Optional[Dict[str, Any]] = None
    ) -> List[RecommendationItem]:
        """
        推荐相关笔记
        """
        return [
            RecommendationItem(
                id=i,
                type=ItemType.NOTE,
                title=f"推荐笔记 {i}",
                content=f"这是推荐笔记 {i} 的内容",
                score=0.85 - i * 0.05,
                reason="热门笔记"
            )
            for i in range(1, min(count + 1, 11))
        ]
    
    async def recommend_villages(
        self,
        user_id: int,
        count: int = 10
    ) -> List[RecommendationItem]:
        """
        推荐相关村落
        """
        return [
            RecommendationItem(
                id=i,
                type=ItemType.VILLAGE,
                title=f"推荐村落 {i}",
                content=f"这是推荐村落 {i} 的描述",
                score=0.8 - i * 0.05,
                reason="可能感兴趣"
            )
            for i in range(1, min(count + 1, 11))
        ]
    
    async def track_behavior(
        self,
        user_id: int,
        item_id: int,
        item_type: ItemType,
        action: str,
        metadata: Optional[Dict[str, Any]] = None
    ) -> bool:
        """
        追踪用户行为，用于改进推荐
        """
        try:
            key = f"user:{user_id}:behavior"
            behavior_data = {
                "item_id": item_id,
                "item_type": item_type.value,
                "action": action,
                "metadata": metadata or {}
            }
            self.redis_client.lpush(key, json.dumps(behavior_data))
            self.redis_client.ltrim(key, 0, 999)  # 保留最近1000条
            return True
        except Exception as e:
            print(f"Error tracking behavior: {e}")
            return False
    
    async def calculate_similarity(self, content1: str, content2: str) -> float:
        """
        计算两段内容的相似度
        """
        try:
            vectors = self.vectorizer.fit_transform([content1, content2])
            similarity = cosine_similarity(vectors[0:1], vectors[1:2])[0][0]
            return float(similarity)
        except Exception as e:
            print(f"Error calculating similarity: {e}")
            return 0.0
    
    async def get_hot_questions(
        self,
        category: Optional[str] = None,
        limit: int = 10
    ) -> List[Dict[str, Any]]:
        """
        获取热门问题
        """
        # TODO: 从Redis或数据库获取真实数据
        return [
            {
                "id": i,
                "title": f"热门问题 {i}",
                "hot_score": 100 - i * 5,
                "category": category or "综合"
            }
            for i in range(1, min(limit + 1, 21))
        ]
    
    async def get_hot_notes(
        self,
        category: Optional[str] = None,
        limit: int = 10
    ) -> List[Dict[str, Any]]:
        """
        获取热门笔记
        """
        return [
            {
                "id": i,
                "title": f"热门笔记 {i}",
                "hot_score": 95 - i * 5,
                "category": category or "综合"
            }
            for i in range(1, min(limit + 1, 21))
        ]
