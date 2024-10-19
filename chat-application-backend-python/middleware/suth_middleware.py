# middleware/auth_middleware.py
from flask_jwt_extended import jwt_required

def auth_middleware():
    return jwt_required()

