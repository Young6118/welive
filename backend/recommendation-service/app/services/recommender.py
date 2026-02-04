import numpy as np
from typing import List, Optional, Dict, Any
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import cosine_similarity
import redis
import json
import mysql.connector
from datetime import datetime, timedelta

from app.config import settings
from app.models.schemas import RecommendationItem, ItemType


class RecommenderService:
    def __init__(self):
        self.vectorizer = TfidfVectorizer(max_features=5000, stop_words='english')
        self.redis_client = redis.Redis(
            host=settings.REDIS_HOST,
            port=settings.REDIS_PORT,
            password=settings.REDIS_PASSWORD,
            db=settings.REDIS_DB,
            decode_responses=True
        )
        self.db_config = {
            'host': settings.MYSQL_HOST,
            'port': settings.MYSQL_PORT,
            'user': settings.MYSQL_USER,
            'password': settings.MYSQL_PASSWORD,
            'database': settings.MYSQL_DATABASE
        }
    
    def _get_db_connection(self):
        return mysql.connector.connect(**self.db_config)
    
    def _get_user_interactions(self, user_id: int) -> Dict[str, Any]:
        """获取用户交互数据"""
        conn = self._get_db_connection()
        cursor = conn.cursor(dictionary=True)
        
        # 获取用户点赞的问题
        cursor.execute("""
            SELECT question_id FROM question_likes WHERE user_id = %s
        """, (user_id,))
        liked_questions = [row['question_id'] for row in cursor.fetchall()]
        
        # 获取用户浏览/互动的问题内容
        cursor.execute("""
            SELECT q.id, q.title, q.content, q.tags 
            FROM questions q 
            JOIN question_likes ql ON q.id = ql.question_id 
            WHERE ql.user_id = %s AND q.status = 1
        """, (user_id,))
        user_questions = cursor.fetchall()
        
        # 获取用户点赞的笔记
        cursor.execute("""
            SELECT note_id FROM note_likes WHERE user_id = %s
        """, (user_id,))
        liked_notes = [row['note_id'] for row in cursor.fetchall()]
        
        cursor.close()
        conn.close()
        
        return {
            'liked_questions': liked_questions,
            'liked_notes': liked_notes,
            'user_questions': user_questions
        }
    
    def _calculate_hot_score(self, likes: int, comments: int, created_at: datetime) -> float:
        """计算热度分数 - 基于点赞、评论和时间衰减"""
        # 时间衰减因子
        hours_since_created = (datetime.now() - created_at).total_seconds() / 3600
        time_decay = np.exp(-hours_since_created / 168)  # 一周衰减
        
        # 互动分数
        engagement_score = likes * 2 + comments * 3
        
        return engagement_score * time_decay
    
    async def recommend_questions(
        self,
        user_id: int,
        count: int = 10,
        filters: Optional[Dict[str, Any]] = None
    ) -> List[RecommendationItem]:
        """
        基于热度、用户兴趣和协同过滤推荐问题
        """
        conn = self._get_db_connection()
        cursor = conn.cursor(dictionary=True)
        
        # 获取用户交互数据
        user_data = self._get_user_interactions(user_id)
        liked_question_ids = set(user_data['liked_questions'])
        
        # 获取所有活跃问题
        cursor.execute("""
            SELECT q.id, q.title, q.content, q.tags, q.likes, q.views, q.created_at,
                   u.username as author_name,
                   (SELECT COUNT(*) FROM answers WHERE question_id = q.id AND status = 1) as answer_count
            FROM questions q
            JOIN users u ON q.author_id = u.id
            WHERE q.status = 1
            ORDER BY q.created_at DESC
            LIMIT 1000
        """)
        questions = cursor.fetchall()
        
        cursor.close()
        conn.close()
        
        if not questions:
            return []
        
        # 计算推荐分数
        scored_questions = []
        for q in questions:
            if q['id'] in liked_question_ids:
                continue  # 跳过已点赞的
            
            # 热度分数
            hot_score = self._calculate_hot_score(
                q['likes'], 
                q['answer_count'], 
                q['created_at']
            )
            
            # 内容相似度分数（如果有用户历史）
            content_score = 0
            if user_data['user_questions']:
                user_texts = [uq['title'] + ' ' + uq['content'] for uq in user_data['user_questions']]
                user_texts.append(q['title'] + ' ' + q['content'])
                try:
                    tfidf_matrix = self.vectorizer.fit_transform(user_texts)
                    similarities = cosine_similarity(tfidf_matrix[-1:], tfidf_matrix[:-1])
                    content_score = np.max(similarities) * 10
                except:
                    content_score = 0
            
            # 综合分数
            total_score = hot_score * 0.6 + content_score * 0.4
            
            scored_questions.append({
                'question': q,
                'score': total_score,
                'hot_score': hot_score,
                'content_score': content_score
            })
        
        # 排序并返回前N个
        scored_questions.sort(key=lambda x: x['score'], reverse=True)
        
        recommendations = []
        for item in scored_questions[:count]:
            q = item['question']
            reason = "热门推荐" if item['hot_score'] > item['content_score'] else "基于您的兴趣"
            
            recommendations.append(RecommendationItem(
                id=q['id'],
                type=ItemType.QUESTION,
                title=q['title'],
                content=q['content'][:200] + '...' if len(q['content']) > 200 else q['content'],
                score=min(item['score'] / 100, 0.99),
                reason=reason,
                metadata={
                    'likes': q['likes'],
                    'author': q['author_name'],
                    'answer_count': q['answer_count']
                }
            ))
        
        return recommendations
    
    async def recommend_notes(
        self,
        user_id: int,
        count: int = 10,
        filters: Optional[Dict[str, Any]] = None
    ) -> List[RecommendationItem]:
        """
        推荐相关笔记
        """
        conn = self._get_db_connection()
        cursor = conn.cursor(dictionary=True)
        
        # 获取用户交互数据
        user_data = self._get_user_interactions(user_id)
        liked_note_ids = set(user_data['liked_notes'])
        
        # 获取所有活跃笔记
        cursor.execute("""
            SELECT n.id, n.title, n.content, n.category, n.tags, n.created_at,
                   u.username as author_name,
                   (SELECT COUNT(*) FROM note_likes WHERE note_id = n.id) as like_count,
                   (SELECT COUNT(*) FROM comments WHERE target_id = n.id AND target_type = 'note' AND status = 1) as comment_count
            FROM notes n
            JOIN users u ON n.author_id = u.id
            WHERE n.status = 1
            ORDER BY n.created_at DESC
            LIMIT 500
        """)
        notes = cursor.fetchall()
        
        cursor.close()
        conn.close()
        
        if not notes:
            return []
        
        # 计算推荐分数
        scored_notes = []
        for n in notes:
            if n['id'] in liked_note_ids:
                continue
            
            hot_score = self._calculate_hot_score(
                n['like_count'], 
                n['comment_count'], 
                n['created_at']
            )
            
            scored_notes.append({
                'note': n,
                'score': hot_score
            })
        
        scored_notes.sort(key=lambda x: x['score'], reverse=True)
        
        recommendations = []
        for item in scored_notes[:count]:
            n = item['note']
            recommendations.append(RecommendationItem(
                id=n['id'],
                type=ItemType.NOTE,
                title=n['title'],
                content=n['content'][:200] + '...' if len(n['content']) > 200 else n['content'],
                score=min(item['score'] / 100, 0.99),
                reason="热门笔记",
                metadata={
                    'category': n['category'],
                    'author': n['author_name'],
                    'likes': n['like_count']
                }
            ))
        
        return recommendations
    
    async def recommend_villages(
        self,
        user_id: int,
        count: int = 10
    ) -> List[RecommendationItem]:
        """
        推荐相关村落
        """
        conn = self._get_db_connection()
        cursor = conn.cursor(dictionary=True)
        
        # 获取用户已加入的村落
        cursor.execute("""
            SELECT village_id FROM village_members WHERE user_id = %s
        """, (user_id,))
        joined_villages = set(row['village_id'] for row in cursor.fetchall())
        
        # 获取活跃村落
        cursor.execute("""
            SELECT v.id, v.name, v.description, v.category, v.member_count, v.post_count
            FROM villages v
            WHERE v.status = 1
            ORDER BY v.member_count DESC, v.post_count DESC
            LIMIT 100
        """)
        villages = cursor.fetchall()
        
        cursor.close()
        conn.close()
        
        recommendations = []
        for v in villages:
            if v['id'] in joined_villages:
                continue
            
            score = min((v['member_count'] + v['post_count'] * 2) / 100, 0.99)
            recommendations.append(RecommendationItem(
                id=v['id'],
                type=ItemType.VILLAGE,
                title=v['name'],
                content=v['description'] or f"{v['category']}村落",
                score=score,
                reason="热门村落",
                metadata={
                    'member_count': v['member_count'],
                    'post_count': v['post_count'],
                    'category': v['category']
                }
            ))
        
        return recommendations[:count]


# 全局推荐服务实例
recommender_service = RecommenderService()
