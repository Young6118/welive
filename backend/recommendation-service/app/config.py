import os
from pydantic_settings import BaseSettings


class Settings(BaseSettings):
    # 服务配置
    HOST: str = "0.0.0.0"
    PORT: int = 8000
    DEBUG: bool = True
    
    # Redis配置
    REDIS_HOST: str = os.getenv("REDIS_HOST", "localhost")
    REDIS_PORT: int = int(os.getenv("REDIS_PORT", "6379"))
    REDIS_PASSWORD: str = os.getenv("REDIS_PASSWORD", "")
    REDIS_DB: int = int(os.getenv("REDIS_DB", "0"))
    
    # 算法配置
    RECOMMENDATION_COUNT: int = 10  # 默认推荐数量
    SIMILARITY_THRESHOLD: float = 0.5  # 相似度阈值
    
    class Config:
        env_file = ".env"


settings = Settings()
