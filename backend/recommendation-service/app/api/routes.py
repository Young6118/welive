from fastapi import APIRouter, Query
from typing import List, Optional

from app.models.schemas import (
    RecommendationRequest,
    RecommendationResponse,
    UserBehaviorRequest,
    SimilarityRequest
)
from app.services.recommender import RecommenderService

router = APIRouter()
recommender_service = RecommenderService()


@router.post("/recommend/questions", response_model=RecommendationResponse)
async def recommend_questions(request: RecommendationRequest):
    """
    推荐相关问题
    """
    recommendations = await recommender_service.recommend_questions(
        user_id=request.user_id,
        count=request.count,
        filters=request.filters
    )
    return RecommendationResponse(
        items=recommendations,
        total=len(recommendations)
    )


@router.post("/recommend/notes", response_model=RecommendationResponse)
async def recommend_notes(request: RecommendationRequest):
    """
    推荐相关笔记
    """
    recommendations = await recommender_service.recommend_notes(
        user_id=request.user_id,
        count=request.count,
        filters=request.filters
    )
    return RecommendationResponse(
        items=recommendations,
        total=len(recommendations)
    )


@router.post("/recommend/villages", response_model=RecommendationResponse)
async def recommend_villages(request: RecommendationRequest):
    """
    推荐相关村落
    """
    recommendations = await recommender_service.recommend_villages(
        user_id=request.user_id,
        count=request.count
    )
    return RecommendationResponse(
        items=recommendations,
        total=len(recommendations)
    )


@router.post("/behavior/track")
async def track_behavior(request: UserBehaviorRequest):
    """
    追踪用户行为
    """
    result = await recommender_service.track_behavior(
        user_id=request.user_id,
        item_id=request.item_id,
        item_type=request.item_type,
        action=request.action,
        metadata=request.metadata
    )
    return {"success": result}


@router.post("/similarity/calculate")
async def calculate_similarity(request: SimilarityRequest):
    """
    计算内容相似度
    """
    similarity = await recommender_service.calculate_similarity(
        content1=request.content1,
        content2=request.content2
    )
    return {"similarity": similarity}


@router.get("/hot/questions")
async def get_hot_questions(
    category: Optional[str] = Query(None, description="分类筛选"),
    limit: int = Query(10, ge=1, le=100)
):
    """
    获取热门问题
    """
    hot_questions = await recommender_service.get_hot_questions(
        category=category,
        limit=limit
    )
    return {"items": hot_questions, "total": len(hot_questions)}


@router.get("/hot/notes")
async def get_hot_notes(
    category: Optional[str] = Query(None, description="分类筛选"),
    limit: int = Query(10, ge=1, le=100)
):
    """
    获取热门笔记
    """
    hot_notes = await recommender_service.get_hot_notes(
        category=category,
        limit=limit
    )
    return {"items": hot_notes, "total": len(hot_notes)}
